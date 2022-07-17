package pow

import (
	"crypto/rand"
	"fmt"
	"io"

	"word-of-wisdom/pkg/hashcash"
)

type Server struct {
	in       *Reader
	out      *Writer
	verifier hashcash.Service
}

func NewServer(rw io.ReadWriter, complexity uint8) *Server {
	return &Server{
		in:       wrapReader(rw),
		out:      wrapWriter(rw),
		verifier: hashcash.New(hashcash.WithComplexity(complexity)),
	}
}

func (s *Server) Verifying() (bool, error) {
	resource, err := randomBytes(defaultSize)
	if err != nil {
		return false, err
	}
	req := request{
		Complexity: s.verifier.Complexity(),
		Resource:   resource,
	}
	if err = s.out.write(req); err != nil {
		return false, fmt.Errorf("send request for solving %w", err)
	}
	resp := response{}
	if err = s.in.read(&resp); err != nil {
		return false, fmt.Errorf("wait response result solve %w", err)
	}
	if !s.verifier.Verify(req.Resource, resp.Result) {
		return false, nil
	}
	if err = s.out.write(accept{}); err != nil {
		return false, fmt.Errorf("send accept responce %w", err)
	}
	return true, nil
}

func randomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
