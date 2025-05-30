package api

import (
	"context"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	devkitv1 "github.com/darwishdev/devkit-api/proto_gen/devkit/v1"
	"github.com/rs/zerolog/log"
)

func (api *Api) WithCookie(resp connect.AnyResponse, name string, value string, maxAge int) {
	resp.Header().Add("Set-Cookie", (&http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		Secure:   api.config.State == "PROD",
		SameSite: http.SameSiteLaxMode,
		MaxAge:   maxAge,
	}).String())
}

func (api *Api) AuthRegister(ctx context.Context, req *connect.Request[devkitv1.AuthRegisterRequest]) (*connect.Response[devkitv1.AuthRegisterResponse], error) {
	response, err := api.accountsUsecase.AuthRegister(ctx, req)
	if err != nil {
		return nil, err
	}

	resp := connect.NewResponse(response)
	maxAge := int(api.config.RefreshTokenDuration.Seconds())
	api.WithCookie(resp, api.config.RefreshTokenCookieName, response.LoginInfo.RefreshToken, maxAge)
	api.WithCookie(resp, api.config.SupabaseTokenCookieName, response.LoginInfo.SupabaseToken, maxAge)
	api.WithCookie(resp, api.config.SupabaseRefreshTokenCookieName, response.LoginInfo.SupabaseRefreshToken, maxAge)
	return resp, nil
}

func (api *Api) AuthLoginProviderCallback(ctx context.Context, req *connect.Request[devkitv1.AuthLoginProviderCallbackRequest]) (*connect.Response[devkitv1.AuthLoginProviderCallbackResponse], error) {
	response, err := api.accountsUsecase.AuthLoginProviderCallback(ctx, req)
	if err != nil {
		return nil, err
	}
	resp := connect.NewResponse(response)
	maxAge := int(api.config.RefreshTokenDuration.Seconds())
	api.WithCookie(resp, api.config.RefreshTokenCookieName, response.LoginInfo.RefreshToken, maxAge)
	api.WithCookie(resp, api.config.SupabaseTokenCookieName, response.LoginInfo.SupabaseToken, maxAge)
	api.WithCookie(resp, api.config.SupabaseRefreshTokenCookieName, response.LoginInfo.SupabaseRefreshToken, maxAge)
	return resp, nil
}

func (api *Api) AuthLoginProvider(ctx context.Context, req *connect.Request[devkitv1.AuthLoginProviderRequest]) (*connect.Response[devkitv1.AuthLoginProviderResponse], error) {
	response, err := api.accountsUsecase.AuthLoginProvider(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(response), nil
}

func (api *Api) AuthResetPasswordEmail(ctx context.Context, req *connect.Request[devkitv1.AuthResetPasswordEmailRequest]) (*connect.Response[devkitv1.AuthResetPasswordEmailResponse], error) {
	response, err := api.accountsUsecase.AuthResetPasswordEmail(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(response), nil
}
func (api *Api) AuthResetPassword(ctx context.Context, req *connect.Request[devkitv1.AuthResetPasswordRequest]) (*connect.Response[devkitv1.AuthResetPasswordResponse], error) {
	err := api.validator.Validate(req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("validation error: %w", err))
	}
	if req.Msg.NewPassword != req.Msg.NewPasswordConfirmation {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("password and confirmation do not match"))
	}
	payload, err := api.authorizeRequestHeader(req.Header())
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, fmt.Errorf("invalid access token: %w", err))
	}
	req.Msg.Email = payload.Username
	response, err := api.accountsUsecase.AuthResetPassword(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to reset password: %w", err)
	}
	return connect.NewResponse(response), nil
}
func (api *Api) AuthLogin(ctx context.Context, req *connect.Request[devkitv1.AuthLoginRequest]) (*connect.Response[devkitv1.AuthLoginResponse], error) {
	response, err := api.accountsUsecase.AuthLogin(ctx, req)
	if err != nil {
		log.Debug().Interface("err is", err).Msg("error is here")
		return nil, err
	}

	maxAge := int(api.config.RefreshTokenDuration.Seconds())
	resp := connect.NewResponse(response)
	api.WithCookie(resp, api.config.RefreshTokenCookieName, response.LoginInfo.RefreshToken, maxAge)
	api.WithCookie(resp, api.config.SupabaseTokenCookieName, response.LoginInfo.SupabaseToken, maxAge)
	api.WithCookie(resp, api.config.SupabaseRefreshTokenCookieName, response.LoginInfo.SupabaseRefreshToken, maxAge)

	return resp, nil
}

func (api *Api) AuthRefreshToken(
	ctx context.Context,
	req *connect.Request[devkitv1.AuthRefreshTokenRequest],
) (*connect.Response[devkitv1.AuthRefreshTokenResponse], error) {
	response, err := api.accountsUsecase.AuthRefreshToken(ctx, req)
	if err != nil {
		return nil, err
	}
	resp := connect.NewResponse(response)
	maxAge := int(api.config.RefreshTokenDuration.Seconds())
	api.WithCookie(resp, api.config.RefreshTokenCookieName, response.LoginInfo.RefreshToken, maxAge)
	return resp, nil
}
func (api *Api) AuthLogout(
	ctx context.Context,
	req *connect.Request[devkitv1.AuthLogoutRequest],
) (*connect.Response[devkitv1.AuthLogoutResponse], error) {
	response, err := api.accountsUsecase.AuthLogout(ctx, req)
	if err != nil {
		return nil, err
	}
	resp := connect.NewResponse(response)
	api.WithCookie(resp, api.config.RefreshTokenCookieName, "", -1)
	api.WithCookie(resp, api.config.SupabaseTokenCookieName, "", -1)
	api.WithCookie(resp, api.config.SupabaseRefreshTokenCookieName, "", -1)
	return resp, nil
}

func (api *Api) AuthSessionList(
	ctx context.Context,
	req *connect.Request[devkitv1.AuthSessionListRequest],
) (*connect.Response[devkitv1.AuthSessionListResponse], error) {
	response, err := api.accountsUsecase.AuthSessionList(ctx, req)
	if err != nil {
		return nil, err
	}
	resp := connect.NewResponse(response)
	resp.Msg.Options = api.getAvailableOptions(req.Header())
	return resp, nil
}

func (api *Api) AuthSessionSetBlocked(
	ctx context.Context,
	req *connect.Request[devkitv1.AuthSessionSetBlockedRequest],
) (*connect.Response[devkitv1.AuthSessionSetBlockedResponse], error) {
	response, err := api.accountsUsecase.AuthSessionSetBlocked(ctx, req)
	if err != nil {
		return nil, err
	}
	resp := connect.NewResponse(response)
	return resp, nil
}
func (api *Api) AuthSessionDelete(
	ctx context.Context,
	req *connect.Request[devkitv1.AuthSessionDeleteRequest],
) (*connect.Response[devkitv1.AuthSessionDeleteResponse], error) {
	response, err := api.accountsUsecase.AuthSessionDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	resp := connect.NewResponse(response)
	return resp, nil
}
