package util

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func FormatNumber(n float64) string {
	no := int64(n)
	p := message.NewPrinter(language.Vietnamese)
	return p.Sprintf("%d", no)
}
