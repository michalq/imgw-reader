package synop_test

import (
	"testing"

	"github.com/michalq/imgw/internal/synop"
	"github.com/stretchr/testify/assert"
)

func TestSynop(t *testing.T) {
	t.Run("should return true for single folder", func(t *testing.T) {
		b := synop.IsYearFolder("https://danepubliczne.imgw.pl/data/dane_pomiarowo_obserwacyjne/dane_meteorologiczne/dobowe/synop/2019/")
		assert.Equal(t, true, b)
	})
	t.Run("should return true for range folder", func(t *testing.T) {
		b := synop.IsYearFolder("https://danepubliczne.imgw.pl/data/dane_pomiarowo_obserwacyjne/dane_meteorologiczne/dobowe/synop/1981_1985/")
		assert.Equal(t, true, b)
	})
}
