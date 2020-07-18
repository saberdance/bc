package main

import (
	"fmt"

	"github.com/suutaku/community-chain/src/lib"
)

func TestWallet() {
	var wa = lib.LoadWallet("./ccw.key")
	// wa.GenPrivKey()
	fmt.Println(wa)
	fmt.Println("wa")
	fmt.Println(wa.Key)
	fmt.Println(wa.Key.PubKey())
	fmt.Println(wa.Key.PubKey().Address())
}

func main() {
	TestWallet()
}
