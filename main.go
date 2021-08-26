package main

import (
	blc "asasi_blockchain/BLC"
	"fmt"
)


func main(){
	blockchain:=blc.CreateBlockChain()
	fmt.Printf("Blockchain: %v", blockchain.Blocks[0])

}