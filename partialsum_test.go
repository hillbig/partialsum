package partialsum

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPartialSum(t *testing.T) {
	ps := New()
	Convey("When an empty PartialSum is examined", t, func() {
		So(ps.Num(), ShouldEqual, 1)
		So(ps.AllSum(), ShouldEqual, 0)
	})

	ps.IncTail(0, 5)
	ps.IncTail(2, 4)
	ps.IncTail(2, 2)
	ps.IncTail(3, 3)
	// V[0]=5, V[1]=0, V[2] = 6, V[3] = 3
	Convey("When Partial Sum is examined", t, func() {
		So(ps.Num(), ShouldEqual, 4)
		So(ps.AllSum(), ShouldEqual, 14)
		So(ps.Lookup(0), ShouldEqual, 5)
		So(ps.Lookup(1), ShouldEqual, 0)
		So(ps.Lookup(2), ShouldEqual, 6)
		So(ps.Lookup(3), ShouldEqual, 3)
		So(ps.Sum(0), ShouldEqual, 0)
		So(ps.Sum(1), ShouldEqual, 5)
		So(ps.Sum(2), ShouldEqual, 5)
		So(ps.Sum(3), ShouldEqual, 11)
		ind, offset := ps.Find(0)
		So(ind, ShouldEqual, 0)
		So(offset, ShouldEqual, 0)
		ind, offset = ps.Find(4)
		So(ind, ShouldEqual, 0)
		So(offset, ShouldEqual, 4)
		ind, offset = ps.Find(5)
		So(ind, ShouldEqual, 2)
		So(offset, ShouldEqual, 0)
		ind, offset = ps.Find(6)
		So(ind, ShouldEqual, 2)
		So(offset, ShouldEqual, 1)
	})
}
