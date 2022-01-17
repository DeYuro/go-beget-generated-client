package main

import (
	"context"
	"errors"
	"fmt"
	vpsClient "github.com/deyuro/go-beget-generated-client/internal/apiclient/vps"
	"github.com/k0kubun/pp"
	"strconv"
)

// https://8gwifi.org/sshfunctions.jsp dummy key
const dummyPublicKey = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCxdhCjn5lo2sLj1aX5aTRXDtGTsDQ/THOIh1dKGSN65ij/JcLrCATo8QBMVER8WlEXLLr2auIWFVp9x3OX0zW/9E58CoihtcKccv9bZfm/cxfM5Z8XugSnkdMJZEYF7yHQmzKKrYsI2S+22ztQ3psM3jXxDClWgV1v0n97ohRL7Vv/kN3lU6PSWhELfa+5SMl1XayZ92kroBAEamgZ7vo7RgJ9PoH9137BeND53SNXWJcY/xWEE505zX67/wP6Wo7TMpbs9VZNt67OaWeUQ2ze9SFXPfQ5HDciYmDU8QnPu7GhoTXzLRKmxq9RN6CvSad18yVS3X69OSKIslslN0QP \n"


func sshTest(ctx context.Context, p provider, uuid string) error {

	pp.Println(fmt.Sprintf("AddSsh request:Expect ssh != nil"))

	rAddSsh, err := p.addSsh(ctx)
	if err != nil {
		return err
	}
	pp.Println("Response")
	pp.Println(rAddSsh)

	if rAddSsh.Key == nil {
		return errors.New(fmt.Sprintf("Expect ssh != nil, got nil"))
	}

	pp.Println(fmt.Sprintf("AttachSshKey request:Expect VPS has key with id == %d", rAddSsh.Key.GetId()))

	rAttachSsh, err := p.attachSsh(ctx,uuid,rAddSsh.Key.GetId())
	if err != nil {
		return err
	}
	pp.Println("Response")
	pp.Println(rAttachSsh)

	if rAttachSsh.Vps.GetSshKeys()[0].GetId() != rAddSsh.Key.GetId() {
		return errors.New(fmt.Sprintf("Expect VPS has key with id == %d: got %d", rAddSsh.Key.GetId(),rAttachSsh.Vps.GetSshKeys()[0].GetId()))
	}

	pp.Println(fmt.Sprintf("detachSshKey request:Expect VPS has no SSH keys"))

	rDetachSsh, err := p.detachSsh(ctx,uuid,rAddSsh.Key.GetId())
	if err != nil {
		return err
	}
	pp.Println("Response")
	pp.Println(rDetachSsh)

	if len(rDetachSsh.Vps.GetSshKeys()) > 0 {
		return errors.New(fmt.Sprintf("Expect VPS has no SSH keys: GOT len(rAttachSsh.Vps.GetSshKeys() == %d", len(rAttachSsh.Vps.GetSshKeys())))
	}

	pp.Println(fmt.Sprintf("RemoveSsh request:Expect removed ssh != nil"))

	rRemoveSsh, err := p.removeSsh(ctx,rAddSsh.Key.GetId())
	if err != nil {
		return err
	}
	pp.Println("Response")
	pp.Println(rRemoveSsh)

	if rRemoveSsh.Key == nil {
		return errors.New(fmt.Sprintf("Expect removed ssh != nil, got nil"))
	}

	return nil
}
// POST api.beget.com/v1/vps/sshKey
func (p *provider) addSsh(ctx context.Context) (*vpsClient.SshKeyAddResponse,error) {
	req := p.vpsClient.SshKeyServiceApi.SshKeyServiceAdd(ctx)
	payload := vpsClient.NewSshKeyAddRequest()

	payload.SetName("go generated client dummy key")
	payload.SetPublicKey(dummyPublicKey)

	req = req.SshKeyAddRequest(*payload)

	resp, httpResp, err := req.Execute()

	if httpResp.StatusCode != 200 {
		return nil,errors.New(fmt.Sprintf("http code is %d", httpResp.StatusCode))
	}

	return &resp, err
}

// DELETE api.beget.com/v1/vps/sshKey/{id}
func (p *provider) removeSsh(ctx context.Context, sshId int32) (*vpsClient.SshKeyRemoveResponse,error)  {
	req := p.vpsClient.SshKeyServiceApi.SshKeyServiceRemove(ctx, strconv.Itoa(int(sshId)))
	req = req.Force("1")

	resp, httpResp, err := req.Execute()

	if httpResp.StatusCode != 200 {
		return nil,errors.New(fmt.Sprintf("http code is %d", httpResp.StatusCode))
	}

	return &resp, err
}
// POST api.beget.com/v1/vps/{id}/sshKey/{ssh_key_id}
func (p *provider) attachSsh(ctx context.Context, uuid string, sshId int32) (*vpsClient.ManageAttachSshKeyResponse, error) {
	req := p.vpsClient.ManageServiceApi.ManageServiceAttachSshKey(ctx, uuid, strconv.Itoa(int(sshId)))

	resp, httpResp, err := req.Execute()

	if httpResp.StatusCode != 200 {
		return nil,errors.New(fmt.Sprintf("http code is %d", httpResp.StatusCode))
	}

	return &resp, err
}

// DELETE api.beget.com/v1/vps/{id}/sshKey/{ssh_key_id}
func (p *provider) detachSsh(ctx context.Context, uuid string, sshId int32) (*vpsClient.ManageDetachSshKeyResponse, error) {
	req := p.vpsClient.ManageServiceApi.ManageServiceDetachSshKey(ctx, uuid, strconv.Itoa(int(sshId)))

	resp, httpResp, err := req.Execute()

	if httpResp.StatusCode != 200 {
		return nil,errors.New(fmt.Sprintf("http code is %d", httpResp.StatusCode))
	}

	return &resp, err
}
