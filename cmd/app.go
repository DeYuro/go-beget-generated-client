package main

import (
	"context"
	"flag"
	authClient "github.com/deyuro/go-beget-generated-client/internal/apiclient/auth"
	vpsClient "github.com/deyuro/go-beget-generated-client/internal/apiclient/vps"
	"github.com/k0kubun/pp"
	"time"
)



type provider struct {
	authClient *authClient.APIClient
	vpsClient *vpsClient.APIClient
}

func app() error  {
	var login, password string
	var delay time.Duration

	flag.StringVar(&login, "login", "", "login")
	flag.StringVar(&password, "password", "", "password")
	flag.DurationVar(&delay, "delay", time.Second * 10, "delay")
	flag.Parse()

	ctx := context.Background()
	authCfg := authClient.NewConfiguration()
	provider := provider{
		authClient: authClient.NewAPIClient(authCfg),
	}

	token, err := authTest(ctx, provider, login, password)
	if err != nil {
		return err
	}
	authCfg.DefaultHeader = map[string]string{
				"Authorization": "Bearer " + token,
	}

	vpsCfg := vpsClient.NewConfiguration()
	vpsCfg.DefaultHeader = map[string]string{
		"Authorization": "Bearer " + token,
	}
	provider.vpsClient = vpsClient.NewAPIClient(vpsCfg)


	err = vpsTest(ctx, provider, delay)
	if err != nil {
		pp.Println("!!!!!!!!!!TESTS NOT PASSED!!!!!!!")
		return err
	}


	// unable to upload client config, need replace
	provider.authClient = authClient.NewAPIClient(authCfg)

	err = logoutTest(ctx,provider)
	if err != nil {
		return err
	}
	pp.Println("!!!!!!!!!!ALL TESTS PASSED!!!!!!!")

	return err
}