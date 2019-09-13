package main

import (
	"reflect"
	"testing"
)

// var tstArgs1 = []string["data.txt", "-r", "-f", "-u", "-o", "blabla.txt"]
// var testText1 = `blabla
// colmondo
// migren
// aloha
// 1234
// `

/*
	For ReadArgs tests.
*/

func TestReadArgs1(t *testing.T) {
	var outFile string
	var dataFile string
	var colNumb int
	var flags flagsMap = map[string]bool{
		"-f": false, // Игнорирвоать регистр букв
		"-u": false, // Выводить только первое среди нескольких равных
		"-r": false, // Сортировка по убыванию
		"-n": false, // Сортировка чисел
		"-k": false, // Сортировать по столбцу
		"-o": false, // Вывод в файл
	}

	outFileOK := ""
	dataFileOK := "data.txt"
	colNumbOK := 0
	var flagsOK flagsMap = map[string]bool{
		"-f": false, // Игнорирвоать регистр букв
		"-u": false, // Выводить только первое среди нескольких равных
		"-r": false, // Сортировка по убыванию
		"-n": false, // Сортировка чисел
		"-k": false, // Сортировать по столбцу
		"-o": false, // Вывод в файл
	}

	args := []string{"data.txt"}

	err := ReadArgs(flags, args, &outFile, &dataFile, &colNumb)
	if err != nil {
		t.Errorf("Test ReadArgs1 failed: %s", err)
	}

	var resultOK bool

	if outFile != outFileOK || dataFile != dataFileOK || colNumb != colNumbOK {
		resultOK = false
	} else {
		if reflect.DeepEqual(flags, flagsOK) {
			resultOK = true
		} else {
			resultOK = false
		}
	}
	if resultOK != true {
		t.Errorf("Test ReadArgs1 failed, result not match")
	}
}

func TestReadArgs2(t *testing.T) {
	var outFile string
	var dataFile string
	var colNumb int
	var flags flagsMap = map[string]bool{
		"-f": false, // Игнорирвоать регистр букв
		"-u": false, // Выводить только первое среди нескольких равных
		"-r": false, // Сортировка по убыванию
		"-n": false, // Сортировка чисел
		"-k": false, // Сортировать по столбцу
		"-o": false, // Вывод в файл
	}

	outFileOK := "blabla.txt"
	dataFileOK := "data.txt"
	colNumbOK := 0
	var flagsOK flagsMap = map[string]bool{
		"-f": true,  // Игнорирвоать регистр букв
		"-u": true,  // Выводить только первое среди нескольких равных
		"-r": true,  // Сортировка по убыванию
		"-n": false, // Сортировка чисел
		"-k": false, // Сортировать по столбцу
		"-o": true,  // Вывод в файл
	}

	args := []string{"data.txt", "-r", "-f", "-u", "-o", "blabla.txt"}

	err := ReadArgs(flags, args, &outFile, &dataFile, &colNumb)
	if err != nil {
		t.Errorf("Test ReadArgs1 failed: %s", err)
	}

	var resultOK bool

	if outFile != outFileOK || dataFile != dataFileOK || colNumb != colNumbOK {
		resultOK = false
	} else {
		if reflect.DeepEqual(flags, flagsOK) {
			resultOK = true
		} else {
			resultOK = false
		}
	}
	if resultOK != true {
		t.Errorf("Test ReadArgs1 failed, result not match")
	}
}

func TestReadArgs3(t *testing.T) {
	var outFile string
	var dataFile string
	var colNumb int
	var flags flagsMap = map[string]bool{
		"-f": false, // Игнорирвоать регистр букв
		"-u": false, // Выводить только первое среди нескольких равных
		"-r": false, // Сортировка по убыванию
		"-n": false, // Сортировка чисел
		"-k": false, // Сортировать по столбцу
		"-o": false, // Вывод в файл
	}

	outFileOK := "null.txt"
	dataFileOK := "data2.txt"
	colNumbOK := 0
	var flagsOK flagsMap = map[string]bool{
		"-f": false, // Игнорирвоать регистр букв
		"-u": false, // Выводить только первое среди нескольких равных
		"-r": true,  // Сортировка по убыванию
		"-n": false, // Сортировка чисел
		"-k": false, // Сортировать по столбцу
		"-o": true,  // Вывод в файл
	}

	args := []string{"data2.txt", "-r", "-r", "-r", "-r", "-o", "null.txt"}

	err := ReadArgs(flags, args, &outFile, &dataFile, &colNumb)
	if err != nil {
		t.Errorf("Test ReadArgs1 failed: %s", err)
	}

	var resultOK bool

	if outFile != outFileOK || dataFile != dataFileOK || colNumb != colNumbOK {
		resultOK = false
	} else {
		if reflect.DeepEqual(flags, flagsOK) {
			resultOK = true
		} else {
			resultOK = false
		}
	}
	if resultOK != true {
		t.Errorf("Test ReadArgs1 failed, result not match")
	}
}

