package blc

import (
	"bytes"
	"crypto/sha256"
	"sync/atomic"
	"time"
)
type Block struct{
	Timestamp int64
	Hash	  []byte
	PreBlockHash	[]byte
	Height	int64
	data	[]byte
}
func NewBlock(preBlockHash []byte, height int64,data []byte) *Block{
	return &Block{
		Timestamp:time.Now().Unix(),
		Hash: nil,
		PreBlockHash: preBlockHash,
		Height:	atomic.AddInt64(&height, 1),
		data: data,
	}
	
}
func (b *Block)SetHash(){
	blockBytes:=bytes.Join([][]byte{
		IntToHex(b.Timestamp),
		IntToHex(b.Height),
		b.PreBlockHash,
		b.data,
	}, []byte{})
	hash:=sha256.Sum256(blockBytes)
	b.Hash=hash[:]
}