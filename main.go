package main

import (
	blc "asasi_blockchain/BLC"
)

func main() {
	bc:=blc.CreateBlockChain()
	// fmt.Printf("Blockchain: %v", blockchain.Blocks[0])
	// blockchain.AddBlock(blockchain.Blocks[len(blockchain.Blocks)-1].Hash, 
	// blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,
	// []byte("alice send 10 btx to bob"))


	// blockchain.AddBlock(blockchain.Blocks[len(blockchain.Blocks)-1].Hash, 
	// blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,
	// []byte("bob send 10 btx to troytan"))

	// for _,block:=range blockchain.Blocks{
	// 	fmt.Printf("preBlockHash: %x, currentHash: %x\n", 
	// 	block.PreBlockHash,block.Hash)
	// }
	bc.AddBlock([]byte("alice send 100 eth to bob"))
	bc.AddBlock([]byte("tom send 100 eth to simon"))

	// bc.DB.View(func(tx *bolt.Tx)error{
	// 	b:=tx.Bucket([]byte("blocks"))
	// 	if nil!=b{
	// 		hash:=b.Get([]byte("1"))
	// 		fmt.Printf("latest hash value%v \n", hash)


	// 	}
	// 	return nil
	// })
	bc.TraverseBlockChain()


}
