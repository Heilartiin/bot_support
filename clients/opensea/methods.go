package opensea

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Heilartin/bot_support/models"
	"io/ioutil"
	"net/http"
	"strings"
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
	collectionInfo := contractInfo.Collection

	res = &models.OpenSeaCollection{
		Address:             contractAddress,
		Name:    		     contractInfo.Name,
		Slug:   			 strings.Replace(strings.ToLower(contractInfo.Name), " ", "", 10),

		EtherscanUrl:        "https://etherscan.io/address/" + contractAddress,
		ImageUrl:            contractInfo.ImageUrl,
		ServiceFee:          contractInfo.OpenseaSellerFeeBasisPoints / 100,
		CreatorFee:          contractInfo.DevSellerFeeBasisPoints / 100,
		TxsEtherscan:  		 "https://etherscan.io/txs?a=" + contractAddress,
		PendingTxsEtherscan: fmt.Sprintf("https://etherscan.io/txsPending?a=%s&m=hf", contractAddress),
		ContractCreated:     c.parseTime(contractInfo.CreatedDate),
		OSCollectionCreated: time.Time{},
	}
	res.OSUrl = "https://opensea.io/collection/" + res.Slug
	if res.ImageUrl == "" {
		res.ImageUrl = "https://sun9-69.userapi.com/impg/dMSu7oaG5M63yaGc_8d2dVvACI-iWa4309qZyg/4HvGJnnbKYc.jpg?size=750x750&quality=96&sign=4711322ea31ba83e81170caa348424c6&type=album"
	}

	if collectionInfo != nil {
		res.Name = collectionInfo.Name
		res.Slug =  collectionInfo.Slug
		res.OSUrl = "https://opensea.io/collection/" + collectionInfo.Slug
		res.OSCollectionCreated = c.parseTime(collectionInfo.CreatedDate)
		res.NFTNerdUrl = "https://nftnerds.ai/collection/" + res.Slug
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
		stats, err := c.RetrievingCollectionStats(collectionInfo.Slug)
		if err != nil {
			c.log.Error(err)
			return nil, err
		}
		if stats.Stats == nil {
			c.log.Error("stats is nil")
			return nil, errors.New("stats is nil")
		}
		res.FloorPrice  =   stats.Stats.FloorPrice
		res.TotalVolume =   stats.Stats.TotalVolume
		res.TotalSales  =   stats.Stats.TotalSales
		res.NumOwners   =   stats.Stats.NumOwners
		res.OneDayVolume = stats.Stats.OneDayVolume
		res.OneDaySales = stats.Stats.OneDaySales
		res.FloorSell = stats.Stats.FloorPrice * (1 - (res.ServiceFee + res.CreatorFee) / 100)
		return res, nil
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
