package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type ChanData struct {
	Data []int
	// хотел хранить еще доп.поля для данных, но не пригодилось в итоге, на матрицы уже не стал переписывать
}

type Params struct {
	ChanData []ChanData
	Channels int
}

func main() {
	res := Input()
	channels := make([]chan int, res.Channels)

	for i := range res.Channels {
		channels[i] = make(chan int)
	}

	for i, cd := range res.ChanData {
		go func(i int, cd ChanData) {
			for _, data := range cd.Data {
				channels[i] <- data
			}
			close(channels[i])
		}(i, cd)
	}

	// трансформируем с двухсторонние каналы в каналы чтения,
	// т.к. наш мультиплексор принимает только каналы чтения
	readChannels := make([]<-chan int, len(channels))
	for i, ch := range channels {
		readChannels[i] = ch
	}

	multiplexedChan := Multiplexer(readChannels)
	fmt.Println("Result:")
	for value := range multiplexedChan {
		fmt.Printf("%d ", value)
	}
	fmt.Println()
}

func Multiplexer(channels []<-chan int) <-chan int {
	// создаем результирующий канал
	res := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(channels))

	// на каждый канал запускаем функцию чтения данных
	for _, c := range channels {
		// описываем функцию, которая будет читать данные и записывать
		// данные в наш канал в горутине
		go func(c <-chan int) {
			for n := range c {
				res <- n
			}
			wg.Done()
		}(c)
	}

	// ожидаем когда с каждого канала будут прочтены данные и закрываем канал
	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}

func Input() Params {
	reader := bufio.NewReader(os.Stdin)
	counter := 0
	res := Params{}
	fmt.Println("Enter number of channels")
	fmt.Scanf("%d", &res.Channels)

	res.ChanData = make([]ChanData, 0, res.Channels)

	for len(res.ChanData) < res.Channels {
		fmt.Println("Enter numbers separated with space for channel #", counter+1)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		temp := ChanData{}

		for t := range strings.SplitSeq(input, " ") {
			n, _ := strconv.Atoi(t)
			temp.Data = append(temp.Data, n)
		}
		res.ChanData = append(res.ChanData, temp)

		counter++
	}

	return res
}
