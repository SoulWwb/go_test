package learnconvey

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAdd(t *testing.T) {
	Convey("将两数相加", t, func() {
		So(Add(1, 2), ShouldEqual, 3)
	})
}

func TestSubtract(t *testing.T) {
	Convey("将两数相加", t, func() {
		So(Subtract(1, 2), ShouldEqual, -1)
	})
}

func TestMultiply(t *testing.T) {
	Convey("将两数相加", t, func() {
		So(Multiply(3, 2), ShouldEqual, 6)
	})
}

func TestDivision(t *testing.T) {
	Convey("将两数相加", t, func() {
		Convey("被除数为 0", func() {
			_, err := Division(10, 0)
			So(err, ShouldNotBeNil)
		})

		Convey("被除数不为 0", func() {
			num, err := Division(10, 2)
			So(err, ShouldBeNil)
			So(num, ShouldEqual, 5)
		})
	})
}
