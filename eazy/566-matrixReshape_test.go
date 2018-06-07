package eazy

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMatrixReshape(t *testing.T) {
	Convey("[[1,2],[3,4]] r=1 c=4 should [[1,2,3,4]]", t, func() {
		in := [][]int{{1, 2}, {3, 4}}
		result := matrixReshape(in, 1, 4)
		So(result, ShouldResemble, [][]int{{1, 2, 3, 4}})
	})

}
