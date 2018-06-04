package array

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIsOneBitCharacter(t *testing.T) {
	Convey("IsOneBitCharacter: [1, 0, 0]", t, func() {
		result := isOneBitCharacter([]int{1, 0, 0})
		So(result, ShouldEqual, true)
	})

	Convey("IsOneBitCharacter: [1, 1, 1, 0]", t, func() {
		result := isOneBitCharacter([]int{1, 1, 1, 0})
		So(result, ShouldEqual, false)
	})

	Convey("IsOneBitCharacter: [0]", t, func() {
		result := isOneBitCharacter([]int{0})
		So(result, ShouldEqual, true)
	})

	Convey("IsOneBitCharacter: [0, 0]", t, func() {
		result := isOneBitCharacter([]int{0, 0})
		So(result, ShouldEqual, true)
	})
}
