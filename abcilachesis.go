package abcilachesis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/rpc/core"
	core_types"github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/rpc/client"
	crypto "github.com/tendermint/tendermint/crypto"
	cmn "github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/libs/pubsub/query"
	tmpubsub "github.com/tendermint/tendermint/libs/pubsub"
//	tmquery "github.com/tendermint/tendermint/libs/pubsub/query"
	rpctypes "github.com/tendermint/tendermint/rpc/lib/types"
	"github.com/tendermint/tendermint/types"
//	"github.com/tendermint/tendermint/p2p"
)

//*********************************************************************************************
//1.		func ClientABCIInfo() error {
//2.		func ClientABCIQuery(path string, data cmn.HexBytes, height int64, trusted bool) error {
//3.		func ClientBlockAtHeight(Height int64) error {
//4.		func ClientBlockResult(heightPtr *int64) error {
//5.		func ClientBlockChainInfo(minHeight, maxHeight int64) error {
//6.		func ClientBroadcastTxAsync(tx types.Tx) error {
//7.		func ClientBroadcastTxSync(tx types.Tx) error {
//8.		func ClientBroadcastTxCommit(tx types.Tx) error {
//9.		func ClientCommit(heightPtr *int64) error {
//10		func ClientConsensusParams(heightPtr *int64) error {
//11		func ClientConsensusState() error {
//12		func ClientDumpConsensusState () error {
//13		func ClientGenesis() error {
//14		func ClientHealth(wsCtx rpctypes.WSRPCContext) error {
//15		func ClientNetInfo() error {
//16		func ClientNumUnconfirmedTxs() error {
//17		func ClientUnsafeDialSeeds(seeds []string) error {
//18		func ClientUnsafeDialPeers(peers []string, persistant bool) error {
//19		func ClientUnsubscribe(wsCtx rpctypes.WSRPCContext, query string) error {
//20		func ClientUnsubscribeAll(wsCtx rpctypes.WSRPCContext) error {
//21		func ClientValidators(heightPtr *int64) error {
//22		func ClientStatus() error {
//23		func ClientSubscribe(????????) error {									------------------------------>>>>> help
//24		func ClienTx(hash []byte, prove bool) error {
//25		func ClientTxSearch(query string) error {
//26		func CoreUnconfirmedTxs(limit int) error {
//*********************************************************************************************


//*********************************************************************************************
//  json layout
//*********************************************************************************************
// {
//   "error": "",
//   "result": {
//     "txs": [],
//     "n_txs": 0
//   },
//   "id": "",
//   "jsonrpc": "2.0"
// }
//*********************************************************************************************
type ABCIUnconfirmedTxsJSON struct {
	Error  		string `json:"error"`
	Result 		struct {
		Txs 	[]types.Tx `json:"txs"`
		N_Txs  	int `json:"n_txs"`
	} `json:"result"`
	ID      	string `json:"id"`
	Jsonrpc 	string `json:"jsonrpc"`
}	

//26
func CoreUnconfirmedTxs(limit int) ([]byte, error) {

	var abciut ABCIUnconfirmedTxsJSON
	//  as seen in Tendermint code
	abciut.Jsonrpc = "2.0"
	//abcini.ID ??v

 	result, err := core.UnconfirmedTxs(limit)

	if err != nil {
		fmt.Println("Handle core.UnconfirmedTxs(limit) error: ", err)
		abciut.Error = err.Error()
	} else {
		abciut.Result.Txs   	=  result.Txs
		abciut.Result.N_Txs   	=  result.N
	}

	MarshalledJson, err := json.Marshal(abciut)

	return MarshalledJson, err
}



//*********************************************************************************************
//  json layout
//*********************************************************************************************
// {
//   "jsonrpc": "2.0",
//   "id": "",
//   "result": {
// 	   "txs": [
//       {
//         "proof": {
//           "Proof": {
//             "aunts": [
//               "J3LHbizt806uKnABNLwG4l7gXCA=",
//               "iblMO/M1TnNtlAefJyNCeVhjAb0=",
//               "iVk3ryurVaEEhdeS0ohAJZ3wtB8=",
//               "5hqMkTeGqpct51ohX0lZLIdsn7Q=",
//               "afhsNxFnLlZgFDoyPpdQSe0bR8g="
//             ]
//           },
//           "Data": "mvZHHa7HhZ4aRT0xMDA=",
//           "RootHash": "F6541223AA46E428CB1070E9840D2C3DF3B6D776",
//           "Total": 32,
//           "Index": 31
//         },
//         "tx": "mvZHHa7HhZ4aRT0xMDA=",
//         "tx_result": {},
//         "index": 31,
//         "height": 12,
//         "hash": "2B8EC32BA2579B3B8606E42C06DE2F7AFA2556EF"
//       }
//     ],
//     "total_count": 1
//   }
// }
//*********************************************************************************************
type ABCITxSearchJSON struct {
	Jsonrpc 			string `json:"jsonrpc"`
	ID      			string `json:"id"`
	Result  			struct {
		Txs 			[]struct {
			Proof 		types.TxProof `json:"proof"`
			Tx       	types.Tx `json:"tx"`
			TxResult 	abci.ResponseDeliverTx  `json:"tx_result"`
			Index  		uint32 `json:"index"`
			Height 		int64 `json:"height"`
			Hash   		cmn.HexBytes `json:"hash"`
		} `json:"txs"`
		TotalCount 		int `json:"total_count"`
	} `json:"result"`
}

type StTx struct {
	Proof 		types.TxProof 			`json:"proof"`
	Tx       	types.Tx 				`json:"tx"`
	TxResult 	abci.ResponseDeliverTx  `json:"tx_result"`
	Index  		uint32 					`json:"index"`
	Height 		int64 					`json:"height"`
	Hash   		cmn.HexBytes 			`json:"hash"`
}

//25
func ClientTxSearch(query string) ([]byte, error) {

	var abcits ABCITxSearchJSON
	//  as seen in Tendermint code
	abcits.Jsonrpc = "2.0"
	//abcini.ID ??

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	tx, err := client.TxSearch(query, true, 1, 1)

	if err != nil {
		fmt.Println("Handle client.TxSearch(q, true) error: ", err)
	}  else {

		var sttx StTx

		abcits.Result.TotalCount 			=	tx.TotalCount

		for i := 0; i < abcits.Result.TotalCount; i++ {
			sttx.Proof 			= tx.Txs[i].Proof
			sttx.Tx       		= tx.Txs[i].Tx
			sttx.TxResult 		= tx.Txs[i].TxResult
			sttx.Index  		= tx.Txs[i].Index
			sttx.Height 		= tx.Txs[i].Height
			sttx.Hash 			= tx.Txs[i].Hash

			abcits.Result.Txs 	=	append(abcits.Result.Txs, sttx)
		}	
	}

	MarshalledJson, err := json.Marshal(abcits)

	return MarshalledJson, err
}



//*********************************************************************************************
//  json layout
//*********************************************************************************************
// {
// 	"error": "",
// 	"result": {
// 		"proof": {
// 			"Proof": {
// 				"aunts": []
// 			},
// 			"Data": "YWJjZA==",
// 			"RootHash": "2B8EC32BA2579B3B8606E42C06DE2F7AFA2556EF",
// 			"Total": 1,
// 			"Index": 0
// 		},
// 		"tx": "YWJjZA==",
// 		"tx_result": {
// 			"log": "",
// 			"data": "",
// 			"code": 0
// 		},
// 		"index": 0,
// 		"height": 52,
//		"hash": "2B8EC32BA2579B3B8606E42C06DE2F7AFA2556EF"
// 	},
// 	"id": "",
// 	"jsonrpc": "2.0"
// }
//*********************************************************************************************
type ABCITxJSON struct {
	Error  				string `json:"error"`
	Result 				struct {
		Proof 			types.TxProof `json:"proof"`
		Tx       		types.Tx `json:"tx"`
		TxResult 		abci.ResponseDeliverTx  `json:"tx_result"`
		Index  			uint32 `json:"index"`
		Height 			int64 `json:"height"`
		Hash   			cmn.HexBytes `json:"hash"`		
	} `json:"result"`
	ID      			string `json:"id"`
	Jsonrpc 			string `json:"jsonrpc"`
}

//*********************************************************************************************
// ### Query Parameters
//
// | Parameter | Type   | Default | Required | Description                                               |
// |-----------+--------+---------+----------+-----------------------------------------------------------|
// | hash      | []byte | nil     | true     | The transaction hash                                      |
// | prove     | bool   | false   | false    | Include a proof of the transaction inclusion in the block |
//
// ### Returns
//
// - `proof`: the `types.TxProof` object
// - `tx`: `[]byte` - the transaction
// - `tx_result`: the `abci.Result` object
// - `index`: `int` - index of the transaction
// - `height`: `int` - height of the block where this transaction was in
// - `hash`: `[]byte` - hash of the transaction
//*********************************************************************************************
//24
func ClienTx(hash []byte, prove bool) ([]byte, error) {
	// Tx allows you to query the transaction results. `nil` could mean the
	// transaction is in the mempool, invalidated, or was not sent in the first
	// place.

	var abcit ABCITxJSON
	//  as seen in Tendermint code
	abcit.Jsonrpc = "2.0"
	//abcini.ID ??

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	tx, err := client.Tx([]byte(hash), prove)

	if err != nil {
		fmt.Println("Handle client.Tx([]byte(hash), prove) error: ", err)
		abcit.Error = err.Error()
	} else {
		abcit.Result.Proof 			= tx.Proof
		abcit.Result.Tx       		= tx.Tx
		abcit.Result.TxResult 		= tx.TxResult
		abcit.Result.Index  		= tx.Index
		abcit.Result.Height 		= tx.Height
		abcit.Result.Hash 			= tx.Hash
	}

	MarshalledJson, err := json.Marshal(abcit)

	return MarshalledJson, err
}


