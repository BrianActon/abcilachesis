package abcilachesis

import (
	"context"
//	"encoding/json"
	"fmt"
	"time"
	
	"github.com/tendermint/tendermint/rpc/core"
	"github.com/tendermint/tendermint/rpc/client"
	cmn "github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/libs/pubsub/query"
	tmpubsub "github.com/tendermint/tendermint/libs/pubsub"
//	tmquery "github.com/tendermint/tendermint/libs/pubsub/query"
	rpctypes "github.com/tendermint/tendermint/rpc/lib/types"
	"github.com/tendermint/tendermint/types"
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
//26		func ClientUnconfirmedTxs(limit int) error {
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
		Txs 	[]string `json:"txs"`
		N_Txs  	int `json:"n_txs"`
	} `json:"result"`
	ID      	string `json:"id"`
	Jsonrpc 	string `json:"jsonrpc"`
}	

//26
func CoreUnconfirmedTxs(limit int) error {

//	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
 	result, err := core.UnconfirmedTxs(limit)

	if err != nil {
		fmt.Println("Handle client.UnconfirmedTxs(limit) error: ", err, result)
		return err
	} 

//	var abciut ABCIUnconfirmedTxsJSON

//	if err := json.Unmarshal(result, &abciut); err != nil {
//		return err
//	}

	return nil
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
			Proof 		struct {
				Proof 	struct {
					Aunts []string `json:"aunts"`
				} `json:"Proof"`
				Data     string `json:"Data"`
				RootHash string `json:"RootHash"`
				Total    string `json:"Total"`
				Index    string `json:"Index"`
			} `json:"proof"`
			Tx       	string `json:"tx"`
			TxResult 	struct {
			} `json:"tx_result"`
			Index  		string `json:"index"`
			Height 		string `json:"height"`
			Hash   		string `json:"hash"`
		} `json:"txs"`
		TotalCount 		string `json:"total_count"`
	} `json:"result"`
}

