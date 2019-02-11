package util

import (
	"encoding/json"

	atypes "github.com/amolabs/amoabci/amo/types"

	cmn "github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/rpc/client"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"
)

var (
	rpcRemote     = "tcp://0.0.0.0:26657"
	rpcWsEndpoint = "/websocket"
)

// MakeMessage handles making tx message
func MakeMessage(t string, payload interface{}) types.Tx {
	raw, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	msg := atypes.Message{
		Type:      t,
		Timestamp: tmtime.Now().Unix(),
		Payload:   raw,
	}
	tx, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return tx
}

// RPCBroadcastTxCommit handles sending transactions
func RPCBroadcastTxCommit(tx types.Tx) (*ctypes.ResultBroadcastTxCommit, error) {
	cli := client.NewHTTP(rpcRemote, rpcWsEndpoint)
	return cli.BroadcastTxCommit(tx)
}

// RPCABCIQuery handles querying
func RPCABCIQuery(path string, data cmn.HexBytes) (*ctypes.ResultABCIQuery, error) {
	cli := client.NewHTTP(rpcRemote, rpcWsEndpoint)
	return cli.ABCIQuery(path, data)
}
