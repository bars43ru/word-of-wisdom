package quotes

import "math/rand"

type Store struct {
}

func New() *Store {
	return &Store{}
}

func (s Store) RandomQuote() string {
	return quotes[rand.Intn(len(quotes))]
}
