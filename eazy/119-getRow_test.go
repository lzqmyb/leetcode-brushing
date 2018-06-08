package eazy

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
Given a non-negative index k where k ≤ 33, return the kth index row of the Pascal's triangle.
Note that the row index starts from 0.
Input: 3
Output: [1,3,3,1]
*/
func TestGetRow(t *testing.T) {
	Convey("1 should be [1,1]", t, func() {
		So(GetRow(1), ShouldResemble, []int{1, 1})
	})

	Convey("2 should be [1,2,1]", t, func() {
		So(GetRow(2), ShouldResemble, []int{1, 2, 1})
	})

	Convey("1 should be [1,3,3,1]", t, func() {
		So(GetRow(3), ShouldResemble, []int{1, 3, 3, 1})
	})
	Convey("1 should be [1,1]", t, func() {
		So(getRowBetter(1), ShouldResemble, []int{1, 1})
	})

	Convey("2 should be [1,2,1]", t, func() {
		So(getRowBetter(2), ShouldResemble, []int{1, 2, 1})
	})

	Convey("3 should be [1,3,3,1]", t, func() {
		So(getRowBetter(3), ShouldResemble, []int{1, 3, 3, 1})
	})
}
