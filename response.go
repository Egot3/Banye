package banye

import (
	"crypto/sha256"
)

type Response struct {
	Headers    map[string][]string
	StatusCode int
	Body       []byte
}

// Hashing response(for comparing)
func (r Response) Hash() [32]uint8 {

	hash := sha256.Sum256(r.Body) //there is really not much to hash tbh

	return hash
}
