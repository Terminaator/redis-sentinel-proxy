package proxy

import (
	"log"
	"net"
	"proxy/proxy/connection"
)

type proxy struct {
	port *string
}

func (p *proxy) listen(listener *net.TCPListener) {
	log.Println("listening proxy", listener.Addr())
	for {
		if conn, err := listener.AcceptTCP(); err == nil {
			c := connection.Connection(conn)
			go c.Read()
		} else {
			log.Fatal("Proxy failed", err)
		}
	}
}

func (p *proxy) Start() {
	if laddr, err := net.ResolveTCPAddr("tcp", *p.port); err == nil {
		if listener, err := net.ListenTCP("tcp", laddr); err == nil {
			p.listen(listener)
		} else {
			log.Fatal("Failed to listen", err)
		}
	} else {
		log.Fatal("Failed to resolve local address", err)
	}
}

func Proxy(port *string) *proxy {
	log.Println("starting proxy")
	return &proxy{port: port}
}
