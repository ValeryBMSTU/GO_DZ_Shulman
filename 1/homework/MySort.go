package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type flagsMap map[string]bool

type DataInterface interface {
	SortData(flagsMap) bool
}

type Data struct {
	Lines []string
}

func (data *Data) SortData(col int, flags flagsMap) error {
	var err error = nil

	if flags["-k"] == false {
		data.Lines = StringSort(data.Lines, flags["-f"])
	} else {
		if flags["-n"] == false {
			data.Lines = ColSort(data.Lines, col, flags["-f"])
		} else {
			if data.Lines, err = ColSortInt(data.Lines, col, flags["-f"]); err != nil {
				return err
			}
		}
	}
	if flags["-r"] {
		reverseStrings(data.Lines)
	}
	if flags["-u"] {
		data.Lines = RemoveDublicates(data.Lines)
	}

	return nil
}

func RemoveDublicates(lines []string) []string {
	for i := 0; i < len(lines)-1; i++ {
		if lines[i] == lines[i+1] {
			lines = append(lines[:i], lines[i+1:]...)
			i = i - 1
		}
	}
	return lines
}

func ColSortInt(lines []string, col int, flagF bool) ([]string, error) {
	var err error = nil
	sort.Slice(lines, func(i, j int) bool {
		var numbI, numbJ int
		numbI, err = strconv.Atoi(strings.Split(lines[i], " ")[col])
		numbJ, err = strconv.Atoi(strings.Split(lines[j], " ")[col])
		return numbI < numbJ
	})
	return lines, err
}

func ColSort(lines []string, col int, flagF bool) []string {
	if flagF == true {
		sort.Slice(lines, func(i, j int) bool {
			return strings.ToUpper(strings.Split(lines[i], " ")[col]) < strings.ToUpper(strings.Split(lines[j], " ")[col])
		})
	} else {
		sort.Slice(lines, func(i, j int) bool { return strings.Split(lines[i], " ")[col] < strings.Split(lines[j], " ")[col] })
	}
	return lines
}

func StringSort(lines []string, flagF bool) []string {
	if flagF == true {
		sort.Slice(lines, func(i, j int) bool { return strings.ToUpper(lines[i]) < strings.ToUpper(lines[j]) })
	} else {
		sort.Slice(lines, func(i, j int) bool { return lines[i] < lines[j] })
	}
	return lines
}

func reverseStrings(lines []string) {
	end := len(lines) - 1
	for i := 0; i < len(lines)/2; i++ {
		lines[i], lines[end-i] = lines[end-i], lines[i]
	}
}

func LinesCorrection(lines []string) []string {

	// if lines[len(lines)-1] == "" {
	// 	lines = lines[:len(lines)-1]
	// }

	spacesOnSide := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	spacesInLine := regexp.MustCompile(`[\s\p{Zs}]{2,}`)

	for i := 0; i < len(lines); i++ {
		lines[i] = spacesOnSide.ReplaceAllString(lines[i], "")
		lines[i] = spacesInLine.ReplaceAllString(lines[i], " ")

		if lines[i] == "" {
			lines = append(lines[:i], lines[i+1:]...)
			i = i - 1
		}
	}
	return lines
}

func GetLines(data *Data, dataFile string) error {

	var err error

	if text, err := ioutil.ReadFile(dataFile); err != nil {
		return err
	} else {
		data.Lines = strings.Split(string(text), "\n")

		if data.Lines[len(data.Lines)-1] == "" {
			data.Lines = data.Lines[:len(data.Lines)-1]
		}
		data.Lines = LinesCorrection(data.Lines)
	}
	return err
}

func OutToFile(lines []string, outFile string) {
	if _, err := os.Create(outFile); err != nil {
		panic(err)
	}
	if file, err := os.OpenFile(outFile, os.O_APPEND|os.O_WRONLY, 0600); err != nil {
		panic(err)
	} else {
		for i := 0; i < len(lines); i++ {
			if _, err = file.WriteString(lines[i] + "\n"); err != nil {
				panic(err)
			}
		}
	}
}

func ReadArgs(flags flagsMap, args []string, outFile *string, dataFile *string, colNumb *int) error {
	for index, value := range args {
		if index > 0 && args[index-1] == "-o" {
			continue
		}
		if index > 0 && args[index-1] == "-k" {
			if number, err := strconv.Atoi(value); err == nil {
				*colNumb = number - 1
				continue
			} else {
				return fmt.Errorf("Col number is not digit")
			}
		}
		switch value {
		case "-f":
			flags["-f"] = true
		case "-u":
			flags["-u"] = true
		case "-r":
			flags["-r"] = true
		case "-n":
			flags["-n"] = true
		case "-k":
			flags["-k"] = true
		case "-o":
			if flags["-o"] != true {
				flags["-o"] = true
				*outFile = args[index+1]
			} else {
				return fmt.Errorf("Dublicated of output file")
			}
		default:
			if *dataFile == "" {
				*dataFile = args[index]
			} else {
				return fmt.Errorf("Dublicated of input file")
			}
		}
	}
	return nil
}

func Sort(args []string) ([]string, string, error) {

	var outFile string
	var dataFile string
	var colNumb int
	var data Data

	var flags flagsMap = map[string]bool{
		"-f": false, // Игнорирвоать регистр букв
		"-u": false, // Выводить только первое среди нескольких равных
		"-r": false, // Сортировка по убыванию
		"-n": false, // Сортировка чисел
		"-k": false, // Сортировать по столбцу
		"-o": false, // Вывод в файл
	}

	if err := ReadArgs(flags, args, &outFile, &dataFile, &colNumb); err != nil {
		return data.Lines, outFile, err
	}

	if err := GetLines(&data, dataFile); err != nil {
		return data.Lines, outFile, err
	}

	if err := data.SortData(colNumb, flags); err != nil {
		return data.Lines, outFile, err
	}

	return data.Lines, outFile, nil
}

func mainPanic() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}

func main() {

	defer mainPanic()

	args := os.Args[1:]
	if len(args) < 1 {
		panic("Count of args is too low")
	}

	lines, outFile, err := Sort(args)

	if err != nil {
		panic(err)
	}

	if outFile != "" {
		OutToFile(lines, outFile)
	} else {
		fmt.Println(lines)
	}
}
