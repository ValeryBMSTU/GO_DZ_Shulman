package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ChoiseOperType(operator string, prevOperType *int) error {
	operators := [...]string{"+", "-", "/", "*", "("}
	newPrevOperType := 0
	for index, oper := range operators {
		if oper == operator {
			newPrevOperType = index + 1
			break
		}
	}
	if newPrevOperType == 0 {
		err := errors.New("Unknown operator")
		return err
	} else {
		*prevOperType = newPrevOperType
	}
	return nil
}

func calcOperands(operatorsStack []string, operandsStack []float64) ([]string, []float64, error) {

	operator := operatorsStack[len(operatorsStack)-1]
	operand1 := operandsStack[len(operandsStack)-2]
	operand2 := operandsStack[len(operandsStack)-1]
	result := 0.0

	switch operator {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "/":
		result = operand1 / operand2
	case "*":
		result = operand1 * operand2
	}

	operandsStack = append(operandsStack[:len(operandsStack)-2], result)
	operatorsStack = operatorsStack[:len(operatorsStack)-1]

	return operatorsStack, operandsStack, nil
}

func AddOperator(oper string, operatorsStack []string, operandsStack []float64, curentOperType *int, prevOperType *int) ([]string, []float64, error) {
	var err error = nil
	oper_table := [6][7]string{
		/*		  +    -    /	 *	  (	   )	<-	*/
		/*	->*/ {"1", "1", "1", "1", "1", "?", "X"},
		/*	+ */ {"2", "2", "1", "1", "1", "4", "4"},
		/*	- */ {"2", "2", "1", "1", "1", "4", "4"},
		/*	/ */ {"4", "4", "2", "2", "1", "4", "4"},
		/*	* */ {"4", "4", "2", "2", "1", "4", "4"},
		/*	( */ {"1", "1", "1", "1", "1", "3", "?"},
	}

	switch oper_table[*prevOperType][*curentOperType] {
	case "1": // Заслать операцию в стек
		operatorsStack = append(operatorsStack, oper)
		*prevOperType = *curentOperType + 1
	case "2": // Произвести над 2 операндами операцию из стека
		if operatorsStack, operandsStack, err = calcOperands(operatorsStack, operandsStack); err != nil {
			return operatorsStack, operandsStack, err
		}
		operatorsStack = append(operatorsStack, oper)
		if len(operatorsStack) > 0 {
			if err = ChoiseOperType(operatorsStack[len(operatorsStack)-1], prevOperType); err != nil {
				return operatorsStack, operandsStack, err
			}
		} else {
			*prevOperType = 0
		}
	case "3": // Удалить вархнюю операцию из стека
		operatorsStack = operatorsStack[:len(operatorsStack)-1]
		if len(operatorsStack) > 0 {
			if err = ChoiseOperType(operatorsStack[len(operatorsStack)-1], prevOperType); err != nil {
				return operatorsStack, operandsStack, err
			}
		} else {
			*prevOperType = 0
		}
	case "4": // Произвести над 2 операндами операцию из стека и повторить с тем же входным операндом
		if operatorsStack, operandsStack, err = calcOperands(operatorsStack, operandsStack); err != nil {
			return operatorsStack, operandsStack, err
		}
		if len(operatorsStack) > 0 {
			if err := ChoiseOperType(operatorsStack[len(operatorsStack)-1], prevOperType); err != nil {
				return operatorsStack, operandsStack, err
			}
			if operatorsStack, operandsStack, err = AddOperator(oper, operatorsStack, operandsStack,
				curentOperType, prevOperType); err != nil {
				return operatorsStack, operandsStack, err
			}
		}
	case "x":
	case "?":
		err := errors.New("Incorrect operator")
		return operatorsStack, operandsStack, err

	}
	return operatorsStack, operandsStack, nil
}

func ExpCorrection(expression *string) error {

	spaces := regexp.MustCompile(`[\s\p{Zs}]{1,}`)
	*expression = spaces.ReplaceAllString(*expression, "")

	return nil
}

func Calc(expression string) (float64, error) {

	var err error = nil

	if expression == "" {
		err := errors.New("Expression is empty")
		return 0.0, err
	}
	if err := ExpCorrection(&expression); err != nil {
		return 0.0, err
	}

	isOperFlag := false
	prevOperType := 0
	curentOperType := 0
	operators := [...]string{"+", "-", "/", "*", "(", ")"}
	operatorsStack := []string{}
	operandsStack := []float64{}
	curentNumber := ""

	for _, symbol := range expression {

		for index, oper := range operators {
			if oper == string(symbol) {

				if curentNumber != "" {
					floatNumber, _ := strconv.ParseFloat(strings.TrimSpace(curentNumber), 64)
					operandsStack = append(operandsStack, floatNumber)
					curentNumber = ""
				}

				curentOperType = index
				isOperFlag = true

				if operatorsStack, operandsStack, err = AddOperator(string(symbol), operatorsStack, operandsStack,
					&curentOperType, &prevOperType); err != nil {
					return 0.0, nil
				}

				break
			}
		}

		if isOperFlag != true {

			if _, err := strconv.Atoi(string(symbol)); err != nil {
				if string(symbol) != "." {
					return 0.0, nil
				} else {
					curentNumber += string(symbol)
				}
			} else {
				curentNumber += string(symbol)
			}

		} else {
			isOperFlag = false
		}
	}

	for len(operatorsStack) > 0 {
		if curentNumber != "" {
			floatNumber, _ := strconv.ParseFloat(strings.TrimSpace(curentNumber), 64)
			operandsStack = append(operandsStack, floatNumber)
			curentNumber = ""
		}

		curentOperType = 6
		if operatorsStack, operandsStack, err = AddOperator("<-", operatorsStack, operandsStack,
			&curentOperType, &prevOperType); err != nil {
			return 0.0, nil
		}
	}

	return operandsStack[0], nil
}

func main() {
	args := os.Args[1]

	result, err := Calc(args)
	fmt.Println(result, err)
}