//*********************************************************************************************
// **Taken from the "https://github.com/tendermint/tendermint/blob/master/rpc/core/events.go#L87" page**
//
// Subscribe for events via WebSocket.
//   
// To tell which events you want, you need to provide a query. query is a
// string, which has a form: "condition AND condition ..." (no OR at the
// moment). condition has a form: "key operation operand". key is a string with
// a restricted set of possible symbols ( \t\n\r\\()"'=>< are not allowed).
// operation can be "=", "<", "<=", ">", ">=", "CONTAINS". operand can be a
// string (escaped with single quotes), number, date or time.
//
// Examples:
//		tm.event = 'NewBlock'								# new blocks
//		tm.event = 'CompleteProposal'				# node got a complete proposal
//		tm.event = 'Tx' AND tx.hash = 'XYZ' # single transaction
//		tm.event = 'Tx' AND tx.height = 5		# all txs of the fifth block
//		tx.height = 5												# all txs of the fifth block
//
// Tendermint provides a few predefined keys: tm.event, tx.hash and tx.height.
// Note for transactions, you can define additional keys by providing tags with
// DeliverTx response.
//
//		DeliverTx{
//			Tags: []*KVPair{
//				"agent.name": "K",
//			}
//	  }
//
//		tm.event = 'Tx' AND agent.name = 'K'
//		tm.event = 'Tx' AND account.created_at >= TIME 2013-05-03T14:45:00Z
//		tm.event = 'Tx' AND contract.sign_date = DATE 2017-01-01
//		tm.event = 'Tx' AND account.owner CONTAINS 'Igor'
//
// See list of all possible events here
// https://godoc.org/github.com/tendermint/tendermint/types#pkg-constants
//
// For complete query syntax, check out
// https://godoc.org/github.com/tendermint/tendermint/libs/pubsub/query.
//
//*********************************************************************************************
//*********************************************************************************************
//  json layout
//*********************************************************************************************
// {
// 	"error": "",
// 	"result": {},
// 	"id": "",
// 	"jsonrpc": "2.0"
// }
//*********************************************************************************************

type ABCISubscribeJSON struct {
	Error  		string `json:"error"`
	Result 		interface{} `json:"result"`
	ID      	string `json:"id"`
	Jsonrpc 	string `json:"jsonrpc"`
}	

//23
//
//   NOTE : "Websocket only" according to https://tendermint.com/rpc/?go#subscribe
//
func ClientSubscribe(clientstring string, querystring string) ([]byte, error) {

//  Only returns err if unsuccessful
	var abcis ABCISubscribeJSON
	//  as seen in Tendermint code
	abcis.Jsonrpc = "2.0"
	//abcis.ID ??

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	ctx, cancel := context.WithTimeout(context.Background(), 1) // timeout)

	defer cancel()

	// TODO:  follow query.MustParse to confirm value to pass  --> More examples in comments above "json layout"

	query := query.MustParse(querystring)
	txs := make(chan interface{})

	//  func Subscribe(wsCtx rpctypes.WSRPCContext, query string) but supplied code below has 4 items.

	err := client.Subscribe(ctx, clientstring, query, txs)

	if err != nil {
		fmt.Println("err= ", err)
		abcis.Error = err.Error()
	} else {
   		abcis.Result = "200 OK"
	}

	// NOTE:: txs is a channel!!!!
	// Nowhere in func Subscribe is any data being passed to txs to come out this side. Defaq? 

	go func() {
	    for e := range txs {
	        fmt.Println("got ", e.(types.EventDataTx))
		}
	}()

	MarshalledJson, err := json.Marshal(abcis)

	return MarshalledJson, err
}

//*********************************************************************************************
//  json layout
//*********************************************************************************************
// {
// "jsonrpc": "2.0",
// "id": "",
// "result": {
//   "node_info": {
//   		"id": "53729852020041b956e86685e24394e0bee4373f",
//   		"listen_addr": "10.0.2.15:26656",
//   		"network": "test-chain-Y1OHx6",
//   		"version": "0.24.0-2ce1abc2",
//   		"channels": "4020212223303800",
//   		"moniker": "ubuntu-xenial",
//   		"other": {
//   			"amino_version": "0.12.0",
//   			"p2p_version": "0.5.0",
//   			"consensus_version": "v1/0.2.2",
//   			"rpc_version": "0.7.0/3",
//   			"tx_index": "on",
//   			"rpc_addr": "tcp://0.0.0.0:26657"
//   		}
//   	},
//   	"sync_info": {
//   		"latest_block_hash": "F51538DA498299F4C57AC8162AAFA0254CE08286",
//   		"latest_app_hash": "0000000000000000",
//   		"latest_block_height": "18",
//   		"latest_block_time": "2018-09-17T11:42:19.149920551Z",
//   		"catching_up": false
//   	},
//   	"validator_info": {
//   		"address": "D9F56456D7C5793815D0E9AF07C3A355D0FC64FD",
//   		"pub_key": {
//   			"type": "tendermint/PubKeyEd25519",
//   			"value": "wVxKNtEsJmR4vvh651LrVoRguPs+6yJJ9Bz174gw9DM="
//   		},
//   		"voting_power": "10"
//   	}
//   }
// }
//*********************************************************************************************
type ABCIStatusJSON struct {
	Jsonrpc 					string `json:"jsonrpc"`
	ID      					string `json:"id"`
	Result  					struct {
//		NodeInfo 				p2p.NodeInfo
		
		NodeInfo 				struct {
			ID         			string `json:"id"`
			ListenAddr 			string `json:"listen_addr"`
			Network    			string `json:"network"`
			Version   		 	string `json:"version"`
			Channels   			cmn.HexBytes `json:"channels"`
			Moniker    			string `json:"moniker"`
			Other      			struct {
				AminoVersion     string `json:"amino_version"`
				P2PVersion       string `json:"p2p_version"`
				ConsensusVersion string `json:"consensus_version"`
				RPCVersion       string `json:"rpc_version"`
				TxIndex          string `json:"tx_index"`
				RPCAddr          string `json:"rpc_addr"`
			} `json:"other"`
		} `json:"node_info"`
		
		SyncInfo 				struct {
			LatestBlockHash   	cmn.HexBytes    `json:"latest_block_hash"`
			LatestAppHash     	cmn.HexBytes    `json:"latest_app_hash"`
			LatestBlockHeight 	int64    `json:"latest_block_height"`
			LatestBlockTime   	time.Time `json:"latest_block_time"`
			CatchingUp        	bool      `json:"catching_up"`
		} `json:"sync_info"`
		ValidatorInfo 			struct {
			Address 			cmn.HexBytes `json:"address"`
			PubKey  			crypto.PubKey `json:"pub_key"`
			VotingPower 		int64 `json:"voting_power"`
		} `json:"validator_info"`
	} `json:"result"`
}

//22
func ClientStatus() ([]byte, error) {

	var abcis ABCIStatusJSON
	//  as seen in Tendermint code
	abcis.Jsonrpc = "2.0"
	//abcis.ID ??

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	result, err := client.Status()

	if err != nil {
		fmt.Println("Handle client.Status() error: ", err, result)
	} else {

		abcis.Result.NodeInfo.ID    	 = string(result.NodeInfo.ID)
		abcis.Result.NodeInfo.ListenAddr = result.NodeInfo.ListenAddr
		abcis.Result.NodeInfo.Network    = result.NodeInfo.Network
		abcis.Result.NodeInfo.Version    = result.NodeInfo.Version
		abcis.Result.NodeInfo.Moniker    = result.NodeInfo.Moniker
		abcis.Result.NodeInfo.Channels   = result.NodeInfo.Channels

		abcis.Result.NodeInfo.Other.AminoVersion    = result.NodeInfo.Other.AminoVersion
		abcis.Result.NodeInfo.Other.P2PVersion   	= result.NodeInfo.Other.P2PVersion
		abcis.Result.NodeInfo.Other.ConsensusVersion = result.NodeInfo.Other.ConsensusVersion
		abcis.Result.NodeInfo.Other.RPCVersion 		= result.NodeInfo.Other.RPCVersion
		abcis.Result.NodeInfo.Other.TxIndex			= result.NodeInfo.Other.TxIndex
		abcis.Result.NodeInfo.Other.RPCAddr			= result.NodeInfo.Other.RPCAddress

		abcis.Result.SyncInfo  		= result.SyncInfo
		abcis.Result.ValidatorInfo 	= result.ValidatorInfo
	}

	MarshalledJson, err := json.Marshal(abcis)

	return MarshalledJson, err
}



//*********************************************************************************************
//  json layout
//*********************************************************************************************
// {
// 	"error": "",
// 	"result": {
// 		"validators": [
// 			{
// 				"accum": 0,
// 				"voting_power": 10,
// 				"pub_key": {
// 					"data": "68DFDA7E50F82946E7E8546BED37944A422CD1B831E70DF66BA3B8430593944D",
// 					"type": "ed25519"
// 				},
// 				"address": "E89A51D60F68385E09E716D353373B11F8FACD62"
// 			}
// 		],
// 		"block_height": 5241
// 	},
// 	"id": "",
// 	"jsonrpc": "2.0"
// }
//*********************************************************************************************
type ABCIValidatorsJSON struct {
	Error  				string `json:"error"`
	Result 				struct {
		Validators 		[]*types.Validator `json:"validators"`
		BlockHeight 	int64 `json:"block_height"`
	} `json:"result"`
	ID      			string `json:"id"`
	Jsonrpc 			string `json:"jsonrpc"`
}	

