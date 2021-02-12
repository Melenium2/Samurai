package imgprocess_test

import (
	"Samurai/internal/pkg/api/imgprocess"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveExtension_ShouldRemoveExtensionRight(t *testing.T) {
	shouldbe := "12345678"
	filename := shouldbe + ".jpg"

	assert.Equal(t, shouldbe, imgprocess.RemoveExtension(filename))
}
