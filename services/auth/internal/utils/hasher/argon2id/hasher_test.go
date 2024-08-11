package argon2id

import (
	"testing"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/utils/hasher"
	"github.com/stretchr/testify/assert"
)

func TestArgonHasher(t *testing.T) {
	cases := []struct {
		name     string
		options  hasher.Options
		password []byte
	}{
		{
			name:     "default options",
			options:  hasher.DefaultOptions,
			password: []byte("password123"),
		},
		{
			name: "high memory",
			options: hasher.Options{
				Memory:      128 * 1024,
				Iterations:  1,
				Parallelism: 2,
				SaltLength:  16,
				KeyLength:   32,
			},
			password: []byte("password123"),
		},
		{
			name: "high iterations",
			options: hasher.Options{
				Memory:      64 * 1024,
				Iterations:  3,
				Parallelism: 2,
				SaltLength:  16,
				KeyLength:   32,
			},
			password: []byte("password123"),
		},
		{
			name: "high parallelism",
			options: hasher.Options{
				Memory:      64 * 1024,
				Iterations:  1,
				Parallelism: 4,
				SaltLength:  16,
				KeyLength:   32,
			},
			password: []byte("password123"),
		},
		{
			name: "empty password",
			options: hasher.Options{
				Memory:      64 * 1024,
				Iterations:  1,
				Parallelism: 2,
				SaltLength:  16,
				KeyLength:   32,
			},
			password: []byte(""),
		},
		{
			name: "very long password",
			options: hasher.Options{
				Memory:      64 * 1024,
				Iterations:  1,
				Parallelism: 2,
				SaltLength:  16,
				KeyLength:   32,
			},
			password: []byte("a very long password that exceeds normal length for testing purposes"),
		},
		{
			name: "different salt length",
			options: hasher.Options{
				Memory:      64 * 1024,
				Iterations:  1,
				Parallelism: 2,
				SaltLength:  32,
				KeyLength:   32,
			},
			password: []byte("password123"),
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			h := New(test.options)

			// Test Hash method
			hash, err := h.Hash(test.password)
			assert.NoError(t, err)
			assert.NotEmpty(t, hash)

			// Test Compare method with correct password
			match, err := h.Compare(test.password, hash)
			assert.NoError(t, err)
			assert.True(t, match)

			// Test Compare method with incorrect password
			wrongPassword := []byte("wrongpassword")
			match, err = h.Compare(wrongPassword, hash)
			assert.NoError(t, err)
			assert.False(t, match)
		})
	}
}