//21
func ClientValidators(heightPtr *int64) ([]byte, error) {

	var abciv ABCIValidatorsJSON
	//  as seen in Tendermint code
	abciv.Jsonrpc = "2.0"
	//abciv.ID ??

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	state, err := client.Validators(heightPtr)

	if err != nil {
		fmt.Println("Handle client.Validators() error: ", err, state)
		abciv.Error = err.Error()
	} else {

		var stValidator *types.Validator	

		for i := 0; i < len(state.Validators); i++ {
			stValidator.Accum 		= state.Validators[i].Accum
			stValidator.PubKey 		= state.Validators[i].PubKey
			stValidator.VotingPower = state.Validators[i].VotingPower
			stValidator.Address 	= state.Validators[i].Address

			abciv.Result.Validators = append(abciv.Result.Validators, stValidator)
		}	

	}

	abciv.Result.BlockHeight 	= state.BlockHeight

	MarshalledJson, err := json.Marshal(abciv)

	return MarshalledJson, err
}


//*********************************************************************************************
//  json layout
//*********************************************************************************************
// {
// 	"error": "",
// 	"result": {},
// 	"id": "",
// 	"jsonrpc": "2.0"
// }
//*********************************************************************************************
type ABCIUnsubscribeAllJSON struct {
	Error  		string `json:"error"`
	Result 		interface{} `json:"result"`
	ID      	string `json:"id"`
	Jsonrpc 	string `json:"jsonrpc"`
}		

//20
func ClientUnsubscribeAll(wsCtx context.Context,  subscriber string) ([]byte, error) {

	var abciusa ABCIUnsubscribeAllJSON
	//  as seen in Tendermint code
	abciusa.Jsonrpc = "2.0"
	//abciusa.ID ??

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	err := client.UnsubscribeAll(wsCtx, subscriber)

	if err != nil {
		fmt.Println("Handle client.UnsubscribeAll(...) error: ", err)
		abciusa.Error = err.Error()
	} else {
   		abciusa.Result = "200 OK"
	}

	MarshalledJson, err := json.Marshal(abciusa)

	return MarshalledJson, err
}


//*********************************************************************************************
//  json layout
//*********************************************************************************************
// {
// 	"error": "",
// 	"result": {},
// 	"id": "",
// 	"jsonrpc": "2.0"
// }
//*********************************************************************************************
type ABCIUnsubscribeJSON struct {
	Error  		string `json:"error"`
	Result 		interface{} `json:"result"`
	ID      	string `json:"id"`
	Jsonrpc 	string `json:"jsonrpc"`
}		

//19
func ClientUnsubscribe(wsCtx context.Context, query tmpubsub.Query) ([]byte, error) {

	var abcius ABCIUnsubscribeJSON
	//  as seen in Tendermint code
	abcius.Jsonrpc = "2.0"
	//abcius.ID ??

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	err := client.Unsubscribe(wsCtx, "test-client", query)

	if err != nil {
		fmt.Println("Handle client.Unsubscribe(...) error: ", err)
		abcius.Error = err.Error()
	} else {
   		abcius.Result = "200 OK"
	}

	MarshalledJson, err := json.Marshal(abcius)

	return MarshalledJson, err
}


//*********************************************************************************************
//   See /net_info for details ...?
//*********************************************************************************************
//18
func CoreUnsafeDialPeers(peers []string, persistant bool) error {

//	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
 	result, err := core.UnsafeDialPeers(peers, persistant)

	if err != nil {
		fmt.Println("Handle client.UnconfirmedTxs() error: ", err, result)
		return err
	} 



	return nil
}

//*********************************************************************************************
//   See /net_info for details ...?
//******************************************************************************************
//17
func CoreUnsafeDialSeeds(seeds []string) error {

//	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
 	result, err := core.UnsafeDialSeeds(seeds)

	if err != nil {
		fmt.Println("Handle client.UnconfirmedTxs() error: ", err, result)
		return err
	} 
	return nil
}


//*********************************************************************************************
//  json layout
//*********************************************************************************************
// {
//   "error": "",
//   "result": {
//     "txs": null,
//     "n_txs": 0
//   },
//   "id": "",
//   "jsonrpc": "2.0"
// }
//*********************************************************************************************
type ABCINumUnconfirmedTxsJSON struct {
	Error  		string `json:"error"`
	Result 		struct {
		Txs 	[]types.Tx `json:"txs"`
		N_Txs  	int `json:"n_txs"`
	} `json:"result"`
	ID      	string `json:"id"`
	Jsonrpc 	string `json:"jsonrpc"`
}	

//16
func ClientNumUnconfirmedTxs() ([]byte, error) {

	var abcinut ABCINumUnconfirmedTxsJSON
	//  as seen in Tendermint code
	abcinut.Jsonrpc = "2.0"
	//abcini.ID ??v

 	result, err := core.UnconfirmedTxs(10)

	if err != nil {
		fmt.Println("Handle client.UnconfirmedTxs() error: ", err)
		abcinut.Error = err.Error()
	} else {
		abcinut.Result.N_Txs    = result.N
		abcinut.Result.Txs    = result.Txs
	}

	MarshalledJson, err := json.Marshal(abcinut)

	return MarshalledJson, err
}


//*********************************************************************************************
//  json layout
//*********************************************************************************************
// {
// 	"error": "",
// 	"result": {
//		"n_peers": 0,
// 		"peers": [],
// 		"listeners": [
// 			"Listener(@10.0.2.15:26656)"
// 		],
// 		"listening": true
// 	},
// 	"id": "",
// 	"jsonrpc": "2.0"
// }
//*********************************************************************************************
type ABCINetInfoJSON struct {
	Error  			string `json:"error"`
	Result 			struct {
		NPeers    	int           `json:"n_peers"`
		Peers     	[]core_types.Peer `json:"peers"`
		Listeners 	[]string      `json:"listeners"`
		Listening 	bool          `json:"listening"`
	} `json:"result"`
	ID      		string `json:"id"`
	Jsonrpc 		string `json:"jsonrpc"`
}

//15
func ClientNetInfo() ([]byte, error) {

	var abcini ABCINetInfoJSON

	//  as seen in Tendermint code
	abcini.Jsonrpc = "2.0"
	//abcini.ID ??v

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
 	info, err := client.NetInfo()

	if err != nil {
		fmt.Println("Handle client.NetInfo() error: ", err)	
		abcini.Error = err.Error()
	} else {
		abcini.Result.NPeers 	=	info.NPeers
		abcini.Result.Peers     =	info.Peers
		abcini.Result.Listeners =	info.Listeners
		abcini.Result.Listening =	info.Listening
	}
	
	MarshalledJson, err := json.Marshal(abcini)

	return MarshalledJson, err
}




//*********************************************************************************************
//  json layout
//*********************************************************************************************
// {
// 	"error": "",
// 	"result": {},
// 	"id": "",
// 	"jsonrpc": "2.0"
// }
//*********************************************************************************************
type ABCIHealthJSON struct {
	Error  		string `json:"error"`
	Result 		interface{} `json:"result"`
	ID      	string `json:"id"`
	Jsonrpc 	string `json:"jsonrpc"`
}		

//14
//
//  NOTE:
// Get node health. Returns empty result (200 OK) on success, no response - in
// case of an error.

func ClientHealth(wsCtx rpctypes.WSRPCContext) ([]byte, error) {

	var abcih ABCIHealthJSON
	//  as seen in Tendermint code
	abcih.Jsonrpc = "2.0"
	//abcini.ID ??v

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	result, err := client.Health()

	if err != nil {
		fmt.Println("Handle client.Health() error: ", err, result)
		abcih.Error = err.Error()
	} else {
   		abcih.Result = "200 OK"
   	}

	MarshalledJson, err := json.Marshal(abcih)

	return MarshalledJson, err
}


//*********************************************************************************************
//  json layout
//*********************************************************************************************
// {
// 	"error": "",
// 	"result": {
// 		"genesis": {
// 			"app_hash": "",
// 			"validators": [
// 				{
// 					"name": "",
// 					"power": 10,
// 					"pub_key": {
// 						"data": "68DFDA7E50F82946E7E8546BED37944A422CD1B831E70DF66BA3B8430593944D",
// 						"type": "ed25519"
// 					}
// 				}
// 			],
// 			"chain_id": "test-chain-6UTNIN",
// 			"genesis_time": "2017-05-29T15:05:41.671Z"
// 		}
// 	},
// 	"id": "",
// 	"jsonrpc": "2.0"
// }
//*********************************************************************************************
type ABCIGenesisJSON struct {
	Error  					string `json:"error"`
	Result 					struct {
		Genesis 			struct {
			AppHash    		cmn.HexBytes `json:"app_hash"`
			Validators 		[]struct {
				Name   		string `json:"name"`
				Power  		int64    `json:"power"`
				PubKey 		crypto.PubKey `json:"pub_key"`
			} `json:"validators"`
			ChainID     	string    `json:"chain_id"`
			GenesisTime 	time.Time `json:"genesis_time"`
		} `json:"genesis"`
	} `json:"result"`
	ID      				string `json:"id"`
	Jsonrpc 				string `json:"jsonrpc"`
}

type GenesisValidator 		struct {
	Name   		string `json:"name"`
	Power  		int64    `json:"power"`
	PubKey 		crypto.PubKey `json:"pub_key"`
}

//13
func ClientGenesis() ([]byte, error) {

	var abcig ABCIGenesisJSON
	//  as seen in Tendermint code
	abcig.Jsonrpc = "2.0"
	//abcig.ID ??v

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	genesis, err := client.Genesis()

	if err != nil {
		fmt.Println("Handle client.UnconfirmedTxs() error: ", err)
		abcig.Error = err.Error()
	} else {

		abcig.Result.Genesis.AppHash 		= genesis.Genesis.AppHash

		var stValidator   GenesisValidator	

		for i := 0; i < len(genesis.Genesis.Validators); i++ {
			stValidator.PubKey 		= genesis.Genesis.Validators[i].PubKey
			stValidator.Power		= genesis.Genesis.Validators[i].Power
			stValidator.Name 		= genesis.Genesis.Validators[i].Name

			abcig.Result.Genesis.Validators = append(abcig.Result.Genesis.Validators, stValidator)
		}	

		abcig.Result.Genesis.ChainID		= genesis.Genesis.ChainID
		abcig.Result.Genesis.GenesisTime	= genesis.Genesis.GenesisTime

	}

	MarshalledJson, err := json.Marshal(abcig)

	return MarshalledJson, err
}


