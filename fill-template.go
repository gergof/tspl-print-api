package main

import (
	"fmt"
)

func fillTemplate(str string, args map[string]string) (err, string) {
	if len(str) > 0 && str[0] == '$' {
		// we need interpolation
		key := str[1:]

		val, ok := args[key]

		if !ok {
			return fmt.Errorf("Arguments map does not contain %v key", key), ""
		}

		return nil, val
	}

	return str
}
