package gowait

import (
	"reflect"
	"time"
)

// AwaitNil will await until the watcher returns a value which is nil or throw an errors if it times out
func AwaitNil(watcher func() interface{}, timeout time.Duration) (err error) {
	return AwaitTrue(func() bool {
		return reflect.ValueOf(watcher()).IsNil()
	}, timeout)
}

// AwaitNotNil will await until the watcher returns a value which is not nil or throw an errors if it times out
func AwaitNotNil(watcher func() interface{}, timeout time.Duration) (err error) {
	return AwaitFalse(func() bool {
		return reflect.ValueOf(watcher()).IsNil()
	}, timeout)
}
