package widgets

import (
	"fmt"
	"math"
	"strings"
)

type progressBar struct {
	progress float64
	max      float64
	bars     int
	text     string
	prefix   string
}

func ProgressBar(progress, max float64) *progressBar {
	return &progressBar{
		progress: progress,
		max:      max,
		bars:     10,
	}
}

func (p *progressBar) WithBars(bars int) *progressBar {
	p.bars = bars
	return p
}

func (p *progressBar) WithText(text string) *progressBar {
	p.text = text
	return p
}

func (p *progressBar) WithPrefix(prefix string) *progressBar {
	p.prefix = prefix
	return p
}

func (p *progressBar) Render() string {
	repeat := int(math.Round(p.progress * float64(p.bars) / p.max))

	bar := strings.Repeat("■", repeat)
	bar += strings.Repeat("□", p.bars-repeat)

	r := fmt.Sprintf("%s %s %s", p.prefix, bar, p.text)
	return strings.TrimSpace(r)
}
