package blc

import (
	"bytes"
	"encoding/gob"
	"log"
	"sync/atomic"
	"time"
)
type Block struct{
	Timestamp int64
	Hash	  []byte
	PreBlockHash	[]byte
	Height	int64
	Data	[]byte
	Nonce	int64 //碰撞次数，也可以是随机数
}
func NewBlock(preBlockHash []byte, height int64,data []byte) *Block{
	block:=Block{
		Timestamp:time.Now().Unix(),
		Hash: nil,
		PreBlockHash: preBlockHash,
		Height:	atomic.AddInt64(&height, 1),
		Data: data,
	}
	pow:=NewProofOfWork(&block)
	//execute pow to get hash
	hash,nonce:=pow.Run()
	block.Hash=hash
	block.Nonce=int64(nonce)
	return &block

	
}

func CreateGenesisBlock(data []byte) *Block{
	return NewBlock(nil, 0, data)
}

func (block *Block)Serialize() []byte{
	var buffer bytes.Buffer
	encoder:=gob.NewEncoder(&buffer)
	if err:=encoder.Encode(block);nil!=err{
		log.Panicf("Block serialize encoding err%v \n", err)
	}

	return buffer.Bytes()
}
func DeserializeBlock(BlockBytes []byte) *Block{
	var block Block
	decoder:=gob.NewDecoder(bytes.NewReader(BlockBytes))
	if err:=decoder.Decode(&block);err!=nil{
		log.Panicf("Block deserialize encoding err%v \n", err)
	}
	return &block
}