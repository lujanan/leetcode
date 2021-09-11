package algorithm_400

import (
	"math/rand"
	"time"
)

func rand10() int {
	for {
		i := rand7()
		j := rand7()
		idx := (i-1)*7 + j
		if idx <= 40 {
			return 1 + (idx-1)%10
		}
	}
}

func rand7() int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(7) + 1
}
