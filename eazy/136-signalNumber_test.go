package eazy

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSingleNumberBetter(t *testing.T) {
	Convey("SingleNumberBetter", t, func() {
		result := SingleNumberBetter([]int{1, 2, 2, 3, 1})
		So(result, ShouldEqual, 3)
	})
}
