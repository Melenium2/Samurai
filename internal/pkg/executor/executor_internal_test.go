package executor

import (
	"Samurai/internal/pkg/api/inhuman"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSamurai_Bundles_ShouldReturnOnlyBundles_NoError(t *testing.T) {
	sam := Samurai{}

	apps := []inhuman.App{
		{Bundle: "1"},
		{Bundle: "2"},
		{Bundle: "3"},
		{Bundle: "4"},
	}
	bundles := sam.bundles(apps)

	assert.Equal(t, []string{"1", "2", "3", "4"}, bundles)
}

func TestSamurai_Position_ShouldReturnCorrectPosition_NoError(t *testing.T) {
	sam := Samurai{}

	bundles := []string {
		"1",
		"2",
		"com.nebundles",
		"jp.japanc",
		"com.bundle1",
		"com.bundle2",
		"eror",
		"empty",
	}
	p := sam.position("com.bundle1", bundles)
	assert.Equal(t, 4, p)
}

func TestSamurai_Position_ShouldReturnMinusOne_NoError(t *testing.T) {
	sam := Samurai{}

	bundles := []string {
		"1",
		"2",
		"com.nebundles",
		"jp.japanc",
		"com.bundle2",
		"eror",
		"empty",
	}
	p := sam.position("com.bundle1", bundles)
	assert.Equal(t, -1, p)
}