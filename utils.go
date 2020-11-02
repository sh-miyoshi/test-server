package main

import (
	"io"
	"math/rand"
	"time"
)

type randReader struct {
	total   int
	current int
}

func newRandReader(size int) *randReader {
	rand.Seed(time.Now().Unix())

	return &randReader{
		total:   size,
		current: 0,
	}
}

func (r *randReader) Read(b []byte) (int, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	remaining := r.total - r.current
	if remaining <= len(b) {
		for i := 0; i < remaining; i++ {
			b[i] = letters[rand.Intn(len(letters))]
		}
		return remaining, io.EOF
	}

	for i := 0; i < len(b); i++ {
		b[i] = letters[rand.Intn(len(letters))]
	}
	r.current += len(b)
	return len(b), nil
}
