package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	//区块编号
	Index int64
	//时间戳
	TimeStamp int64
	//上一区块哈希值
	PrevBlockHash string
	//当前区块哈希值
	Hash string
	//区块数据
	Data string
}

/*
计算区块hash值
*/
func calculateHash(b Block) string {
	blockData := string(b.Index) + string(b.TimeStamp) + b.PrevBlockHash + b.Data
	hashInBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashInBytes[:])
}

/*
创建新区块
 */
func generateNewBlock(perBlock Block,data string) Block {
	newBlock := Block{}
	newBlock.Index = perBlock.Index + 1
	newBlock.TimeStamp = time.Now().Unix()
	newBlock.PrevBlockHash = perBlock.Hash
	newBlock.Data = data
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

/*
创建创世区块
 */
func generaFirstBlock() Block{
		firstBlock := Block{}
		firstBlock.Index = -1
		firstBlock.Hash = ""
		return generateNewBlock(firstBlock,"First Block")
}