//25
func ClientTxSearch(query string) error {

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
//	q, err := tmquery.New("account.owner='Ivan'")
//	q, err := tmquery.New(query)
	tx, err := client.TxSearch(query, true, 1, 1)

	if err != nil {
		fmt.Println("Handle client.TxSearch(q, true) error: ", err, tx)
		return err
	} 

//	var abcits ABCITxSearchJSON

//	if err := json.Unmarshal(tx, &abcits); err != nil {
//		return err
//	}
	return nil
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
		Proof 			struct {
			Proof 		struct {
				Aunts 	[]interface{} `json:"aunts"`
			} `json:"Proof"`
			Data     	string `json:"Data"`
			RootHash 	string `json:"RootHash"`
			Total    	string `json:"Total"`
			Index    	string `json:"Index"`
		} `json:"proof"`
		Tx       		string `json:"tx"`
		TxResult 		struct {
			Log  		string `json:"log"`
			Data 		string `json:"data"`
			Code 		string `json:"code"`
		} `json:"tx_result"`
		Index  			string `json:"index"`
		Height 			string `json:"height"`
		Hash   			string `json:"hash"`
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
func ClienTx(hash []byte, prove bool) error {
	// Tx allows you to query the transaction results. `nil` could mean the
	// transaction is in the mempool, invalidated, or was not sent in the first
	// place.
	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	tx, err := client.Tx([]byte(hash), prove)

	if err != nil {
		fmt.Println("Handle client.Tx([]byte(hash), prove) error: ", err, tx)
		return err
	} 

//	var abcit ABCITxJSON

//	if err := json.Unmarshal(tx, &abcit); err != nil {
//		return err
//	}
	return nil
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
func ClientSubscribe() error {

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	ctx, cancel := context.WithTimeout(context.Background(), 1) // timeout)

	defer cancel()

	// TODO:  follow query.MustParse to confirm value to pass  --> MOre examples in comments above "json layout"

	query := query.MustParse("tm.event = 'Tx' AND tx.height = 3")
	txs := make(chan interface{})

	//  test-client?? txs?? 
	//  func Subscribe(wsCtx rpctypes.WSRPCContext, query string) but supplied code below has 4 items.

//	result, err := client.Subscribe(ctx,  query)
	err := client.Subscribe(ctx, "test-client", query, txs)

	if err != nil {
		fmt.Println("err= ", err)
	}
	
//	var abcis ABCISubscribeJSON

//	if err := json.Unmarshal(result, &abcis); err != nil {
//		return err
//	}  

	// txs is a channel!!!!
	// Nowhere in func Subscribe is any data being passed to txs to come out this side. Defaq? 

	go func() {
	    for e := range txs {
	        fmt.Println("got ", e.(types.EventDataTx))
		}
	}()
	return nil
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
		NodeInfo 				struct {
			ID         			string `json:"id"`
			ListenAddr 			string `json:"listen_addr"`
			Network    			string `json:"network"`
			Version   		 	string `json:"version"`
			Channels   			string `json:"channels"`
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
			LatestBlockHash   	string    `json:"latest_block_hash"`
			LatestAppHash     	string    `json:"latest_app_hash"`
			LatestBlockHeight 	string    `json:"latest_block_height"`
			LatestBlockTime   	time.Time `json:"latest_block_time"`
			CatchingUp        	bool      `json:"catching_up"`
		} `json:"sync_info"`
		ValidatorInfo 			struct {
			Address 			string `json:"address"`
			PubKey  			struct {
				Type  			string `json:"type"`
				Value 			string `json:"value"`
			} `json:"pub_key"`
			VotingPower 		string `json:"voting_power"`
		} `json:"validator_info"`
	} `json:"result"`
}

//22
func ClientStatus() error {

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	result, err := client.Status()

	if err != nil {
		fmt.Println("Handle client.Status() error: ", err, result)
		return err
	} 

//	var abcis ABCIStatusJSON

//	if err := json.Unmarshal(result, &abcis); err != nil {
//		return err
//	}

	return nil
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
		Validators 		[]struct {
			Accum       int `json:"accum"`
			VotingPower int `json:"voting_power"`
			PubKey      struct {
				Data 	string `json:"data"`
				Type 	string `json:"type"`
			} `json:"pub_key"`
			Address 	string `json:"address"`
		} `json:"validators"`
		BlockHeight 	int `json:"block_height"`
	} `json:"result"`
	ID      			string `json:"id"`
	Jsonrpc 			string `json:"jsonrpc"`
}	

//21
func ClientValidators(heightPtr *int64) error {

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	state, err := client.Validators(heightPtr)

	if err != nil {
		fmt.Println("Handle client.Validators() error: ", err, state)
		return err
	} 

//	var abciv ABCIValidatorsJSON

//	if err := json.Unmarshal(state, &abciv); err != nil {
//		return err
//	}

	return nil
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
func ClientUnsubscribeAll(wsCtx context.Context,  subscriber string) error {

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	err := client.UnsubscribeAll(wsCtx, subscriber)
//result, 
	if err != nil {
		fmt.Println("Handle client.UnsubscribeAll(...) error: ", err)
		return err
	} 

//	var abciusa ABCIUnsubscribeAllJSON

//	if err := json.Unmarshal(result, &abciusa); err != nil {
//		return err
//	}
//
	return nil
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
func ClientUnsubscribe(wsCtx context.Context, query tmpubsub.Query) error {

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	err := client.Unsubscribe(wsCtx, "test-client", query)

	if err != nil {
		fmt.Println("Handle client.Unsubscribe(...) error: ", err)
		return err
	} 


 //  *****  client.Unsubscribe  only return error. 
//   *****  Documentation is wrong... AGAIN!!!!!
	//
//	var abcius ABCIUnsubscribeJSON
//
//	if err := json.Unmarshal(result, &abcius); err != nil {
//		return err
//	}
//
	return nil
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
		Txs 	string `json:"txs"`
		N_Txs  	int `json:"n_txs"`
	} `json:"result"`
	ID      	string `json:"id"`
	Jsonrpc 	string `json:"jsonrpc"`
}	

//16
func ClientNumUnconfirmedTxs() error {

//	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
 	result, err := core.UnconfirmedTxs(10)

	if err != nil {
		fmt.Println("Handle client.UnconfirmedTxs() error: ", err, result)
		return err
	} 

//	var abcinut ABCINumUnconfirmedTxsJSON
//
//	if err := json.Unmarshal(result, &abcinut); err != nil {
//		return err
//	}
	return nil
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
		Peers     	[]interface{} `json:"peers"`
		Listeners 	[]string      `json:"listeners"`
		Listening 	bool          `json:"listening"`
	} `json:"result"`
	ID      		string `json:"id"`
	Jsonrpc 		string `json:"jsonrpc"`
}

