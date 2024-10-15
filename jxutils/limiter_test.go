package jxutils

import (
	"context"
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
		go func(wg *sync.WaitGroup) {
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

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
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

func TestWeight(t *testing.T) {
	wg := &sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		k := i
		go DoWithMaxCurrentNum(context.Background(), func() error {
			defer wg.Done()
			fmt.Println(k)
			time.Sleep(2 * time.Second)
			return nil
		})
	}
	wg.Wait()
}
