package eazy

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIsUgly(t *testing.T) {
	Convey("6 is true", t, func() {
		So(IsUgly(6), ShouldBeTrue)
	})

	Convey("8 is true", t, func() {
		So(IsUgly(8), ShouldBeTrue)
	})

	Convey("14 is false", t, func() {
		So(IsUgly(14), ShouldBeFalse)
	})
}
