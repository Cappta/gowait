package gowait

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/Cappta/gofixture"
	. "github.com/smartystreets/goconvey/convey"
)

type Whatever struct {
}

var nilWhatever *Whatever
var whatever = &Whatever{}

func TestNil(t *testing.T) {
	Convey("When awaiting nil and getting only not nil", t, func() {
		err := AwaitNil(func() interface{} {
			return whatever
		}, time.Millisecond)
		Convey("Then error should not be nil", func() {
			So(err, ShouldNotBeNil)
		})
	})
	Convey("When awaiting nil and getting nil instantly", t, func() {
		err := AwaitNil(func() interface{} {
			return nilWhatever
		}, time.Nanosecond)
		Convey("Then error should be nil", func() {
			So(err, ShouldBeNil)
		})
	})
	Convey("When awaiting not nil and getting only nil", t, func() {
		err := AwaitNotNil(func() interface{} {
			return nilWhatever
		}, time.Millisecond)
		Convey("Then error should not be nil", func() {
			So(err, ShouldNotBeNil)
		})
	})
	Convey("When awaiting not nil and getting not nil instantly", t, func() {
		err := AwaitNotNil(func() interface{} {
			return whatever
		}, time.Nanosecond)
		Convey("Then error should be nil", func() {
			So(err, ShouldBeNil)
		})
	})

	seed := time.Now().UTC().UnixNano()
	// Only pass t into top-level Convey calls
	Convey(fmt.Sprintf("Given the random seed %d", seed), t, func() {
		rand.Seed(seed)
		Convey("Given an integer between 1 and 100", func() {
			reps := gofixture.AnyIntBetween(1, 100)
			Convey(fmt.Sprintf("When awaiting nil after returning not nill %d times", reps), func() {
				i := 0
				err := AwaitNil(func() interface{} {
					i++
					if i >= reps {
						return nilWhatever
					}
					return whatever
				}, time.Second)
				Convey("Then error should be nil", func() {
					So(err, ShouldBeNil)
				})
			})
			Convey(fmt.Sprintf("When awaiting not nil after returning nil %d times", reps), func() {
				i := 0
				err := AwaitNotNil(func() interface{} {
					i++
					if i >= reps {
						return whatever
					}
					return nilWhatever
				}, time.Second)
				Convey("Then error should be nil", func() {
					So(err, ShouldBeNil)
				})
			})
		})
	})
}
