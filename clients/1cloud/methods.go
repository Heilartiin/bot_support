package _cloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Heilartin/bot_support/models"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

func (v *OneCClient) doReq(method, apiToken, path string, body []byte) (*http.Response, error)  {
	url := fmt.Sprintf("%s/%s", v.cfg.ApiUrl, path)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		v.log.Error(errors.WithStack(err))
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json'")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiToken))

	resp, err := v.c.Do(req)
	if err != nil {
		v.log.Error(errors.WithStack(err))
		return nil, err
	}
	return resp, nil
}


func (v *OneCClient) GetBalance(apiToken string) (*Balance, error)  {
	path := "account"
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

func (v *OneCClient) GetAllServers(apiToken string) ([]*Server, error)  {
	path := "Server"
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
	var servers []*Server
	err = json.Unmarshal(br, &servers)
	if err != nil {
		v.log.Error(errors.WithStack(err))
		return nil, err
	}
	return servers, nil
}


func (v *OneCClient) GetAllCredential(apiToken string) ([]*models.Server, error) {
	resp, err := v.GetAllServers(apiToken)
	if err != nil {
		v.log.Error(errors.WithStack(err))
		return nil, err
	}
	if len(resp) == 0 {
		v.log.Error(errors.WithStack(errors.New("Servers is empty")))
		return nil, errors.New("Servers is empty")
	}
	var servers []*models.Server
	for index, s := range resp {
		serv :=  models.Server{
			Name:     			fmt.Sprintf("1C A-%d", index + 1),
			Host:     			s.IP,
			CredentialPassword: s.AdminPassword,
			CredentialUserName: "",
			SubMode:            "",
			ConnectionType:     "Microsoft RDP",
		}
		servers = append(servers, &serv)
	}
	return servers, nil
}