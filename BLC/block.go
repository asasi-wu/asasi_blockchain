package blc

import (
	"bytes"
	"crypto/sha256"
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
	Txs     []*Transaction //交易数据
	Nonce	int64 //碰撞次数，也可以是随机数
}
func NewBlock(preBlockHash []byte, height int64,txs []*Transaction) *Block{
	block:=Block{
		Timestamp:time.Now().Unix(),
		Hash: nil,
		PreBlockHash: preBlockHash,
		Height:	atomic.AddInt64(&height, 1),
		Txs: txs,
	}
	pow:=NewProofOfWork(&block)
	//execute pow to get hash
	hash,nonce:=pow.Run()
	block.Hash=hash
	block.Nonce=int64(nonce)
	return &block

	
}

func CreateGenesisBlock(txs []*Transaction) *Block{
	return NewBlock(nil, 0, txs)
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
func (block *Block)HashTransaction()[]byte{
	var txHashes [][]byte
	for _,tx:=range block.Txs{
		txHashes=append(txHashes, tx.TxHash)
	}
	txHash:=sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}