func TestReadArgs4(t *testing.T) {
	var outFile string
	var dataFile string
	var colNumb int
	var flags flagsMap = map[string]bool{
		"-f": false, // Игнорирвоать регистр букв
		"-u": false, // Выводить только первое среди нескольких равных
		"-r": false, // Сортировка по убыванию
		"-n": false, // Сортировка чисел
		"-k": false, // Сортировать по столбцу
		"-o": false, // Вывод в файл
	}

	outFileOK := "null.txt"
	dataFileOK := "data2.txt"
	colNumbOK := 3
	var flagsOK flagsMap = map[string]bool{
		"-f": false, // Игнорирвоать регистр букв
		"-u": false, // Выводить только первое среди нескольких равных
		"-r": true,  // Сортировка по убыванию
		"-n": false, // Сортировка чисел
		"-k": true,  // Сортировать по столбцу
		"-o": true,  // Вывод в файл
	}

	args := []string{"data2.txt", "-r", "-k", "4", "-o", "null.txt"}

	err := ReadArgs(flags, args, &outFile, &dataFile, &colNumb)
	if err != nil {
		t.Errorf("Test ReadArgs1 failed: %s", err)
	}

	var resultOK bool

	if outFile != outFileOK || dataFile != dataFileOK || colNumb != colNumbOK {
		resultOK = false
	} else {
		if reflect.DeepEqual(flags, flagsOK) {
			resultOK = true
		} else {
			resultOK = false
		}
	}
	if resultOK != true {
		t.Errorf("Test ReadArgs1 failed, result not match")
	}
}

func TestReadArgs5(t *testing.T) {
	var outFile string
	var dataFile string
	var colNumb int
	var flags flagsMap = map[string]bool{
		"-f": false, // Игнорирвоать регистр букв
		"-u": false, // Выводить только первое среди нескольких равных
		"-r": false, // Сортировка по убыванию
		"-n": false, // Сортировка чисел
		"-k": false, // Сортировать по столбцу
		"-o": false, // Вывод в файл
	}

	args := []string{"data2.txt", "-r", "-n", "-k", "A", "-o", "null.txt"}

	err := ReadArgs(flags, args, &outFile, &dataFile, &colNumb)
	if err == nil {
		t.Errorf("Test ReadArgs5 failed: err is nil")
	}
}

func TestReadArgs6(t *testing.T) {
	var outFile string
	var dataFile string
	var colNumb int
	var flags flagsMap = map[string]bool{
		"-f": false, // Игнорирвоать регистр букв
		"-u": false, // Выводить только первое среди нескольких равных
		"-r": false, // Сортировка по убыванию
		"-n": false, // Сортировка чисел
		"-k": false, // Сортировать по столбцу
		"-o": false, // Вывод в файл
	}

	args := []string{"data2.txt", "data3.txt", "-o", "newFile.txt"}

	err := ReadArgs(flags, args, &outFile, &dataFile, &colNumb)
	if err == nil {
		t.Errorf("Test ReadArgs6 failed: err is nil")
	}
}

func TestLinesCorrection1(t *testing.T) {
	var data Data
	data.Lines = []string{" BAndora   curiosity  34234   2  ", " abrcadabra  some more   text"}

	var dataOK Data
	dataOK.Lines = []string{"BAndora curiosity 34234 2", "abrcadabra some more text"}

	data.Lines = LinesCorrection(data.Lines)

	var resultOK bool

	if reflect.DeepEqual(data.Lines, dataOK.Lines) {
		resultOK = true
	} else {
		resultOK = false
	}
	if resultOK != true {
		t.Errorf("Test ReadArgs1 failed, result not match")
	}
}

func TestLinesCorrection2(t *testing.T) {
	var data Data
	data.Lines = []string{"tik &%$#@&#@$     /p\\ntok", " 1 text"}

	var dataOK Data
	dataOK.Lines = []string{"tik &%$#@&#@$ /p\\ntok", "1 text"}

	data.Lines = LinesCorrection(data.Lines)

	var resultOK bool

	if reflect.DeepEqual(data.Lines, dataOK.Lines) {
		resultOK = true
	} else {
		resultOK = false
	}
	if resultOK != true {
		t.Errorf("Test ReadArgs1 failed, result not match")
	}
}

