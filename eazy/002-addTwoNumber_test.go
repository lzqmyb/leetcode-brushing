package eazy

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	Convey("SingleNumber", t, func() {
		result := SingleNumber([]int{2, 2, 1})
		So(result, ShouldEqual, 1)
	})
}
