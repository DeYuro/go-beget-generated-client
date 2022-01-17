package main

import (
	"context"
	"errors"
	"fmt"
	vpsClient "github.com/deyuro/go-beget-generated-client/internal/apiclient/vps"
	"github.com/k0kubun/pp"
	"time"
)

const (
	error404             = "404 Not Found"
	invalidHostnameError = "INVALID_HOSTNAME"
	removingStatus		 = "REMOVING"
	creatingStatus		 = "CREATING"
	runningStatus    = "RUNNING"
	restartingStatus = "RESTARTING"
	reinstallingStatus = "REINSTALLING"
)

func vpsTest(ctx context.Context, p provider, delay time.Duration) error {
	pp.Println("GetVpsList request: Expect customer vps list")
	rGetList, err := p.getVpsList(ctx)

	if err != nil {
		return err
	}

	pp.Println("Response:")
	pp.Println(rGetList)

	pp.Println(fmt.Sprintf("GetVpsList request with wrong id: Expect err with text %s, Vps == nil", error404))
	rGetInfo, err := p.getInfo(ctx, "qwerty")
	pp.Println("Response")
	pp.Println(rGetInfo)
	pp.Println("Error",err)

	if err == nil {
		return errors.New(fmt.Sprintf("Expect err %s got nil",error404))
	}

	if err.Error() != error404 {
		return errors.New(fmt.Sprintf("Expect err %s got %s",error404, err.Error()))
	}

	if rGetInfo.Vps != nil {
		return errors.New(fmt.Sprintf("Expect Vps == nil, got != nil"))
	}

	pp.Println("GetVpsInfo request correct id: Expect VpsInfo != nil")
	rGetInfo, err = p.getInfo(ctx, rGetList.GetVps()[0].GetId())
	if err != nil {
		return err
	}
	pp.Println("Response")
	pp.Println(rGetInfo)

	if rGetInfo.Vps == nil {
		return errors.New(fmt.Sprintf("Expect Vps != nil, got == nil"))
	}

	pp.Println("GetAvailableConfiguration request: Expect VpsConfiguration list and OperatingSystem list")

	rGetAvailableConfiguration, err := p.getAvailableConfiguration(ctx)
	if err != nil {
		return err
	}

	pp.Println("Response")
	pp.Println(rGetAvailableConfiguration)

	if rGetAvailableConfiguration.GetConfigurations() == nil {
		return errors.New(fmt.Sprintf("Expect VpsConfiguration != nil, got == nil"))
	}

	if rGetAvailableConfiguration.GetOperatingSystems() == nil {
		return errors.New(fmt.Sprintf("Expect OperatingSystem != nil, got == nil"))
	}

	pp.Println(fmt.Sprintf("CreateVps request with invalid hostname: Expect: %s error", invalidHostnameError))

	rCreateVps, err := p.createVps(ctx,
	"Invalid:Host?Name",
		rGetAvailableConfiguration.GetConfigurations()[10].GetId(),
		rGetAvailableConfiguration.GetOperatingSystems()[0].GetId())
	if err != nil {
		return err
	}
	pp.Println("Response")
	pp.Println(rCreateVps)

	if rCreateVps.Error.GetCode() != invalidHostnameError {
		return  errors.New(fmt.Sprintf("Expect: %s: Got %s", invalidHostnameError, rCreateVps.Error.GetCode()))
	}

	pp.Println(fmt.Sprintf("CreateVps request with valid data: Expect: VPS status == %s", creatingStatus))

	rCreateVps, err = p.createVps(ctx,
		"validhostname",
		rGetAvailableConfiguration.GetConfigurations()[10].GetId(),
		rGetAvailableConfiguration.GetOperatingSystems()[0].GetId())
	if err != nil {
		return err
	}
	pp.Println("Response")
	pp.Println(rCreateVps)
	if rCreateVps.Vps.GetStatus() != creatingStatus {
		return errors.New(fmt.Sprintf("Expect vps status == %s, got == %s", creatingStatus, rCreateVps.Vps.GetStatus()))
	}
	// Wait for creating
	time.Sleep(delay)

	pp.Println(fmt.Sprintf("Check VpsInfo on created vps by uuid: Expect Vps with %s exist and status is %s", rCreateVps.Vps.GetId(), runningStatus))

	rGetInfo, err = p.getInfo(ctx, rCreateVps.Vps.GetId())
	if err != nil {
		return err
	}
	pp.Println("Response")
	pp.Println(rGetInfo)

	if rGetInfo.Vps.GetStatus() != runningStatus {
		return errors.New(fmt.Sprintf("Expect vps status == %s, got == %s", runningStatus, rGetInfo.Vps.GetStatus()))
	}

	pp.Println(fmt.Sprintf("ResetPassword request:Expect vps != nil"))

	rResetPassword, err := p.resetPassword(ctx, rCreateVps.Vps.GetId())
	if err != nil {
		return err
	}
	pp.Println("Response")
	pp.Println(rResetPassword)

	if rResetPassword.Vps == nil {
		return errors.New(fmt.Sprintf("Expect vps != nil, got nil"))
	}

	hostName := "newhostname"
	pp.Println(fmt.Sprintf("UpdateInfo request with new hostname: Expect: VPS hostname == %s", hostName))

	rUpdateInfo, err := p.updateInfo(ctx, rCreateVps.Vps.GetId(), hostName)
	if err != nil {
		return err
	}
	pp.Println("Response")
	pp.Println(rUpdateInfo)

	if rUpdateInfo.Vps.GetHostname() != hostName {
		return errors.New(fmt.Sprintf("Expect vps.hostname == %s, got %s", hostName, rUpdateInfo.Vps.GetHostname()))
	}

	pp.Println(fmt.Sprintf("RebootVps request: Expect vps status == %s", restartingStatus))
	rRebootVps, err := p.rebootVps(ctx, rCreateVps.Vps.GetId())
	pp.Println("Response")
	pp.Println(rRebootVps)

	if err != nil {
		return err
	}

	if rRebootVps.Vps.GetStatus() != restartingStatus {
		return errors.New(fmt.Sprintf("Expect vps status == %s, got == %s", restartingStatus, rRebootVps.Vps.GetStatus()))
	}
	//wait rebooting
	time.Sleep(delay)

	pp.Println(fmt.Sprintf("Check VpsInfo on rebooting vps by uuid: Expect Vps with %s exist and status is %s", rCreateVps.Vps.GetId(), runningStatus))

	rGetInfo, err = p.getInfo(ctx, rCreateVps.Vps.GetId())
	if err != nil {
		return err
	}
	pp.Println("Response")
	pp.Println(rGetInfo)

	if rGetInfo.Vps.GetStatus() != runningStatus {
		return errors.New(fmt.Sprintf("Expect vps status == %s, got == %s", runningStatus, rGetInfo.Vps.GetStatus()))
	}

	pp.Println(fmt.Sprintf("Reinstall request: Expect vps status == %s", reinstallingStatus))
	rReinstall, err := p.reinstall(ctx, rCreateVps.Vps.GetId(),rGetAvailableConfiguration.GetOperatingSystems()[1].GetId())
	pp.Println("Response")
	pp.Println(rReinstall)

	if err != nil {
		return err
	}

	if rReinstall.Vps.GetStatus() != reinstallingStatus {
		return errors.New(fmt.Sprintf("Expect vps status == %s, got == %s", reinstallingStatus, rRebootVps.Vps.GetStatus()))
	}
	//wait reinstall
	time.Sleep(delay)

	pp.Println(fmt.Sprintf("Check VpsInfo on reinstall vps by uuid: Expect Vps with %s os id == %d",
		rCreateVps.Vps.GetId(), rGetAvailableConfiguration.GetOperatingSystems()[1].GetId()))

	rGetInfo, err = p.getInfo(ctx, rCreateVps.Vps.GetId())
	if err != nil {
		return err
	}
	pp.Println("Response")
	pp.Println(rGetInfo)

	if rGetInfo.Vps.OperatingSystem.GetId()!= rGetAvailableConfiguration.GetOperatingSystems()[1].GetId() {
		return errors.New(fmt.Sprintf("Expect Vps os id == %d: Got %d",rGetAvailableConfiguration.GetOperatingSystems()[1].GetId(),rGetInfo.Vps.OperatingSystem.GetId()))
	}

	err = sshTest(ctx,p,rCreateVps.Vps.GetId())
	if err != nil {
		return err
	}

	pp.Println(fmt.Sprintf("RemoveVps request: Expect vps status == %s", removingStatus))
	rRemoveVps, err := p.removeVps(ctx, rCreateVps.Vps.GetId())
	pp.Println("Response")
	pp.Println(rRemoveVps)

	if err != nil {
		return err
	}

	if rRemoveVps.Vps.GetStatus() != removingStatus {
		return errors.New(fmt.Sprintf("Expect vps status == %s, got == %s", removingStatus, rRemoveVps.Vps.GetStatus()))
	}

	pp.Println(fmt.Sprintf("GetVpsInfo request with removed uuid %s: Expect err with text %s", rCreateVps.Vps.GetId(), error404))
	rGetInfo, err = p.getInfo(ctx, rCreateVps.Vps.GetId())
	pp.Println("Response")
	pp.Println(rGetInfo)
	pp.Println("Error",err)

	if err == nil {
		return errors.New(fmt.Sprintf("Expect err %s got nil",error404))
	}

	if err.Error() != error404 {
		return errors.New(fmt.Sprintf("Expect err %s got %s",error404, err.Error()))
	}



	return nil
}

