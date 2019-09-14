package main

import (
	"math"
	"testing"
)

func TestChoiseOperType1(t *testing.T) {
	var prevOperType int
	operatorsStack := []string{"+"}

	prevOperTypeOK := 1

	if err := ChoiseOperType(operatorsStack[len(operatorsStack)-1], &prevOperType); err != nil {
		t.Errorf("Test failed: %s", err)
	}

	if prevOperType != prevOperTypeOK {
		t.Errorf("Test ReadArgs1 failed, result not match")
	}

}

func TestChoiseOperType2(t *testing.T) {
	var prevOperType int
	operatorsStack := []string{"+", "-", "/"}

	prevOperTypeOK := 3

	if err := ChoiseOperType(operatorsStack[len(operatorsStack)-1], &prevOperType); err != nil {
		t.Errorf("Test failed: %s", err)
	}

	if prevOperType != prevOperTypeOK {
		t.Errorf("Test ReadArgs1 failed, result not match")
	}

}

func TestChoiseOperType3(t *testing.T) {
	var prevOperType int
	operatorsStack := []string{"+", "-", "/", "+", "-", "*", "-"}

	prevOperTypeOK := 2

	if err := ChoiseOperType(operatorsStack[len(operatorsStack)-1], &prevOperType); err != nil {
		t.Errorf("Test failed: %s", err)
	}

	if prevOperType != prevOperTypeOK {
		t.Errorf("Test ReadArgs1 failed, result not match")
	}

}

func TestChoiseOperType4(t *testing.T) {
	var prevOperType int
	operatorsStack := []string{"+", "("}

	prevOperTypeOK := 5

	if err := ChoiseOperType(operatorsStack[len(operatorsStack)-1], &prevOperType); err != nil {
		t.Errorf("Test failed: %s", err)
	}

	if prevOperType != prevOperTypeOK {
		t.Errorf("Test ReadArgs1 failed, result not match")
	}

}

func TestChoiseOperType5(t *testing.T) {
	var prevOperType int
	operatorsStack := []string{"+", ")"}

	if err := ChoiseOperType(operatorsStack[len(operatorsStack)-1], &prevOperType); err == nil {
		t.Errorf("Test failed: err is nil")
	}
}

func TestCalcOperands1(t *testing.T) {
	var err error
	operatorsStack := []string{"+"}
	operandsStack := []float64{1, 2}

	operatorsStackOK := []string{}
	operandsStackOK := []float64{3}

	if operatorsStack, operandsStack, err = CalcOperands(operatorsStack, operandsStack); err != nil {
		t.Errorf("Test failed: %s", err)
	}

	if len(operatorsStack) != len(operatorsStackOK) {
		t.Errorf("Test failed, result not match")
	}
	for i := 0; i < len(operatorsStack); i++ {
		if operatorsStack[i] != operatorsStackOK[i] {
			t.Errorf("Test failed, result not match")
		}
	}

	if len(operandsStack) != len(operandsStackOK) {
		t.Errorf("Test failed, result not match")
	}
	for i := 0; i < len(operandsStack); i++ {
		if operandsStack[i] != operandsStackOK[i] {
			t.Errorf("Test failed, result not match")
		}
	}
}

func TestCalcOperands2(t *testing.T) {
	var err error
	operatorsStack := []string{"+", "-"}
	operandsStack := []float64{1, 2, 1}

	operatorsStackOK := []string{"+"}
	operandsStackOK := []float64{1, 1}

	if operatorsStack, operandsStack, err = CalcOperands(operatorsStack, operandsStack); err != nil {
		t.Errorf("Test failed: %s", err)
	}

	if len(operatorsStack) != len(operatorsStackOK) {
		t.Errorf("Test failed, result not match")
	}
	for i := 0; i < len(operatorsStack); i++ {
		if operatorsStack[i] != operatorsStackOK[i] {
			t.Errorf("Test failed, result not match")
		}
	}

	if len(operandsStack) != len(operandsStackOK) {
		t.Errorf("Test failed, result not match")
	}
	for i := 0; i < len(operandsStack); i++ {
		if operandsStack[i] != operandsStackOK[i] {
			t.Errorf("Test failed, result not match")
		}
	}
}

