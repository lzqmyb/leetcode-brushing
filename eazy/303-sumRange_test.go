package eazy

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNumArray_SumRange(t *testing.T) {
	Convey("[]", t, func() {
		in := []int{-2, 0, 3, -5, 2, -1}
		na := NumArray{in}
		So(1, ShouldEqual, na.SumRange(0, 2))
		So(-1, ShouldEqual, na.SumRange(2, 5))
		So(-3, ShouldEqual, na.SumRange(0, 5))
	})
}
