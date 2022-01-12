package main

import (
	"context"
	"flag"
	authClient "github.com/deyuro/go-beget-generated-client/internal/apiclient/auth"
	vpsClient "github.com/deyuro/go-beget-generated-client/internal/apiclient/vps"
)



type provider struct {
	authClient *authClient.APIClient
	vpsClient *vpsClient.APIClient
}

func app() error  {
	var login, password string

	flag.StringVar(&login, "login", "", "login")
	flag.StringVar(&password, "password", "", "password")
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

	// unable to upload client config, need replace
	provider.authClient = authClient.NewAPIClient(authCfg)

	err = logoutTest(ctx,provider)
	if err != nil {
		return err
	}

	return err
}





//func createVps(ctx context.Context, client *openapi.APIClient, vpsCfg config.Vps) (*openapi.ManageVpsInfo, error) {
//
//	req := client.ManageServiceApi.ManageServiceCreateVps(ctx)
//	payload := openapi.NewManageCreateVpsRequest()
//
//	payload.BegetSshAccessAllowed = &vpsCfg.BegetSshAccessAllowed
//	payload.ConfigurationId = &vpsCfg.ConfigurationId
//	payload.Description = &vpsCfg.Description
//	payload.DisplayName = &vpsCfg.DisplayName
//	payload.Hostname = &vpsCfg.Hostname
//	payload.OperatingSystemId = &vpsCfg.OperationSystemId
//	payload.Password = &vpsCfg.Password
//
//	req = req.ManageCreateVpsRequest(*payload)
//
//	resp, httpResp, err := req.Execute()
//
//	_ = err
//	if httpResp.StatusCode != 200 {
//		return nil,errors.New(fmt.Sprintf("http code is %d", httpResp.StatusCode))
//	}
//	vpsInfo := resp.GetVps()
//
//	return &vpsInfo, nil
//}


//func app() error {
//
//
//
//	ctx := context.Background()
//	apiCfg := openapi.NewConfiguration()
//	apiCfg.DefaultHeader = map[string]string{
//		"Authorization": "Bearer " + cfg.Auth.Jwt,
//	}
//
//	client := openapi.NewAPIClient(apiCfg)
//
//	list, err := getVpsList(ctx, client)
//	if err != nil {
//		return err
//	}
//
//	for _, v := range list {
//		fmt.Printf("VpsID %s  has slug %s\n", *v.Id, *v.Slug)
//	}
//
//
//	vpsInfo, err := createVps(ctx, client, cfg.Vps)
//	if err != nil {
//		return err
//	}
//	fmt.Printf("VpsID %s with slug %s has been created\n", *vpsInfo.Id, *vpsInfo.Slug)
//
//	list, err = getVpsList(ctx, client)
//	if err != nil {
//		return err
//	}
//	for _, v := range list {
//		if *vpsInfo.Id == *v.Id {
//			fmt.Printf("VpsID %s  has returned from getVpsList request", *v.Id)
//		}
//	}
//
//	for i:=0;i<5;i++ {
//		time.Sleep(1 * time.Second)
//		switch i {
//		case 0:
//			fmt.Println("Vps is began to work")
//		case 1:
//			fmt.Println("Doing some cool stuffs")
//		case 2:
//			fmt.Println("Makes millions")
//		case 3:
//			fmt.Println("Almost finish")
//		case 4:
//			fmt.Println("And....done!")
//		}
//	}
//	removedVps, err := removeVps(ctx, client, *vpsInfo)
//	if err != nil {
//		return err
//	}
//
//	fmt.Printf("VpsID %s has end it life cylcle and deleted", *removedVps.Id)
//
//	list, err = getVpsList(ctx, client)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func getVpsList(ctx context.Context, client *openapi.APIClient) ([]openapi.ManageVpsInfo, error) {
//
//	req := client.ManageServiceApi.ManageServiceGetList(ctx)
//
//	resp, httpResp, _ := req.Execute()
//
//	if httpResp.StatusCode != 200 {
//		return nil, errors.New(fmt.Sprintf("http code is %d", httpResp.StatusCode))
//	}
//
//	return resp.GetVps(), nil
//}
//
//func createVps(ctx context.Context, client *openapi.APIClient, vpsCfg config.Vps) (*openapi.ManageVpsInfo, error) {
//
//	req := client.ManageServiceApi.ManageServiceCreateVps(ctx)
//	payload := openapi.NewManageCreateVpsRequest()
//
//	payload.BegetSshAccessAllowed = &vpsCfg.BegetSshAccessAllowed
//	payload.ConfigurationId = &vpsCfg.ConfigurationId
//	payload.Description = &vpsCfg.Description
//	payload.DisplayName = &vpsCfg.DisplayName
//	payload.Hostname = &vpsCfg.Hostname
//	payload.OperatingSystemId = &vpsCfg.OperationSystemId
//	payload.Password = &vpsCfg.Password
//
//	req = req.ManageCreateVpsRequest(*payload)
//
//	resp, httpResp, err := req.Execute()
//
//	_ = err
//	if httpResp.StatusCode != 200 {
//		return nil,errors.New(fmt.Sprintf("http code is %d", httpResp.StatusCode))
//	}
//	vpsInfo := resp.GetVps()
//
//	return &vpsInfo, nil
//}
//
//func removeVps(ctx context.Context, client *openapi.APIClient, vpsInfo openapi.ManageVpsInfo) (*openapi.ManageVpsInfo, error)  {
//	req := client.ManageServiceApi.ManageServiceRemoveVps(ctx, *vpsInfo.Id)
//	resp, httpResp, err := req.Execute()
//	_ = err
//
//	if httpResp.StatusCode != 200 {
//		return nil,errors.New(fmt.Sprintf("http code is %d", httpResp.StatusCode))
//	}
//	r := resp.GetVps()
//	return &r, nil
//}