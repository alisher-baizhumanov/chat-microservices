package argon2id

import (
	"github.com/alexedwards/argon2id"
	hasher2 "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/utils/hasher"
)

// argonHasher implements the PasswordHasher interface using the Argon2id algorithm.
type argonHasher struct {
	params *argon2id.Params
}

// Hash generates a hashed password from the given plain text password using the Argon2id algorithm.
func (a *argonHasher) Hash(password []byte) ([]byte, error) {
	hash, err := argon2id.CreateHash(string(password), a.params)
	if err != nil {
		return nil, err
	}

	return []byte(hash), nil
}

// Compare checks if the given plain text password matches the hashed password using the Argon2id algorithm.
func (a *argonHasher) Compare(password, hashedPassword []byte) (bool, error) {
	return argon2id.ComparePasswordAndHash(string(password), string(hashedPassword))
}

// New creates a new instance of argonHasher with the provided options.
func New(opts hasher2.Options) hasher2.PasswordHasher {
	return &argonHasher{
		params: &argon2id.Params{
			Memory:      opts.Memory,
			Iterations:  opts.Iterations,
			Parallelism: opts.Parallelism,
			SaltLength:  opts.SaltLength,
			KeyLength:   opts.KeyLength,
		},
	}
}
