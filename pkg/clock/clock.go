package clock

import "time"

// Clock defines the interface for getting the current time.
type Clock interface {
	Now() time.Time
}

// RealClock implements the Clock interface using the system time.
type RealClock struct{}

// Now return system time.
func (c RealClock) Now() time.Time {
	return time.Now()
}

// MockClock implements the Clock interface for testing purposes.
type MockClock struct {
	CurrentTime time.Time
}

// Now return defined time.
func (c MockClock) Now() time.Time {
	return c.CurrentTime
}
