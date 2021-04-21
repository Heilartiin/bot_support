package gs_nike_products

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Heilartin/bot_support/models"
	"io/ioutil"
	"net/http"
)

func (c *Client) getClient() *http.Client  {
	return c.c
}

func (c *Client) doReq(method, path string, body []byte) (*http.Response, error)  {
	req, err := http.NewRequest(method, c.ApiUrl + path, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Host", "gs-product.nike.com")
	req.Header.Add("Origin", "https://gs.nike.com")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36")
	req.Header.Add("Accept-Language", "ru")
	req.Header.Add("Connection", "keep-alive")
	return c.getClient().Do(req)
}

func (c *Client) GetProductByID(productID string) (result *models.NikeProduct, err error)  {
	path := fmt.Sprintf("/api/v1/Product/%s", productID)
	resp, err := c.doReq("GET", path, nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case http.StatusOK:
		br, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.log.Error(err)
			return nil, err
		}
		var productOrigin *ProductResponse
		err = json.Unmarshal(br, &productOrigin)
		if err != nil {
			c.log.Error(err)
			return nil, err
		}
		var sizes []string
		for _, v := range productOrigin.ProductMerchSkus {

			sizes = append(sizes, v.NikeSize)
		}
		countryDetail, err := c.getCountryDetail(productOrigin, "RU")
		if err != nil {
			c.log.Error(err)
			return nil, err
		}
		productCharge, err := c.getProductCharge(countryDetail)
		if err != nil {
			c.log.Error(err)
			return nil, err
		}
		product := models.NikeProduct{
			ProductID:    productID,
			Sizes:        sizes,
			Name:         productOrigin.Name,
			Color:        countryDetail.ColorDescription,
			StyleID:      productOrigin.StyleColor,
			Price:        productCharge.Value.Value,
			CurrencyCode: productCharge.Value.CurrencyCode,
			ProductImage: fmt.Sprintf("https://imageresize.24i.com/?w=300&url=%s", countryDetail.ProductImageURL),
		}
		return &product, nil
	case http.StatusServiceUnavailable, http.StatusInternalServerError:
		return nil, errors.New(resp.Status)
	default:
		return nil, errors.New(resp.Status)
	}

}

func (c *Client) getCountryDetail(productOrigin *ProductResponse, countryCode string) (*ProductCountryDetail, error)  {
	if productOrigin.CountryDetails == nil || len(productOrigin.CountryDetails) == 0 {
		return nil, errors.New("country details is empty")
	}
	for _, v := range productOrigin.CountryDetails {
		if v.CountryCode == countryCode {
			return v, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("not fount country detail for country code: %s", countryCode))
}

func (c *Client) getProductCharge(countryDetail *ProductCountryDetail) (*ProductCharge, error)  {
	if countryDetail.Charges == nil || len(countryDetail.Charges) == 0 {
		return nil, errors.New("product charge is empty")
	}
	for _, v := range countryDetail.Charges {
		if v.Name == "Item" {
			if v.Value == nil {
				return nil, errors.New("product charge value is nil")
			}
			return v, nil
		}
	}
	return nil, errors.New("not found product charge")
}

func (c *Client) getSizeInfo(productOrigin *ProductResponse, sizeName string) (*ProductSku, error)  {
	if productOrigin.ProductMerchSkus == nil || len(productOrigin.ProductMerchSkus) == 0 {
		return nil, errors.New("product sizes is nil")
	}
	for _, v := range productOrigin.ProductMerchSkus {
		if v.NikeSize == sizeName {
			return v, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("size name %s not found", sizeName))
}

func (c *Client) getSizeInfoByID(productOrigin *ProductResponse, sizeID string) (*ProductSku, error)  {
	if productOrigin.ProductMerchSkus == nil || len(productOrigin.ProductMerchSkus) == 0 {
		return nil, errors.New("product sizes is nil")
	}
	for _, v := range productOrigin.ProductMerchSkus {
		if v.ID == sizeID {
			return v, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("size id %s not found", sizeID))
}