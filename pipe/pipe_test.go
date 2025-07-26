package pipe

import (
	"fmt"
	"gofrs"
	"testing"
)

func TestPipe(t *testing.T) {
	res := Pipe(
		gofrs.FromSlice([]int{1, 2, 3, 4, 5}),
		Filter(func(i int) bool { return i%2 == 0 }),
		Map(func(i int) int { return i + 78 }),
	)

	for v := range res {
		fmt.Println(v)
	}
}
