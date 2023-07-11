package jxutils

import "golang.org/x/sync/singleflight"


var group = &singleflight.Group{}

func DoOnlyOnceAtSameTime(key string, f func() error) error{
	_, err, _ := group.Do(key, func() (any, error) {
		err := f()
		if err != nil {
			return nil, err
		}
		return nil, nil
	})

	return err
}
