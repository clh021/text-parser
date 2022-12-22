package env

import (
	"os"
	"strings"
)

func GetEnv() map[string]string {
	envStrs := os.Environ() // strings.Split(str, "\n")
	m := make(map[string]string)
	for _, e := range envStrs {
		if i := strings.Index(e, "="); i >= 0 {
			m[e[:i]] = e[i+1:]
		}
	}
	return m
}
