package repo

import (
	"context"

	"github.com/darwishdev/devkit-api/db"
	"github.com/darwishdev/devkit-api/pkg/contextkeys"
)

func (repo *TenantRepo) TenantCreateUpdate(ctx context.Context, req *db.TenantCreateUpdateParams) (*db.TenantsSchemaTenant, error) {
	resp, err := repo.store.TenantCreateUpdate(ctx, *req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}
	return &resp, nil
}

func (repo *TenantRepo) TenantList(ctx context.Context, tenantId int32) (*[]db.TenantsSchemaTenant, error) {
	tenateID, _ := contextkeys.TenantID(ctx)
	if tenateID > 0 {
		tenantId = tenateID
	}
	resp, err := repo.store.TenantList(ctx, tenantId)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}
	return &resp, nil
}

func (repo *TenantRepo) TenantDeleteRestore(ctx context.Context, req []int32) (*[]db.TenantsSchemaTenant, error) {
	resp, err := repo.store.TenantDeleteRestore(ctx, req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}
	return &resp, nil
}