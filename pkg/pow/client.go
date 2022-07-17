package pow

import (
	"context"
	"io"

	"quote-book/pkg/hashcash"
)

type Client struct {
	in  *Reader
	out *Writer
}

func NewClient(rw io.ReadWriter) *Client {
	return &Client{
		in:  wrapReader(rw),
		out: wrapWriter(rw),
	}
}

func (c *Client) Verify(ctx context.Context) (bool, error) {
	req := request{}
	if err := c.in.read(&req); err != nil {
		return false, err
	}
	solver := hashcash.New(hashcash.WithComplexity(req.Complexity))
	result, err := solver.Solve(ctx, req.Resource)
	if err != nil {
		return false, err
	}
	if err := c.out.write(response{Result: result}); err != nil {
		return false, err
	}
	accept := accept{}
	if err := c.in.read(&accept); err != nil {
		return false, err
	}
	return true, nil
}
