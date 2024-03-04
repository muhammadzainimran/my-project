
package blockchain

import "fmt"

type MyBlock struct {
	Data string
}

var MyBlockchain []MyBlock

func DisplayAllMyBlocks() {
	for _, block := range MyBlockchain {
		fmt.Println(block.Data)
	}
}

func CreateMyBlock(data string) {
	newBlock := MyBlock{Data: data}
	MyBlockchain = append(MyBlockchain, newBlock)
}

func ModifyMyBlock(index int, newData string) {
	if index < 0 || index >= len(MyBlockchain) {
		fmt.Println("Invalid index")
		return
	}
	MyBlockchain[index].Data = newData
}