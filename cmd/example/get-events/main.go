package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ContractABI = "{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Example\",\"type\":\"event\"}"

func main() {
	// define and parse params to get from terminal
	urlParam := flag.String("url", "", "The URL of the node for connection to the network")
	privateKeyParam := flag.String("privateKey", "", "The private key of the walle owns the event")
	contractAddrParam := flag.String("contractAddr", "", "The contract addres for deployment contract")
	flag.Parse()

	// validate params
	if *urlParam == "" {
		log.Fatal(errors.New("invalid url param"))
	}
	if *privateKeyParam == "" {
		log.Fatal(errors.New("invlid privateKey"))
	}
	if *contractAddrParam == "" {
		log.Fatal(errors.New("invalid contractAddr param"))
	}

	// initialize eth client instance
	client, err := ethclient.Dial(*urlParam)
	if err != nil {
		log.Fatal(err)
	}

	// instance smart contract
	// instance, err := contract.NewContract(common.HexToAddress(*contractAddrParam), client)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	query := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock: nil,
		Addresses: []common.Address{
			common.HexToAddress(*contractAddrParam),
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	// TODO(ca): get abi from event struct
	// contractABI := contract.ContractABI

	contractAbi, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		// fmt.Println("vLog.BlockHash.Hex()", vLog.BlockHash.Hex())
		// fmt.Println("vLog.BlockNumber", vLog.BlockNumber)
		// fmt.Println("vLog.TxHash.Hex()", vLog.TxHash.Hex())

		e, err := contractAbi.Unpack("Example", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(reflect.TypeOf(e), e)
		fmt.Println("-----")
	}
}