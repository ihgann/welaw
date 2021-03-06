package services

import (
	"context"
	"strconv"
	"strings"

	apiv1 "github.com/welaw/welaw/api/v1"
	"github.com/welaw/welaw/pkg/errs"
	"github.com/welaw/welaw/pkg/permissions"
)

func (svc service) CreateVote(ctx context.Context, vote *apiv1.Vote, opts *apiv1.CreateVoteOptions) (v *apiv1.Vote, err error) {
	_, ok := ctx.Value("user_id").(string)
	if !ok {
		return nil, errs.ErrUnauthorized
	}
	//perm, err := svc.hasPermission(uid, permissions.OpVoteCreate)
	//if err != nil {
	//return nil, err
	//}
	//if !perm {
	//return nil, errs.ErrUnauthorized
	//}
	switch {
	// case opts.ReqType == apiv1.CreateVoteOptions_BY_LATEST:
	case vote.Username != "" && vote.Branch != "" && opts.Version == "latest":
		v, err = svc.db.CreateVoteByLatest(vote)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value") {
				return nil, errs.ErrConflict
			}
			return nil, err
		}
		return v, nil
	case vote.Username != "" && vote.Branch != "" && opts.Version != "":
		version, err := strconv.Atoi(opts.Version)
		if err != nil {
			return nil, err
		}
		vote.Version = uint32(version)
		v, err = svc.db.CreateVoteByVersion(vote)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value") {
				return nil, errs.ErrConflict
			}
			return nil, err
		}
		return v, nil
	default:
		svc.logger.Log("bad create_vote_opts", opts)
	}
	return nil, errs.ErrBadRequest
}

func (svc service) CreateVotes(ctx context.Context, votes []*apiv1.Vote, opts *apiv1.CreateVotesOptions) (v []*apiv1.Vote, err error) {
	username, usernameFound := ctx.Value("username").(string)
	password, passwordFound := ctx.Value("password").(string)
	if usernameFound && username != "" && passwordFound {
		pass, err := svc.db.AuthorizeUser(username, password, permissions.OpVotesCreate)
		if err != nil {
			return nil, err
		}
		if !pass {
			return nil, errs.ErrUnauthorized
		}
	} else {
		return nil, errs.ErrUnauthorized
	}
	for _, vote := range votes {
		_, err = svc.db.CreateVoteByLatest(vote)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value") {
				svc.logger.Log("warning", "duplicate vote")
				//return nil, errs.ErrConflict
				continue
			}
			return nil, err
		}
	}
	if opts.GetVoteResult() != nil {
		err := svc.db.CreateVoteResult(opts.GetVoteResult())
		if err != nil {
			return votes, err
		}
	}
	return votes, nil
}

func (svc service) DeleteVote(ctx context.Context, upstream, ident string, opts *apiv1.DeleteVoteOptions) (err error) {
	uid, ok := ctx.Value("user_id").(string)
	if !ok {
		return errs.ErrUnauthorized
	}
	perm, err := svc.hasPermission(uid, permissions.OpVoteDelete)
	if err != nil {
		return err
	}
	if !perm {
		return errs.ErrUnauthorized
	}
	switch {
	// case opts.ReqType == apiv1.DeleteVoteOptions_BY_USER_VERSION:
	case opts.Username != "" && opts.Branch != "" && opts.Version != "":
		version, err := strconv.Atoi(opts.Version)
		if err != nil {
			return err
		}
		err = svc.db.DeleteVote(opts.Username, upstream, ident, opts.Branch, uint32(version))
		if err != nil {
			return err
		}
	default:
		return errs.ErrBadRequest
	}
	return nil
}

func (svc service) GetVote(ctx context.Context, upstream, ident string, opts *apiv1.GetVoteOptions) (*apiv1.Vote, error) {
	switch {
	case opts == nil:
		return nil, errs.ErrBadRequest
	case opts.ReqType == apiv1.GetVoteOptions_BY_USER_VERSION:
		v, err := svc.getVote(opts.Username, upstream, ident, opts.Branch, opts.Version)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
	return nil, errs.ErrBadRequest
}

func (svc service) ListVotes(_ context.Context, opts *apiv1.ListVotesOptions) (*apiv1.ListVotesReply, error) {
	var votes []*apiv1.Vote
	var total int32
	var err error
	switch {
	case opts.Upstream != "" && opts.Ident != "" && opts.Branch != "" && opts.Version != 0:
		votes, total, err = svc.db.ListVersionVotes(opts.Upstream, opts.Ident, opts.Branch, opts.Version, opts.PageNum, opts.PageSize)
	case opts.Username != "":
		votes, total, err = svc.db.ListUserVotes(opts.Username, opts.PageNum, opts.PageSize)
	default:
		return nil, errs.ErrBadRequest
	}
	if err != nil {
		return nil, err
	}
	return &apiv1.ListVotesReply{
		Votes: votes,
		Total: total,
	}, nil
}

func (svc service) UpdateVote(ctx context.Context, vote *apiv1.Vote, opts *apiv1.UpdateVoteOptions) (v *apiv1.Vote, err error) {
	uid, ok := ctx.Value("user_id").(string)
	if !ok {
		return nil, errs.ErrUnauthorized
	}
	perm, err := svc.hasPermission(uid, permissions.OpVoteUpdate)
	if err != nil {
		return nil, err
	}
	if !perm {
		return nil, errs.ErrUnauthorized
	}
	v, err = svc.getVote(vote.Username, vote.Upstream, vote.Ident, vote.Branch, vote.VersionId)
	if err != nil {
		return
	}
	copyVote(v, vote)
	v, err = svc.db.UpdateVote(v.Uid, v)
	if err != nil {
		return
	}
	return
}

func (svc service) getVote(username, upstream, ident, branch, version string) (vote *apiv1.Vote, err error) {
	if hasBlank(username, upstream, ident, branch) {
		return nil, errs.ErrBadRequest
	}
	if version == "" || version == "latest" {
		vote, err = svc.db.GetVoteByLatest(username, upstream, ident, branch)
		if err != nil {
			return
		}
		return
	}
	var v int
	v, err = strconv.Atoi(version)
	if err != nil {
		return
	}
	vote, err = svc.db.GetVoteByVersion(username, upstream, ident, branch, uint32(v))
	return
}

func hasBlank(vars ...string) bool {
	for _, v := range vars {
		if v == "" {
			return true
		}
	}
	return false
}

func copyVote(to, from *apiv1.Vote) {
	if from.LawId != "" {
		to.LawId = from.LawId
	}
	if from.UserId != "" {
		to.UserId = from.UserId
	}
	if from.Vote != "" {
		to.Vote = from.Vote
	}
	if from.Comment != "" {
		to.Comment = from.Comment
	}
	if from.Upstream != "" {
		to.Upstream = from.Upstream
	}
	if from.Ident != "" {
		to.Ident = from.Ident
	}
	if from.Branch != "" {
		to.Branch = from.Branch
	}
	if from.Version != 0 {
		to.Version = from.Version
	}
	if from.Username != "" {
		to.Username = from.Username
	}
	return
}
