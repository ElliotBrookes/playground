package pkg

import (
	"fmt"
	"time"
)

func BubbleSort() {
	l := [20]int{1, 20, 14, 11, 18, 6, 4, 5, 9, 2, 7, 3, 8, 10, 13, 12, 16, 15, 17, 19}
	end := len(l) - 1

	t := time.Now()

	for end > 2 {

		for i := 0; i < end; i++ {
			v := &l[i]
			n := &l[i+1]

			if *v > *n {
				h := *v
				*v = *n
				*n = h
			}

		}
		end--
	}
	fmt.Println("time: ", time.Since(t))
	fmt.Println("mine: ", l)

	j := [20]int{1, 20, 14, 11, 18, 6, 4, 5, 9, 2, 7, 3, 8, 10, 13, 12, 16, 15, 17, 19}
	t = time.Now()

	for i := 0; i < len(j); i++ {
		for k := 0; k < len(j)-1-i; k++ {
			v := &j[k]
			n := &j[k+1]

			if *v > *n {
				h := *v
				*v = *n
				*n = h
			}
		}
	}
	fmt.Println("time: ", time.Since(t))
	fmt.Println("mine: ", j)

}
