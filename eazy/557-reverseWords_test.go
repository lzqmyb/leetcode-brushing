package eazy

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestReverseWords(t *testing.T) {
	Convey("reverse words", t, func() {

		in := "Let's take LeetCode contest"
		result := reverseWords(in)
		So(result, ShouldEqual, "s'teL ekat edoCteeL tsetnoc")

	})

}

//import (
//	"testing"
//	. "github.com/smartystreets/goconvey/convey"
//)
//
//func TestReverseWords(t *testing.T) string {
//
//	Convey("SingleNumber", t,
//	result := reverseWords(in)
//
//}
