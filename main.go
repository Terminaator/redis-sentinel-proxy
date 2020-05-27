package main

import (
	"log"
	"proxy/proxy"
	"proxy/resp"
	"proxy/resp/client/sentinel"
)

func main() {
	log.Println("starting redis sentinel proxy")

	name := "mymaster"
	addr := "127.0.0.1:26379"

	resp.REDIS_MASTER = &name
	resp.SENTINEL_ADDRESS = &addr
	//sentinel.StartSentinel(initializer.NewInitializer("./configuration/clients.json"))
	sentinel := sentinel.Sentinel()
	go sentinel.Start()

	port := ":9999"

	proxy := proxy.Proxy(&port)
	go proxy.Start()
	//proxy.StartProxy(":9999")

	for {

	}
}
