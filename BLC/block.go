package blc

import (
	"sync/atomic"
	"time"
)
type Block struct{
	Timestamp int64
	Hash	  []byte
	PreBlockHash	[]byte
	Height	int64
	data	[]byte
	nonce	int64
}
func NewBlock(preBlockHash []byte, height int64,data []byte) *Block{
	block:=Block{
		Timestamp:time.Now().Unix(),
		Hash: nil,
		PreBlockHash: preBlockHash,
		Height:	atomic.AddInt64(&height, 1),
		data: data,
	}
	block.SetHash()
	pow:=NewProofOfWork(&block)
	hash,nonce:=pow.Run()
	block.Hash=hash
	block.nonce=int64(nonce)
	return &block

	
}
func (b *Block)SetHash(){

}

func CreateGenesisBlock(data []byte) *Block{
	return NewBlock(nil, 1, data)
}