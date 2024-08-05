package widgets

import (
	"fmt"
	"math"
	"strings"
)

type ProgressBar struct {
	progress float32
	min      float32
	max      float32
	text     string
	prefix   string
}

func NewProgressBar(min, max float32) *ProgressBar {
	return &ProgressBar{
		progress: 0,
		min:      min,
		max:      max,
	}
}

func (p *ProgressBar) WithText(text string) *ProgressBar {
	p.text = text
	return p
}

func (p *ProgressBar) WithPrefix(prefix string) *ProgressBar {
	p.prefix = prefix
	return p
}

func (p *ProgressBar) Set(progress float32) *ProgressBar {
	p.progress = progress
	return p
}

func (p *ProgressBar) Render() any {
	const maxLen = 25

	var charset = []rune("⣀⣄⣤⣦⣶⣷⣿")

	length := maxLen * ((p.progress - p.min) / (p.max - p.min))
	if length < 0 {
		length = 0
	}

	_, frac := math.Modf(float64(length))
	char := charset[int(frac*float64(len(charset)))]

	full := strings.Repeat(string(charset[len(charset)-1]), int(length))
	empty := strings.Repeat(string(charset[0]), maxLen-int(length)-1)
	percentage := 100 * float32(p.progress-p.min) / float32(p.max-p.min)

	return fmt.Sprintf("%s %s%c%s %.1f%%\t%s",
		p.prefix,
		full, char, empty,
		percentage,
		p.text)
}
