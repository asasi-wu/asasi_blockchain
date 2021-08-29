package blc

import "github.com/boltdb/bolt"
type BlockChainIterator struct{
	DB *bolt.DB
	CurrentHash []byte
}
func (blc *BlockChain) Iterator() *BlockChainIterator{
	return &BlockChainIterator{blc.DB,blc.Tip}
}
func (bcit *BlockChainIterator) Next() *Block{
	var block *Block
	bcit.DB.View(func(tx *bolt.Tx)error{
		b:=tx.Bucket([]byte(blockTableName))
		if b!=nil{
			currentBlockBytes:=b.Get(bcit.CurrentHash)
			block=DeserializeBlock(currentBlockBytes)
			bcit.CurrentHash=block.PreBlockHash

		}
		return nil
	})
	return block
}