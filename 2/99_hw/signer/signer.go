package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
	"time"
)

// сюда писать код

var mutex = &sync.Mutex{}

func OnlyCrc(in, out chan interface{}) {
	data := (<-in).(string)
	result := DataSignerCrc32(data)
	out <- result
}

func CrcAndM(in, out chan interface{}) {
	data := (<-in).(string)

	mutex.Lock()
	data = DataSignerMd5(data)
	mutex.Unlock()

	result := DataSignerCrc32(data)

	out <- result
}

func DoHash(in, out chan interface{}) {

	data := strconv.Itoa((<-in).(int))

	OnlyCrcChIn := make(chan interface{})
	OnlyCrcChOut := make(chan interface{})

	CrcAndMChIn := make(chan interface{})
	CrcAndMChOut := make(chan interface{})

	go OnlyCrc(OnlyCrcChIn, OnlyCrcChOut)
	go CrcAndM(CrcAndMChIn, CrcAndMChOut)

	OnlyCrcChIn <- data

	CrcAndMChIn <- data

	OnlyCrcRes := (<-OnlyCrcChOut).(string)
	CrcAndMRes := (<-CrcAndMChOut).(string)

	result := OnlyCrcRes + "~" + CrcAndMRes

	fmt.Println(data, "SingleHash result: ", result)

	out <- result

	return
}

func SingleHash(in, out chan interface{}) {

	for {
		data := (<-in).(int)

		dataCh := make(chan interface{})

		go DoHash(dataCh, out)

		dataCh <- data
	}

	return
}

func DoMultiHash(in, out chan interface{}) {
	data := (<-in).(string)

	result := ""

	var dataSignerChIn = make([]chan interface{}, 6)
	for i := 0; i < 6; i++ {
		dataSignerChIn[i] = make(chan interface{}, 1)
	}
	var dataSignerChOut = make([]chan interface{}, 6)
	for i := 0; i < 6; i++ {
		dataSignerChOut[i] = make(chan interface{}, 1)
	}

	for i := 0; i < 6; i++ {
		go OnlyCrc(dataSignerChIn[i], dataSignerChOut[i])
	}

	for i := 0; i < 6; i++ {
		dataSignerChIn[i] <- (strconv.Itoa(i) + data)
	}

	iterRes := ""
	for i := 0; i < 6; i++ {
		iterRes = (<-(dataSignerChOut[i])).(string)
		fmt.Println(data, i, "Multihash iter result: ", iterRes)
		result = result + iterRes
	}

	fmt.Println(data, "MultiHash result: ", result)

	out <- result

	return
}

func MultiHash(in, out chan interface{}) {
	for {
		data := (<-in).(string)

		dataCh := make(chan interface{})

		go DoMultiHash(dataCh, out)

		dataCh <- data
	}
	return
}

func CombineResults(in, out chan interface{}) {

	waiter := (<-in).(*sync.WaitGroup)
	defer waiter.Done()

	sortedData := []string{}

	data := <-in

	sortedData = append(sortedData, data.(string))

	sort.Slice(sortedData, func(i, j int) bool {
		return sortedData[i] < sortedData[j]
	})

	waitTime := time.Duration(30 * time.Millisecond)

LOOP:
	for {
		select {

		case data := <-in:

			sortedData = append(sortedData, data.(string))

			sort.Slice(sortedData, func(i, j int) bool {
				return sortedData[i] < sortedData[j]
			})
		case <-time.After(waitTime):
			break LOOP
		}
	}

	result := ""

	for i := 0; i < len(sortedData); i++ {
		result = result + sortedData[i]
		if i < len(sortedData)-1 {
			result = result + "_"
		}
	}

	fmt.Println(sortedData, "CombineResults result: ", result)

	out <- result
}

func ExecutePipeline(workers ...job) {

	Ch1 := make(chan interface{})
	Ch2 := make(chan interface{})
	wg := &sync.WaitGroup{}

	for i := 0; i < len(workers); i++ {
		go workers[i](Ch1, Ch2)
		if i == 3 {
			wg.Add(1)
			Ch1 <- wg
		}
		Ch1 = Ch2
		Ch2 = make(chan interface{})
	}

	if len(workers) == 5 {
		wg.Wait()

		waitTime := time.Duration(1 * time.Millisecond)

		<-time.After(waitTime)
	} else {

		waitTime := time.Duration(320 * time.Millisecond)

		<-time.After(waitTime)
	}
	return
}

func main() {

}
