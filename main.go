package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"register/hospital"
)

func main() {

	startTime := hospital.StartRunTime
	fmt.Printf("定时抢号时间：%v\n", startTime.String())
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(hospital.BaseEnv.GoNum)
	fmt.Printf("倒计时执行时间：%v\n", startTime.Sub(time.Now()).String())

	ctx, cancelFunc := context.WithCancel(context.Background())
	time.AfterFunc(startTime.Sub(time.Now()), func() {
		for i := 0; i < hospital.BaseEnv.GoNum; i++ {
			go func(i int) {
				time.Sleep(time.Duration(i*10)*time.Millisecond)
				err := hospital.RunZhongRi(ctx)
				if err != nil {
					fmt.Printf("run res err=%v,%v\n", err.Error(),i)
				}
				waitGroup.Done()
				cancelFunc()
				return
			}(i)
		}
	},
	)

	waitGroup.Wait()
}
