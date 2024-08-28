package confhandler

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config holds the configuration data loaded from a YAML file.
type Config struct {
	data map[string]interface{}
}

// LoadConfig loads a YAML configuration file into a Config object.
func LoadConfig(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var configData map[string]interface{}
	err = yaml.Unmarshal(data, &configData)
	if err != nil {
		return nil, err
	}

	return &Config{data: configData}, nil
}

// Get retrieves the value by key, resolves environment variables, and parses to the desired type.
func (c *Config) Get(key string, expectedType SupportedType) (interface{}, error) {
	rawValue, ok := c.data[key]
	if !ok {
		return nil, errors.New("key not found in config")
	}

	stringValue := fmt.Sprintf("%v", rawValue)
	resolvedValue := ResolveEnvVars(stringValue)

	return ParseValue(resolvedValue, expectedType)
}
