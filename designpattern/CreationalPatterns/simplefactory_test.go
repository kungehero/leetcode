package CreationalPatterns

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSimpleFactory(t *testing.T) {
	Convey("", t, func() {
		So(simpleFactory(1), ShouldEqual, "你好")
	})
}
