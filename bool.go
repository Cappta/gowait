package gowait

import (
	"errors"
	"time"
)

var (
	errTimeout = errors.New("The operation has timed out")
)

// AwaitTrue will await until the watcher returns true or throw an errors if it times out
func AwaitTrue(watcher func() bool, timeout time.Duration) (err error) {
	timeoutUnixNano := time.Now().Add(timeout).UnixNano()
	for watcher() == false {
		if time.Now().UnixNano() >= timeoutUnixNano {
			return errTimeout
		}
		time.Sleep(time.Millisecond)
	}
	return
}

// AwaitFalse will await until the watcher returns false or throw an errors if it times out
func AwaitFalse(watcher func() bool, timeout time.Duration) (err error) {
	return AwaitTrue(func() bool {
		return watcher() == false
	}, timeout)
}
