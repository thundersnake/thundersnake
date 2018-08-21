package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAlphaNumString(t *testing.T) {
	assert.Equal(t, "B3XakKZx", AlphaNumString(8))
	assert.Equal(t, "OBrqkaS47Gh35X9aN6thyrrjEWhx84KRBFMqLlNa137nSKQ8zitH6dXbRJlFQqLU9mJ6dTbMij0CVKZjh5pHhI4f0MOvfRgTRAzrjJ3He0t915RY6LpB4fKoTZf3XUasZ0KY5uW1oVJmlWS2sTcE3PaXMCc4xoGRk0OFrTVBwZIg8ttetm7KWkd0bikZYyzKSztwlHlX0sG0IzXzujVsfQwI9x6k8mzbvg8xToOx39jvsBy1cWu4YhGcAl8BuUpi", AlphaNumString(256))
	assert.Equal(t, "", AlphaNumString(0))
}

func TestLetterString(t *testing.T) {
	assert.Equal(t, "OIaVXSoj", LetterString(8))
	assert.Equal(t, "PxDFPgrTjgjiSefrcMlQOKzQGdJfkBzTfoDmzciSosNXxiMEPvgZUTEYhBMgnjwSsgaBXwFDQXjRIFyoBxfFrdWqdlKsXWSAhPvzkyLSxjGpMTEhYxOCTZwXQTsDLWxfHSQUtAbhXbNZDRbqpAUQTcNhWTrADLvVBXVpUKJhlWzRegIIfYFIFRztukdjREfgTMtjnDAPYziWbWnysckODyKAAZIwBEmAxCSmdBsMXZFGyzvboyXXrHpqzJUUESUW", LetterString(256))
	assert.Equal(t, "", LetterString(0))
}
