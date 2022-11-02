package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"math/big"

	"github.com/darchlabs/synchronizer-v2/internal/blockchain"
	"github.com/darchlabs/synchronizer-v2/internal/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main () {
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

	// initialice eth client instance
	client, err := ethclient.Dial(*urlParam)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(11111, *contractAddrParam)
	
	// intance smart contract
	instance, err := contract.NewContract(common.HexToAddress(*contractAddrParam), client)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(11111)
	
	// initialize auth
	auth, err := blockchain.NewAuth(client, *privateKeyParam, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(22222)
	
	// call example contract method
	tx, err := instance.Example(auth, big.NewInt(int64(444)))
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(333333)
	fmt.Printf("tx sent hash: %s", tx.Hash().Hex())

	// get example number from contract
	num, err := instance.PerfomUpkeep(/* ... */)
	if err != nil {
		fmt.Println("jajaja")
		fmt.Println(err.Error())
		log.Fatal(err)
	}
	
	fmt.Println(4444)
	fmt.Println("current number", num.String())
}