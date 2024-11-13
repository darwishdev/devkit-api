package api

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	"github.com/darwishdev/devkit-api/proto_gen/devkit/v1"
)

func (api *Api) RoleList(ctx context.Context, req *connect.Request[devkitv1.RoleListRequest]) (*connect.Response[devkitv1.RoleListResponse], error) {
	response, err := api.accountsUsecase.RoleList(ctx)
	if err != nil {
		return nil, err
	}
	response.Options = api.getAvailableOptions(req.Header())
	return connect.NewResponse(response), nil
}

func (api *Api) RoleCreateUpdate(ctx context.Context, req *connect.Request[devkitv1.RoleCreateUpdateRequest]) (*connect.Response[devkitv1.RoleCreateUpdateResponse], error) {
	if req.Msg.GetRoleId() == 0 && req.Msg.GetRoleName() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("role name is required if role ID not passed (create scenario)"))
	}
	response, err := api.accountsUsecase.RoleCreateUpdate(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(response), nil
}

func (api *Api) RoleDelete(ctx context.Context, req *connect.Request[devkitv1.RoleDeleteRequest]) (*connect.Response[devkitv1.RoleDeleteResponse], error) {
	resp, err := api.accountsUsecase.RoleDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&devkitv1.RoleDeleteResponse{Record: resp}), nil
}
func (api *Api) RoleDeleteRestore(ctx context.Context, req *connect.Request[devkitv1.RoleDeleteRestoreRequest]) (*connect.Response[devkitv1.RoleDeleteRestoreResponse], error) {
	resp, err := api.accountsUsecase.RoleDeleteRestore(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
func (api *Api) RoleFindForUpdate(ctx context.Context, req *connect.Request[devkitv1.RoleFindForUpdateRequest]) (*connect.Response[devkitv1.RoleFindForUpdateResponse], error) {
	response, err := api.accountsUsecase.RoleFindForUpdate(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve role list: %w", err)
	}
	return connect.NewResponse(response), nil
}
func (api *Api) RoleListInput(ctx context.Context, req *connect.Request[devkitv1.RoleListInputRequest]) (*connect.Response[devkitv1.RoleListInputResponse], error) {
	response, err := api.accountsUsecase.RoleListInput(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve roles list input: %w", err)
	}
	return connect.NewResponse(response), nil
}
