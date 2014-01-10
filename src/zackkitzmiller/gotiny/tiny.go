package tiny

import (
	"math"
	"strings"
)

type tiny struct {
	set string
}

func NewTiny(set string) *tiny {
	return &tiny{set}
}

func (t *tiny) To(id int) string {
	hexn := ""
	radix := len(t.set)
	for {
		r := id % radix
		hexn = t.set[r:r+1] + hexn
		id = (id - r) / radix
		if id == 0 {
			break
		}
	}
	return hexn
}

func (t *tiny) From(tinystring string) int {

	radix := len(t.set)
	strlen := len(tinystring)
	n := 0
	for i := 0; i < strlen; i++ {
		n += strings.Index(t.set, tinystring[i:i+1]) * int(math.Pow(float64(radix), float64((strlen-i-1))))
	}
	return n
}