//15
func ClientNetInfo() error {

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
 	info, err := client.NetInfo()

	if err != nil {
		fmt.Println("Handle client.NetInfo() error: ", err, info)
		return err
	} 

//	var abcini ABCINetInfoJSON

//	if err := json.Unmarshal(info, &abcini); err != nil {
//		return err
//	}
	return nil
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
// Get node health. Returns empty result (200 OK) on success, no response - in
// case of an error.
func ClientHealth(wsCtx rpctypes.WSRPCContext) error {

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	result, err := client.Health()

	if err != nil {
		fmt.Println("Handle client.UnsubscribeAll(...) error: ", err, result)
		return err
	} 

//	var abcih ABCIHealthJSON
//
//	if err := json.Unmarshal(result, &abcih); err != nil {
//		return err
//	}

	return nil
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
			AppHash    		string `json:"app_hash"`
			Validators 		[]struct {
				Name   		string `json:"name"`
				Power  		int    `json:"power"`
				PubKey 		struct {
					Data 	string `json:"data"`
					Type 	string `json:"type"`
				} `json:"pub_key"`
			} `json:"validators"`
			ChainID     	string    `json:"chain_id"`
			GenesisTime 	time.Time `json:"genesis_time"`
		} `json:"genesis"`
	} `json:"result"`
	ID      				string `json:"id"`
	Jsonrpc 				string `json:"jsonrpc"`
}

//13
func ClientGenesis() error {

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	genesis, err := client.Genesis()

	if err != nil {
		fmt.Println("Handle client.UnconfirmedTxs() error: ", err, genesis)
		return err
	} 

//	var abcig ABCIGenesisJSON
//
//	if err := json.Unmarshal(genesis, &abcig); err != nil {
//		return err
//	}
	return nil
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
func ClientDumpConsensusState () error {

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
    state, err := client.DumpConsensusState()

	if err != nil {
		fmt.Println("Handle client.UnconfirmedTxs() error: ", err, state)
		return err
	} 

//	var abcidcs ABCIDumpConsensusStateJSON
//
//	if err := json.Unmarshal(state, &abcidcs); err != nil {
//		return err
//	}
	return nil
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
		RoundState 					struct {
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
		} `json:"round_state"`
	} `json:"result"`
}

//11
func ClientConsensusState() error {

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	state, err := client.ConsensusState()

	if err != nil {
		fmt.Println("Handle client.BlockResults(heightPtr) error: ", err, state)
		return err
	} 

//	var abcics ABCIConsensusStateJSON
//
//	if err := json.Unmarshal(state, &abcics); err != nil {
//		return err
//	}
	return nil
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
		BlockHeight     	string `json:"block_height"`
		ConsensusParams 	struct {
			BlockSizeParams struct {
				MaxTxsBytes string `json:"max_txs_bytes"`
				MaxGas      string `json:"max_gas"`
			} `json:"block_size_params"`
			EvidenceParams 	struct {
				MaxAge 		string `json:"max_age"`
			} `json:"evidence_params"`
		} `json:"consensus_params"`
	} `json:"result"`
}

