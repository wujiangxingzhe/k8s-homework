package main

import (
	"log"

	"github.com/wujiangxingzhe/k8s-homework/server"
)

func main() {
	s := server.NewServer(":8080")
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
