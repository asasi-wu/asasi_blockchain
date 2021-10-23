package blc

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/boltdb/bolt"
)


const dbName="block.db"
const blockTableName="blocks"

type BlockChain struct{
	DB *bolt.DB	//数据库
	Tip []byte
}


func (bc *BlockChain)AddBlock(txs []*Transaction){
	err:=bc.DB.Update(func(tx *bolt.Tx) error{
		b:=tx.Bucket([]byte(blockTableName))
		if nil!=b{
			blockBytes:=b.Get(bc.Tip)
			latest_block:=DeserializeBlock(blockBytes)
			newBlock:=NewBlock(latest_block.Hash, latest_block.Height, txs)
			err:=b.Put(newBlock.Hash, newBlock.Serialize())
			if err!=nil{
				log.Panicf("Insert new block to DB failed %v\n", err)
			}
			err=b.Put([]byte("1"), newBlock.Hash)
			if err!=nil{
				log.Panicf("Update the latest block hash failed %v", err)
			}
			bc.Tip=newBlock.Hash

		}
		return nil
	})
	if err!=nil{
		log.Panicf("Adding block updated err %v\n", err)
	}

}
func dbExist() bool{
	if _,err:=os.Stat(dbName);os.IsNotExist(err){
		return false
	}
	return true
}
func CreateBlockChain(txs []*Transaction)*BlockChain{
	if dbExist(){
		fmt.Println("创世区块存在")
		os.Exit(1)
	}
	var blockHash []byte
	db, err:=bolt.Open(dbName,0600,nil)
	if err!=nil{
		log.Panicf("create db [%s]failed %v\n",dbName,err)
	}
	db.Update(func(tx *bolt.Tx) error {
		b:=tx.Bucket([]byte(blockTableName))
		if b==nil{
			b,err:=tx.CreateBucket([]byte(blockTableName))
			if err!=nil{
				log.Panicf("create bucket [%s] failed %v\n", blockTableName,err)
			}
			genesisBlock:=CreateGenesisBlock(txs)
			err=b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err!=nil{
				log.Panicf("insert the genesis block failed%v", err)

			}

			blockHash=genesisBlock.Hash


			err=b.Put([]byte("1"), genesisBlock.Hash)
			if err!=nil{
				log.Panicf("insert the genesis block failed %v\n", err)
			}

		}
		return nil
	})
	return &BlockChain{db,blockHash}

}
func (bc *BlockChain) TraverseBlockChain(){
	var curBlock *Block
	bcit:=bc.Iterator()
	fmt.Println("traversing the blcokchain and print all blocks")

	for{
		fmt.Println("---------------------------------------")
		curBlock=bcit.Next()
		fmt.Printf("Hash: %x\n", curBlock.Hash)
		fmt.Printf("Height: %v\n", curBlock.Height)
		fmt.Println("Timestamp:", time.Now().Format(fmt.Sprint(curBlock.Timestamp)))
		fmt.Printf("PreHash: %x\n", curBlock.PreBlockHash)
		fmt.Printf("Nonce: %v\n", curBlock.Nonce)
		var hashInt big.Int
		hashInt.SetBytes(curBlock.PreBlockHash)
		if big.NewInt(0).Cmp(&hashInt)==0{
			break
		}

	}

}
func BlockChainObject() *BlockChain{
	db,err:=bolt.Open(dbName,0600,nil)
	var tip []byte
	if err!=nil{
		log.Panicf("Open db failed [%s]", err)
	}
	err=db.View(func(tx *bolt.Tx) error{
		b:=tx.Bucket([]byte(blockTableName))
		if b!=nil{
			tip=b.Get([]byte("1"))
		}
		return nil
	})
	if err!=nil{
		log.Panicf("get the blockchain object failed %v\n", err)
	}
	return &BlockChain{db,tip}
}