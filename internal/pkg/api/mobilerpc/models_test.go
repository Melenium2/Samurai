package mobilerpc_test

import (
	"Samurai/internal/pkg/api/mobilerpc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProxy_ShouldReturnCorrectInstance(t *testing.T) {
	p := mobilerpc.NewProxy("http://1d213:asddasd@14.23.51.22:1322")
	assert.Equal(t, "http://1d213:asddasd@14.23.51.22:1322", p.Http)
	assert.Equal(t, "https://1d213:asddasd@14.23.51.22:1322", p.Https)

	p = mobilerpc.NewProxy("https://1d213:asddasd@14.23.51.22:1322")
	assert.Equal(t, "http://1d213:asddasd@14.23.51.22:1322", p.Http)
	assert.Equal(t, "https://1d213:asddasd@14.23.51.22:1322", p.Https)
}
