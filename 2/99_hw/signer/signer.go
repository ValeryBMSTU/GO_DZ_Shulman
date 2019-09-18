package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
	"time"
)

// сюда писать код

func OnlyCrc(in, out chan interface{}) {
	data := (<-in).(string)
	result := DataSignerCrc32(data)
	out <- result
}

func CrcAndM(in, out chan interface{}) {
	mu := (<-in).(*sync.Mutex)
	data := (<-in).(string)

	mu.Lock()
	data = DataSignerMd5(data)
	mu.Unlock()

	result := DataSignerCrc32(data)

	out <- result
}

func SingleHash(in, out chan interface{}) {

	mu := (<-in).(*sync.Mutex)
	data := strconv.Itoa((<-in).(int))

	OnlyCrcChIn := make(chan interface{})
	OnlyCrcChOut := make(chan interface{})

	CrcAndMChIn := make(chan interface{})
	CrcAndMChOut := make(chan interface{})

	go OnlyCrc(OnlyCrcChIn, OnlyCrcChOut)
	go CrcAndM(CrcAndMChIn, CrcAndMChOut)

	OnlyCrcChIn <- data

	CrcAndMChIn <- mu
	CrcAndMChIn <- data

	OnlyCrcRes := (<-OnlyCrcChOut).(string)
	CrcAndMRes := (<-CrcAndMChOut).(string)

	result := OnlyCrcRes + "~" + CrcAndMRes

	fmt.Println(data, "SingleHash result: ", result)

	out <- result

	return
}

func MultiHash(in, out chan interface{}) {
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

func ReeadInput(inputWorker job, out chan interface{}) {

	inputCh := make(chan interface{})
	nilCh := make(chan interface{})

	waitTime := time.Duration(800 * time.Millisecond)

	go inputWorker(nilCh, inputCh)

LOOP:
	for {
		select {
		case data := <-inputCh:
			out <- data
		case <-time.After(waitTime):
			close(out)
			break LOOP
		}
	}

}

func WriteOutput(OutputWorker job, in, out chan interface{}) {
	data := (<-in).(string)

	outCh := make(chan interface{})
	nilCh := make(chan interface{})

	waitTime := time.Duration(1 * time.Millisecond)

	go OutputWorker(outCh, nilCh)

	outCh <- data
LOOP:
	for {
		select {
		case <-time.After(waitTime):
			close(out)
			break LOOP
		}
	}

}

func CombineResults(in, out chan interface{}) {
	data := (<-in).([]string)

	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})

	result := ""

	for i := 0; i < len(data); i++ {
		result = result + data[i]
		if i < len(data)-1 {
			result = result + "_"
		}
	}

	fmt.Println(data, "CombineResults result: ", result)

	out <- result
}

func ExecutePipeline(workers ...job) {

	inputWorker := workers[0]
	SingleHashWorker := workers[1]
	MultiHashWorker := workers[2]
	CombineResultsWorker := workers[3]
	Outputworker := workers[4]

	mu := &sync.Mutex{}

	inputCh := make(chan interface{})
	outCh := make(chan interface{})

	dataCount := 0

	CombineChIn := make(chan interface{})
	CombineChOut := make(chan interface{})

	MultiHashChOut := make(chan interface{})

	go ReeadInput(inputWorker, inputCh)

	for inputData := range inputCh {
		data := inputData.(int)
		dataCount++

		SinglChIn := make(chan interface{})
		SinglChOut := make(chan interface{})

		go SingleHashWorker(SinglChIn, SinglChOut)
		go MultiHashWorker(SinglChOut, MultiHashChOut)

		SinglChIn <- mu
		SinglChIn <- data
	}

	hashes := []string{}

	for i := 0; i < dataCount; i++ {
		hash := (<-MultiHashChOut).(string)
		hashes = append(hashes, hash)
	}

	go CombineResultsWorker(CombineChIn, CombineChOut)

	CombineChIn <- hashes

	result := (<-CombineChOut).(string)

	OkCh := make(chan interface{})

	go WriteOutput(Outputworker, outCh, OkCh)

	outCh <- result
	for request := range OkCh {
		fmt.Println(request)
	}

}

func main() {

}
