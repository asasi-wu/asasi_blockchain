package blc
type BlockChain struct{
	Blocks []*Block
}
func (chain *BlockChain)AddBlock(preBlockHash []byte, height int64,data []byte){
	newblock:=NewBlock(preBlockHash, height, data)
	chain.Blocks=append(chain.Blocks, newblock)

}
func CreateBlockChain()*BlockChain{
	block:=CreateGenesisBlock([]byte("init blockchain"))
	return &BlockChain{[]*Block{block}}

}