//*********************************************************************************************
//  json layout
//*********************************************************************************************
// {
//   "jsonrpc": "2.0",
//   "id": "",
//   "result": {
//     "round_state": {
//       "height": 7185,
//       "round": 0,
//       "step": 1,
//       "start_time": "2018-05-12T13:57:28.440293621-07:00",
//       "commit_time": "2018-05-12T13:57:27.440293621-07:00",
//       "validators": {
//         "validators": [
//           {
//             "address": "B5B3D40BE53982AD294EF99FF5A34C0C3E5A3244",
//             "pub_key": {
//               "type": "tendermint/PubKeyEd25519",
//               "value": "SBctdhRBcXtBgdI/8a/alTsUhGXqGs9k5ylV1u5iKHg="
//             },
//             "voting_power": 10,
//             "accum": 0
//           }
//         ],
//         "proposer": {
//           "address": "B5B3D40BE53982AD294EF99FF5A34C0C3E5A3244",
//           "pub_key": {
//             "type": "tendermint/PubKeyEd25519",
//             "value": "SBctdhRBcXtBgdI/8a/alTsUhGXqGs9k5ylV1u5iKHg="
//           },
//           "voting_power": 10,
//           "accum": 0
//         }
//       },
//       "proposal": null,
//       "proposal_block": null,
//       "proposal_block_parts": null,
//       "locked_round": 0,
//       "locked_block": null,
//       "locked_block_parts": null,
//       "valid_round": 0,
//       "valid_block": null,
//       "valid_block_parts": null,
//       "votes": [
//         {
//           "round": 0,
//           "prevotes": "_",
//           "precommits": "_"
//         }
//       ],
//       "commit_round": -1,
//       "last_commit": {
//         "votes": [
//           "Vote{0:B5B3D40BE539 7184/00/2(Precommit) 14F946FA7EF0 /702B1B1A602A.../ @ 2018-05-12T20:57:27.342Z}"
//         ],
//         "votes_bit_array": "x",
//         "peer_maj_23s": {}
//       },
//       "last_validators": {
//         "validators": [
//           {
//             "address": "B5B3D40BE53982AD294EF99FF5A34C0C3E5A3244",
//             "pub_key": {
//               "type": "tendermint/PubKeyEd25519",
//               "value": "SBctdhRBcXtBgdI/8a/alTsUhGXqGs9k5ylV1u5iKHg="
//             },
//             "voting_power": 10,
//             "accum": 0
//           }
//         ],
//         "proposer": {
//           "address": "B5B3D40BE53982AD294EF99FF5A34C0C3E5A3244",
//           "pub_key": {
//             "type": "tendermint/PubKeyEd25519",
//             "value": "SBctdhRBcXtBgdI/8a/alTsUhGXqGs9k5ylV1u5iKHg="
//           },
//           "voting_power": 10,
//           "accum": 0
//         }
//       }
//     },
//     "peers": [
//       {
//         "node_address": "30ad1854af22506383c3f0e57fb3c7f90984c5e8@172.16.63.221:26656",
//         "peer_state": {
//           "round_state": {
//             "height": 7185,
//             "round": 0,
//             "step": 1,
//             "start_time": "2018-05-12T13:57:27.438039872-07:00",
//             "proposal": false,
//             "proposal_block_parts_header": {
//               "total": 0,
//               "hash": ""
//             },
//             "proposal_block_parts": null,
//             "proposal_pol_round": -1,
//             "proposal_pol": "_",
//             "prevotes": "_",
//             "precommits": "_",
//             "last_commit_round": 0,
//             "last_commit": "x",
//             "catchup_commit_round": -1,
//             "catchup_commit": "_"
//           },
//           "stats": {
//             "last_vote_height": 7184,
//             "votes": 255,
//             "last_block_part_height": 7184,
//             "block_parts": 255
//           }
//         }
//       }
//     ]
//   }
// }
//*********************************************************************************************
type ABCIDumpConsensusStateJSON struct {
	Jsonrpc 						string `json:"jsonrpc"`
	ID      						string `json:"id"`
	Result  						struct {
		RoundState 					struct {
			Height     				int    `json:"height"`
			Round      				int    `json:"round"`
			Step       				int    `json:"step"`
			StartTime  				string `json:"start_time"`
			CommitTime 				string `json:"commit_time"`
			Validators 				struct {
				Validators 			[]struct {
					Address 		string `json:"address"`
					PubKey  		struct {
						Type  		string `json:"type"`
						Value 		string `json:"value"`
					} `json:"pub_key"`
					VotingPower 	int `json:"voting_power"`
					Accum       	int `json:"accum"`
				} `json:"validators"`
				Proposer 			struct {
					Address 		string `json:"address"`
					PubKey  		struct {
						Type  		string `json:"type"`
						Value 		string `json:"value"`
					} `json:"pub_key"`
					VotingPower 	int `json:"voting_power"`
					Accum       	int `json:"accum"`
				} `json:"proposer"`
			} `json:"validators"`
			Proposal           		interface{} `json:"proposal"`
			ProposalBlock      		interface{} `json:"proposal_block"`
			ProposalBlockParts 		interface{} `json:"proposal_block_parts"`
			LockedRound        		int         `json:"locked_round"`
			LockedBlock        		interface{} `json:"locked_block"`
			LockedBlockParts   		interface{} `json:"locked_block_parts"`
			ValidRound         		int         `json:"valid_round"`
			ValidBlock         		interface{} `json:"valid_block"`
			ValidBlockParts    		interface{} `json:"valid_block_parts"`
			Votes              		[]struct {
				Round      			int    `json:"round"`
				Prevotes   			string `json:"prevotes"`
				Precommits 			string `json:"precommits"`
			} `json:"votes"`
			CommitRound 			int `json:"commit_round"`
			LastCommit  			struct {
				Votes         		[]time.Time `json:"votes"`
				VotesBitArray 		string      `json:"votes_bit_array"`
				PeerMaj23S    		struct {
				} `json:"peer_maj_23s"`
			} `json:"last_commit"`
			LastValidators 			struct {
				Validators 			[]struct {
					Address 		string `json:"address"`
					PubKey  		struct {
						Type  		string `json:"type"`
						Value 		string `json:"value"`
					} `json:"pub_key"`
					VotingPower 	int `json:"voting_power"`
					Accum       	int `json:"accum"`
				} `json:"validators"`
				Proposer 			struct {
					Address 		string `json:"address"`
					PubKey  		struct {
						Type  		string `json:"type"`
						Value 		string `json:"value"`
					} `json:"pub_key"`
					VotingPower 	int `json:"voting_power"`
					Accum       	int `json:"accum"`
				} `json:"proposer"`
			} `json:"last_validators"`
		} `json:"round_state"`
		Peers 						[]struct {
			NodeAddress 			string `json:"node_address"`
			PeerState  				struct {
				RoundState 			struct {
					Height          int    `json:"height"`
					Round           int    `json:"round"`
					Step            int    `json:"step"`
					StartTime       string `json:"start_time"`
					Proposal        bool   `json:"proposal"`
					ProposalBlockPartsHeader struct {
						Total 			int    `json:"total"`
						Hash  			string `json:"hash"`
					} `json:"proposal_block_parts_header"`
					ProposalBlockParts interface{} `json:"proposal_block_parts"`
					ProposalPolRound   int         `json:"proposal_pol_round"`
					ProposalPol        string      `json:"proposal_pol"`
					Prevotes           string      `json:"prevotes"`
					Precommits         string      `json:"precommits"`
					LastCommitRound    int         `json:"last_commit_round"`
					LastCommit         string      `json:"last_commit"`
					CatchupCommitRound int         `json:"catchup_commit_round"`
					CatchupCommit      string      `json:"catchup_commit"`
				} `json:"round_state"`
				Stats 					struct {
					LastVoteHeight      int `json:"last_vote_height"`
					Votes               int `json:"votes"`
					LastBlockPartHeight int `json:"last_block_part_height"`
					BlockParts          int `json:"block_parts"`
				} `json:"stats"`
			} `json:"peer_state"`
		} `json:"peers"`
	} `json:"result"`
}

//12
func ClientDumpConsensusState () ([]byte, error) {

	var abcidcs ABCIDumpConsensusStateJSON
	//  as seen in Tendermint code
	abcidcs.Jsonrpc = "2.0"
	//abcidcs.ID ??v


	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
    state, err := client.DumpConsensusState()

	if err != nil {
		fmt.Println("Handle client.UnconfirmedTxs() error: ", err, state)
//		abcidcs.Error = err.Error()
	} else {

		//  Place in loop **********
		//	
		//for i := 0; i < len(state.Peers); i++ {
			abcidcs.Result.Peers[0].NodeAddress = state.Peers[0].NodeAddress
			
		//}	
	}
	MarshalledJson, err := json.Marshal(abcidcs)

	return MarshalledJson, err
}