//10
func ClientConsensusParams(heightPtr *int64) error {

//	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	state, err := core.ConsensusParams(heightPtr)

	if err != nil {
		fmt.Println("Handle client.BlockResults(heightPtr) error: ", err, state)
		return err
	} 

//	var abcicp ABCIConsensusParamsJSON
//
//	if err := json.Unmarshal(state, &abcicp); err != nil {
//		return err
//	}
	return nil
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
				Signature 		struct {
					Data 		string `json:"data"`
					Type 		string `json:"type"`
				} `json:"signature"`
				BlockID 		struct {
					Parts 		struct {
						Hash  	string `json:"hash"`
						Total 	int    `json:"total"`
					} `json:"parts"`
					Hash string `json:"hash"`
				} `json:"block_id"`
				Type             int    `json:"type"`
				Round            int    `json:"round"`
				Height           int    `json:"height"`
				ValidatorIndex   int    `json:"validator_index"`
				ValidatorAddress string `json:"validator_address"`
			} `json:"precommits"`
			BlockID 			struct {
				Parts 			struct {
					Hash  		string `json:"hash"`
					Total 		int    `json:"total"`
				} `json:"parts"`
				Hash 			string `json:"hash"`
			} `json:"blockID"`
		} `json:"commit"`
		Header 					struct {
			AppHash     		string    `json:"app_hash"`
			ChainID     		string    `json:"chain_id"`
			Height      		int       `json:"height"`
			Time        		time.Time `json:"time"`
			NumTxs      		int       `json:"num_txs"`
			LastBlockID 		struct {
				Parts 			struct {
					Hash  		string `json:"hash"`
					Total 		int    `json:"total"`
				} `json:"parts"`
				Hash 			string `json:"hash"`
			} `json:"last_block_id"`
			LastCommitHash 		string `json:"last_commit_hash"`
			DataHash       		string `json:"data_hash"`
			ValidatorsHash 		string `json:"validators_hash"`
		} `json:"header"`
	} `json:"result"`
	ID      					string `json:"id"`
	Jsonrpc 					string `json:"jsonrpc"`
}

//9
func ClientCommit(heightPtr *int64) error {

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	info, err := client.BlockResults(heightPtr)

	if err != nil {
		fmt.Println("Handle client.BlockResults(heightPtr) error: ", err, info)
		return err
	} 

//	var abciresults ABCIBlockResultsJSON
//
//	if err := json.Unmarshal(info, &abciresults); err != nil {
//		return err
//	}
	return nil
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
type ABCIbctcJSON struct {
	Error  		string `json:"error"`
	Result 		struct {
		Hash 	string `json:"hash"`
		Log  	string `json:"log"`
		Data 	string `json:"data"`
		Code 	int    `json:"code"`
	} `json:"result"`
	ID      	string `json:"id"`
	Jsonrpc 	string `json:"jsonrpc"`
}	

//8   
func ClientBroadcastTxCommit(tx types.Tx) error {

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	result, err := client.BroadcastTxCommit(tx)

	if err != nil {
		fmt.Println("Handle client.BroadcastTxCommit(tx) error: ", err, result)
		return err
	} 

//	var abcibctc ABCIbctcJSON
//
//	if err := json.Unmarshal(result, &abcibctc); err != nil {
//		return err
//	}
	return nil
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
		Code 	int    `json:"code"`
		Data 	string `json:"data"`
		Log  	string `json:"log"`
		Hash 	string `json:"hash"`
	} `json:"result"`
	Error 	string `json:"error"`
}

