package eazy

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func  TestRotate(t *testing.T){
	Convey("6 is true", t, func() {
		So(rotate([] int {1,2,3,4,5,6,7}, 3), ShouldResemble, []int{5,6,7,1,2,3,4})
	})
	Convey("6 is true", t, func() {
		So(rotate([] int {-1}, 2), ShouldResemble, []int{0, 0, -1})
	})
}

//
// func TestIsUgly(t *testing.T) {
// 	Convey("6 is true", t, func() {
// 		So(IsUgly(6), ShouldBeTrue)
// 	})
//
// 	Convey("8 is true", t, func() {
// 		So(IsUgly(8), ShouldBeTrue)
// 	})
//
// 	Convey("14 is false", t, func() {
// 		So(IsUgly(14), ShouldBeFalse)
// 	})
// }