//*********************************************************************************************
//  json layout
//*********************************************************************************************
//{
//  "jsonrpc": "2.0",
//  "id": "",
//  "result": {
//    "round_state": {
//      "height/round/step": "9336/0/1",
//      "start_time": "2018-05-14T10:25:45.72595357-04:00",
//      "proposal_block_hash": "",
//      "locked_block_hash": "",
//      "valid_block_hash": "",
//      "height_vote_set": [
//        {
//          "round": 0,
//          "prevotes": [
//            "nil-Vote"
//          ],
//          "prevotes_bit_array": "BA{1:_} 0/10 = 0.00",
//          "precommits": [
//            "nil-Vote"
//          ],
//          "precommits_bit_array": "BA{1:_} 0/10 = 0.00"
//        }
//      ]
//    }
//  }
//}
//*********************************************************************************************
type ABCIConsensusStateJSON struct {
	Jsonrpc 						string `json:"jsonrpc"`
	ID      						string `json:"id"`
	Result  						struct {
/*		RoundState 					struct {
			HeightRoundStep   		string `json:"height/round/step"`
			StartTime         		string `json:"start_time"`
			ProposalBlockHash 		string `json:"proposal_block_hash"`
			LockedBlockHash   		string `json:"locked_block_hash"`
			ValidBlockHash    		string `json:"valid_block_hash"`
			HeightVoteSet     		[]struct {
				Round              	int      `json:"round"`
				Prevotes           	[]string `json:"prevotes"`
				PrevotesBitArray   	string   `json:"prevotes_bit_array"`
				Precommits         	[]string `json:"precommits"`
				PrecommitsBitArray 	string   `json:"precommits_bit_array"`
			} `json:"height_vote_set"`
		} `json:"round_state"` */
		RoundState 					json.RawMessage `json:"round_state"` 
	} `json:"result"`
}

//11
func ClientConsensusState() ([]byte, error) {

	var abcics ABCIConsensusStateJSON
	//  as seen in Tendermint code
	abcics.Jsonrpc = "2.0"
	//abcics.ID ??v

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	state, err := client.ConsensusState()

	if err != nil {
		fmt.Println("Handle client.ConsensusState() error: ", err, state)
//		abcics.Error = err.Error()
	} else {
		abcics.Result.RoundState = state.RoundState
	}

	// state = *core_types.ResultConsensusState)
// UNSTABLE
//type ResultConsensusState struct {
//	RoundState json.RawMessage `json:"round_state"`               ------->> already json format!!!
//}
//	if err := json.Unmarshal(state.RoundState, &abcics.Result.RoundState); err != nil {
//		return err
//	}
	MarshalledJson, err := json.Marshal(abcics)

	return MarshalledJson, err
}


//*********************************************************************************************
//  json layout
//*********************************************************************************************
//{
//   "jsonrpc": "2.0",
//   "id": "",
//   "result": {
//     "block_height": "1",
//     "consensus_params": {
//       "block_size_params": {
//         "max_txs_bytes": "22020096",
//         "max_gas": "-1"
//       },
//       "evidence_params": {
//         "max_age": "100000"
//       }
//     }
//   }
// }
//*********************************************************************************************
type ABCIConsensusParamsJSON struct {
	Jsonrpc 				string `json:"jsonrpc"`
	ID      				string `json:"id"`
	Result  				struct {
		BlockHeight     	int64 `json:"block_height"`
		ConsensusParams 	struct {
			BlockSizeParams struct {
				MaxTxsBytes int64 `json:"max_txs_bytes"`
				MaxGas      int64 `json:"max_gas"`
			} `json:"block_size_params"`
			EvidenceParams 	struct {
				MaxAge 		int64 `json:"max_age"`
			} `json:"evidence_params"`
		} `json:"consensus_params"`
	} `json:"result"`
}

//10
func ClientConsensusParams(heightPtr *int64) ([]byte, error) {

	var abcicp ABCIConsensusParamsJSON
	//  as seen in Tendermint code
	abcicp.Jsonrpc = "2.0"
	//abcicp.ID ??

	state, err := core.ConsensusParams(heightPtr)

	if err != nil {
		fmt.Println("Handle core.ConsensusParams(heightPtr) error: ", err)
	//	abcicp.Error = err.Error()
	} else {

		abcicp.Result.BlockHeight		=	state.BlockHeight

		abcicp.Result.ConsensusParams.BlockSizeParams.MaxTxsBytes	=	state.ConsensusParams.BlockSize.MaxBytes
		abcicp.Result.ConsensusParams.BlockSizeParams.MaxGas		=	state.ConsensusParams.BlockSize.MaxGas

		abcicp.Result.ConsensusParams.EvidenceParams.MaxAge		=	state.ConsensusParams.EvidenceParams.MaxAge

	}

	MarshalledJson, err := json.Marshal(abcicp)

	return MarshalledJson, err
}


//*********************************************************************************************
//  json layout
//*********************************************************************************************
//{
//   "error": "",
//   "result": {
//     "canonical": true,
//     "commit": {
//       "precommits": [
//         {
//           "signature": {
//             "data": "00970429FEC652E9E21D106A90AE8C5413759A7488775CEF4A3F44DC46C7F9D941070E4FBE9ED54DF247FA3983359A0C3A238D61DE55C75C9116D72ABC9CF50F",
//             "type": "ed25519"
//           },
//           "block_id": {
//             "parts": {
//               "hash": "9E37CBF266BC044A779E09D81C456E653B89E006",
//               "total": 1
//             },
//             "hash": "CC6E861E31CA4334E9888381B4A9137D1458AB6A"
//           },
//           "type": 2,
//           "round": 0,
//           "height": 11,
//           "validator_index": 0,
//           "validator_address": "E89A51D60F68385E09E716D353373B11F8FACD62"
//         }
//       ],
//       "blockID": {
//         "parts": {
//           "hash": "9E37CBF266BC044A779E09D81C456E653B89E006",
//           "total": 1
//         },
//         "hash": "CC6E861E31CA4334E9888381B4A9137D1458AB6A"
//       }
//     },
//     "header": {
//       "app_hash": "",
//       "chain_id": "test-chain-6UTNIN",
//       "height": 11,
//       "time": "2017-05-29T15:05:54.893Z",
//       "num_txs": 0,
//       "last_block_id": {
//         "parts": {
//           "hash": "277A4DBEF91483A18B85F2F5677ABF9694DFA40F",
//           "total": 1
//         },
//         "hash": "96B1D2F2D201BA4BC383EB8224139DB1294944E5"
//       },
//       "last_commit_hash": "3CE0C9727CE524BA9CB7C91E28F08E2B94001087",
//       "data_hash": "",
//       "validators_hash": "9365FC80F234C967BD233F5A3E2AB2F1E4B0E5AA"
//     }
//   },
//   "id": "",
//   "jsonrpc": "2.0"
// }
//*********************************************************************************************
type ABCICommitJSON struct {
	Error  						string `json:"error"`
	Result 						struct {
		Canonical 				bool `json:"canonical"`
		Commit    				struct {
			Precommits 			[]struct {
				Signature 		[]byte `json:"signature"`
				BlockID 		types.BlockID `json:"block_id"`
				Type             byte    `json:"type"`
				Round            int    `json:"round"`
				Height           int64   `json:"height"`
				ValidatorIndex   int    `json:"validator_index"`
				ValidatorAddress cmn.HexBytes `json:"validator_address"`
			} `json:"precommits"`
			BlockID 			types.BlockID `json:"blockID"`
		} `json:"commit"`
		Header 					struct {
			AppHash     		cmn.HexBytes    `json:"app_hash"`
			ChainID     		string    `json:"chain_id"`
			Height      		int64       `json:"height"`
			Time        		time.Time `json:"time"`
			NumTxs      		int64       `json:"num_txs"`
			LastBlockID 		struct {
				Parts 			struct {
					Hash  		cmn.HexBytes `json:"hash"`
					Total 		int    `json:"total"`
				} `json:"parts"`
				Hash 			string `json:"hash"`
			} `json:"last_block_id"`
			LastCommitHash 		cmn.HexBytes `json:"last_commit_hash"`
			DataHash       		cmn.HexBytes `json:"data_hash"`
			ValidatorsHash 		cmn.HexBytes `json:"validators_hash"`
		} `json:"header"`
	} `json:"result"`
	ID      					string `json:"id"`
	Jsonrpc 					string `json:"jsonrpc"`
}


type Vote struct {
	Signature        []byte    `json:"signature"`
	BlockID          types.BlockID   `json:"block_id"` // zero if vote is nil.
	Type             byte      `json:"type"`
	Round            int       `json:"round"`
	Height           int64     `json:"height"`
	ValidatorIndex   int       `json:"validator_index"`
	ValidatorAddress cmn.HexBytes   `json:"validator_address"`
//	Timestamp        time.Time `json:"timestamp"`
}

//9
func ClientCommit(heightPtr *int64) ([]byte, error) {

	var abcic ABCICommitJSON
	//  as seen in Tendermint code
	abcic.Jsonrpc = "2.0"
	//abcic.ID ??

	var StVote Vote

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	info, err := client.Commit(heightPtr)

	if err != nil {
		fmt.Println("Handleclient.Commit(heightPtr) error: ", err, info)
		abcic.Error = err.Error()
	} else {

//canonical

		abcic.Result.Canonical  	= info.CanonicalCommit

//commit

		for i := 0; i < len(info.SignedHeader.Commit.Precommits); i++ {

 			StVote.ValidatorAddress = 	info.SignedHeader.Commit.Precommits[i].ValidatorAddress
 			StVote.ValidatorIndex 	= 	info.SignedHeader.Commit.Precommits[i].ValidatorIndex
 			StVote.Height 			= 	info.SignedHeader.Commit.Precommits[i].Height
 			StVote.Round 			= 	info.SignedHeader.Commit.Precommits[i].Round
 			StVote.Type 			= 	info.SignedHeader.Commit.Precommits[i].Type
 			StVote.BlockID 			= 	info.SignedHeader.Commit.Precommits[i].BlockID
 			StVote.Signature 		= 	info.SignedHeader.Commit.Precommits[i].Signature

			abcic.Result.Commit.Precommits  	= append(abcic.Result.Commit.Precommits , StVote)
		}

		abcic.Result.Commit.BlockID =  abcic.Result.Commit.BlockID 

//header

		abcic.Result.Header.AppHash  		= info.SignedHeader.Header.AppHash
		abcic.Result.Header.ChainID  		= info.SignedHeader.Header.ChainID
		abcic.Result.Header.Height  		= info.SignedHeader.Header.Height
		abcic.Result.Header.Time  			= info.SignedHeader.Header.Time
		abcic.Result.Header.NumTxs  		= info.SignedHeader.Header.NumTxs
		abcic.Result.Header.LastCommitHash  = info.SignedHeader.Header.LastCommitHash
		abcic.Result.Header.DataHash  		= info.SignedHeader.Header.DataHash
		abcic.Result.Header.ValidatorsHash  = info.SignedHeader.Header.ValidatorsHash

		abcic.Result.Header.LastBlockID.Parts.Hash  	= info.SignedHeader.Header.LastBlockID.PartsHeader.Hash
		abcic.Result.Header.LastBlockID.Parts.Total  	= info.SignedHeader.Header.LastBlockID.PartsHeader.Total

	}

	MarshalledJson, err := json.Marshal(abcic)

	return MarshalledJson, err
}



