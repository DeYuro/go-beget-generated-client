package main

import (
	"context"
	"errors"
	"fmt"
	authClient "github.com/deyuro/go-beget-generated-client/internal/apiclient/auth"
	"github.com/k0kubun/pp"
)
const IncorrectCredentialsError = "INCORRECT_CREDENTIALS"
const EmptyLoginError = "EMPTY_LOGIN"

func authTest(ctx context.Context, p provider, login, password string) (string,error) {
	pp.Println(fmt.Sprintf("Auth wrong credential request: Expect %s error", IncorrectCredentialsError))
	r, err := p.auth(ctx, "wrong", "credential")
	if err != nil {
		return "", err
	}
	pp.Println("Response:")
	pp.Println(r)

	if r.GetError() != IncorrectCredentialsError {
		return  "",errors.New(fmt.Sprintf("Expect: %s: Got %s", IncorrectCredentialsError, r.GetError()))
	}

	pp.Println(fmt.Sprintf("Empty login auth request: Expect %s error", EmptyLoginError))
	r, err = p.auth(ctx, "", "emptylogin")
	if err != nil {
		return "",err
	}
	pp.Println("Response:")
	pp.Println(r)

	if r.GetError() != EmptyLoginError {
		return  "",errors.New(fmt.Sprintf("Expect: %s: Got %s", EmptyLoginError, r.GetError()))
	}
	pp.Println("Auth correct credential request: Expect r.Token != nil and r.Error == nil")

	r, err = p.auth(ctx, login, password)
	if err != nil {
		return "",err
	}
	pp.Println("Response:")
	pp.Println(r)

	if len(r.GetError()) > 0 || len(r.GetToken()) < 1  {
		return  "",errors.New(fmt.Sprintf("Expect: Expect r.Token != nil and r.Error == nil: Got r.Token == %s, r.Error == %s", r.GetToken(), r.GetError()))
	}

	return r.GetToken(), nil
}

func logoutTest(ctx context.Context, p provider) error {
	pp.Println("Logout request: Expect httpResp == 200, no response")

	req := p.authClient.AuthServiceApi.AuthServiceLogout(ctx)

	// func (r ApiAuthServiceLogoutRequest) Execute() (interface{}, *_nethttp.Response, error)  No response
	_, httpResp, err := req.Execute()
	pp.Println(fmt.Sprintf("httpResp == %d", httpResp.StatusCode))

	if httpResp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("http code is %d", httpResp.StatusCode))
	}

	return err
}

// auth request POST api.beget.com/v1/auth
func (p *provider) auth(ctx context.Context, login, password string) (*authClient.AuthAuthResponse, error) {
	req := p.authClient.AuthServiceApi.AuthServiceAuth(ctx)
	payload := authClient.NewAuthAuthRequest()

	payload.Login = &login
	payload.Password = &password

	req = req.AuthAuthRequest(*payload)

	resp, httpResp, err := req.Execute()

	if httpResp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("http code is %d", httpResp.StatusCode))
	}

	return &resp, err
}