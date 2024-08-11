package hasher

// PasswordHasher defines the interface for hashing and comparing passwords.
type PasswordHasher interface {
	Hash(password []byte) ([]byte, error)
	Compare(password, hashedPassword []byte) (bool, error)
}
