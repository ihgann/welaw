package database

import (
	"database/sql"

	apiv1 "github.com/welaw/welaw/api/v1"
	"github.com/welaw/welaw/pkg/errs"
)

func (db *_database) AuthorizeUser(username, password, operation string) (bool, error) {
	q := `
	SELECT 1
	FROM users
	INNER JOIN user_roles ON user_roles.user_id = users.uid
	INNER JOIN roles ON roles.uid = user_roles.role_id
	INNER JOIN role_operations ON role_operations.role_id = roles.uid
	INNER JOIN operations ON operations.uid = role_operations.operation_id
	WHERE users.password = crypt($2, users.password)
		AND users.username = $1
		AND operations.name = $3
		AND users.deleted_at IS NULL
		AND user_roles.deleted_at IS NULL
	`
	var r bool
	err := db.conn.QueryRow(q, username, password, operation).Scan(&r)
	if err == sql.ErrNoRows {
		return false, errs.ErrNotFound
	}
	return r, nil
}

func (db *_database) SetUserPassword(uid, password string) (err error) {
	q := `
	UPDATE users
	SET password = crypt($2, gen_salt('bf'))
	WHERE uid = $1
		AND deleted_at IS NULL
	`
	_, err = db.conn.Exec(q, uid, password)
	return
}

func (db *_database) CreateUserRoles(username string, role []string) error {
	q := `
	INSERT INTO user_roles (
		user_id,
		role_id
	)
	VALUES (
		(SELECT uid FROM users WHERE username = $1 AND deleted_at IS NULL),
		(SELECT uid FROM roles WHERE name = $2 AND deleted_at IS NULL)
	)
	`
	_, err := db.conn.Exec(
		q,
		username,
		role[0],
	)
	return err
}

func (db *_database) DeleteUserRoles(username string, role []string) error {
	q := `
	INSERT INTO user_roles (
		user_id,
		role_id
	)
	VALUES (
		(SELECT uid FROM users WHERE users.username = $1),
		(SELECT uid FROM roles WHERE roles.name = $2)
	)
	`
	_, err := db.conn.Exec(
		q,
		username,
		role[0],
	)
	return err
}

func (db *_database) ListUserRoles(username string) ([]*apiv1.UserRole, error) {
	q := `
	SELECT roles.name
	FROM roles
	INNER JOIN user_roles ON user_roles.role_id = roles.uid
	INNER JOIN users ON users.uid = user_roles.user_id
	WHERE users.username = $1
		AND users.deleted_at IS NULL
		AND user_roles.deleted_at IS NULL
	`
	rows, err := db.conn.Query(q, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var roles []*apiv1.UserRole
	for rows.Next() {
		r := new(apiv1.UserRole)
		err = rows.Scan(
			&r.Name,
		)
		if err != nil {
			return nil, err
		}
		roles = append(roles, r)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (db *_database) HasRole(user_id, role string) (bool, error) {
	q := `
	SELECT 1
	FROM roles
	INNER JOIN user_roles ON user_roles.role_id = roles.uid AND user_roles.deleted_at IS NULL
	INNER JOIN users ON users.uid = user_roles.user_id AND users.deleted_at IS NULL
	WHERE users.uid = $1
		AND roles.name = $2
		AND roles.deleted_at IS NULL
	`
	rows, err := db.conn.Query(q, user_id, role)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	err = rows.Err()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (db *_database) HasPermission(user_id, operation string) (bool, error) {
	db.logger.Log("method", "has_permission", "user_id", user_id, "operation", operation)
	q := `
	SELECT 1
	FROM roles
	INNER JOIN user_roles ON user_roles.role_id = roles.uid AND user_roles.deleted_at IS NULL
	INNER JOIN users ON users.uid = user_roles.user_id AND users.deleted_at IS NULL
	INNER JOIN role_operations ON role_operations.role_id = roles.uid AND role_operations.deleted_at IS NULL
	INNER JOIN operations ON operations.uid = role_operations.operation_id AND operations.deleted_at IS NULL
	WHERE users.uid = $1
		AND operations.name = $2
		AND roles.deleted_at IS NULL
	`
	rows, err := db.conn.Query(q, user_id, operation)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	err = rows.Err()
	if err != nil {
		return false, err
	}
	// TODO
	for rows.Next() {
		return true, nil
	}
	return false, nil
}