func TestLinesCorrection3(t *testing.T) {
	var data Data
	data.Lines = []string{"tik &%$#@&#@$  q p    [  ][[]] /p\\ntok", " ", ""}

	var dataOK Data
	dataOK.Lines = []string{"tik &%$#@&#@$ q p [ ][[]] /p\\ntok"}

	data.Lines = LinesCorrection(data.Lines)

	if len(data.Lines) != len(dataOK.Lines) {
		t.Errorf("Test ReadArgs1 failed, result not match")
	}
	for i := 0; i < len(data.Lines); i++ {
		if data.Lines[i] != dataOK.Lines[i] {
			t.Errorf("Test ReadArgs1 failed, result not match")
		}
	}
}

func TestStringSort1(t *testing.T) {
	var data Data
	data.Lines = []string{"afjejf", "fnwofe"}

	var dataOK Data
	dataOK.Lines = []string{"afjejf", "fnwofe"}

	data.Lines = StringSort(data.Lines, false)

	var resultOK bool

	if reflect.DeepEqual(data.Lines, dataOK.Lines) {
		resultOK = true
	} else {
		resultOK = false
	}
	if resultOK != true {
		t.Errorf("Test ReadArgs1 failed, result not match")
	}
}

func TestStringSort2(t *testing.T) {
	var data Data
	data.Lines = []string{"12", "2", "6", "18", "22"}

	var dataOK Data
	dataOK.Lines = []string{"12", "18", "2", "22", "6"}

	data.Lines = StringSort(data.Lines, false)

	var resultOK bool

	if reflect.DeepEqual(data.Lines, dataOK.Lines) {
		resultOK = true
	} else {
		resultOK = false
	}
	if resultOK != true {
		t.Errorf("Test ReadArgs1 failed, result not match")
	}
}

func TestStringSort3(t *testing.T) {
	var data Data
	data.Lines = []string{}

	var dataOK Data
	dataOK.Lines = []string{}

	data.Lines = StringSort(data.Lines, true)

	var resultOK bool

	if reflect.DeepEqual(data.Lines, dataOK.Lines) {
		resultOK = true
	} else {
		resultOK = false
	}
	if resultOK != true {
		t.Errorf("Test ReadArgs1 failed, result not match")
	}
}

func TestColSort1(t *testing.T) {
	var data Data
	data.Lines = []string{"12 b", "2 k", "6 a", "18 b", "22 c"}
	col := 0

	var dataOK Data
	dataOK.Lines = []string{"12 b", "18 b", "2 k", "22 c", "6 a"}

	data.Lines = ColSort(data.Lines, col, true)
	var resultOK bool

	if reflect.DeepEqual(data.Lines, dataOK.Lines) {
		resultOK = true
	} else {
		resultOK = false
	}
	if resultOK != true {
		t.Errorf("Test ReadArgs1 failed, result not match")
	}
}

func TestColSort2(t *testing.T) {
	var data Data
	data.Lines = []string{"12 b", "2 k", "6 a", "18 b", "22 c"}
	col := 1

	var dataOK Data
	dataOK.Lines = []string{"6 a", "12 b", "18 b", "22 c", "2 k"}

	data.Lines = ColSort(data.Lines, col, true)
	var resultOK bool

	if reflect.DeepEqual(data.Lines, dataOK.Lines) {
		resultOK = true
	} else {
		resultOK = false
	}
	if resultOK != true {
		t.Errorf("Test ReadArgs1 failed, result not match")
	}
}

func TestColSort3(t *testing.T) {
	var data Data
	data.Lines = []string{"12 A", "2 A", "6 a", "18 a", "22 A"}
	col := 1

	var dataOK Data
	dataOK.Lines = []string{"12 A", "2 A", "22 A", "6 a", "18 a"}

	data.Lines = ColSort(data.Lines, col, false)
	var resultOK bool

	if reflect.DeepEqual(data.Lines, dataOK.Lines) {
		resultOK = true
	} else {
		resultOK = false
	}
	if resultOK != true {
		t.Errorf("Test ReadArgs1 failed, result not match")
	}
}

