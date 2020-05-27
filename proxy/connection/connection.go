package connection

import (
	"log"
	"net"
	"proxy/resp/client/redis"
	"proxy/resp/client/tcp/util"
)

type connection struct {
	conn   *net.TCPConn
	reader *util.Reader
	writer *util.Writer
	redis  *redis.Redis
}

func (r *connection) error(err error) *[]byte {
	return &([]byte("-" + err.Error() + "\r\n"))
}

func (c *connection) out(buf *[]byte) {
	c.writer.Write(buf)
}

func (c *connection) in(buf *[]byte) {
	buf, err := c.redis.Do(buf)

	if err == nil {
		c.out(buf)
	} else {
		c.out(c.error(err))
	}
}

func (c *connection) Read() {
	log.Println("proxy reading from a connection", c.conn.RemoteAddr().String())

	for {
		if buf, err := c.reader.Read(); err == nil {
			c.in(&buf)
		} else {
			break
		}
	}

	c.close()
}

func (c *connection) close() {
	log.Println("proxy connection close ip", c.conn.RemoteAddr().String())
	c.conn.Close()
}

func Connection(c *net.TCPConn) *connection {
	log.Println("new proxy connection", c.RemoteAddr().String())
	return &connection{conn: c, writer: util.NewWriter(c), reader: util.NewReader(c), redis: redis.GetRedis()}
}
