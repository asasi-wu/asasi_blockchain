package blc

import (
	"asasi_blockchain/BLC/utils"
	"bytes"
	"crypto/sha256"
	"math/big"
)

const targetBit = 16

type ProofOfWork struct {
	Block  *Block
	target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	//数据总长度8位
	//需要满足前两位为0
	//
	target = target.Lsh(target, 256-targetBit)
	return &ProofOfWork{Block: block, target: target}

}

//返回哈希值与碰撞次数
func (proofOfWork *ProofOfWork) Run() ([]byte, int) {
	var nonce = 0
	var hashInt big.Int
	var hash [32]byte
	for {
		dataBytes := proofOfWork.prepareData(int64(nonce))
		hash = sha256.Sum256(dataBytes)
		hashInt.SetBytes(hash[:])
		if proofOfWork.target.Cmp(&hashInt) == 1 {
			break
		}
		nonce++

	}
	return hash[:], nonce

}

func (pow *ProofOfWork) prepareData(nonce int64) []byte {

	data := bytes.Join([][]byte{
		utils.IntToHex(pow.Block.Timestamp),
		utils.IntToHex(pow.Block.Height),
		pow.Block.PreBlockHash,
		pow.Block.HashTransaction(),
		utils.IntToHex(nonce),
		utils.IntToHex(targetBit),
	}, []byte{})
	return data

}