//*********************************************************************************************
//  json layout
//*********************************************************************************************
// {
//    "error": "",
//    "result": {
//        "height": "26682",
//        "hash": "75CA0F856A4DA078FC4911580360E70CEFB2EBEE",
//        "deliver_tx": {
//            "log": "",
//            "data": "",
//            "code": "0"
//        },
//        "check_tx": {
//            "log": "",
//            "data": "",
//            "code": "0"
//        }
//    },
//    "id": "",
//    "jsonrpc": "2.0"
//}
//*********************************************************************************************


type ABCIbctcJSON struct {
	Error  string `json:"error"`
	Result struct {
		Height    int64 `json:"height"`
		Hash      cmn.HexBytes `json:"hash"`
		DeliverTx struct {
			Log  string `json:"log"`
			Data []byte `json:"data"`
			Code uint32 `json:"code"`
		} `json:"deliver_tx"`
		CheckTx struct {
			Log  string `json:"log"`
			Data []byte `json:"data"`
			Code uint32 `json:"code"`
		} `json:"check_tx"`
	} `json:"result"`
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
}

//8   
func ClientBroadcastTxCommit(tx types.Tx) ([]byte, error) {

	var abcibctc ABCIbctcJSON
	//  as seen in Tendermint code
	abcibctc.Jsonrpc = "2.0"
	//abcibctc.ID ??

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	result, err := client.BroadcastTxCommit(tx)

	if err != nil {
		fmt.Println("Handle client.BroadcastTxCommit(tx) error: ", err, result)
		abcibctc.Error = err.Error()
	} else {

		abcibctc.Result.Height  	= result.Height
		abcibctc.Result.Hash  		= result.Hash

		abcibctc.Result.DeliverTx.Log 	= result.DeliverTx.Log
		abcibctc.Result.DeliverTx.Data 	= result.DeliverTx.Data
		abcibctc.Result.DeliverTx.Code 	= result.DeliverTx.Code

		abcibctc.Result.CheckTx.Log 	= result.CheckTx.Log
		abcibctc.Result.CheckTx.Data 	= result.CheckTx.Data
		abcibctc.Result.CheckTx.Code 	= result.CheckTx.Code

	}

	MarshalledJson, err := json.Marshal(abcibctc)

	return MarshalledJson, err
}




//*********************************************************************************************
//  json layout
//*********************************************************************************************
// {
// 	"jsonrpc": "2.0",
// 	"id": "",
// 	"result": {
// 		"code": 0,
// 		"data": "",
// 		"log": "",
// 		"hash": "0D33F2F03A5234F38706E43004489E061AC40A2E"
// 	},
// 	"error": ""
// }
//*********************************************************************************************
type ABCIbctsJSON struct {
	Jsonrpc 	string `json:"jsonrpc"`
	ID      	string `json:"id"`
	Result  	struct {
		Code 	uint32    `json:"code"`
		Data 	cmn.HexBytes `json:"data"`
		Log  	string `json:"log"`
		Hash 	cmn.HexBytes `json:"hash"`
	} `json:"result"`
	Error 	string `json:"error"`
}

//7    *** Just  SYNC ...  No A! *** 
func ClientBroadcastTxSync(tx types.Tx) ([]byte, error) {

	var abcibcts ABCIbctsJSON
	//  as seen in Tendermint code
	abcibcts.Jsonrpc = "2.0"
	//abcibcts.ID ??

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	result, err := client.BroadcastTxSync(tx)

	if err != nil {
		fmt.Println("Handle client.BroadcastTxSync(tx) error: ", err, result)
		abcibcts.Error = err.Error()
	} else {
		abcibcts.Result.Code  	= result.Code
		abcibcts.Result.Data  	= result.Data
		abcibcts.Result.Log  	= result.Log
		abcibcts.Result.Hash  	= result.Hash
	}

	MarshalledJson, err := json.Marshal(abcibcts)

	return MarshalledJson, err
}




//*********************************************************************************************
//  json layout
//*********************************************************************************************
// {
// 	"error": "",
// 	"result": {
// 		"hash": "E39AAB7A537ABAA237831742DCE1117F187C3C52",
// 		"log": "",
// 		"data": "",
// 		"code": 0
// 	},
// 	"id": "",
// 	"jsonrpc": "2.0"
// }
//*********************************************************************************************
type ABCIbctaJSON struct {
	Error  		string `json:"error"`
	Result 		struct {
		Hash 	cmn.HexBytes `json:"hash"`
		Log  	string `json:"log"`
		Data 	cmn.HexBytes `json:"data"`
		Code 	uint32    `json:"code"`
	} `json:"result"`
	ID      	string `json:"id"`
	Jsonrpc 	string `json:"jsonrpc"`
}	

//6  ***  Async...   A ... sync *** 
func ClientBroadcastTxAsync(tx types.Tx) ([]byte, error) {

	var abcibcta ABCIbctaJSON
	//  as seen in Tendermint code
	abcibcta.Jsonrpc = "2.0"
	//abcibcta.ID ??

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
 	result, err := client.BroadcastTxAsync(tx)

	if err != nil {
		fmt.Println("Handle client.BroadcastTxAsync(tx) error: ", err, result)
		abcibcta.Error = err.Error()
	} else {
		abcibcta.Result.Code  	= result.Code
		abcibcta.Result.Data  	= result.Data
		abcibcta.Result.Log  	= result.Log
		abcibcta.Result.Hash  	= result.Hash
	}

	MarshalledJson, err := json.Marshal(abcibcta)

	return MarshalledJson, err
}




//*********************************************************************************************
//  json layout
//*********************************************************************************************
// {
// 	"error": "",
// 	"result": {
// 		"block_metas": [
// 			{
// 				"header": {
// 					"app_hash": "",
// 					"chain_id": "test-chain-6UTNIN",
// 					"height": 10,
// 					"time": "2017-05-29T15:05:53.877Z",
// 					"num_txs": 0,
// 					"last_block_id": {
// 						"parts": {
// 							"hash": "3C78F00658E06744A88F24FF97A0A5011139F34A",
// 							"total": 1
// 						},
// 						"hash": "F70588DAB36BDA5A953D548A16F7D48C6C2DFD78"
// 					},
// 					"last_commit_hash": "F31CC4282E50B3F2A58D763D233D76F26D26CABE",
// 					"data_hash": "",
// 					"validators_hash": "9365FC80F234C967BD233F5A3E2AB2F1E4B0E5AA"
// 				},
// 				"block_id": {
// 					"parts": {
// 						"hash": "277A4DBEF91483A18B85F2F5677ABF9694DFA40F",
// 						"total": 1
// 					},
// 					"hash": "96B1D2F2D201BA4BC383EB8224139DB1294944E5"
// 				}
//             "type": 2,
//             "round": 0,
//             "height": 9,
//             "validator_index": 0,
//             "validator_address": "E89A51D60F68385E09E716D353373B11F8FACD62"
//           }
//         ],
//         "blockID": {
//           "parts": {
//             "hash": "3C78F00658E06744A88F24FF97A0A5011139F34A",
//             "total": 1
//           },
//           "hash": "F70588DAB36BDA5A953D548A16F7D48C6C2DFD78"
//         }
//       },
//       "header": {
//         "app_hash": "",
//         "chain_id": "test-chain-6UTNIN",
//         "height": 10,
//         "time": "2017-05-29T15:05:53.877Z",
//         "num_txs": 0,
//         "last_block_id": {
//           "parts": {
//             "hash": "3C78F00658E06744A88F24FF97A0A5011139F34A",
//             "total": 1
//           },
//           "hash": "F70588DAB36BDA5A953D548A16F7D48C6C2DFD78"
//         },
//         "last_commit_hash": "F31CC4282E50B3F2A58D763D233D76F26D26CABE",
//         "data_hash": "",
//         "validators_hash": "9365FC80F234C967BD233F5A3E2AB2F1E4B0E5AA"
//       }
//     },
// 		"last_height": 5493
// 	},
// 	"id": "",
// 	"jsonrpc": "2.0"
// }
//*********************************************************************************************
type ABCIbcinfoJSON struct {
	Error  					string `json:"error"`
	Result 					struct {
		BlockMetas 			[]struct {
			Header 			struct {
				AppHash     cmn.HexBytes    `json:"app_hash"`
				ChainID     string    `json:"chain_id"`
				Height      int64    `json:"height"`
				Time        time.Time `json:"time"`
				NumTxs      int64    `json:"num_txs"`
		//		LastBlockID types.BlockID `json:"last_block_id"`
				LastCommitHash cmn.HexBytes `json:"last_commit_hash"`
				DataHash       cmn.HexBytes `json:"data_hash"`
				ValidatorsHash cmn.HexBytes `json:"validators_hash"`
			} `json:"header"`
			BlockID 			types.BlockID `json:"block_id"`
		} `json:"block_metas"`
		LastHeight 				int64 `json:"last_height"`
	} `json:"result"`
	ID      					string `json:"id"`
	Jsonrpc 					string `json:"jsonrpc"`
}