func TestColSortInt1(t *testing.T) {
	var data Data
	data.Lines = []string{"12 A", "2 A", "6 a", "18 a", "22 A"}
	col := 0

	var dataOK Data
	dataOK.Lines = []string{"2 A", "6 a", "12 A", "18 a", "22 A"}

	var err error
	data.Lines, err = ColSortInt(data.Lines, col, true)

	if err != nil {
		t.Errorf("Test ReadArgs1 failed: %s", err)
	}

	var resultOK bool

	if reflect.DeepEqual(data.Lines, dataOK.Lines) {
		resultOK = true
	} else {
		resultOK = false
	}
	if resultOK != true {
		t.Errorf("Test ReadArgs1 failed, result not match")
	}
}

func TestColSortInt2(t *testing.T) {
	var data Data
	data.Lines = []string{"12 A 342234", "2 A 222", "6 a 90", "18 a 2000", "22 A 0"}
	col := 2

	var dataOK Data
	dataOK.Lines = []string{"22 A 0", "6 a 90", "2 A 222", "18 a 2000", "12 A 342234"}

	var err error
	data.Lines, err = ColSortInt(data.Lines, col, true)

	if err != nil {
		t.Errorf("Test ReadArgs1 failed: %s", err)
	}

	var resultOK bool

	if reflect.DeepEqual(data.Lines, dataOK.Lines) {
		resultOK = true
	} else {
		resultOK = false
	}
	if resultOK != true {
		t.Errorf("Test ReadArgs1 failed, result not match")
	}
}

func TestColSortInt3(t *testing.T) {
	var data Data
	data.Lines = []string{"12 A 1", "2 A 1", "6 a 1", "18 a 1", "22 A 1"}
	col := 2

	var dataOK Data
	dataOK.Lines = []string{"12 A 1", "2 A 1", "6 a 1", "18 a 1", "22 A 1"}

	var err error
	data.Lines, err = ColSortInt(data.Lines, col, true)

	if err != nil {
		t.Errorf("Test ReadArgs1 failed: %s", err)
	}

	var resultOK bool

	if reflect.DeepEqual(data.Lines, dataOK.Lines) {
		resultOK = true
	} else {
		resultOK = false
	}
	if resultOK != true {
		t.Errorf("Test ReadArgs1 failed, result not match")
	}
}

func TestReverseStrings1(t *testing.T) {
	var data Data
	data.Lines = []string{"a", "b", "c", "d"}

	var dataOK Data
	dataOK.Lines = []string{"d", "c", "b", "a"}

	reverseStrings(data.Lines)

	var resultOK bool

	if reflect.DeepEqual(data.Lines, dataOK.Lines) {
		resultOK = true
	} else {
		resultOK = false
	}
	if resultOK != true {
		t.Errorf("Test ReadArgs1 failed, result not match")
	}
}

func TestReverseStrings2(t *testing.T) {
	var data Data
	data.Lines = []string{"abc", "kaew", "amigo", "c///qrer", "[][][][[]]"}

	var dataOK Data
	dataOK.Lines = []string{"[][][][[]]", "c///qrer", "amigo", "kaew", "abc"}

	reverseStrings(data.Lines)

	var resultOK bool

	if reflect.DeepEqual(data.Lines, dataOK.Lines) {
		resultOK = true
	} else {
		resultOK = false
	}
	if resultOK != true {
		t.Errorf("Test ReadArgs1 failed, result not match")
	}
}

func TestReverseStrings3(t *testing.T) {
	var data Data
	data.Lines = []string{"1"}

	var dataOK Data
	dataOK.Lines = []string{"1"}

	reverseStrings(data.Lines)

	var resultOK bool

	if reflect.DeepEqual(data.Lines, dataOK.Lines) {
		resultOK = true
	} else {
		resultOK = false
	}
	if resultOK != true {
		t.Errorf("Test ReadArgs1 failed, result not match")
	}
}

func TestRemoveDublicates1(t *testing.T) {
	var data Data
	data.Lines = []string{"1", "1", "1"}

	var dataOK Data
	dataOK.Lines = []string{"1"}

	data.Lines = RemoveDublicates(data.Lines)

	if len(data.Lines) != len(dataOK.Lines) {
		t.Errorf("Test ReadArgs1 failed, result not match")
	}
	for i := 0; i < len(data.Lines); i++ {
		if data.Lines[i] != dataOK.Lines[i] {
			t.Errorf("Test ReadArgs1 failed, result not match")
		}
	}
}

