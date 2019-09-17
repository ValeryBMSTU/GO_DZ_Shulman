package main

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		t.Errorf("Test LinesCorrection1 failed, result not match")
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
		t.Errorf("Test LinesCorrection2 failed, result not match")
	}
}

func TestLinesCorrection3(t *testing.T) {
	var data Data
	data.Lines = []string{"tik &%$#@&#@$  q p    [  ][[]] /p\\ntok", " ", ""}

	var dataOK Data
	dataOK.Lines = []string{"tik &%$#@&#@$ q p [ ][[]] /p\\ntok"}

	data.Lines = LinesCorrection(data.Lines)

	assert.Equal(t, data.Lines, dataOK.Lines, "The two slices should be the same.")
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
		t.Errorf("Test StringSort1 failed, result not match")
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
		t.Errorf("Test StringSort2 failed, result not match")
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
		t.Errorf("Test StringSort3 failed, result not match")
	}
}

func TestColSort1(t *testing.T) {
	var data Data
	data.Lines = []string{"12 b", "2 k", "6 a", "18 b", "22 c"}
	col := 1

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
		t.Errorf("Test ColSort1 failed, result not match")
	}
}

func TestColSort2(t *testing.T) {
	var data Data
	data.Lines = []string{"12 b", "2 k", "6 a", "18 b", "22 c"}
	col := 2

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
		t.Errorf("Test ColSort2 failed, result not match")
	}
}

func TestColSort3(t *testing.T) {
	var data Data
	data.Lines = []string{"12 A", "2 A", "6 a", "18 a", "22 A"}
	col := 2

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
		t.Errorf("Test ColSort3 failed, result not match")
	}
}

func TestColSortInt1(t *testing.T) {
	var data Data
	data.Lines = []string{"12 A", "2 A", "6 a", "18 a", "22 A"}
	col := 1

	var dataOK Data
	dataOK.Lines = []string{"2 A", "6 a", "12 A", "18 a", "22 A"}

	var err error
	data.Lines, err = ColSortInt(data.Lines, col, true)

	if err != nil {
		t.Errorf("Test ColSortInt1 failed: %s", err)
	}

	var resultOK bool

	if reflect.DeepEqual(data.Lines, dataOK.Lines) {
		resultOK = true
	} else {
		resultOK = false
	}
	if resultOK != true {
		t.Errorf("Test ColSortInt1 failed, result not match")
	}
}

