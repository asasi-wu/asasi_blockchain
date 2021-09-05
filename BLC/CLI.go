package blc

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct{

}
func PrintUsage(){
	fmt.Println("Usage: ")
	fmt.Println("\tcreateblockchain --创建区块链")
	fmt.Println("\taddblock --添加区块")
	fmt.Println("\tprintchain --输出区块链信息")

}

func (cli *CLI) create(){
	CreateBlockChain()
}

func (cli *CLI) addBlock(data string){
	if !dbExist(){
		fmt.Println("Database not exists")
		os.Exit(1)
	}
	blockchain:=BlockChainObject()
	blockchain.AddBlock([]byte(data))
}

func (cli *CLI) printchain(){
	if !dbExist(){
		fmt.Println("Database not exists")
		os.Exit(1)
	}
	
	blockchain:=BlockChainObject()
	blockchain.TraverseBlockChain()
}

func IsValidArgs(){
	if len(os.Args)<2{
		PrintUsage()
		os.Exit(1)
	}
}
func (cli *CLI) Run(){
	IsValidArgs()
	addBlockCmd:=flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd:=flag.NewFlagSet("printchain", flag.ExitOnError)
	createChainCmd:=flag.NewFlagSet("createblockchain", flag.ExitOnError)


	flagAddBlockArg:=addBlockCmd.String("data", "send 100 btc to alice", "添加区块数据")
	switch os.Args[1]{
	case "addblock":
		if err:=addBlockCmd.Parse(os.Args[2:]);err!=nil{
			log.Panicf("parse addblockCmd failed %v\n", err)
		}
	case "printchain":
		if err:=printChainCmd.Parse(os.Args[2:]);err!=nil{
			log.Panicf("parse printChainCmd failed %v\n", err)
		}
	case "createblockchain":
		if err:=createChainCmd.Parse(os.Args[2:]);err!=nil{
			log.Panicf("parse createblockchainCmd failed %v\n", err)
		}
	default:
		PrintUsage()
		os.Exit(1)

	}
	if addBlockCmd.Parsed(){
		if *flagAddBlockArg == ""{
			PrintUsage()
			os.Exit(1)
		}
		cli.addBlock(*flagAddBlockArg)

	}
	if printChainCmd.Parsed(){
		cli.printchain()
	}
	if createChainCmd.Parsed(){
		cli.create()
	}


}