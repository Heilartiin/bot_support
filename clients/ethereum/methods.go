package ethereum

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func (c *Client) TransactionByHash(hash string) (response *types.Transaction, pending bool, err error)  {
	response, pending, err = c.ethClient.TransactionByHash(context.Background(), common.HexToHash(hash))
	if err != nil {
		c.log.Error(err)
		return
	}
	return
}

func (c *Client) GetCurrentBlockNumber() (blockNumber uint64, err error) {
	blockNumber, err = c.ethClient.BlockNumber(context.Background())
	if err != nil {
		c.log.Error(err)
		return
	}
	return
}

func (c *Client) GetBlockByNumber(int2 *big.Int) (block *types.Block, err error) {
	block, err = c.ethClient.BlockByNumber(context.Background(), int2)
	if err != nil {
		c.log.Error(err)
		return
	}
	return
}
