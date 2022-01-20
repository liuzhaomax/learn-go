/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/1/20 17:26
 * @version     v1.0
 * @filename    AnimalGo1.go
 * @description
 ***************************************************************************/
package main

import (
	"fmt"
	"sync"
	"time"
)

func AnimalGo1() {
	chCat := make(chan bool)
	chDog := make(chan bool)
	chFish := make(chan bool)
	wg := sync.WaitGroup{}
	WorkCat1(chFish, chCat, &wg)
	WorkDog1(chCat, chDog, &wg)
	WorkFish1(chDog, chFish, &wg)
	wg.Wait()
}

func WorkCat1(chPrev chan bool, chCurr chan bool, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for {
			fmt.Println("cat")
			chCurr <- true
			<-chPrev
		}
		wg.Done()
	}()
}

func WorkDog1(chPrev chan bool, chCurr chan bool, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for {
			<-chPrev
			fmt.Println("dog")
			chCurr <- true
		}
		wg.Done()
	}()
}

func WorkFish1(chPrev chan bool, chCurr chan bool, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for {
			<-chPrev
			fmt.Println("fish")
			time.Sleep(time.Second)
			chCurr <- true
		}
		wg.Done()
	}()
}
