package cmd

import "demochain/core"

func main() {
	bc := core.NewBlockChain()
	bc.SendData("b1")
	bc.SendData("b2")
	bc.PrintBlock()


}
