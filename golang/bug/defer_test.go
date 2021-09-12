package bug

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Defer() (err error){
	defer func() {
		fmt.Printf("%v\n",err) // err 值为 defer0
	}()
	defer func() {
		fmt.Printf("%v\n",err) // err 值为 defer1
		err = fmt.Errorf("defer0")
	}()
	defer func() {
		fmt.Printf("%v\n",err) // err 值为 nil
		err = fmt.Errorf("defer1")
	}()

	return
}

func Defer2() (err error){
	err = fmt.Errorf("begin")
	defer fmt.Printf("%v\n",err) // err 值为 begin
	defer func(err error) {
		fmt.Printf("%v\n",err) // err 值为 begin
		err = fmt.Errorf("defer") // 不会改变外层 err 的实际值
	}(err)

	err = fmt.Errorf("end")
	return
}

// 要点：
// defer 后进先出。
// defer 函数的参数，如果被作为参数传递，那么传递的值是 defer 定义时的参数值；如果被作为值直接使用，那么传递的值是 defer 执行的时候的数值。
func TestDeferFunc(t *testing.T) {
	Convey("当参数被 defer 函数中使用，值为 return 值",t, func() {
		err := Defer()
		So(err,ShouldNotBeNil)
		So(err.Error(),ShouldEqual,"defer0")
	}, )

	Convey("当参数被 defer 函数传入，值为 defer 定义时的值",t, func() {
		err := Defer2()
		So(err,ShouldNotBeNil)
		So(err.Error(),ShouldEqual,"end")
	}, )
}
