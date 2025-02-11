// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: accounts_role.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const roleCreateUpdate = `-- name: RoleCreateUpdate :one
SELECT
	role_id, tenant_id, role_name, role_security_level, role_description, created_at, updated_at, deleted_at
FROM
	accounts_schema.role_create_update (in_role_id => $1, in_role_name => $2, in_tenant_id => $3, in_role_security_level => $4, in_caller_id => $5, in_role_description => $6, in_permissions => $7::int[])
`

type RoleCreateUpdateParams struct {
	RoleID            int32   `json:"role_id"`
	RoleName          string  `json:"role_name"`
	TenantID          int32   `json:"tenant_id"`
	RoleSecurityLevel int32   `json:"role_security_level"`
	CallerID          int32   `json:"caller_id"`
	RoleDescription   string  `json:"role_description"`
	Permissions       []int32 `json:"permissions"`
}

func (q *Queries) RoleCreateUpdate(ctx context.Context, arg RoleCreateUpdateParams) (AccountsSchemaRole, error) {
	row := q.db.QueryRow(ctx, roleCreateUpdate,
		arg.RoleID,
		arg.RoleName,
		arg.TenantID,
		arg.RoleSecurityLevel,
		arg.CallerID,
		arg.RoleDescription,
		arg.Permissions,
	)
	var i AccountsSchemaRole
	err := row.Scan(
		&i.RoleID,
		&i.TenantID,
		&i.RoleName,
		&i.RoleSecurityLevel,
		&i.RoleDescription,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const roleDelete = `-- name: RoleDelete :one
SELECT
	role_id, tenant_id, role_name, role_security_level, role_description, created_at, updated_at, deleted_at
FROM
	accounts_schema.role_delete (in_role_id => $1, in_caller_id => $2)
`

type RoleDeleteParams struct {
	RoleID   int32 `json:"role_id"`
	CallerID int32 `json:"caller_id"`
}

func (q *Queries) RoleDelete(ctx context.Context, arg RoleDeleteParams) (AccountsSchemaRole, error) {
	row := q.db.QueryRow(ctx, roleDelete, arg.RoleID, arg.CallerID)
	var i AccountsSchemaRole
	err := row.Scan(
		&i.RoleID,
		&i.TenantID,
		&i.RoleName,
		&i.RoleSecurityLevel,
		&i.RoleDescription,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const roleDeleteRestore = `-- name: RoleDeleteRestore :one
SELECT
	role_id, tenant_id, role_name, role_security_level, role_description, created_at, updated_at, deleted_at
FROM
	accounts_schema.role_delete_restore (in_role_id => $1, in_caller_id => $2)
`

type RoleDeleteRestoreParams struct {
	RoleID   int32 `json:"role_id"`
	CallerID int32 `json:"caller_id"`
}

func (q *Queries) RoleDeleteRestore(ctx context.Context, arg RoleDeleteRestoreParams) (AccountsSchemaRole, error) {
	row := q.db.QueryRow(ctx, roleDeleteRestore, arg.RoleID, arg.CallerID)
	var i AccountsSchemaRole
	err := row.Scan(
		&i.RoleID,
		&i.TenantID,
		&i.RoleName,
		&i.RoleSecurityLevel,
		&i.RoleDescription,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const roleFindForUpdate = `-- name: RoleFindForUpdate :one
WITH permissions AS (
	SELECT
		p.role_id,
		array_agg(p.permission_id)::int[] permissions
	FROM
		accounts_schema.role_permission p
	WHERE
		p.role_id = $1
	GROUP BY
		p.role_id
)
SELECT
	r.role_id,
	r.role_name,
	r.tenant_id,
	r.role_security_level,
	r.role_description,
	p.permissions permissions
FROM
	accounts_schema.role r
	JOIN permissions p ON r.role_id = p.role_id
`

type RoleFindForUpdateRow struct {
	RoleID            int32       `json:"role_id"`
	RoleName          string      `json:"role_name"`
	TenantID          pgtype.Int4 `json:"tenant_id"`
	RoleSecurityLevel int32       `json:"role_security_level"`
	RoleDescription   pgtype.Text `json:"role_description"`
	Permissions       []int32     `json:"permissions"`
}

func (q *Queries) RoleFindForUpdate(ctx context.Context, roleID int32) (RoleFindForUpdateRow, error) {
	row := q.db.QueryRow(ctx, roleFindForUpdate, roleID)
	var i RoleFindForUpdateRow
	err := row.Scan(
		&i.RoleID,
		&i.RoleName,
		&i.TenantID,
		&i.RoleSecurityLevel,
		&i.RoleDescription,
		&i.Permissions,
	)
	return i, err
}

const roleList = `-- name: RoleList :many
SELECT
	role_id::int,
	role_name::varchar(200),
	role_description::varchar(200),
	created_at::timestamptz,
	total_count::bigint
FROM
	execute_dynamic_pagination (primary_key => 'role_id', query_base => CONCAT(FORMAT('SELECT role_id, role_name,role_description, created_at
             FROM accounts_schema.role
             WHERE role_name LIKE CONCAT(''%%'', %L, ''%%'')
               AND role_description LIKE CONCAT(''%%'', %L, ''%%'')
AND deleted_at IS ', $1::varchar(200), $2::text),  iif($3::boolean , 'not null'::varchar , 'null'::varchar)), sort_func => $4, page_number => $5, -- Page number
		sort_column => is_null_replace($6::varchar, 'role_id'), page_size => $7) AS result (role_id int,
		role_name varchar(200),
		role_description varchar(200),
		created_at timestamptz,
		total_count bigint)
`

type RoleListParams struct {
	InRoleName        string `json:"in_role_name"`
	InRoleDescription string `json:"in_role_description"`
	InIsDeleted       bool   `json:"in_is_deleted"`
	SortFunction      string `json:"sort_function"`
	PageNumber        int32  `json:"page_number"`
	SortColumn        string `json:"sort_column"`
	PageSize          int32  `json:"page_size"`
}

type RoleListRow struct {
	RoleID          int32              `json:"role_id"`
	RoleName        string             `json:"role_name"`
	RoleDescription string             `json:"role_description"`
	CreatedAt       pgtype.Timestamptz `json:"created_at"`
	TotalCount      int64              `json:"total_count"`
}

func (q *Queries) RoleList(ctx context.Context, arg RoleListParams) ([]RoleListRow, error) {
	rows, err := q.db.Query(ctx, roleList,
		arg.InRoleName,
		arg.InRoleDescription,
		arg.InIsDeleted,
		arg.SortFunction,
		arg.PageNumber,
		arg.SortColumn,
		arg.PageSize,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []RoleListRow{}
	for rows.Next() {
		var i RoleListRow
		if err := rows.Scan(
			&i.RoleID,
			&i.RoleName,
			&i.RoleDescription,
			&i.CreatedAt,
			&i.TotalCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const roleListInput = `-- name: RoleListInput :many
SELECT
	role_id value,
	role_name label,
	concat("level: ", role_security_level::varchar)::varchar note
FROM
	accounts_schema.role
WHERE
	role_security_level <= accounts_schema.user_security_level_find ($1)
`

type RoleListInputRow struct {
	Value int32  `json:"value"`
	Label string `json:"label"`
	Note  string `json:"note"`
}

func (q *Queries) RoleListInput(ctx context.Context, callerID int32) ([]RoleListInputRow, error) {
	rows, err := q.db.Query(ctx, roleListInput, callerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []RoleListInputRow{}
	for rows.Next() {
		var i RoleListInputRow
		if err := rows.Scan(&i.Value, &i.Label, &i.Note); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
