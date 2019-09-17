package main

import (
	"fmt"
	"strconv"
	"sync"
)

// сюда писать код

type SingleHashType struct {
	md5      int
	crc32    int
	crc32Md5 int
}

// var (
// 	crc32 = 1
// 	md5   = 2
// )

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

	OnlyCrcRes := (<-OnlyCrcChOut).(string)
	CrcAndMRes := (<-CrcAndMChOut).(string)

	result := OnlyCrcRes + "~" + CrcAndMRes

	out <- result

	// hashType := (<-in).(int)

	// if hashType == 1 {

	// 	data := strconv.Itoa((<-in).(int))

	// } else if hashType == 2 {

	// 	mu := (<-in).(*sync.Mutex)
	// 	data := strconv.Itoa((<-in).(int))

	// 	mu.Lock()
	// 	data = DataSignerMd5(data)
	// 	mu.Unlock()

	// 	data =

	// }

	// if value, ok := d.(string); !ok {
	// 	if value, ok := d.(int); !ok {
	// 		return
	// 	} else {
	// 		data = string(value)
	// 	}
	// } else {
	// 	data = value
	// }

	// fmt.Println("SingleHash data ", data)

	//result := DataSignerCrc32(data) + "~" + DataSignerCrc32(DataSignerMd5(data))
	// md5Data := DataSignerMd5(data)
	// md5CrcData := DataSignerCrc32(md5Data)
	// crcData := DataSignerCrc32(data)

	// fmt.Println("SingleHash md5(data) ", md5Data)
	// fmt.Println("SingleHash crc32(md5(data)) ", md5CrcData)
	// fmt.Println("SingleHash crc32(data) ", crcData)

	// result := crcData + "~" + md5CrcData

	// fmt.Println("SingleHash result ", result)

	// out <- result

	return
}

func MultiHash(in, out chan interface{}) {
	d := <-in

	data := ""

	if value, ok := d.(string); !ok {
		if value, ok := d.(int); !ok {
			return
		} else {
			data = strconv.Itoa(value)
		}
	} else {
		data = value
	}

	result := ""

	for i := 0; i < 6; i++ {
		iterRes := DataSignerCrc32(strconv.Itoa(i) + data)
		fmt.Println(data, "MultiHash: crc32(th+step1)) ", i, " ", iterRes)
		result += iterRes
	}

	fmt.Println(data, "MultiHash result: ", result)

	out <- result

	return
}

func CombineResults(in, out chan interface{}) {

}

func ExecutePipeline(workers ...job) {

	inputWorker := workers[0]

	SingleHashWorker := workers[1]

	MultiHashWorker := workers[2]
	CombineResultsWorker := workers[3]
	Outputworker := workers[4]

	mu := &sync.Mutex{}

	inputCh := make(chan interface{})
	dataCount := 0

	//Md5ChIn := make(chan interface{})
	//Md5ChOut := make(chan interface{})

	//go DataSignerMd5(Md5ChIn, Md5ChOut)

	CombineChIn := make(chan interface{})
	CombineChOut := make(chan interface{})

	for inputData := range inputCh {

		dataCount++

		// Singl32ChIn := make(chan interface{})
		// Singl5ChIn := make(chan interface{})

		// SinglChOut := make(chan interface{})

		// go SingleHashWorker(Singl32ChIn, Singl32ChOut)
		// go SingleHashWorker(Singl5ChIn, Singl5ChOut)

		// Singl32ChIn <- 1
		// Singl5ChIn <- 2

		// Singl32ChIn <- inputData

		// Singl5ChIn <- mu
		// Singl5ChIn <- inputData

		dataCount++

		SinglChIn := make(chan interface{})
		SinglChOut := make(chan interface{})

		go SingleHashWorker(SinglChIn, SinglChOut)

		go MultiHashWorker(SinglChOut, MultiHashChOut)
	}

	hashes := []string{}

	for i := 0; i < dataCount; i++ {
		hash := <-MultiHashChOut
		hashes = append(hashes, hash)
	}

	go CombineResultsWorker(CombineChIn, CombineChOut)

	CombineChIn <- hashes

	result := <-CombineChOut

	// myWork := work[i]

	// in0 := make(chan interface{})
	// out0 := make(chan interface{})

	// go myWork(in0, out0)

	// data := <-out0

	// in1 := make(chan interface{})
	// out1 := make(chan interface{})

	// go SingleHash(in1, out1)

	// in1 <- data
	// step1 := <-out1

	// go MultiHash(in1, out1)

	// in1 <- step1
	// result := <-out1

	// in0 <- result

	// myWork := work
	// in1 := make(chan interface{})
	// out1 := make(chan interface{})
	// myWork(in1, out1)

}

func main() {

}
