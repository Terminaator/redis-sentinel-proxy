package redis

import (
	"errors"
	"proxy/resp"
	"proxy/resp/client/tcp"
	"time"
)

type Redis struct {
	connection *tcp.TcpConnection
}

func (r *Redis) Do(buf *[]byte) (*[]byte, error) {
	if err := r.write(buf); err != nil {
		return nil, err
	}

	return r.connection.Read()
}

func (r *Redis) write(buf *[]byte) error {
	for i := 0; i < 10; i++ {
		if err := r.connection.Write(buf); err == nil {
			return nil
		}
		time.Sleep(time.Second * 1)
	}

	return errors.New("redis writing failed")
}

func GetRedis() *Redis {
	return &Redis{connection: tcp.Connection(resp.REDIS_ADDRESS)}
}