func TestColSortInt2(t *testing.T) {
	var data Data
	data.Lines = []string{"12 A 342234", "2 A 222", "6 a 90", "18 a 2000", "22 A 0"}
	col := 3

	var dataOK Data
	dataOK.Lines = []string{"22 A 0", "6 a 90", "2 A 222", "18 a 2000", "12 A 342234"}

	var err error
	data.Lines, err = ColSortInt(data.Lines, col, true)

	if err != nil {
		t.Errorf("Test ColSortInt2 failed: %s", err)
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
	col := 3

	var dataOK Data
	dataOK.Lines = []string{"12 A 1", "2 A 1", "6 a 1", "18 a 1", "22 A 1"}

	var err error
	data.Lines, err = ColSortInt(data.Lines, col, true)

	if err != nil {
		t.Errorf("Test ColSortInt3 failed: %s", err)
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
		t.Errorf("Test ReverseStrings1 failed, result not match")
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
		t.Errorf("Test ReverseStrings2 failed, result not match")
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
		t.Errorf("Test ReverseStrings3 failed, result not match")
	}
}

func TestRemoveDublicates1(t *testing.T) {
	var data Data
	data.Lines = []string{"1", "1", "1"}

	var dataOK Data
	dataOK.Lines = []string{"1"}

	data.Lines = RemoveDublicates(data.Lines)

	assert.Equal(t, data.Lines, dataOK.Lines, "The two slices should be the same.")
}

func TestRemoveDublicates2(t *testing.T) {
	var data Data
	data.Lines = []string{"1", "2", "3"}

	var dataOK Data
	dataOK.Lines = []string{"1", "2", "3"}

	data.Lines = RemoveDublicates(data.Lines)

	assert.Equal(t, data.Lines, dataOK.Lines, "The two slices should be the same.")
}

func TestRemoveDublicates3(t *testing.T) {
	var data Data
	data.Lines = []string{"AAAA 4", "BB 6", "CCC 8", "AA 12", "AAA 55", "AAA 55"}

	var dataOK Data
	dataOK.Lines = []string{"AAAA 4", "BB 6", "CCC 8", "AA 12", "AAA 55"}

	data.Lines = RemoveDublicates(data.Lines)

	assert.Equal(t, data.Lines, dataOK.Lines, "The two slices should be the same.")
}

func TestSortData1(t *testing.T) {
	var data Data
	data.Lines = []string{"AAAA 14", "BB 6", "CCC 14", "AA 22", "AAA 19", "AAA 25"}
	flags := Options{
		caseIgnoreF: false,
		uniqueF:     false,
		reverseF:    false,
		numericF:    false,
		keyPos:      2,
		outputFile:  "",
	}

	var dataOK Data
	dataOK.Lines = []string{"AAAA 14", "CCC 14", "AAA 19", "AA 22", "AAA 25", "BB 6"}

	data.SortData(flags)

	assert.Equal(t, data.Lines, dataOK.Lines, "The two slices should be the same.")
}

func TestSortData2(t *testing.T) {
	var data Data
	data.Lines = []string{"AAAA 14", "BB 6", "CCC 14", "AA 22", "AAA 19", "AAA 25"}
	flags := Options{
		caseIgnoreF: false,
		uniqueF:     false,
		reverseF:    true,
		numericF:    false,
		keyPos:      1,
		outputFile:  "",
	}

	var dataOK Data
	dataOK.Lines = []string{"CCC 14", "BB 6", "AAAA 14", "AAA 25", "AAA 19", "AA 22"}

	data.SortData(flags)

	assert.Equal(t, data.Lines, dataOK.Lines, "The two slices should be the same.")
}

func TestSortData3(t *testing.T) {
	var data Data
	data.Lines = []string{"AAAA 14", "BB 6", "CCC 14", "AA 22", "AAA 19", "AAA 25", "AAA 19"}
	flags := Options{
		caseIgnoreF: false,
		uniqueF:     true,
		reverseF:    false,
		numericF:    true,
		keyPos:      2,
		outputFile:  "",
	}

	var dataOK Data
	dataOK.Lines = []string{"BB 6", "AAAA 14", "CCC 14", "AAA 19", "AA 22", "AAA 25"}

	data.SortData(flags)

	assert.Equal(t, data.Lines, dataOK.Lines, "The two slices should be the same.")
}

func TestSort1(t *testing.T) {
	var data Data

	file, err := os.OpenFile("testInput.txt", os.O_RDONLY, 0600)

	if err != nil {
		t.Errorf("Test Sort1 failed: %s", err)
	}

	if err := GetLines(&data, file); err != nil {
		t.Errorf("Test Sort1 failed: %s", err)
	}

	flags := Options{
		caseIgnoreF: false,
		uniqueF:     false,
		reverseF:    false,
		numericF:    false,
		keyPos:      -1,
		outputFile:  "",
	}

	linesOK := []string{"Apple perametr int pp 26",
		"BOOK perametr int pp 24",
		"Book perametr int pp 21",
		"Go perametr int gg 20",
		"Hauptbahnhof perametr int gg 22",
		"January perametr int gg 25",
		"January perametr int gg 25",
		"Napkin perametr int gg 29",
	}

	lines, err := Sort(data, flags)
	if err != nil {
		t.Errorf("Test Sort1 failed: %s", err)
	}

	assert.Equal(t, lines, linesOK, "The two slices should be the same.")
}

func TestMain1(t *testing.T) {
	os.Args = []string{os.Args[0], "-o", "testOutput.txt", "testInput.txt"}
	main()
}
