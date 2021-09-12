package bug

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// 要点， map 中的数据是无序的
func TestRangeMap(t *testing.T) {
	Convey("", t, func() {
		var mp = map[int]int{}
		for i := 0; i < 10; i++ {
			mp[i] = i
		}

		for k, v := range mp {
			fmt.Println(k, v)
		}
	})
}