type StHeader struct {
		AppHash     cmn.HexBytes    `json:"app_hash"`
		ChainID     string    `json:"chain_id"`
		Height      int64    `json:"height"`
		Time        time.Time `json:"time"`
		NumTxs      int64    `json:"num_txs"`
		LastBlockID types.BlockID `json:"last_block_id"`
		LastCommitHash cmn.HexBytes `json:"last_commit_hash"`
		DataHash       cmn.HexBytes `json:"data_hash"`
		ValidatorsHash cmn.HexBytes `json:"validators_hash"`
	}

type StBlockMeta struct {
	StHeader StHeader `json:"header"`
	BlockID types.BlockID `json:"block_id"`
}


//5
func ClientBlockChainInfo(minHeight, maxHeight int64) ([]byte, error) {

	var abcibci ABCIbcinfoJSON
	//  as seen in Tendermint code
	abcibci.Jsonrpc = "2.0"
	//abcibci.ID ??

	var stbm StBlockMeta
	var sth StHeader

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	info, err := client.BlockchainInfo(minHeight, maxHeight)

	if err != nil {
		fmt.Println("Handle client.BlockchainInfo(minHeight, maxHeight) error: ", err, info)
		abcibci.Error = err.Error()
	} else {


//	info = *core_types.ResultBlockchainInfo
// List of blocks

//type ResultBlockchainInfo struct {
//	LastHeight int64              `json:"last_height"`
//	BlockMetas []*types.BlockMeta `json:"block_metas"`
//}

// BlockMeta contains meta information about a block - namely, it's ID and Header.
//type BlockMeta struct {
//	BlockID BlockID `json:"block_id"` // the block hash and partsethash
//	Header  Header  `json:"header"`   // The block's Header
//}

// BlockID defines the unique ID of a block as its Hash and its PartSetHeader
//type BlockID struct {
//	Hash        cmn.HexBytes  `json:"hash"`
//	PartsHeader PartSetHeader `json:"parts"`
//}	

//type Header struct {
//	ChainID  string    `json:"chain_id"`
//	Height   int64     `json:"height"`
//	Time     time.Time `json:"time"`
//	NumTxs   int64     `json:"num_txs"`
//	TotalTxs int64     `json:"total_txs"`
//	LastBlockID BlockID `json:"last_block_id"`
//	LastCommitHash cmn.HexBytes `json:"last_commit_hash"` 
//	DataHash       cmn.HexBytes `json:"data_hash"`        
//	ValidatorsHash     cmn.HexBytes `json:"validators_hash"`      
//	NextValidatorsHash cmn.HexBytes `json:"next_validators_hash"` 
//	ConsensusHash      cmn.HexBytes `json:"consensus_hash"`       
//	AppHash            cmn.HexBytes `json:"app_hash"`             
//	LastResultsHash    cmn.HexBytes `json:"last_results_hash"`    
///	EvidenceHash    cmn.HexBytes `json:"evidence_hash"`    
//	ProposerAddress Address      `json:"proposer_address"` 
//}	

	// nope.  Need to build up Header and BlockID per occurence, and append THAT BlockMeta to BlockMetas



		for i := 0; i < len(info.BlockMetas); i++ {

// WAAAY more info according to header struct.

			sth.AppHash  			= info.BlockMetas[i].Header.AppHash
			sth.ChainID  			= info.BlockMetas[i].Header.ChainID
			sth.Height  			= info.BlockMetas[i].Header.Height
			sth.Time  				= info.BlockMetas[i].Header.Time
			sth.NumTxs  			= info.BlockMetas[i].Header.NumTxs
			sth.LastCommitHash  	= info.BlockMetas[i].Header.LastCommitHash
			sth.LastBlockID  		= info.BlockMetas[i].Header.LastBlockID
			sth.DataHash  			= info.BlockMetas[i].Header.DataHash
			sth.ValidatorsHash  	= info.BlockMetas[i].Header.ValidatorsHash
			
			stbm.StHeader 			= sth
			stbm.BlockID  			= info.BlockMetas[i].BlockID

			abcibci.Result.BlockMetas		= append(abcibci.Result.BlockMetas, stbm)

		}

		abcibci.Result.LastHeight 	= info.LastHeight
	}
	
	MarshalledJson, err := json.Marshal(abcibci)

	return  MarshalledJson, err
}


//*********************************************************************************************
//  json layout
//*********************************************************************************************
//{
//  "height": 10,
//  "results": [
//   {
//    "code": 0,
//    "data": "CAFE00F00D"
//   },
//   {
//    "code": 102,
//    "data": ""
//   }
//  ]
// }
//*********************************************************************************************
type ABCIBlockResultsJSON struct {
	Height  int64 `json:"height"`
	Results []struct {
		Code uint32   `json:"code"`
		Data []byte   `json:"data"`
	} `json:"results"`
}

type StResults struct {
	Code uint32    	`json:"code"`
	Data []byte  	`json:"data"`
} 

//4
func ClientBlockResults(heightPtr *int64)  ([]byte, error) {

	var abciresults ABCIBlockResultsJSON

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	info, err := client.BlockResults(heightPtr)

	if err != nil {
		fmt.Println("Handle client.BlockResults(heightPtr) error: ", err, info)
	//	abciresults.Error = err.Error()
	} else {

		var sr StResults

		abciresults.Height = info.Height
			
		var i int64 = 0
		for ; i < info.Height; i++ {  
			sr.Code = info.Results.DeliverTx[i].Code        
			sr.Data = info.Results.DeliverTx[i].Data   
			abciresults.Results = append(abciresults.Results, sr)      
		}
	}	

	MarshalledJson, err := json.Marshal(abciresults)

	return MarshalledJson, err
}



//*********************************************************************************************
//  json layout
//*********************************************************************************************
// {
//   "error": "",
//   "result": {
//     "block": {
//       "last_commit": {
//         "precommits": [
//           {
//             "signature": {
//               "data": "12C0D8893B8A38224488DC1DE6270DF76BB1A5E9DB1C68577706A6A97C6EC34FFD12339183D5CA8BC2F46148773823DE905B7F6F5862FD564038BB7AE03BF50D",
//               "type": "ed25519"
//             },
//             "block_id": {
//               "parts": {
//                 "hash": "3C78F00658E06744A88F24FF97A0A5011139F34A",
//                 "total": 1
//               },
//               "hash": "F70588DAB36BDA5A953D548A16F7D48C6C2DFD78"
//             },
//             "type": 2,
//             "round": 0,
//             "height": 9,
//             "validator_index": 0,
//             "validator_address": "E89A51D60F68385E09E716D353373B11F8FACD62"
//           }
//         ],
//         "blockID": {
//           "parts": {
//             "hash": "3C78F00658E06744A88F24FF97A0A5011139F34A",
//             "total": 1
//           },
//           "hash": "F70588DAB36BDA5A953D548A16F7D48C6C2DFD78"
//         }
//       },
//       "data": {
//         "txs": []
//       },
//       "header": {
//         "app_hash": "",
//         "chain_id": "test-chain-6UTNIN",
//         "height": 10,
//         "time": "2017-05-29T15:05:53.877Z",
//         "num_txs": 0,
//         "last_block_id": {
//           "parts": {
//             "hash": "3C78F00658E06744A88F24FF97A0A5011139F34A",
//             "total": 1
//           },
//           "hash": "F70588DAB36BDA5A953D548A16F7D48C6C2DFD78"
//         },
//         "last_commit_hash": "F31CC4282E50B3F2A58D763D233D76F26D26CABE",
//         "data_hash": "",
//         "validators_hash": "9365FC80F234C967BD233F5A3E2AB2F1E4B0E5AA"
//       }
//     },
//     "block_meta": {
//       "header": {
//         "app_hash": "",
//         "chain_id": "test-chain-6UTNIN",
//         "height": 10,
//         "time": "2017-05-29T15:05:53.877Z",
//         "num_txs": 0,
//         "last_block_id": {
//           "parts": {
//             "hash": "3C78F00658E06744A88F24FF97A0A5011139F34A",
//             "total": 1
//           },
//           "hash": "F70588DAB36BDA5A953D548A16F7D48C6C2DFD78"
//         },
//         "last_commit_hash": "F31CC4282E50B3F2A58D763D233D76F26D26CABE",
//         "data_hash": "",
//         "validators_hash": "9365FC80F234C967BD233F5A3E2AB2F1E4B0E5AA"
//       },
//       "block_id": {
//         "parts": {
//           "hash": "277A4DBEF91483A18B85F2F5677ABF9694DFA40F",
//           "total": 1
//         },
//         "hash": "96B1D2F2D201BA4BC383EB8224139DB1294944E5"
//       }
//     }
//   },
//   "id": "",
//   "jsonrpc": "2.0"
// }
//*********************************************************************************************
type ABCIBlockAtHeightJSON struct {
	Error  						string `json:"error"`
	Result 						struct {

		Block 					struct {

			LastCommit 			struct {

				Precommits 		[]struct {

					Signature 	[]byte `json:"signature"`

					BlockID 	types.BlockID `json:"block_id"`

					Type             byte `json:"type"`
					Round            int `json:"round"`
					Height           int64 `json:"height"`
					ValidatorIndex   int `json:"validator_index"`
					ValidatorAddress cmn.HexBytes `json:"validator_address"`
				} `json:"precommits"`

				BlockID 		types.BlockID `json:"blockID"`

			} `json:"last_commit"`

			Data 				struct {
				Txs 			[]byte `json:"txs"`
			} `json:"data"`

			Header 			struct {

				AppHash     cmn.HexBytes    `json:"app_hash"`
				ChainID     string    `json:"chain_id"`
				Height      int64    `json:"height"`
				Time        time.Time `json:"time"`
				NumTxs      int64    `json:"num_txs"`
				LastBlockID 	types.BlockID `json:"last_block_id"`
				LastCommitHash 	cmn.HexBytes `json:"last_commit_hash"`
				DataHash       	cmn.HexBytes `json:"data_hash"`
				ValidatorsHash 	cmn.HexBytes `json:"validators_hash"`
			} `json:"header"`

		} `json:"block"`

		BlockMeta 			struct {

			Header 			struct {

				AppHash     cmn.HexBytes    `json:"app_hash"`
				ChainID     string    `json:"chain_id"`
				Height      int64    `json:"height"`
				Time        time.Time `json:"time"`
				NumTxs      int64    `json:"num_txs"`
				LastBlockID types.BlockID `json:"last_block_id"`
				LastCommitHash cmn.HexBytes `json:"last_commit_hash"`
				DataHash       cmn.HexBytes `json:"data_hash"`
				ValidatorsHash cmn.HexBytes `json:"validators_hash"`
			} `json:"header"`
			BlockID 		types.BlockID `json:"block_id"`
		} `json:"block_meta"`
	} `json:"result"`

	ID      				string `json:"id"`
	Jsonrpc 				string `json:"jsonrpc"`
}


