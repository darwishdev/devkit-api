package api

import (
	"bytes"
	"context"
	"strings"

	"connectrpc.com/connect"
	"github.com/darwishdev/devkit-api/proto_gen/devkit/v1"
)

func (api *Api) FileDelete(ctx context.Context, req *connect.Request[devkitv1.FileDeleteRequest]) (*connect.Response[devkitv1.FileDeleteResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	response, err := api.publicUsecase.FileDelete(ctx, req.Msg)
	return connect.NewResponse(response), err
}

func (api *Api) FileList(ctx context.Context, req *connect.Request[devkitv1.FileListRequest]) (*connect.Response[devkitv1.FileListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	response, err := api.publicUsecase.FileList(ctx, req.Msg)
	return connect.NewResponse(response), err
}
func (api *Api) BucketCreateUpdate(ctx context.Context, req *connect.Request[devkitv1.BucketCreateUpdateRequest]) (*connect.Response[devkitv1.BucketCreateUpdateResponse], error) {
	_, err := api.checkForAccess(req.Header(), "bucket", "create")
	if err != nil {
		return nil, err
	}
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	response, err := api.publicUsecase.BucketCreateUpdate(ctx, req.Msg)
	return connect.NewResponse(response), err
}

func (api *Api) BucketList(ctx context.Context, req *connect.Request[devkitv1.BucketListRequest]) (*connect.Response[devkitv1.BucketListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	response, err := api.publicUsecase.BucketList(ctx, req.Msg)
	return connect.NewResponse(response), err
}

func (api *Api) FileCreate(ctx context.Context, req *connect.Request[devkitv1.FileCreateRequest]) (*connect.Response[devkitv1.FileCreateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.publicUsecase.FileCreate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) FileCreateBulk(ctx context.Context, req *connect.Request[devkitv1.FileCreateBulkRequest]) (*connect.Response[devkitv1.FileCreateBulkResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.publicUsecase.FileCreateBuilk(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) ImportTable(ctx context.Context, req *connect.Request[devkitv1.ImportTableRequest]) (*connect.Response[devkitv1.ImportTableResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	permissionName := strings.Replace(req.Msg.SchemaName, "_schema", "", 1)
	_, err := api.checkForAccess(req.Header(), permissionName, "create")
	if err != nil {
		_, err := api.checkForAccess(req.Header(), permissionName, "create_update")
		if err != nil {
			return nil, err
		}
	}
	buffer := bytes.NewBuffer(req.Msg.Reader)
	_, err = api.sqlSeeder.SeedFromExcel(*buffer, req.Msg.SchemaName, req.Msg.TableName, req.Msg.SchemaName)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&devkitv1.ImportTableResponse{
		Message: "imported",
	}), nil
}
