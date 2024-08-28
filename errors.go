package confhandler

import "fmt"

// ConfigError is a custom error type for configuration-related errors.
type ConfigError struct {
	Key string
	Err error
}

func (e *ConfigError) Error() string {
	return fmt.Sprintf("error retrieving config for key '%s': %v", e.Key, e.Err)
}

// NewConfigError creates a new ConfigError.
func NewConfigError(key string, err error) error {
	return &ConfigError{Key: key, Err: err}
}
