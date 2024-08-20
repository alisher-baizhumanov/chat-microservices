package hasher

import "runtime"

// DefaultOptions provides a set of default parameters for the Argon2id hashing algorithm.
// These parameters are chosen to balance security and performance for most use cases.
var DefaultOptions = Options{
	Memory:      64 * 1024,
	Iterations:  1,
	Parallelism: uint8(runtime.NumCPU()),
	SaltLength:  16,
	KeyLength:   32,
}

// Options defines the configurable parameters for the hashing algorithm.
type Options struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}
