package pkg

import "fmt"

func RecursiveSum(num int) int {
	if num == 1 {
		return 1
	}
	// all recursion has a base and recursive case, above is base and below is recursive
	// the recursive case has 3 parts - pre, recurse and post
	// pre is, in this case, the num + before the recursive part calling function again
	// post is important, this is the returning of count after recursion - this is done to
	// pop the function calls off the stack so we arent generating and holding lots of memory.

	// pre and recursive
	count := num + RecursiveSum(num-1)

	// post
	fmt.Println(num)
	return count
}
