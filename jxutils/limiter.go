package jxutils

import (
	"context"

	"golang.org/x/sync/semaphore"
	"golang.org/x/sync/singleflight"
)

var group = &singleflight.Group{}

var weighted = semaphore.NewWeighted(100)

var pipeline = make(chan struct{}, 100)

func DoOnlyOnceAtSameTime(key string, f func() error) error {
	_, err, _ := group.Do(key, func() (any, error) {
		err := f()
		if err != nil {
			return nil, err
		}
		return nil, nil
	})

	return err
}

func DoWithMaxCurrentNum(ctx context.Context,f func() error) error {
	weighted.Acquire()
}
