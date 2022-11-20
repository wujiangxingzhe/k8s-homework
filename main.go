package main

import (
	"github.com/wujiangxingzhe/k8s-homework/server"
)

func main() {
	s := server.NewServer(":8080")
	s.ListenAndServe()
}
