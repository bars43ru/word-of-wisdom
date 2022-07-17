package tcpclient

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"quote-book/pkg/pow"
)

const tcpTimeout = 3 * time.Second

type Client struct {
	host    string
	timeout time.Duration
}

func New(host string) *Client {
	return &Client{
		host:    host,
		timeout: tcpTimeout,
	}
}

func (c *Client) Run(ctx context.Context) error {
	ctxConn, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	dialer := &net.Dialer{}
	log.Println("connecting to", c.host)
	conn, err := dialer.DialContext(ctxConn, "tcp", c.host)
	if err != nil {
		return fmt.Errorf("connect: %w", err)
	}
	defer conn.Close()

	log.Println("proof of work verification")
	ok, err := pow.NewClient(conn).Verify(ctx)
	if err != nil {
		err := fmt.Errorf("verification failed: %w", err)
		log.Println(err)
		return err
	}
	log.Println("verification:", ok)
	if !ok {
		return errors.New("verification failed")
	}

	b, err := io.ReadAll(conn)
	if err != nil {
		return fmt.Errorf("read data: %w", err)
	}
	log.Println(string(b))
	return nil
}
