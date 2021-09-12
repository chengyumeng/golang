package bug

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)


// 要点，切片的扩容，就是当切片添加元素时，切片容量不够了，就会扩容，扩容的大小遵循下面的原则：
// 如果切片的容量小于 1024 个元素，那么扩容的时候 slice 的 cap 就翻番，乘以 2；
// 如果扩容之后，还没有触及原数组的容量，那么，切片中的指针指向的位置，就还是原数组（这就是产生 bug 的原因）；
// 如果扩容之后，超过了原数组的容量，那么，Go 就会开辟一块新的内存，把原来的值拷贝过来，这种情况丝毫不会影响到原数组。
// 存在两个例外，1、如果数组的容量超过了 1024，扩容超过 1 就拷贝，不涉及上述问题；
// 2、如果切片拷贝了原数组完整的数组项，那么直接会拷贝一份内存，在新地址扩容，与原数组容量超过 1024 的处理类似。
// 此测试在 golang 1.6 下测试验证，网络上也有其他说法，可能涉及到不同的语言版本，建议实现上绕过该问题。
func TestAppendSlice(t *testing.T) {
	Convey("在 1024 个元素以下，切片扩容不足原切片容量一倍，直接在原地扩容", t, func() {
		var arr = [4]int{10, 20, 30, 40}
		slice := arr[0:2]
		So(slice[1],ShouldEqual,20)
		newSlice := append(slice, 50)
		newSlice[1] += 1
		So(slice[1],ShouldEqual,21)
		So(arr[1],ShouldEqual,21)
	})

	Convey("在 1024 个元素以下，切片扩容超过原切片一倍，复制一份新内存扩容", t, func() {
		var arr = [4]int{10, 20, 30, 40}
		slice := arr[0:2]
		So(slice[1],ShouldEqual,20)
		newSlice := append(append(append(slice, 50), 100), 150)
		newSlice[1] += 1
		So(slice[1],ShouldEqual,20)
		So(arr[1],ShouldEqual,20)
	})

	Convey("在 1024 个元素以上，切片扩容 1 个，原地扩容", t, func() {
		var arr = [1026]int{}
		for i:= 0 ;i<1026;i++ {
			arr[i] = i
		}
		slice := arr[0:1025]

		newSlice := append(slice,0)
		So(slice[1],ShouldEqual,1)
		newSlice[1] = 2
		So(slice[1],ShouldEqual,2)
	})

	Convey("在 1024 个元素以上，切片扩容超过 1 个，复制内存扩容", t, func() {
		var arr = [1026]int{}
		for i:= 0 ;i<1026;i++ {
			arr[i] = i
		}
		slice := arr[0:1025]

		newSlice := append(slice,0,1)
		So(slice[1],ShouldEqual,1)
		newSlice[1] = 2
		So(slice[1],ShouldEqual,1)
		So(newSlice[1],ShouldEqual,2)
	})

	Convey("在 1024 个元素以上，复制了完整的切片，扩容即复制内存", t, func() {
		var arr = [1026]int{}
		for i:= 0 ;i<1026;i++ {
			arr[i] = i
		}
		slice := arr[0:1026]

		newSlice := append(slice,0)
		So(slice[1],ShouldEqual,1)
		newSlice[1] = 2
		So(slice[1],ShouldEqual,1)
		So(newSlice[1],ShouldEqual,2)
	})

	Convey("在 1024 个元素以下，复制了完整的数组，扩容即复制内存", t, func() {
		var arr = [100]int{}
		for i:= 0 ;i<100;i++ {
			arr[i] = i
		}
		slice := arr[0:100]

		newSlice := append(slice,0)
		So(slice[1],ShouldEqual,1)
		newSlice[1] = 2
		So(slice[1],ShouldEqual,1)
		So(newSlice[1],ShouldEqual,2)
	})
}
