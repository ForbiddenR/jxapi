package jxutils

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGroup(t *testing.T) {
	wg := &sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func (wg *sync.WaitGroup)  {
			defer wg.Done()
			err := DoOnlyOnceAtSameTime("key", func() error {
				fmt.Println("do do do")
				time.Sleep(time.Second)
				return errors.New("have error")
			})
			if err != nil {
				t.Log(err)
			}
		}(wg)
	}
	wg.Wait()
}
