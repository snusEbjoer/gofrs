package gofrs

import (
	"fmt"
	"slices"
	"testing"
	"time"
)

func TestFilter(t *testing.T) {
	// arr := make([]int, 100000)

	start := time.Now()
	seq := FromSlice([]int{1, 2, 3, 5})
	seq = Drop(seq, 1)

	fmt.Println(len(slices.Collect(seq)))

	fmt.Println(time.Since(start))

	// arr2 := make([]int, 100000)

	// start := time.Now()
	// for i, v := range arr2 {
	// 	arr2[i] = v + 78
	// }

	// fmt.Println(len(arr2))

	// fmt.Println(time.Since(start))
}
