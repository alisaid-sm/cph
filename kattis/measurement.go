package main

import (
	"fmt"
)

type Unit struct {
	name   string
	amount int
}

func main() {
	var A float64
	var B, C, D string

	fmt.Scan(&A, &B, &C, &D)

	// fmt.Println(A, B, C, D)

	aliases := map[string]string{
		"thou":    "th",
		"inch":    "in",
		"foot":    "ft",
		"yard":    "yd",
		"chain":   "ch",
		"furlong": "fur",
		"mile":    "mi",
		"league":  "lea",
	}

	if len([]rune(B)) > 3 {
		B = aliases[B]
	}

	if len([]rune(D)) > 3 {
		D = aliases[D]
	}

	units := map[string]Unit{
		"th":  {name: "", amount: 0},
		"in":  {name: "th", amount: 1000},
		"ft":  {name: "in", amount: 12},
		"yd":  {name: "ft", amount: 3},
		"ch":  {name: "yd", amount: 22},
		"fur": {name: "ch", amount: 10},
		"mi":  {name: "fur", amount: 8},
		"lea": {name: "mi", amount: 3},
	}

	// 43*12 = 504
	// 10*1/(3*8) = 0.41

	multiplication := 1
	var choosenMultiplication string

	unitB := B

	for unitB != D {
		if units[unitB].name != D {
			multiplication *= units[unitB].amount
			unitB = units[unitB].name
		} else {
			multiplication *= units[unitB].amount
			choosenMultiplication = "B"
			break
		}

		if units[unitB].name == "" {
			break
		}
	}

	if choosenMultiplication == "" {
		unitD := D
		multiplication = 1

		for unitD != B {
			if units[unitD].name != B {
				multiplication *= units[unitD].amount
				unitD = units[unitD].name
			} else {
				multiplication *= units[unitD].amount
				choosenMultiplication = "D"
				break
			}

			if units[unitD].name == "" {
				break
			}
		}
	}

	// fmt.Println(multiplication)
	// fmt.Println(choosenMultiplication)

	if choosenMultiplication == "B" {
		// fmt.Printf("%.13f\n", A*float64(multiplication))
		fmt.Println(A * float64(multiplication))
	} else {
		fmt.Printf("%.13f\n", A*(1/float64(multiplication)))
		// fmt.Println(A * (1 / float64(multiplication)))
	}
}
