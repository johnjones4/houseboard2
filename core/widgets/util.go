package widgets

import (
	"fmt"
	"strings"
	"time"
)

func ageStr(d time.Duration) string {
	ds := d / time.Second
	return ageStrSeconds(int(ds))
}

func ageStrSeconds(ds int) string {
	h := int(ds / 3600)
	m := int(ds % 3600 / 60)
	s := int(ds % 3600 % 60)

	parts := make([]string, 0)

	if h > 0 {
		parts = append(parts, fmt.Sprintf("%dh", h))
	}

	if m > 0 {
		parts = append(parts, fmt.Sprintf("%dm", m))
	}

	if s > 0 {
		parts = append(parts, fmt.Sprintf("%ds", s))
	}

	return strings.Join(parts, " ")
}