type StPrecommits 	struct {

	Signature 		[]byte `json:"signature"`

	BlockID 		types.BlockID `json:"block_id"`
	Type             byte `json:"type"`

	Round            int `json:"round"`
	Height           int64 `json:"height"`
	ValidatorIndex   int `json:"validator_index"`
	ValidatorAddress cmn.HexBytes `json:"validator_address"`
} 

// Get block at a given height. If no height is provided, it will fetch the latest block.
//3
func ClientBlockAtHeight(Height *int64) ([]byte, error) {

	var abcibah ABCIBlockAtHeightJSON
	//  as seen in Tendermint code
	abcibah.Jsonrpc = "2.0"
	//abcibah.ID ??

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
 	info, err := client.Block(Height)

	if err != nil {
		fmt.Println("Handle client.Block(Height) error: ", err)
		abcibah.Error = err.Error()
	} else {

		var rblp 	StPrecommits
		for i := 0; i < len(info.Block.LastCommit.Precommits); i ++ {

			rblp.Signature 				= 	info.Block.LastCommit.Precommits[i].Signature
			rblp.BlockID 				= 	info.Block.LastCommit.Precommits[i].BlockID
			rblp.BlockID.Hash 			= 	info.Block.LastCommit.Precommits[i].BlockID.Hash
			rblp.Type 					= 	info.Block.LastCommit.Precommits[i].Type
			rblp.Round 					= 	info.Block.LastCommit.Precommits[i].Round
			rblp.Height 				=	info.Block.LastCommit.Precommits[i].Height
			rblp.ValidatorIndex 		= 	info.Block.LastCommit.Precommits[i].ValidatorIndex 
			rblp.ValidatorAddress 		= 	info.Block.LastCommit.Precommits[i].ValidatorAddress

			abcibah.Result.Block.LastCommit.Precommits	= append(abcibah.Result.Block.LastCommit.Precommits, rblp)
		}

	// -- Block  --  Data
		//  ??????????????????????????????????????????????????????
		//
		//abcibah.Result.Block.Data.Txs	=   info.Block.Data.Txs

	// -- Block  -- Header

		abcibah.Result.Block.Header.AppHash  		= info.Block.Header.AppHash
		abcibah.Result.Block.Header.ChainID  		= info.Block.Header.ChainID
		abcibah.Result.Block.Header.Height  		= info.Block.Header.Height
		abcibah.Result.Block.Header.Time  			= info.Block.Header.Time
		abcibah.Result.Block.Header.NumTxs  		= info.Block.Header.NumTxs
		abcibah.Result.Block.Header.LastBlockID  	= info.Block.Header.LastBlockID
		abcibah.Result.Block.Header.LastCommitHash  = info.Block.Header.LastCommitHash
		abcibah.Result.Block.Header.DataHash  		= info.Block.Header.DataHash
		abcibah.Result.Block.Header.ValidatorsHash  = info.Block.Header.ValidatorsHash

	// -- BlockMeta
	// -- BlockMeta  -- Header

		abcibah.Result.BlockMeta.Header.AppHash 		= info.BlockMeta.Header.AppHash 
		abcibah.Result.BlockMeta.Header.ChainID 		= info.BlockMeta.Header.ChainID 
		abcibah.Result.BlockMeta.Header.Height 			= info.BlockMeta.Header.Height 
		abcibah.Result.BlockMeta.Header.Time 			= info.BlockMeta.Header.Time 
		abcibah.Result.BlockMeta.Header.LastBlockID 	= info.BlockMeta.Header.LastBlockID 
		abcibah.Result.BlockMeta.Header.LastBlockID.Hash 		= info.BlockMeta.Header.LastBlockID.Hash 
		abcibah.Result.BlockMeta.Header.LastCommitHash 	= info.BlockMeta.Header.LastCommitHash 
		abcibah.Result.BlockMeta.Header.DataHash 		= info.BlockMeta.Header.DataHash 		
		abcibah.Result.BlockMeta.Header.ValidatorsHash 	= info.BlockMeta.Header.ValidatorsHash

	// -- BlockMeta  -- BlockID

		abcibah.Result.BlockMeta.BlockID 				= info.BlockMeta.BlockID	
		abcibah.Result.BlockMeta.BlockID.Hash 			= info.BlockMeta.BlockID.Hash 		

	}

	MarshalledJson, err := json.Marshal(abcibah)

	return MarshalledJson, err
}


//*********************************************************************************************
//  json layout
//*********************************************************************************************
// {
// 	"error": "",
// 	"result": {
// 		"response": {
// 			"log": "exists",
// 			"height": 0,
// 			"proof": "010114FED0DAD959F36091AD761C922ABA3CBF1D8349990101020103011406AA2262E2F448242DF2C2607C3CDC705313EE3B0001149D16177BC71E445476174622EA559715C293740C",
// 			"value": "61626364",
// 			"key": "61626364",
// 			"index": -1,
// 			"code": 0
// 		}
// 	},
// 	"id": "",
// 	"jsonrpc": "2.0"
// }
//*********************************************************************************************
type ABCIQueryJSON struct {
	Error 	string  `json:"error"`
	Result  struct  {
		Response 	struct {
			Log 	string `json:"log"`
			Height 	int64 `json:"height"`
			Proof 	[]byte `json:"proof"`
			Value 	[]byte `json:"value"`
			Key 	[]byte `json:"key"`
			Index 	int64 `json:"index"`
			Code 	uint32 `json:"code"`
		} `json:"response"` 
	} `json:"result"`
	ID 		string `json:"id"`
	JsonRPC	string `json:"jsonrpc"`
}
// ### Query Parameters
//*********************************************************************************************
// | Parameter | Type   | Default | Required | Description                                    |
// |-----------+--------+---------+----------+------------------------------------------------|
// | path      | string | false   | false    | Path to the data ("/a/b/c")                    |
// | data      | []byte | false   | true     | Data                                           |
// | height    | int64  | 0       | false    | Height (0 means latest)                        |
// | trusted   | bool   | false   | false    | Does not include a proof of the data inclusion |
//*********************************************************************************************
//2
func ClientABCIQuery(path string, data cmn.HexBytes, height int64, trusted bool) ([]byte, error) {

	var abciquery ABCIQueryJSON
	//  as seen in Tendermint code
	abciquery.JsonRPC = "2.0"
	//abciquery.ID ??

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")

// ***also different: https://github.com/tendermint/tendermint/blob/master/rpc/core/abci.go#L89
//	func ABCIQuery(path string, data cmn.HexBytes, height int64, trusted bool)
// see line 52	result, err := client.ABCIQuery(path, data, height, trusted)
// see Line 20   result, err := client.ABCIQuery("", "abcd", true)

//func (c *HTTP) ABCIQuery(path string, data cmn.HexBytes) (*ctypes.ResultABCIQuery, error) {
//	return c.ABCIQueryWithOptions(path, data, DefaultABCIQueryOptions)
//} 

	result, err := client.ABCIQuery(path, data)

	if err != nil {
		fmt.Println("Handle client.ABCIQuery(path, data, height, trusted) error: ", err, result)
		abciquery.Error = err.Error()
	} else {
		abciquery.Result.Response.Log 			= result.Response.Log
		abciquery.Result.Response.Height 		= result.Response.Height
		abciquery.Result.Response.Proof 		= result.Response.Proof
		abciquery.Result.Response.Value 		= result.Response.Value
		abciquery.Result.Response.Key 			= result.Response.Key
		abciquery.Result.Response.Index 		= result.Response.Index
		abciquery.Result.Response.Code 			= result.Response.Code
	}

	MarshalledJson, err := json.Marshal(abciquery)

	return MarshalledJson, err
}


//*********************************************************************************************
//  json layout
//*********************************************************************************************
// {
// 	"error": "",
// 	"result": {
// 		"response": {
// 			"data": "{\"size\":3}"     ---------------------->>>>> Is this how its returned? slice maybe?
// 		}
// 	},
// 	"id": "",
// 	"jsonrpc": "2.0"
// }
//*********************************************************************************************
type ABCIInfoJSON struct {
	Error  string `json:"error"`
	Result struct {
		Response struct {
			Data string `json:"data"`
		} `json:"response"`
	} `json:"result"`
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
}

//1
func ClientABCIInfo() ([]byte, error) {


	var abciinfo ABCIInfoJSON
	//  as seen in Tendermint code
	abciinfo.Jsonrpc = "2.0"
	//abciinfo.ID ??

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	info, err := client.ABCIInfo()

	if err != nil {
		fmt.Println("Handle client.ABCIInfo() error: ", err, info)
		abciinfo.Error = err.Error()
	} else {
		abciinfo.Result.Response.Data  = info.Response.Data
	}

	MarshalledJson, err := json.Marshal(abciinfo)

	return MarshalledJson, err
}