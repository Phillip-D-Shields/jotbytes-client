/*
Copyright © 2024 PHILLIP DAVID SHIELDS matua.phillip.shields@gmail.com
*/
package main

import (
	"example.com/jotbytes-client/cmd"
	"math/rand"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {
	cmd.Execute()
}
