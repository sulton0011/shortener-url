package helper

import (
	"strconv"
	"strings"
)

func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    int = 1
		args []interface{}
	)

	for k, v := range params {
		if k != "" {
			oldsize := len(namedQuery)
			namedQuery = strings.ReplaceAll(namedQuery, ":"+k, "$"+strconv.Itoa(i))

			if oldsize != len(namedQuery) {
				args = append(args, v)
				i++
			}
		}
	}

	return namedQuery, args
}

func IsEmtpy(s interface{}) bool {
	switch v := s.(type) {
	case int:
		return v == 0
	case int32:
		return v == 0
	case int64:
		return v == 0
	case string:
		return v == ""
	case float32:
		return v == 0.0
	case float64:
		return v == 0.0
	case []int:
		return len(v) == 0
	case []string:
		return len(v) == 0
	case []float32:
		return len(v) == 0
	case []float64:
		return len(v) == 0
	}

	return true
}
