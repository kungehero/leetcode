package tree

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

//smarty streets  聪明的街道
func TestStringSliceEqual(t *testing.T) {
	Convey("TestStringSliceEqual should return true when a != nil  && b != nil", t, func() {
		a := []string{"hello", "goconvey"}
		b := []string{"hello", "goconvey"}
		So(StringSliceEqual(a, b), ShouldBeTrue)
	})
}
