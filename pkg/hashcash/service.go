package hashcash

import (
	"context"
	"fmt"
	"math"
)

type Opt func(*Service)

type Service struct {
	complexity uint8
}

func New(opts ...Opt) Service {
	s := Service{complexity: defaultComplexity}
	for _, f := range opts {
		f(&s)
	}
	return s
}

func (s *Service) Solve(ctx context.Context, resource []byte) ([]byte, error) {
	for count := uint(0); ctx.Err() == nil && s.complexity != 0; {
		ending := base64EncodeUInt(count)
		hash := sha1Hash(s.makeValue(resource, ending))
		if s.acceptHash(hash) {
			return ending, nil
		}
		if math.MaxUint == count {
			return []byte{}, ErrSolutionNotFound
		}
		count++
	}
	return []byte{}, ctx.Err()
}

func (s Service) Verify(resource []byte, ending []byte) bool {
	hash := sha1Hash(s.makeValue(resource, ending))
	return s.acceptHash(hash)
}

func (s Service) Complexity() uint8 {
	return s.complexity
}

func (s Service) acceptHash(hash string) bool {
	for _, val := range hash[:s.complexity] {
		if val != defaultPrefix {
			return false
		}
	}
	return true
}

func (s Service) makeValue(resource []byte, ending []byte) string {
	return fmt.Sprintf("%s:%s", resource, ending)
}
