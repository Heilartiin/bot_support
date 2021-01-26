package mrporter

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

func (c *MrpClient) doReq(path string) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.cfg.ApiUrl + path, nil)
	if err != nil {
		c.log.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	req.Header.Add("authority", "www.mrporter.com")
	req.Header.Add("x-ibm-client-id", c.cfg.ClientID)
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36")
	resp, err := c.client.Do(req)
	if err != nil {
		c.log.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

func (c *MrpClient) GetProductInfo(storeID, productID string) (*MrpProduct, error) {
	path := fmt.Sprintf("/api/inseason/search/resources/store/mrp_%s/productview/%s?locale=en_GB", storeID, productID)
	resp, err := c.doReq(path)
	if err != nil {
		c.log.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	if resp.StatusCode != http.StatusOK {
		t := fmt.Sprintf("Bad status code: %d StoreID: %s", resp.StatusCode, storeID)
		c.log.Error(errors.WithStack(errors.New(t)))
		return nil, errors.WithStack(errors.New(t))
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.log.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	var mrpResp *MrpResp
	err = json.Unmarshal(b, &mrpResp)
	if err != nil {
		c.log.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	if mrpResp.RecordSetTotal == 0 {
		return nil, errors.New("Product Not Found")
	}
	return mrpResp.Products[0], nil
}