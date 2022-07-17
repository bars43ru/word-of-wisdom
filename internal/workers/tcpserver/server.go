package tcpserver

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"quote-book/pkg/pow"
)

const (
	powTimeout = 15 * time.Second
)

type Quotes interface {
	RandomQuote() string
}

type Server struct {
	addr       *net.TCPAddr
	quotes     Quotes
	complexity uint8
}

func New(addr *net.TCPAddr, quotes Quotes, complexity uint8) *Server {
	return &Server{
		addr:       addr,
		quotes:     quotes,
		complexity: complexity,
	}
}

func (s *Server) Run(ctx context.Context) error {
	listener, err := net.ListenTCP("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("opening TCP listener: %w", err)
	}
	go func() {
		<-ctx.Done()
		log.Println("closing listen tcp")
		listener.Close()
	}()

	for conn := range s.loopAcceptingConn(ctx, listener) {
		s.handleConnection(ctx, conn)
	}
	return ctx.Err()
}

func (s *Server) loopAcceptingConn(ctx context.Context, listener *net.TCPListener) <-chan *net.TCPConn {
	log.Println("listen accept tcp connection")
	ch := make(chan *net.TCPConn)
	go func() {
		for ctx.Err() == nil {
			tcpConn, err := listener.AcceptTCP()
			if err != nil {
				log.Println(err)
				continue
			}
			log.Println("accepted: ", tcpConn.RemoteAddr().String())
			ch <- tcpConn
		}
	}()
	return ch
}

func (s *Server) handleConnection(ctx context.Context, conn *net.TCPConn) {
	defer conn.Close()

	ctx, cancel := context.WithTimeout(ctx, powTimeout)
	defer cancel()

	log.Println("proof of work verification")
	ok, err := pow.NewServer(conn, s.complexity).Verifying()
	if err != nil || !ok {
		log.Println("verification failed (err:", err, ")")
		return
	}

	quote := s.quotes.RandomQuote()
	io.WriteString(conn, quote)
}
