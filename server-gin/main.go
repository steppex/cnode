package main

import (
	"server-gin/server"
)

func main() {
	go server.HttpListen()
	select {}
}
