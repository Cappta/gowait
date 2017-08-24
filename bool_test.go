package gowait

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/Cappta/gofixture"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBool(t *testing.T) {
	Convey("When awaiting true and getting only false", t, func() {
		err := AwaitTrue(func() bool {
			return false
		}, time.Millisecond)
		Convey("Then error should not be nil", func() {
			So(err, ShouldNotBeNil)
		})
	})
	Convey("When awaiting true and getting true instantly", t, func() {
		err := AwaitTrue(func() bool {
			return true
		}, time.Nanosecond)
		Convey("Then error should be nil", func() {
			So(err, ShouldBeNil)
		})
	})
	Convey("When awaiting false and getting only true", t, func() {
		err := AwaitFalse(func() bool {
			return true
		}, time.Millisecond)
		Convey("Then error should not be nil", func() {
			So(err, ShouldNotBeNil)
		})
	})
	Convey("When awaiting false and getting false instantly", t, func() {
		err := AwaitFalse(func() bool {
			return false
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
			Convey(fmt.Sprintf("When awaiting true after returning false %d times", reps), func() {
				i := 0
				err := AwaitTrue(func() bool {
					i++
					if i >= reps {
						return true
					}
					return false
				}, time.Second)
				Convey("Then error should be nil", func() {
					So(err, ShouldBeNil)
				})
			})
			Convey(fmt.Sprintf("When awaiting false after returning true %d times", reps), func() {
				i := 0
				err := AwaitFalse(func() bool {
					i++
					if i >= reps {
						return false
					}
					return true
				}, time.Second)
				Convey("Then error should be nil", func() {
					So(err, ShouldBeNil)
				})
			})
		})
	})
}
