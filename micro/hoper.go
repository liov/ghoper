package main

import (
	"hoper/client"
	"hoper/server"
)

func main() {
	go server.Service()
	client.Client()
}

//查看日志
//tail -n 10 *.log
