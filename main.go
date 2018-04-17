package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

//in memory database of blocks
var blocks []*Block

func main() {

	go func() {
		err := serveHttp()
		if err != nil {
			log.Printf("could not start web server %v", err)
			os.Exit(2)
		}
	}()

	if err := listenForBlocks(); err != nil {
		log.Printf("failed listening for blocks %v", err)
		os.Exit(2)
	}
}

//opens a web socket connection to infura's hosted Ethereum nodes and listens for new block events
func listenForBlocks() error {
	conn, err := ethclient.Dial("wss://" + os.Getenv("ENVIRONMENT") + ".infura.io/ws")
	if err != nil {
		return fmt.Errorf("Cannot dial websocket connection %v", err)
	}

	ch := make(chan *types.Header)
	sub, err := conn.SubscribeNewHead(context.Background(), ch)
	if err != nil {
		return fmt.Errorf("Cannot subscribe to head %v", err)
	}

	for {
		select {
		case msg := <-ch:
			log.Println("New Block:", msg.Number.String())
			getBlock(msg.Number)
		case err := <-sub.Err():
			log.Println("Connection error:", err)
		case <-time.After(15 * time.Second):
			fmt.Println("Waiting for blocks...")
		}
	}
}

//Gets block information from chain, takes awhile for the block to show up after it is announced.
func getBlock(msg *big.Int) {

	//cannot re-use existing websocket connection
	conn, err := ethclient.Dial("https://" + os.Getenv("ENVIRONMENT") + ".infura.io/")
	if err != nil {
		log.Println("Could not open connection to chain")
		return
	}

	for {
		block, err := conn.BlockByNumber(context.Background(), msg)
		if err != nil {
			if err.Error() == "not found" {
				time.Sleep(5 * time.Second)
				continue
			}
			fmt.Println("Error Getting Block:", err)
			break
		} else {
			b := &Block{
				Block: block.Number().String(),
			}
			b.Transactions = make([]string, block.Transactions().Len())
			for i, t := range block.Transactions() {
				b.Transactions[i] = t.Hash().String()
			}

			blocks = append([]*Block{b}, blocks...)
			break
		}
	}
}
