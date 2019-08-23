package core

import (
	"demochain/utils"
	"fmt"
	"github.com/satori/go.uuid"
	"log"
)

type BlockChain struct {
	Blocks []*Block
	uid string
}

/*
请求生成区块
 */
func (bc *BlockChain) SendData(data string) {
	preBlock := bc.Blocks[len(bc.Blocks) - 1]
	newBlock := generateNewBlock(*preBlock,data)
	bc.AppendBlock(&newBlock)
}

/*
创建链
 */
func NewBlockChain() *BlockChain {

	firstBlock := generaFirstBlock()
	blockChain := BlockChain{}

	uid,_ := uuid.NewV4()
	blockChain.uid= uid.String()
	log.Printf("Generator BlockChain UUID: %s\n", blockChain.uid)
	blockChain.AppendBlock(&firstBlock)
	return &blockChain
}

/*
向链上增加区块节点
 */
func (bc *BlockChain) AppendBlock(newBlock *Block) {
	//新创建的链为空 添加创世区块
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks,newBlock)
		//写入数据库
		writeBlock2Db(newBlock,bc.uid)
		return
	}
	if isValid(*newBlock,*bc.Blocks[len(bc.Blocks) - 1]) {
		bc.Blocks = append(bc.Blocks,newBlock)
		//写入数据库
		writeBlock2Db(newBlock,bc.uid)
	}else {
		log.Fatal("invalid block")
	}

}

/*
验证区块有效性
 */
func isValid(newBlock Block,oldBlock Block) bool {
	if newBlock.Index - 1 != oldBlock.Index {
		return false
	}

	if newBlock.PrevBlockHash != oldBlock.Hash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

func (bc *BlockChain) PrintBlock() {
	for _,block := range bc.Blocks {
		fmt.Printf("Index:%d\n",block.Index)
		fmt.Printf("PrevHash:%s\n",block.PrevBlockHash)
		fmt.Printf("TimeStrmp:%d\n",block.TimeStamp)
		fmt.Printf("Hash:%s\n",block.Hash)
		fmt.Printf("Data:%s\n",block.Data)
	}

}

/*
区块写入库
 */
func writeBlock2Db(block *Block,uid string){

	db := utils.GetConnection()
	cnt := utils.Insert(*db,"INSERT INTO t_block(block_id,crt_time,prev_block_hash,block_hash,block_data,chain_uid) values(?,?,?,?,?,?)",
					block.Index,block.TimeStamp,block.PrevBlockHash,block.Hash,block.Data,uid)
	log.Printf("write block to db cnt:%d ,block index:%d\n",cnt,block.Index)
	defer db.Close()
}