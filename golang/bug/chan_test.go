package bug

import (
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

// 无缓冲的 chan 一旦被写入，当前 goroutines 会被阻塞
// 当所有 goroutines 被阻塞，会直接 panic，recover 都无效。
func TestChanNoBuffer(t *testing.T) {
	Convey("Print 01234",t, func() {
		ch := make(chan int)
		go func() {
			fmt.Println("1")
			fmt.Println(<-ch)
			fmt.Println("3")
		}()
		fmt.Println("0")
		ch<-2
		fmt.Println("4")
		time.Sleep(time.Second)
	})

	Convey("fatal error: all goroutines are asleep - deadlock!",t, func() {
		ch := make(chan int)
		ch<-2
	})
}

// 带 buffer 的 chan 直到写满之前都不会阻塞
func TestChanWithBuffer(t *testing.T)  {
	Convey("带缓冲，good",t, func() {
		ch := make(chan int,1)
		ch<-2
		So(<-ch,ShouldEqual,2)
	})
}
