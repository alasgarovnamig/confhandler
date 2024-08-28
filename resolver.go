package confhandler

import (
	"os"
	"regexp"
)

// ResolveEnvVars dynamically resolves environment variables with a fallback to default values.
func ResolveEnvVars(value string) string {
	re := regexp.MustCompile(`\$\{([^:}]+)(?::([^}]+))?\}`)
	return re.ReplaceAllStringFunc(value, func(sub string) string {
		matches := re.FindStringSubmatch(sub)
		if len(matches) == 3 {
			envVar, defaultValue := matches[1], matches[2]
			if envValue, exists := os.LookupEnv(envVar); exists {
				return envValue
			}
			return defaultValue
		}
		return sub
	})
}
