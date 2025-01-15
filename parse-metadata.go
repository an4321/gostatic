package main

import (
	"fmt"
	"strings"
)

var requiredFields = []string{"title", "description", "author", "date"}

func ParseMetadata(metadata string) (map[string]string, error) {
	result := make(map[string]string)

	lines := strings.Split(strings.TrimSpace(metadata), "\n")

	for _, line := range lines {
		// skip empty lines
		if strings.TrimSpace(line) == "" {
			continue
		}

		// split into key and value
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid format in line: %s", line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// remove quotes if they exist
		value = strings.Trim(value, "\"")

		// add to map
		result[key] = value
	}

	// validate required fields
	for _, field := range requiredFields {
		value, exists := result[field]
		if !exists {
			return nil, fmt.Errorf("missing required field: %s", field)
		}
		if strings.TrimSpace(value) == "" {
			return nil, fmt.Errorf("empty required field: %s", field)
		}
	}

	return result, nil
}
