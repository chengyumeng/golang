package bug

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type s struct {
	v int
}

type s1 struct {
	v *int
}

// 要点， 在 range 循环中，赋值传递的值永远是准确的，如果赋值传递的指针，那么到最后所有的项目都会是 range 中最后一项的指针
func TestRangeArray(t *testing.T) {
	var (
		one = 0
		two = 1
		three = 2
	)

	Convey("point is right", t, func() {
		var mp = []s1{s1{&one}, s1{&two}, s1{&three}}
		var rt = []*int{}
		for _, v := range mp {
			rt = append(rt, v.v)
		}
		So(rt[one], ShouldEqual, &one)
		So(rt[two], ShouldEqual, &two)
		So(rt[three], ShouldEqual, &three)
	})

	Convey("point will always be the last one", t, func() {
		var mp = []s{s{one}, s{two}, s{three}}
		var rt = []*int{}
		for _, v := range mp {
			rt = append(rt, &v.v) // but here will always be 3
		}
		So(*rt[one], ShouldEqual, three)
		So(*rt[two], ShouldEqual, three)
		So(*rt[three], ShouldEqual, three)
	})

	Convey("value is right", t, func() {
		var mp = []s{s{one}, s{two}, s{three}}
		var rt = []int{}
		for _, v := range mp {
			rt = append(rt, v.v) // but here will always be 3
		}
		So(rt[one], ShouldEqual, one)
		So(rt[two], ShouldEqual, two)
		So(rt[three], ShouldEqual, three)
	})
}
