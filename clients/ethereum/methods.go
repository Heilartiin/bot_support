package ethereum

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (c *Client) TransactionByHash(hash string) (response *types.Transaction, pending bool, err error)  {
	response, pending, err = c.ethClient.TransactionByHash(context.Background(), common.HexToHash(hash))
	if err != nil {
		c.log.Error(err)
		return
	}
	return
}