func TestRemoveDublicates2(t *testing.T) {
	var data Data
	data.Lines = []string{"1", "2", "3"}

	var dataOK Data
	dataOK.Lines = []string{"1", "2", "3"}

	data.Lines = RemoveDublicates(data.Lines)

	if len(data.Lines) != len(dataOK.Lines) {
		t.Errorf("Test ReadArgs1 failed, result not match")
	}
	for i := 0; i < len(data.Lines); i++ {
		if data.Lines[i] != dataOK.Lines[i] {
			t.Errorf("Test ReadArgs1 failed, result not match")
		}
	}
}

func TestRemoveDublicates3(t *testing.T) {
	var data Data
	data.Lines = []string{"AAAA 4", "BB 6", "CCC 8", "AA 12", "AAA 55", "AAA 55"}

	var dataOK Data
	dataOK.Lines = []string{"AAAA 4", "BB 6", "CCC 8", "AA 12", "AAA 55"}

	data.Lines = RemoveDublicates(data.Lines)

	if len(data.Lines) != len(dataOK.Lines) {
		t.Errorf("Test ReadArgs1 failed, result not match")
	}
	for i := 0; i < len(data.Lines); i++ {
		if data.Lines[i] != dataOK.Lines[i] {
			t.Errorf("Test ReadArgs1 failed, result not match")
		}
	}
}

func TestSortData1(t *testing.T) {
	var data Data
	data.Lines = []string{"AAAA 14", "BB 6", "CCC 14", "AA 22", "AAA 19", "AAA 25"}
	col := 1
	var flags flagsMap = map[string]bool{
		"-f": true,  // Игнорирвоать регистр букв
		"-u": true,  // Выводить только первое среди нескольких равных
		"-r": false, // Сортировка по убыванию
		"-n": false, // Сортировка чисел
		"-k": true,  // Сортировать по столбцу
		"-o": false, // Вывод в файл
	}

	var dataOK Data
	dataOK.Lines = []string{"AAAA 14", "CCC 14", "AAA 19", "AA 22", "AAA 25", "BB 6"}

	data.SortData(col, flags)

	if len(data.Lines) != len(dataOK.Lines) {
		t.Errorf("Test ReadArgs1 failed, result not match")
	}
	for i := 0; i < len(data.Lines); i++ {
		if data.Lines[i] != dataOK.Lines[i] {
			t.Errorf("Test ReadArgs1 failed, result not match")
		}
	}
}

func TestSortData2(t *testing.T) {
	var data Data
	data.Lines = []string{"AAAA 14", "BB 6", "CCC 14", "AA 22", "AAA 19", "AAA 25"}
	col := 0
	var flags flagsMap = map[string]bool{
		"-f": false, // Игнорирвоать регистр букв
		"-u": false, // Выводить только первое среди нескольких равных
		"-r": true,  // Сортировка по убыванию
		"-n": true,  // Сортировка чисел
		"-k": false, // Сортировать по столбцу
		"-o": true,  // Вывод в файл
	}

	var dataOK Data
	dataOK.Lines = []string{"CCC 14", "BB 6", "AAAA 14", "AAA 25", "AAA 19", "AA 22"}

	data.SortData(col, flags)

	if len(data.Lines) != len(dataOK.Lines) {
		t.Errorf("Test ReadArgs1 failed, result not match")
	}
	for i := 0; i < len(data.Lines); i++ {
		if data.Lines[i] != dataOK.Lines[i] {
			t.Errorf("Test ReadArgs1 failed, result not match")
		}
	}
}

func TestSort1(t *testing.T) {
	args := []string{"data.txt", "-o", "null.txt"}

	linesOK := []string{"Apple perametr int pp 26",
		"BOOK perametr int pp 24",
		"Book perametr int pp 21",
		"Go perametr int gg 20",
		"Hauptbahnhof perametr int gg 22",
		"January perametr int gg 25",
		"January perametr int gg 25",
		"Napkin perametr int gg 29",
	}
	outFileOK := "null.txt"

	lines, outFile, err := Sort(args)
	if err != nil {
		t.Errorf("Test failed: %s", err)
	}
	if outFile != outFileOK {
		t.Errorf("Test failed, result not match")
	}

	if len(lines) != len(linesOK) {
		t.Errorf("Test ReadArgs1 failed, result not match")
	}
	for i := 0; i < len(lines); i++ {
		if lines[i] != linesOK[i] {
			t.Errorf("Test ReadArgs1 failed, result not match")
		}
	}
}

// func TestFail(t *testing.T) {
// 	in := bytes.NewBufferString(testFailInput)
// 	out := bytes.NewBuffer(nil)
// 	err := uniq(in, out)
// 	if err == nil {
// 		t.Errorf("Test FAIL failed: expected error")
// 	}
// }
