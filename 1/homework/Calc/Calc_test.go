package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChoiseOperType1(t *testing.T) {
	var prevOperType int
	operatorsStack := []string{"+"}

	prevOperTypeOK := 1

	if err := ChoiseOperType(operatorsStack[len(operatorsStack)-1], &prevOperType); err != nil {
		t.Errorf("Test ChoiseOperType1 failed: %s", err)
	}

	if prevOperType != prevOperTypeOK {
		t.Errorf("Test ChoiseOperType1 failed, result not match")
	}

}

func TestChoiseOperType2(t *testing.T) {
	var prevOperType int
	operatorsStack := []string{"+", "-", "/"}

	prevOperTypeOK := 3

	if err := ChoiseOperType(operatorsStack[len(operatorsStack)-1], &prevOperType); err != nil {
		t.Errorf("Test ChoiseOperType2 failed: %s", err)
	}

	if prevOperType != prevOperTypeOK {
		t.Errorf("Test ChoiseOperType2 failed, result not match")
	}

}

func TestChoiseOperType3(t *testing.T) {
	var prevOperType int
	operatorsStack := []string{"+", "-", "/", "+", "-", "*", "-"}

	prevOperTypeOK := 2

	if err := ChoiseOperType(operatorsStack[len(operatorsStack)-1], &prevOperType); err != nil {
		t.Errorf("Test ChoiseOperType3 failed: %s", err)
	}

	if prevOperType != prevOperTypeOK {
		t.Errorf("Test ChoiseOperType3 failed, result not match")
	}

}

func TestChoiseOperType4(t *testing.T) {
	var prevOperType int
	operatorsStack := []string{"+", "("}

	prevOperTypeOK := 5

	if err := ChoiseOperType(operatorsStack[len(operatorsStack)-1], &prevOperType); err != nil {
		t.Errorf("Test ChoiseOperType4 failed: %s", err)
	}

	if prevOperType != prevOperTypeOK {
		t.Errorf("Test ChoiseOperType4 failed, result not match")
	}

}

func TestChoiseOperType5(t *testing.T) {
	var prevOperType int
	operatorsStack := []string{"+", ")"}

	if err := ChoiseOperType(operatorsStack[len(operatorsStack)-1], &prevOperType); err == nil {
		t.Errorf("Test ChoiseOperType5 failed: err is nil")
	}
}

func TestCalcOperands1(t *testing.T) {
	var err error
	operatorsStack := []string{"+"}
	operandsStack := []float64{1, 2}

	operatorsStackOK := []string{}
	operandsStackOK := []float64{3}

	if operatorsStack, operandsStack, err = CalcOperands(operatorsStack, operandsStack); err != nil {
		t.Errorf("Test CalcOperands1 failed: %s", err)
	}

	assert.Equal(t, operatorsStack, operatorsStackOK, "The two slices should be the same.")
	assert.Equal(t, operandsStack, operandsStackOK, "The two slices should be the same.")
}

func TestCalcOperands2(t *testing.T) {
	var err error
	operatorsStack := []string{"+", "-"}
	operandsStack := []float64{1, 2, 1}

	operatorsStackOK := []string{"+"}
	operandsStackOK := []float64{1, 1}

	if operatorsStack, operandsStack, err = CalcOperands(operatorsStack, operandsStack); err != nil {
		t.Errorf("Test CalcOperands2 failed: %s", err)
	}

	assert.Equal(t, operatorsStack, operatorsStackOK, "The two slices should be the same.")
	assert.Equal(t, operandsStack, operandsStackOK, "The two slices should be the same.")
}

func TestCalcOperands3(t *testing.T) {
	var err error
	operatorsStack := []string{"+", "-", "/"}
	operandsStack := []float64{1, 2, 1, 10}

	operatorsStackOK := []string{"+", "-"}
	operandsStackOK := []float64{1, 2, 0.1}

	if operatorsStack, operandsStack, err = CalcOperands(operatorsStack, operandsStack); err != nil {
		t.Errorf("Test CalcOperands3 failed: %s", err)
	}

	assert.Equal(t, operatorsStack, operatorsStackOK, "The two slices should be the same.")
	assert.Equal(t, operandsStack, operandsStackOK, "The two slices should be the same.")
}

func TestCalcOperands4(t *testing.T) {
	var err error
	operatorsStack := []string{"*"}
	operandsStack := []float64{1, 25, 0.2}

	operatorsStackOK := []string{}
	operandsStackOK := []float64{1, 5}

	if operatorsStack, operandsStack, err = CalcOperands(operatorsStack, operandsStack); err != nil {
		t.Errorf("Test CalcOperands4 failed: %s", err)
	}

	assert.Equal(t, operatorsStack, operatorsStackOK, "The two slices should be the same.")
	assert.Equal(t, operandsStack, operandsStackOK, "The two slices should be the same.")
}

func TestCalcOperands5(t *testing.T) {
	var err error
	operatorsStack := []string{"/"}
	operandsStack := []float64{1, 0}

	operatorsStackOK := []string{}
	operandsStackOK := []float64{math.Inf(1)}

	if operatorsStack, operandsStack, err = CalcOperands(operatorsStack, operandsStack); err != nil {
		t.Errorf("Test CalcOperands5 failed: %s", err)
	}
	assert.Equal(t, operatorsStack, operatorsStackOK, "The two slices should be the same.")
	assert.Equal(t, operandsStack, operandsStackOK, "The two slices should be the same.")
}

