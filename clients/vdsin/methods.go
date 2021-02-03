package vdsin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Heilartin/bot_support/models"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

func (v *VDSinClient) doReq(method, apiToken, path string, body []byte) (*http.Response, error)  {
	url := fmt.Sprintf("%s/%s", v.cfg.ApiUrl, path)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		v.log.Error(errors.WithStack(err))
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json'")
	req.Header.Add("Authorization", apiToken)

	resp, err := v.c.Do(req)
	if err != nil {
		v.log.Error(errors.WithStack(err))
		return nil, err
	}
	return resp, nil
}

func (v *VDSinClient) GetBalance(apiToken string) (*Balance, error)  {
	path := "v1/account.balance"
	resp, err := v.doReq("GET", apiToken, path, nil)
	if err != nil {
		v.log.Error(errors.WithStack(err))
		return nil, err
	}
	if resp.StatusCode != 200 {
		v.log.Error(errors.WithStack(errors.New(resp.Status)))
		return nil, errors.New(resp.Status)
	}
	br, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		v.log.Error(errors.WithStack(err))
		return nil, err
	}
	var balance *Balance
	err = json.Unmarshal(br, &balance)
	if err != nil {
		v.log.Error(errors.WithStack(err))
		return nil, err
	}
	return balance, nil
}

func (v *VDSinClient) GetAllServers(apiToken string) (*ServerResponse, error)  {
	path := "v1/server"
	resp, err := v.doReq("GET", apiToken, path, nil)
	if err != nil {
		v.log.Error(errors.WithStack(err))
		return nil, err
	}
	if resp.StatusCode != 200 {
		v.log.Error(errors.WithStack(errors.New(resp.Status)))
		return nil, errors.New(resp.Status)
	}
	br, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		v.log.Error(errors.WithStack(err))
		return nil, err
	}
	var servers *ServerResponse
	err = json.Unmarshal(br, &servers)
	if err != nil {
		v.log.Error(errors.WithStack(err))
		return nil, err
	}
	return servers, nil
}


func (v *VDSinClient) GetServerByID(apiToken string, serverID int) (*ServerCredential, error)  {
	path := fmt.Sprintf("v1/server.vnc/%d", serverID)
	resp, err := v.doReq("GET", apiToken, path, nil)
	if err != nil {
		v.log.Error(errors.WithStack(err))
		return nil, err
	}
	if resp.StatusCode != 200 {
		v.log.Error(errors.WithStack(errors.New(resp.Status)))
		return nil, errors.New(resp.Status)
	}
	br, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		v.log.Error(errors.WithStack(err))
		return nil, err
	}
	var serv *ServerCredential
	err = json.Unmarshal(br, &serv)
	if err != nil {
		v.log.Error(errors.WithStack(err))
		return nil, err
	}
	return serv, nil
}

func (v *VDSinClient) DeleteServerByID(apiToken string, serverID int) (*ServerCredential, error)  {
	path := fmt.Sprintf("v1/server/%d", serverID)
	resp, err := v.doReq("DELETE", apiToken, path, nil)
	if err != nil {
		v.log.Error(errors.WithStack(err))
		return nil, err
	}
	if resp.StatusCode != 200 {
		v.log.Error(errors.WithStack(errors.New(resp.Status)))
		return nil, errors.New(resp.Status)
	}
	br, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		v.log.Error(errors.WithStack(err))
		return nil, err
	}
	var serv *ServerCredential
	err = json.Unmarshal(br, &serv)
	if err != nil {
		v.log.Error(errors.WithStack(err))
		return nil, err
	}
	return serv, nil
}

func (v *VDSinClient) GetAllCredential(apiToken string) ([]*models.Server, error) {
	resp, err := v.GetAllServers(apiToken)
	if err != nil {
		v.log.Error(errors.WithStack(err))
		return nil, err
	}
	if len(resp.Data) == 0 {
		v.log.Error(errors.WithStack(errors.New("Servers is empty")))
		return nil, errors.New("Servers is empty")
	}
	var servers []*models.Server
	for index, s := range resp.Data {
		creds, err := v.GetServerByID(apiToken, s.ID)
		if err != nil {
			v.log.Error(errors.WithStack(err))
			return nil, err
		}
		serv :=  models.Server{
			Name:     			fmt.Sprintf("A-%d", index + 1),
			Host:     			s.IP.IP,
			CredentialPassword: creds.Data.Password,
			CredentialUserName: "",
			SubMode:            "",
			ConnectionType:     "Microsoft RDP",
		}
		servers = append(servers, &serv)
	}
	return servers, nil
}

func (v *VDSinClient) DeleteAllServers(apiToken string) error {
	resp, err := v.GetAllServers(apiToken)
	if err != nil {
		v.log.Error(errors.WithStack(err))
		return err
	}
	if len(resp.Data) == 0 {
		v.log.Error(errors.WithStack(errors.New("Servers is empty")))
		return errors.New("Servers is empty")
	}
	for _, s := range resp.Data {
		_, err := v.DeleteServerByID(apiToken, s.ID)
		if err != nil {
			v.log.Error(errors.WithStack(err))
			return err
		}
	}
	return nil
}