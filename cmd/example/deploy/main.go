package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/darchlabs/synchronizer-v2/internal/blockchain"
	"github.com/darchlabs/synchronizer-v2/internal/contract"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// define and parse params to get from terminal
	urlParam := flag.String("url", "", "The URL of the node for connection to the network")
	privateKeyParam := flag.String("privateKey", "", "The private key of the wallet owns the deploy")
	flag.Parse()

	// validate params
	if *urlParam == "" {
		log.Fatal(errors.New("invalid url param"))
	}
	if *privateKeyParam == "" {
		log.Fatal(errors.New("invalid privateKey param"))
	}
	
	// initialize eth client instance
	client, err := ethclient.Dial(*urlParam)
	if err != nil {
		log.Fatal(err)
	}

	// intialize auth
	auth, err := blockchain.NewAuth(client, *privateKeyParam, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	// deploy the contract
	address, tx, _, err := contract.DeployContract(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	// show values
	fmt.Println("contract address:", address.Hex())
	fmt.Println("tx address:", tx.Hash().Hex())
}