func TestCalcOperands3(t *testing.T) {
	var err error
	operatorsStack := []string{"+", "-", "/"}
	operandsStack := []float64{1, 2, 1, 10}

	operatorsStackOK := []string{"+", "-"}
	operandsStackOK := []float64{1, 2, 0.1}

	if operatorsStack, operandsStack, err = CalcOperands(operatorsStack, operandsStack); err != nil {
		t.Errorf("Test failed: %s", err)
	}

	if len(operatorsStack) != len(operatorsStackOK) {
		t.Errorf("Test failed, result not match")
	}
	for i := 0; i < len(operatorsStack); i++ {
		if operatorsStack[i] != operatorsStackOK[i] {
			t.Errorf("Test failed, result not match")
		}
	}

	if len(operandsStack) != len(operandsStackOK) {
		t.Errorf("Test failed, result not match")
	}
	for i := 0; i < len(operandsStack); i++ {
		if operandsStack[i] != operandsStackOK[i] {
			t.Errorf("Test failed, result not match")
		}
	}
}

func TestCalcOperands4(t *testing.T) {
	var err error
	operatorsStack := []string{"*"}
	operandsStack := []float64{1, 25, 0.2}

	operatorsStackOK := []string{}
	operandsStackOK := []float64{1, 5}

	if operatorsStack, operandsStack, err = CalcOperands(operatorsStack, operandsStack); err != nil {
		t.Errorf("Test failed: %s", err)
	}

	if len(operatorsStack) != len(operatorsStackOK) {
		t.Errorf("Test failed, result not match")
	}
	for i := 0; i < len(operatorsStack); i++ {
		if operatorsStack[i] != operatorsStackOK[i] {
			t.Errorf("Test failed, result not match")
		}
	}

	if len(operandsStack) != len(operandsStackOK) {
		t.Errorf("Test failed, result not match")
	}
	for i := 0; i < len(operandsStack); i++ {
		if operandsStack[i] != operandsStackOK[i] {
			t.Errorf("Test failed, result not match")
		}
	}
}

func TestCalcOperands5(t *testing.T) {
	var err error
	operatorsStack := []string{"/"}
	operandsStack := []float64{1, 0}

	operatorsStackOK := []string{}
	operandsStackOK := []float64{math.Inf(1)}

	if operatorsStack, operandsStack, err = CalcOperands(operatorsStack, operandsStack); err != nil {
		t.Errorf("Test failed: %s", err)
	}

	if len(operatorsStack) != len(operatorsStackOK) {
		t.Errorf("Test failed, result not match")
	}
	for i := 0; i < len(operatorsStack); i++ {
		if operatorsStack[i] != operatorsStackOK[i] {
			t.Errorf("Test failed, result not match")
		}
	}

	if len(operandsStack) != len(operandsStackOK) {
		t.Errorf("Test failed, result not match")
	}
	for i := 0; i < len(operandsStack); i++ {
		if operandsStack[i] != operandsStackOK[i] {
			t.Errorf("Test failed, result not match")
		}
	}
}

func TestExpCorrection1(t *testing.T) {
	expression := "(  3+ 2) "
	expressionOK := "(3+2)"

	ExpCorrection(&expression)

	if expressionOK != expression {
		t.Errorf("Test failed, result not match")
	}
}

func TestExpCorrection2(t *testing.T) {
	expression := "    "
	expressionOK := ""

	ExpCorrection(&expression)

	if expressionOK != expression {
		t.Errorf("Test failed, result not match")
	}
}

func TestAddOperator1(t *testing.T) {
	var err error
	oper := "+"
	curentOperType := 0
	prevOperType := 0
	operatorsStack := []string{}
	operandsStack := []float64{1}

	operatorsStackOK := []string{"+"}
	operandsStackOK := []float64{1}
	prevOperTypeOK := 1

	if operatorsStack, operandsStack, err = AddOperator(oper, operatorsStack, operandsStack,
		&curentOperType, &prevOperType); err != nil {
		t.Errorf("Test failed: %s", err)
	}

	if prevOperTypeOK != prevOperType {
		t.Errorf("Test failed, result not match")
	}

	if len(operatorsStack) != len(operatorsStackOK) {
		t.Errorf("Test failed, result not match")
	}
	for i := 0; i < len(operatorsStack); i++ {
		if operatorsStack[i] != operatorsStackOK[i] {
			t.Errorf("Test failed, result not match")
		}
	}

	if len(operandsStack) != len(operandsStackOK) {
		t.Errorf("Test failed, result not match")
	}
	for i := 0; i < len(operandsStack); i++ {
		if operandsStack[i] != operandsStackOK[i] {
			t.Errorf("Test failed, result not match")
		}
	}
}