// GET api.beget.com/v1/vps/server/list
func (p *provider) getVpsList(ctx context.Context) (*vpsClient.ManageGetListResponse, error) {
	req := p.vpsClient.ManageServiceApi.ManageServiceGetList(ctx)

	resp, httpResp, err := req.Execute()

	if httpResp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("http code is %d", httpResp.StatusCode))
	}

	return &resp, err
}

// GET api.beget.com/v1/vps/server/{id}
func (p *provider) getInfo(ctx context.Context, uuid string) (*vpsClient.ManageGetInfoResponse, error) {
	req := p.vpsClient.ManageServiceApi.ManageServiceGetInfo(ctx, uuid)

	resp, _, err := req.Execute()

	return &resp, err
}
// GET api.beget.com/v1/vps/server/v1/vps/configuration
func (p *provider) getAvailableConfiguration(ctx context.Context) (*vpsClient.ManageGetAvailableConfigurationResponse, error) {
	req := p.vpsClient.ManageServiceApi.ManageServiceGetAvailableConfiguration(ctx)

	resp, httpResp, err := req.Execute()

	if httpResp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("http code is %d", httpResp.StatusCode))
	}

	return &resp, err
}

// POST api.beget.com/v1/vps/server
func (p *provider) createVps(ctx context.Context, hostname, configurationId string, operationSystemId int32) (*vpsClient.ManageCreateVpsResponse, error) {

	var displayName = "Go Generated Client Test"
	var begetSshAccessAllowed =  true
	var description = "Test client from generated code"
	var password = "L%v2sZ2j"

	req := p.vpsClient.ManageServiceApi.ManageServiceCreateVps(ctx)
	payload := vpsClient.NewManageCreateVpsRequest()

	payload.BegetSshAccessAllowed = &begetSshAccessAllowed
	payload.ConfigurationId = &configurationId
	payload.Description = &description
	payload.DisplayName = &displayName
	payload.Hostname = &hostname
	payload.OperatingSystemId = &operationSystemId
	payload.Password = &password

	req = req.ManageCreateVpsRequest(*payload)

	resp, httpResp, err := req.Execute()

	if httpResp.StatusCode != 200 {
		return nil,errors.New(fmt.Sprintf("http code is %d", httpResp.StatusCode))
	}

	return &resp, err
}

