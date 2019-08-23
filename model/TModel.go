package model

/*
T_BLOCK表
 */
type TBlock struct {
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
