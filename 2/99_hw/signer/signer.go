package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
)

// сюда писать код

type MD5Hasher struct {
	mu sync.Mutex
}

func (m *MD5Hasher) Hash(data string) string {
	m.mu.Lock()
	defer m.mu.Unlock()

	return DataSignerMd5(data)
}

var md5Hasher = MD5Hasher{}

func OnlyCrc(in, out chan interface{}) {
	data := (<-in).(string)
	result := DataSignerCrc32(data)
	out <- result
}

func CrcAndM(in, out chan interface{}) {
	data := (<-in).(string)

	result := DataSignerCrc32(md5Hasher.Hash(data))

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
	wg := &sync.WaitGroup{}
	for val := range in {
		data := val.(int)
		dataCh := make(chan interface{})

		wg.Add(1)
		go func(waiter *sync.WaitGroup, in, out chan interface{}) {
			DoHash(dataCh, out)
			waiter.Done()
		}(wg, dataCh, out)

		dataCh <- data
	}
	wg.Wait()
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
	wg := &sync.WaitGroup{}
	for val := range in {
		data := val.(string)

		dataCh := make(chan interface{})

		wg.Add(1)
		go func(waiter *sync.WaitGroup, in, out chan interface{}) {
			DoMultiHash(dataCh, out)
			waiter.Done()
		}(wg, dataCh, out)

		dataCh <- data
	}
	wg.Wait()
	return
}

func CombineResults(in, out chan interface{}) {

	sortedData := []string{}

	for data := range in {

		sortedData = append(sortedData, data.(string))

		sort.Slice(sortedData, func(i, j int) bool {
			return sortedData[i] < sortedData[j]
		})
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
		wg.Add(1)
		go func(waiter *sync.WaitGroup, i int, in, out chan interface{}) {
			workers[i](in, out)
			close(out)
			waiter.Done()
		}(wg, i, Ch1, Ch2)
		Ch1 = Ch2
		Ch2 = make(chan interface{})
	}

	wg.Wait()
	return
}

func main() {

}
