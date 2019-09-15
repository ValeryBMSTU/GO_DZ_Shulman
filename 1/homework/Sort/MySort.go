package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	spacesOnSide = regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	spacesInLine = regexp.MustCompile(`[\s\p{Zs}]{2,}`)
)

type Options struct {
	caseIgnoreF bool
	uniqueF     bool
	reverseF    bool
	numericF    bool
	keyPos      int
	outputFile  string
}

type flagsMap map[string]bool

type DataInterface interface {
	SortData(flagsMap) bool
}

type Data struct {
	Lines []string
}

func (data *Data) SortData(flags Options) error {
	var err error

	if flags.keyPos == -1 {
		data.Lines = StringSort(data.Lines, flags.caseIgnoreF)
	} else {
		if flags.numericF == false {
			data.Lines = ColSort(data.Lines, flags.keyPos, flags.caseIgnoreF)
		} else {
			if data.Lines, err = ColSortInt(data.Lines, flags.keyPos, flags.caseIgnoreF); err != nil {
				return err
			}
		}
	}
	if flags.reverseF {
		reverseStrings(data.Lines)
	}
	if flags.uniqueF {
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
	var err error
	sort.Slice(lines, func(i, j int) bool {
		var numbI, numbJ int
		numbI, err = strconv.Atoi(strings.Split(lines[i], " ")[col-1])
		numbJ, err = strconv.Atoi(strings.Split(lines[j], " ")[col-1])
		return numbI < numbJ
	})
	return lines, err
}

func ColSort(lines []string, col int, flagF bool) []string {
	if flagF == true {
		sort.Slice(lines, func(i, j int) bool {
			return strings.ToUpper(strings.Split(lines[i], " ")[col-1]) < strings.ToUpper(strings.Split(lines[j], " ")[col-1])
		})
	} else {
		sort.Slice(lines, func(i, j int) bool { return strings.Split(lines[i], " ")[col-1] < strings.Split(lines[j], " ")[col-1] })
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

func OutToFile(lines []string, outFile string) error {
	if _, err := os.Create(outFile); err != nil {
		return err
	}
	if file, err := os.OpenFile(outFile, os.O_APPEND|os.O_WRONLY, 0600); err != nil {
		return err
	} else {
		for i := 0; i < len(lines); i++ {
			if _, err = file.WriteString(lines[i] + "\n"); err != nil {
				return err
			}
		}
	}
	return nil
}

func ReadArgs(flags *Options) ([]string, error) {

	flag.BoolVar(&flags.caseIgnoreF, "f", false, "a bool")
	flag.BoolVar(&flags.uniqueF, "u", false, "a bool")
	flag.BoolVar(&flags.reverseF, "r", false, "a bool")
	flag.BoolVar(&flags.numericF, "n", false, "a bool")
	flag.IntVar(&flags.keyPos, "k", -1, "a bool")
	flag.StringVar(&flags.outputFile, "o", "", "a bool")

	flag.Parse()

	return flag.Args(), nil
}

func ChoseOutInputFilesFile(data Data, outOption string, notFlags []string) (Data, error) {

	if len(notFlags) != 1 {
		return data, errors.New("Incorrect input")
	}
	if err := GetLines(&data, notFlags[0]); err != nil {
		return data, err
	}

	return data, nil
}

func Sort(data Data, flags Options) ([]string, error) {

	if err := data.SortData(flags); err != nil {
		return data.Lines, err
	}

	return data.Lines, nil
}

func main() {
	var err error
	var flags Options
	var data Data
	var notFlags []string

	dat := os.Args[:1]

	fmt.Println(dat)

	if notFlags, err = ReadArgs(&flags); err != nil {
		log.Fatal(err)
	}

	if data, err = ChoseOutInputFilesFile(data, flags.outputFile, notFlags); err != nil {
		log.Fatal(err)
	}

	lines, err := Sort(data, flags)

	if err != nil {
		log.Fatal(err)
	}

	if flags.outputFile != "" {
		if err = OutToFile(lines, flags.outputFile); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println(lines)
	}
}
