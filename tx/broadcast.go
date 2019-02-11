package tx

import (
	ctypes "github.com/tendermint/tendermint/rpc/core/types"

	atypes "github.com/amolabs/amoabci/amo/types"
	"github.com/amolabs/amoconsole/util"
)

// Transfer handles transfer transaction
func Transfer(from, to atypes.Address, amount *uint64) (*ctypes.ResultBroadcastTxCommit, error) {
	return util.RPCBroadcastTxCommit(util.MakeMessage(atypes.TxTransfer, atypes.Transfer{
		From:   from,
		To:     to,
		Amount: *amount,
	}))
}

// Purchase handles purchase transaction
func Purchase(from atypes.Address, fileHash atypes.Hash) (*ctypes.ResultBroadcastTxCommit, error) {
	return util.RPCBroadcastTxCommit(util.MakeMessage(atypes.TxPurchase, atypes.Purchase{
		From:     from,
		FileHash: fileHash,
	}))
}
