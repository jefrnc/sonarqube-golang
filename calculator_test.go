package main

import "testing"

func TestCalculate(t *testing.T) {
	// Prueba de suma
	sumResult := calculate(4, 2, "+")
	if sumResult != 6 {
		t.Errorf("Sum test failed. Expected 6, got %.2f", sumResult)
	}

	// Prueba de resta
	subResult := calculate(4, 2, "-")
	if subResult != 2 {
		t.Errorf("Subtraction test failed. Expected 2, got %.2f", subResult)
	}

	// Prueba de multiplicación
	mulResult := calculate(4, 2, "*")
	if mulResult != 8 {
		t.Errorf("Multiplication test failed. Expected 8, got %.2f", mulResult)
	}

	// Prueba de división
	divResult := calculate(4, 2, "/")
	if divResult != 2 {
		t.Errorf("Division test failed. Expected 2, got %.2f", divResult)
	}

	// Prueba de operador inválido
	invalidResult := calculate(4, 2, "%")
	if invalidResult != 0 {
		t.Errorf("Invalid operator test failed. Expected 0, got %.2f", invalidResult)
	}
}
