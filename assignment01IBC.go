package assignment01IBC

import (
	"crypto/sha256"
	"fmt"
)

type BlockData struct {
	Transactions []string
}
type Block struct {
	Data        BlockData
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}

func CalculateHash(inputBlock *Block) string {
	var c1 string
	//fmt.Println("calc hash entered")
	if inputBlock.PrevPointer == nil {
		v := fmt.Sprintf("%v", inputBlock.Data.Transactions)
		c1 = fmt.Sprintf("%x", sha256.Sum256([]byte(v)))
		return c1
	} else {
		n := fmt.Sprintf("%v", inputBlock.Data.Transactions)
		str := n + inputBlock.PrevHash
		//v:=fmt.Sprintf("%v",str)
		//fmt.Printf(v,"hiiiii\n")
		c1 = fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
		//fmt.Printf(c1)
		return c1
	}

}
func InsertBlock(dataToInsert BlockData, chainHead *Block) *Block {

	if chainHead == nil {
		var newBlock Block
		newBlock.Data = dataToInsert
		newBlock.CurrentHash = CalculateHash(&newBlock)
		//fmt.Printf(newBlock.CurrentHash,"\n")
		chainHead = &newBlock
		return chainHead

	} else {
		var newBlock Block
		newBlock.Data = dataToInsert
		newBlock.PrevPointer = chainHead
		newBlock.PrevHash = chainHead.CurrentHash
		newBlock.CurrentHash = CalculateHash(&newBlock)
		//fmt.Printf(newBlock.CurrentHash)
		chainHead = &newBlock
		return chainHead
	}
}
func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {

	for chainHead != nil {
		i := 0
		for i < len(chainHead.Data.Transactions) {
			//fmt.Printf("..................")
			//fmt.Printf(chainHead.Data.Transactions[i])
			if chainHead.Data.Transactions[i] == oldTrans {
				//fmt.Printf("..................")
				chainHead.Data.Transactions[i] = newTrans
			}
			i++
		}
		//fmt.Printf(".........********.........")
		chainHead = chainHead.PrevPointer
	}
}
func ListBlocks(chainHead *Block) {
	for chainHead != nil {
		fmt.Printf("%v", chainHead.Data.Transactions)
		fmt.Printf("<-")
		chainHead = chainHead.PrevPointer
	}
	fmt.Printf("\n")
}
func VerifyChain(chainHead *Block) {
	var change bool = false
	for chainHead != nil {
		hash := CalculateHash(chainHead)
		if chainHead.CurrentHash != hash {
			change = true
		}
		chainHead = chainHead.PrevPointer
	}
	if change == false {
		fmt.Printf("Blockchain is not compromised")
	} else {
		fmt.Printf("Blockchain is compromised")
	}
}

/*
import (
a1 "github.com/ehteshamz/assignment01IBC"
)
*/
func main() {
	var chainHead *Block
	genesis := BlockData{Transactions: []string{"S2E", "S2Z"}}
	//fmt.Printf("helllooo")
	//fmt.Println(genesis)
	chainHead = InsertBlock(genesis, chainHead)
	//var x string=CalculateHash(chainHead)
	//fmt.Printf("%x\n", x)

	secondBlock := BlockData{Transactions: []string{"E2Alice", "E2Bob", "S2John"}}
	chainHead = InsertBlock(secondBlock, chainHead)

	ListBlocks(chainHead)

	//ChangeBlock("S2E", "S2Trudy", chainHead)

	ListBlocks(chainHead)

	VerifyChain(chainHead)

}
