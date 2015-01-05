package tiny

import (
	"math"
	"math/rand"
	"strings"
	"time"
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
		n += strings.Index(t.set, tinystring[i:i+1]) *
			int(math.Pow(float64(radix), float64((strlen-i-1))))
	}
	return n
}

func GenerateSet() string {
	rand.Seed(time.Now().Unix())
	list := []rune{}
	for i := 65; i < 122; i++ {
		if i < 91 || i > 96 {
			list = append(list, rune(i))
		}
	}

	for i := 48; i <= 57; i++ {
		list = append(list, rune(i))
	}

	set := make([]rune, len(list))
	perm := rand.Perm(len(list))
	for i, v := range perm {
		set[v] = list[i]
	}

	return string(set)
}