// POST api.beget.com/v1/vps/server/{id}/remove
func (p *provider) removeVps(ctx context.Context, uuid string) (*vpsClient.ManageRemoveVpsResponse, error) {
	req := p.vpsClient.ManageServiceApi.ManageServiceRemoveVps(ctx, uuid)

	resp, _, err := req.Execute()

	return &resp, err
}


// PUT api.beget.com/v1/vps/{id}/password
func (p *provider) resetPassword(ctx context.Context, uuid string) (*vpsClient.ManageResetPasswordResponse, error) {
	req := p.vpsClient.ManageServiceApi.ManageServiceResetPassword(ctx, uuid)

	resp, _, err := req.Execute()

	return &resp, err
}

// PUT api.beget.com/v1/vps/server/{id}/info
func (p *provider) updateInfo(ctx context.Context, uuid, hostname string) (*vpsClient.ManageUpdateInfoResponse, error) {
	var displayName = "Go Generated Client Test"
	var description = "Test client from generated code"

	req := p.vpsClient.ManageServiceApi.ManageServiceUpdateInfo(ctx, uuid)
	payload := vpsClient.NewManageUpdateInfoRequest()

	payload.Hostname = &hostname
	payload.Description = &description
	payload.DisplayName = &displayName

	req = req.ManageUpdateInfoRequest(*payload)

	resp, httpResp, err := req.Execute()

	if httpResp.StatusCode != 200 {
		return nil,errors.New(fmt.Sprintf("http code is %d", httpResp.StatusCode))
	}

	return &resp, err
}

// POST api.beget.com/v1/vps/server/{id}/reboot
func (p *provider) rebootVps(ctx context.Context, uuid string) (*vpsClient.ManageRebootVpsResponse, error) {
	req := p.vpsClient.ManageServiceApi.ManageServiceRebootVps(ctx, uuid)

	resp, _, err := req.Execute()

	return &resp, err
}

// POST api.beget.com/v1/vps/server/{id}/reinstall
func (p *provider) reinstall(ctx context.Context, uuid string, operationSystemId int32) (*vpsClient.ManageReinstallResponse, error) {
	var password = "L%v2sZ2j"

	req := p.vpsClient.ManageServiceApi.ManageServiceReinstall(ctx, uuid)
	payload := vpsClient.NewManageReinstallRequest()

	payload.Password = &password
	payload.OperatingSystemId = &operationSystemId

	req = req.ManageReinstallRequest(*payload)

	resp, httpResp, err := req.Execute()

	if httpResp.StatusCode != 200 {
		return nil,errors.New(fmt.Sprintf("http code is %d", httpResp.StatusCode))
	}

	return &resp, err
}
