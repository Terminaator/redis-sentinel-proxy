package tcp

import (
	"errors"
	"net"
	"proxy/resp/client/tcp/util"
)

type TcpConnection struct {
	addr       *string
	connection *net.TCPConn
	writer     *util.Writer
	reader     *util.Reader
}

func (t *TcpConnection) Write(buf *[]byte) error {
	if t.connection == nil {
		if err := t.connect(); err != nil {
			return err
		}
	}

	if err := t.writer.Write(buf); err == nil {
		return nil
	} else {
		t.connection = nil
		return err
	}
}

func (t *TcpConnection) connect() error {
	if t.addr == nil {
		return errors.New("addr is null")
	}

	a, _ := net.ResolveTCPAddr("tcp", *t.addr)

	if conn, err := net.DialTCP("tcp", nil, a); err == nil {
		t.connection = conn
		t.writer = util.NewWriter(conn)
		t.reader = util.NewReader(conn)
		return nil
	} else {
		return err
	}
}

func (t *TcpConnection) Read() (*[]byte, error) {
	if buf, err := t.reader.Read(); err == nil {
		return &buf, err
	} else {
		t.connection = nil
		return nil, err
	}
}

func Connection(addr *string) *TcpConnection {
	return &TcpConnection{addr: addr}
}