func TestExpCorrection1(t *testing.T) {
	expression := "(  3+ 2) "
	expressionOK := "(3+2)"

	expression = ExpCorrection(expression)

	if expressionOK != expression {
		t.Errorf("Test ExpCorrection1 failed, result not match")
	}
}

func TestExpCorrection2(t *testing.T) {
	expression := "    "
	expressionOK := ""

	expression = ExpCorrection(expression)

	if expressionOK != expression {
		t.Errorf("Test ExpCorrection2 failed, result not match")
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
		t.Errorf("Test AddOperator1 failed: %s", err)
	}

	if prevOperTypeOK != prevOperType {
		t.Errorf("Test AddOperator1 failed, result not match")
	}

	assert.Equal(t, operatorsStack, operatorsStackOK, "The two slices should be the same.")
	assert.Equal(t, operandsStack, operandsStackOK, "The two slices should be the same.")
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
		t.Errorf("Test AddOperator2 failed: %s", err)
	}

	if prevOperTypeOK != prevOperType {
		t.Errorf("Test AddOperator2 failed, result not match")
	}

	assert.Equal(t, operatorsStack, operatorsStackOK, "The two slices should be the same.")
	assert.Equal(t, operandsStack, operandsStackOK, "The two slices should be the same.")
}

func TestCalc1(t *testing.T) {
	resultOK := 5.0
	expression := "3 + 2"
	if result, err := Calc(expression); err != nil {
		t.Errorf("Test Calc1 failed: %s", err)
	} else {
		if resultOK != result {
			t.Errorf("Test Calc1 failed, result not match")
		}
	}

}

func TestCalc2(t *testing.T) {
	resultOK := 2.5
	expression := "(3 + 2  )/2"
	if result, err := Calc(expression); err != nil {
		t.Errorf("Test Calc2 failed: %s", err)
	} else {
		if resultOK != result {
			t.Errorf("Test Calc2 failed, result not match")
		}
	}

}

func TestCalc3(t *testing.T) {
	resultOK := 10.0
	expression := " 4 * (3 + 2  )/2"
	if result, err := Calc(expression); err != nil {
		t.Errorf("Test Calc3 failed: %s", err)
	} else {
		if resultOK != result {
			t.Errorf("Test Calc3 failed, result not match")
		}
	}

}

func TestCalc4(t *testing.T) {
	resultOK := 20.0
	expression := "((( 4 * (3 + 2  )/2 )+ 7) + 3 -1 + 1)"
	if result, err := Calc(expression); err != nil {
		t.Errorf("Test Calc4 failed: %s", err)
	} else {
		if resultOK != result {
			t.Errorf("Test Calc4 failed, result not match")
		}
	}
}

func TestCalc5(t *testing.T) {
	resultOK := 20.0
	expression := "((( 4 * (3 + 2  )/2 )+ 7) + 3 -1 + 1)"
	if result, err := Calc(expression); err != nil {
		t.Errorf("Test Calc5 failed: %s", err)
	} else {
		if resultOK != result {
			t.Errorf("Test Calc5 failed, result not match")
		}
	}
}

func TestCalc6(t *testing.T) {
	resultOK := 9.9
	expression := "5.0 + 4.9"
	if result, err := Calc(expression); err != nil {
		t.Errorf("Test Calc6 failed: %s", err)
	} else {
		if resultOK != result {
			t.Errorf("Test Calc6 failed, result not match")
		}
	}
}

func TestCalc7(t *testing.T) {
	expression := "5.0 + P4.9"
	if _, err := Calc(expression); err == nil {
		t.Errorf("Test Calc7 failed: err is nil")
	}
}

func TestCalc8(t *testing.T) {
	expression := ""
	if _, err := Calc(expression); err == nil {
		t.Errorf("Test Calc8 failed: err is nil")
	}
}

func TestCalc9(t *testing.T) {
	expression := "(( 5 + 7"
	if _, err := Calc(expression); err == nil {
		t.Errorf("Test Calc9 failed: err is nil")
	}
}

func TestCalc10(t *testing.T) {
	expression := " 5 + +7"
	if _, err := Calc(expression); err == nil {
		t.Errorf("Test Calc10 failed: err is nil")
	}
}

func TestCalc11(t *testing.T) {
	expression := " 5 & 7"
	if _, err := Calc(expression); err == nil {
		t.Errorf("Test Calc11 failed: err is nil")
	}
}

func TestCalc12(t *testing.T) {
	expression := " (5 + 7) | 1"
	if _, err := Calc(expression); err == nil {
		t.Errorf("Test Calc12 failed: err is nil")
	}
}

func TestCalc13(t *testing.T) {
	resultOK := -1.0
	expression := " (2.5 + (((  (5 +((4  *4)+ 4 )  -20 )/2.0 ))) - 6.00000) "
	if result, err := Calc(expression); err != nil {
		t.Errorf("Test Calc13 failed: %s", err)
	} else {
		if resultOK != result {
			t.Errorf("Test Calc13 failed, result not match")
		}
	}
}
