package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	funcA = func(next HandlerFunc) HandlerFunc {
		return func(env interface{}) {
			data := env.(map[string]string)
			data["Order"] = data["Order"] + "A"
			next(env)
			data["Order"] = data["Order"] + "a"
		}
	}

	funcB = func(next HandlerFunc) HandlerFunc {
		return func(env interface{}) {
			data := env.(map[string]string)
			data["Order"] = data["Order"] + "B"
			next(env)
			data["Order"] = data["Order"] + "b"
		}
	}
)

func TestRunnerCall_WithOneMiddleware(t *testing.T) {
	env := make(map[string]string)
	env["Order"] = ""
	runner := New()
	runner.Use(funcA)
	runner.Call(env)

	assert.Equal(t, "Aa", env["Order"])
}

func TestRunnerCall(t *testing.T) {
	env := make(map[string]string)
	env["Order"] = ""
	runner := New()
	runner.Use(funcA)
	runner.Use(funcB)
	runner.Call(env)

	assert.Equal(t, "ABba", env["Order"])
}