//7    *** Just  SYNC ...  No A! *** 
func ClientBroadcastTxSync(tx types.Tx) error {

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	result, err := client.BroadcastTxSync(tx)

	if err != nil {
		fmt.Println("Handle client.BroadcastTxSync(tx) error: ", err, result)
		return err
	} 

//	var abcibcts ABCIbctsJSON
//
//	if err := json.Unmarshal(result, &abcibcts); err != nil {
//		return err
//	}
	return nil
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
		Hash 	string `json:"hash"`
		Log  	string `json:"log"`
		Data 	string `json:"data"`
		Code 	int    `json:"code"`
	} `json:"result"`
	ID      	string `json:"id"`
	Jsonrpc 	string `json:"jsonrpc"`
}	

//6  ***  Async...   A ... sync *** 
func ClientBroadcastTxAsync(tx types.Tx) error {

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
 	result, err := client.BroadcastTxAsync(tx)

	if err != nil {
		fmt.Println("Handle client.BroadcastTxAsync(tx) error: ", err, result)
		return err
	} 

//	var abcibcta ABCIbctaJSON
//
//	if err := json.Unmarshal(result, &abcibcta); err != nil {
//		return err	
//	}
	return nil
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
				AppHash     string    `json:"app_hash"`
				ChainID     string    `json:"chain_id"`
				Height      string    `json:"height"`
				Time        time.Time `json:"time"`
				NumTxs      string    `json:"num_txs"`
				LastBlockID struct {
					Parts 	struct {
						Hash  string `json:"hash"`
						Total string `json:"total"`
					} `json:"parts"`
					Hash 		string `json:"hash"`
				} `json:"last_block_id"`
				LastCommitHash string `json:"last_commit_hash"`
				DataHash       string `json:"data_hash"`
				ValidatorsHash string `json:"validators_hash"`
			} `json:"header"`
			BlockID 			struct {
				Parts 			struct {
					Hash  		string `json:"hash"`
					Total 		string `json:"total"`
				} `json:"parts"`
				Hash 			string `json:"hash"`
			} `json:"block_id"`
		} `json:"block_metas"`
		LastHeight 				string `json:"last_height"`
	} `json:"result"`
	ID      					string `json:"id"`
	Jsonrpc 					string `json:"jsonrpc"`
}

//5
func ClientBlockChainInfo(minHeight, maxHeight int64) error {

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	info, err := client.BlockchainInfo(minHeight, maxHeight)

	if err != nil {
		fmt.Println("Handle client.BlockchainInfo(minHeight, maxHeight) error: ", err, info)
		return err
	} 

//	var abciquery ABCIQueryJSON
//
//	if err := json.Unmarshal(info, &abciquery); err != nil {
//		return err
//	}
	return nil
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
	Height  string `json:"height"`
	Results []struct {
		Code string `json:"code"`
		Data string `json:"data"`
	} `json:"results"`
}

