package sentinel

import (
	"fmt"
	"log"
	"proxy/resp"
	"proxy/resp/client/tcp"
	"proxy/resp/constant"
	"strings"
	"time"
)

type sentinel struct {
	connection *tcp.TcpConnection
}

func (s *sentinel) master(buf *[]byte) {
	parts := strings.Split(string(*buf), "\r\n")

	if len(parts) == 6 {
		addr := parts[2] + ":" + parts[4]
		if resp.REDIS_ADDRESS == nil || *resp.REDIS_ADDRESS != addr {
			log.Println("sentinel redis master", addr)
			resp.REDIS_ADDRESS = &addr
		}
	}
}

func (s *sentinel) Start() {
	buf := []byte(fmt.Sprintf(constant.SENTINEL_COMMAND, *resp.REDIS_MASTER))

	if err := s.connection.Write(&buf); err == nil {
		if buf, err := s.connection.Read(); err == nil {
			s.master(buf)
		}
	}

	time.Sleep(10 * time.Second)
	s.Start()
}

func Sentinel() *sentinel {
	log.Println("creating sentinel client", *resp.SENTINEL_ADDRESS, "master name", *resp.REDIS_MASTER)
	return &sentinel{connection: tcp.Connection(resp.SENTINEL_ADDRESS)}
}