func TestAddOperator2(t *testing.T) {
	var err error
	oper := "("
	curentOperType := 4
	prevOperType := 3
	operatorsStack := []string{"/"}
	operandsStack := []float64{1}

	operatorsStackOK := []string{"/", "("}
	operandsStackOK := []float64{1}
	prevOperTypeOK := 5

	if operatorsStack, operandsStack, err = AddOperator(oper, operatorsStack, operandsStack,
		&curentOperType, &prevOperType); err != nil {
		t.Errorf("Test failed: %s", err)
	}

	if prevOperTypeOK != prevOperType {
		t.Errorf("Test failed, result not match")
	}

	if len(operatorsStack) != len(operatorsStackOK) {
		t.Errorf("Test failed, result not match")
	}
	for i := 0; i < len(operatorsStack); i++ {
		if operatorsStack[i] != operatorsStackOK[i] {
			t.Errorf("Test failed, result not match")
		}
	}

	if len(operandsStack) != len(operandsStackOK) {
		t.Errorf("Test failed, result not match")
	}
	for i := 0; i < len(operandsStack); i++ {
		if operandsStack[i] != operandsStackOK[i] {
			t.Errorf("Test failed, result not match")
		}
	}
}

func TestCalc1(t *testing.T) {
	resultOK := 5.0
	expression := "3 + 2"
	if result, err := Calc(expression); err != nil {
		t.Errorf("Test failed: %s", err)
	} else {
		if resultOK != result {
			t.Errorf("Test failed, result not match")
		}
	}

}

func TestCalc2(t *testing.T) {
	resultOK := 2.5
	expression := "(3 + 2  )/2"
	if result, err := Calc(expression); err != nil {
		t.Errorf("Test failed: %s", err)
	} else {
		if resultOK != result {
			t.Errorf("Test failed, result not match")
		}
	}

}

func TestCalc3(t *testing.T) {
	resultOK := 10.0
	expression := " 4 * (3 + 2  )/2"
	if result, err := Calc(expression); err != nil {
		t.Errorf("Test failed: %s", err)
	} else {
		if resultOK != result {
			t.Errorf("Test failed, result not match")
		}
	}

}

func TestCalc4(t *testing.T) {
	resultOK := 20.0
	expression := "((( 4 * (3 + 2  )/2 )+ 7) + 3 -1 + 1)"
	if result, err := Calc(expression); err != nil {
		t.Errorf("Test failed: %s", err)
	} else {
		if resultOK != result {
			t.Errorf("Test failed, result not match")
		}
	}
}

func TestCalc5(t *testing.T) {
	resultOK := 20.0
	expression := "((( 4 * (3 + 2  )/2 )+ 7) + 3 -1 + 1)"
	if result, err := Calc(expression); err != nil {
		t.Errorf("Test failed: %s", err)
	} else {
		if resultOK != result {
			t.Errorf("Test failed, result not match")
		}
	}
}

func TestCalc6(t *testing.T) {
	resultOK := 9.9
	expression := "5.0 + 4.9"
	if result, err := Calc(expression); err != nil {
		t.Errorf("Test failed: %s", err)
	} else {
		if resultOK != result {
			t.Errorf("Test failed, result not match")
		}
	}
}

func TestCalc7(t *testing.T) {
	expression := "5.0 + P4.9"
	if _, err := Calc(expression); err == nil {
		t.Errorf("Test failed: err is nil")
	}
}

func TestCalc8(t *testing.T) {
	expression := ""
	if _, err := Calc(expression); err == nil {
		t.Errorf("Test failed: err is nil")
	}
}

func TestCalc9(t *testing.T) {
	expression := "(( 5 + 7"
	if _, err := Calc(expression); err == nil {
		t.Errorf("Test failed: err is nil")
	}
}

func TestCalc10(t *testing.T) {
	expression := " 5 + +7"
	if _, err := Calc(expression); err == nil {
		t.Errorf("Test failed: err is nil")
	}
}

func TestCalc11(t *testing.T) {
	expression := " 5 & 7"
	if _, err := Calc(expression); err == nil {
		t.Errorf("Test failed: err is nil")
	}
}

func TestCalc12(t *testing.T) {
	expression := " (5 + 7) | 1"
	if _, err := Calc(expression); err == nil {
		t.Errorf("Test failed: err is nil")
	}
}

func TestCalc13(t *testing.T) {
	resultOK := -1.0
	expression := " (2.5 + (((  (5 +((4  *4)+ 4 )  -20 )/2.0 ))) - 6.00000) "
	if result, err := Calc(expression); err != nil {
		t.Errorf("Test failed: %s", err)
	} else {
		if resultOK != result {
			t.Errorf("Test failed, result not match")
		}
	}
}
