package functions

import (
	"fmt"
)

func BilanganGanjil(N []int) {
	for _, value := range N {
		if value%2 != 0 {
			fmt.Println(value)
		}
	}
}
