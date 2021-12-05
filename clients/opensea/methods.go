package opensea

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Heilartin/bot_support/models"
	"io/ioutil"
	"net/http"
	"time"
)

func (c *Client) doReq(method, path string, body []byte) (response *http.Response, err error)  {
	req, err := http.NewRequest(method, c.cfg.ApiUrl + path, bytes.NewBuffer(body))
	if err != nil {
		c.log.Error(err)
		return
	}
	req.Header.Add("authority","api.opensea.io")
	req.Header.Add("x-api-key",c.cfg.ApiKey)
	req.Header.Add("user-agent",c.cfg.UserAgent)
	req.Header.Add("accept-language","ru")

	response, err = c.httpClient.Do(req)
	if err != nil {
		c.log.Error(err)
		return
	}
	switch response.StatusCode {
	case http.StatusOK:
		return
	default:
		return nil, errors.New(fmt.Sprintf("error: %s", response.Status))
	}
}

func (c *Client) RetrievingSingleContract(address string) (res *ContractResponse, err error)  {
	resp, err := c.doReq(http.MethodGet, "/api/v1/asset_contract/" + address, nil)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.jsonUnmarshal(resp, &res)
	if err != nil {
		c.log.Error(err)
		return
	}
	return
}

func (c *Client) RetrievingSingleCollection(collectionName string) (res *ContractResponse, err error)  {
	resp, err := c.doReq(http.MethodGet, "/api/v1/collection/" + collectionName, nil)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.jsonUnmarshal(resp, &res)
	if err != nil {
		c.log.Error(err)
		return
	}
	return
}

func (c *Client) RetrievingCollectionStats(collectionName string) (res *StatsResponse, err error) {
	resp, err := c.doReq(http.MethodGet, fmt.Sprintf("/api/v1/collection/%s/stats", collectionName), nil)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.jsonUnmarshal(resp, &res)
	if err != nil {
		c.log.Error(err)
		return
	}
	return
}

func (c *Client) jsonUnmarshal(resp *http.Response, res interface{}) (err error) {
	br, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = json.Unmarshal(br, res)
	if err != nil {
		c.log.Error(err)
		return
	}
	return
}

func (c *Client) GetInformationByContract(contractAddress string) (res *models.OpenSeaCollection, err error) {
	contractInfo, err := c.RetrievingSingleContract(contractAddress)
	if err != nil {
		c.log.Error(err)
		return
	}
	if contractInfo.Collection == nil {
		c.log.Error("collection is nil")
		return nil, errors.New("collection is nil")
	}
	collectionInfo := contractInfo.Collection

	stats, err := c.RetrievingCollectionStats(contractInfo.Collection.Slug)
	if err != nil {
		c.log.Error(err)
		return
	}
	if stats.Stats == nil {
		c.log.Error("stats is nil")
		return nil, errors.New("stats is nil")
	}
	res = &models.OpenSeaCollection{
		Address:             contractAddress,
		EtherscanUrl:        "https://etherscan.io/address/" + contractAddress,
		Name:                collectionInfo.Name,
		Slug:                collectionInfo.Slug,
		OSUrl:               "https://opensea.io/collection/" + collectionInfo.Slug,
		ImageUrl:            contractInfo.ImageUrl,
		ServiceFee:          contractInfo.OpenseaSellerFeeBasisPoints / 100,
		CreatorFee:          contractInfo.DevSellerFeeBasisPoints / 100,
		FloorPrice:          stats.Stats.FloorPrice,
		TotalVolume:         stats.Stats.TotalVolume,
		TotalSales:          stats.Stats.TotalSales,
		TxsEtherscan:  		 "https://etherscan.io/txs?a=" + contractAddress,
		PendingTxsEtherscan: fmt.Sprintf("https://etherscan.io/txsPending?a=%s&m=hf", contractAddress),
		NumOwners:           stats.Stats.NumOwners,
		ContractCreated:     c.parseTime(contractInfo.CreatedDate),
		OSCollectionCreated: c.parseTime(collectionInfo.CreatedDate),
	}
	if collectionInfo.TelegramUrl != nil {
		res.TelegramUrl = *collectionInfo.TelegramUrl
	}
	if collectionInfo.TwitterUsername != nil {
		res.TwitterUrl = "https://twitter.com/" + *collectionInfo.TwitterUsername
	}
	if collectionInfo.InstagramUsername != nil {
		res.InstagramUrl ="https://www.instagram.com/" + *collectionInfo.InstagramUsername
	}
	if collectionInfo.DiscordUrl != nil {
		res.DiscordUrl = *collectionInfo.DiscordUrl
	}
	if collectionInfo.ExternalUrl != nil {
		res.ExternalLink = *collectionInfo.ExternalUrl
	}
	return
}

func (c *Client) parseTime(t string) time.Time  {
	res, err := time.Parse("2006-01-02T15:04:05.999999", t)
	if err != nil {
		return time.Time{}
	}
	return res
}
