package blc
type TxInput struct{
	//上一笔交易的索引
	Vout int
	ScriptSig string
	TxHash []byte
}