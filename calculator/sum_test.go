package calculator_test

import (
	"test/calculator"
	"testing"
)

var create_error = false

func TestSum(t *testing.T) {

	t.Run("SUCCESS CASE", func(t *testing.T) {
		a, b := 4, 5
		expected := 9

		result := calculator.Sum(a, b)

		if expected != result {
			t.Fatal("SUM ERROR [SUCCESS CASE]")
		}
	})

	t.Run("ERROR CASE", func(t *testing.T) {

		a, b := 4, 5
		expected := 10
		result := calculator.Sum(a, b)
		if expected != result && create_error {
			t.Fatal("SUM ERROR [ERROR CASE]")
		}
	})

}
