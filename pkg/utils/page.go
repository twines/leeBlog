package utils

import (
	"math"
)

func Page(total int, limit int) (totalPage int) {
	return int(math.Ceil(float64(total) / float64(limit)))
}
