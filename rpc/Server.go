package main

import (
	"demochain/core"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var blockchain *core.BlockChain

func run() {
	http.HandleFunc("/blockchain/get",blockchainGetHandler)
	http.HandleFunc("/blockchain/write",blockchainWriteHandler)
	http.ListenAndServe("localhost:8888",nil)
}

func blockchainGetHandler(w http.ResponseWriter,r *http.Request) {
	//跨域设置
	w.Header().Set("Access-Control-Allow-Origin", "*")
	bytes,error := json.Marshal(blockchain)
	if error != nil {
		http.Error(w,error.Error(),http.StatusInternalServerError)
	}
	io.WriteString(w,string(bytes))
}
func blockchainWriteHandler(w http.ResponseWriter,r *http.Request) {
	//跨域设置
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers","Content-Type")//header的类型
	w.Header().Set("content-type","application/json")//返回数据格式是json

	if r.Method == http.MethodOptions {
		log.Println("request options ")
		io.WriteString(w,"success")
		return
	}
	blockData := r.URL.Query().Get("data")
	blockchain.SendData(blockData)
	blockchainGetHandler(w,r)
}

func main() {

	blockchain = core.NewBlockChain()
	run()

	 /*db := utils.GetConnection()
	 rows := utils.Query(*db,"select * from t_block")
	rs := utils.HandleRls(rows)
	fmt.Printf("%x",rs)*/
	// fmt.Printf("query rs:%d \n",len(rs))

	//cnt := utils.Insert(*db,"INSERT INTO t_block(block_id,crt_time,block_hash,block_data) values(?,?,?,?)","2","123","111","222")
	//fmt.Println("cnt=",cnt)
}