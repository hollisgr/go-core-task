package main

import (
	"fmt"
	"math/rand"
	"time"
)

type MyWaitGroup struct {
	count    int
	routines chan int // тут храним счетчик наших горутин

	wait chan struct{} // канал ожидания, закроется когда активные рутины отработают
	stop chan struct{} // канал остановки, закроется когда получит сигнал на остановку

	// используем пустые структуры, т.к. по факту никакие данные передавать не будем,
	// также это самый дешевый вариант по памяти

	isInit bool // статус структуры, инициализированная или нет
}

func (wg *MyWaitGroup) init() {
	wg.count = 0
	wg.routines = make(chan int)
	wg.wait = make(chan struct{})
	wg.stop = make(chan struct{})
	wg.isInit = true
	go wg.worker()
}

func (wg *MyWaitGroup) Add(n int) {
	if !wg.isInit {
		wg.init()
	}
	wg.routines <- n
	return
}

func (wg *MyWaitGroup) Done() {
	wg.routines <- -1
}

func (wg *MyWaitGroup) Wait() {
	// пытаемся читать данные с канала, прочитанные данные нигде не храним,
	// т.к. единственное, что мы можем с него прочитать - это ивент закрытия канала
	// как только прочитаем данные, мы закроем канал остановки
	<-wg.wait
	close(wg.stop) // при закрытии отправляется специальный ивент в канал, его можно прочитать
}

func (wg *MyWaitGroup) worker() {
	for {
		select {
		case n := <-wg.routines:
			// получили данные в канал, прибавляем n к нашему счетчику рутин
			wg.count += n
			if wg.count <= 0 {
				// активные рутины закончились, отключаем канал ожидания
				close(wg.wait) // при закрытии отправляется специальный ивент в канал, его можно прочитать
			}
		case <-wg.stop:
			// как только канал остановки отключился, прерываем работу воркера,
			// единственное, что мы можем с него прочитать - это ивент закрытия канала
			return
		}
	}
}

func main() {
	params := Input()

	wg := MyWaitGroup{}
	wg.Add(params.RoutinesCount)

	for i := range params.RoutinesCount {
		go func(index int) {
			defer wg.Done()
			for range params.NumberCount {
				n := GetRandInt(params.MinValue, params.MaxValue)
				t := time.Now().Format("15:04:05.000")
				fmt.Printf("goroutine #%d, value: %d, time: %v\n", index, n, t)
				time.Sleep(time.Millisecond * 500)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("All goroutines done")
}

func GetRandInt(min int, max int) int {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	res := min + rng.Intn(max-min+1)
	return res
}

type Params struct {
	RoutinesCount int
	MaxValue      int
	MinValue      int
	NumberCount   int
}

func Input() Params {
	res := Params{}

	fmt.Println("Enter count of routines")
	fmt.Scanf("%d", &res.RoutinesCount)

	fmt.Println("Enter number count")
	fmt.Scanf("%d", &res.NumberCount)

	fmt.Println("Enter min num")
	fmt.Scanf("%d", &res.MinValue)

	fmt.Println("Enter max num")
	fmt.Scanf("%d", &res.MaxValue)

	return res
}