//4
func ClientBlockResult(heightPtr *int64) error {

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	info, err := client.BlockResults(heightPtr)

	if err != nil {
		fmt.Println("Handle client.BlockResults(heightPtr) error: ", err, info)
		return err
	} 

//	var abciresults ABCIBlockResultsJSON
//
//	if err := json.Unmarshal(info, &abciresults); err != nil {
//		return err
//	}
	return nil
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
					Signature 	struct {
						Data 	string `json:"data"`
						Type 	string `json:"type"`
					} `json:"signature"`
					BlockID 	struct {
						Parts 	struct {
							Hash  string `json:"hash"`
							Total string `json:"total"`
						} `json:"parts"`
						Hash 	string `json:"hash"`
					} `json:"block_id"`
					Type             string `json:"type"`
					Round            string `json:"round"`
					Height           string `json:"height"`
					ValidatorIndex   string `json:"validator_index"`
					ValidatorAddress string `json:"validator_address"`
				} `json:"precommits"`
				BlockID 		struct {
					Parts 		struct {
						Hash  	string `json:"hash"`
						Total 	string `json:"total"`
					} `json:"parts"`
					Hash 		string `json:"hash"`
				} `json:"blockID"`
			} `json:"last_commit"`
			Data 				struct {
				Txs 			[]interface{} `json:"txs"`
			} `json:"data"`
			Header 			struct {
				AppHash     string    `json:"app_hash"`
				ChainID     string    `json:"chain_id"`
				Height      string    `json:"height"`
				Time        time.Time `json:"time"`
				NumTxs      string    `json:"num_txs"`
				LastBlockID struct {
					Parts 	struct {
						Hash  string `json:"hash"`
						Total string `json:"total"`
					} `json:"parts"`
					Hash 	string `json:"hash"`
				} `json:"last_block_id"`
				LastCommitHash string `json:"last_commit_hash"`
				DataHash       string `json:"data_hash"`
				ValidatorsHash string `json:"validators_hash"`
			} `json:"header"`
		} `json:"block"`
		BlockMeta 			struct {
			Header 			struct {
				AppHash     string    `json:"app_hash"`
				ChainID     string    `json:"chain_id"`
				Height      string    `json:"height"`
				Time        time.Time `json:"time"`
				NumTxs      string    `json:"num_txs"`
				LastBlockID struct {
					Parts 	struct {
						Hash  string `json:"hash"`
						Total string `json:"total"`
					} `json:"parts"`
					Hash 	string `json:"hash"`
				} `json:"last_block_id"`
				LastCommitHash string `json:"last_commit_hash"`
				DataHash       string `json:"data_hash"`
				ValidatorsHash string `json:"validators_hash"`
			} `json:"header"`
			BlockID 		struct {
				Parts 		struct {
					Hash  	string `json:"hash"`
					Total 	string `json:"total"`
				} `json:"parts"`
				Hash 		string `json:"hash"`
			} `json:"block_id"`
		} `json:"block_meta"`
	} `json:"result"`
	ID      				string `json:"id"`
	Jsonrpc 				string `json:"jsonrpc"`
}

// Get block at a given height.
// If no height is provided, it will fetch the latest block.
//3
func ClientBlockAtHeight(Height *int64) error {

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
 	info, err := client.Block(Height)

	if err != nil {
		fmt.Println("Handle client.Block(Height) error: ", err, info)
		return err
	} 

//	var abcibah ABCIBlockAtHeightJSON
//
//	if err := json.Unmarshal(info, &abcibah); err != nil {
//		return err
//	}
	return nil
}



//*********************************************************************************************
//*********************************************************************************************
//*********************************************************************************************
//*********************************************************************************************
//*********************************************************************************************
//*********************************************************************************************
//*********************************************************************************************
//*********************************************************************************************


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
	Error 	error  `json:"error"`
	Result  struct  {
		Response 	struct {
			Log 	string `json:"log"`
			Height 	string `json:"height"`
			Proof 	string `json:"proof"`
			Value 	string `json:"value"`
			Key 	string `json:"key"`
			Index 	int `json:"index"`
			Code 	int `json:"code"`
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
func ClientABCIQuery(path string, data cmn.HexBytes, height int64, trusted bool) error {

	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
//	result, err := client.ABCIQuery(path, data, height, trusted)
	result, err := client.ABCIQuery(path, data)

	if err != nil {
		fmt.Println("Handle client.ABCIQuery(path, data, height, trusted) error: ", err, result)
		return err
	} 

//	var abciquery ABCIQueryJSON
//
//	if err := json.Unmarshal(result, &abciquery); err != nil {
//		return err
//	}
	return nil
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
func ClientABCIInfo() error {
	client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
	info, err := client.ABCIInfo()

	if err != nil {
		fmt.Println("Handle client.ABCIInfo() error: ", err, info)
		return err
	} 

///	var abciinfo ABCIInfoJSON

///	if err := json.Unmarshal(info, &abciinfo); err != nil {
///		return err
///	}
	//  return info, nil
	return nil
}