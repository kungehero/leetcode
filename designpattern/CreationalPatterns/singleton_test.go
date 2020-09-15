package CreationalPatterns

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSingleton(t *testing.T) {
	Convey("", t, func() {
		/* ins1 := GetInstance()
		ins2 := GetInstance() */
		So(GetInstance(), ShouldNotBeNil)
	})
}
