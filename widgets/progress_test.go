package widgets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProgressBar(t *testing.T) {
	assert.Equal(t,
		"⏳ ■■□□ 2h10m",
		ProgressBar(130, 240).
			WithBars(4).
			WithPrefix("⏳").
			WithText("2h10m").
			Render(),
	)
}
