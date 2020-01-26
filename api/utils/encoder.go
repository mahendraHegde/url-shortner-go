package utils

import (
	"fmt"
	"math"
	"strings"
)

const charSeq = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Base62Encode(n int) string {
	res := ""
	for n > 0 {
		rem := math.Mod(float64(n), float64(62))
		res = string(charSeq[int(rem)]) + res
		n /= 62
	}
	return res
}

func Base62Decode(s string) (int64, error) {
	var res int64
	for _, ch := range []byte(s) {
		idx := strings.IndexByte(charSeq, ch)
		if idx < 0 {
			return 0, fmt.Errorf("unexpected value %c in base62 literal", ch)
		}
		res = 62*res + int64(idx)
	}
	return res, nil
}
