package main

import (
	"fmt"
	"strconv"
)

// Exercise #1

func getCharsWordsLines(text string) (int, int, int) {
	var (
		chars = 0
		words = 1
		lines = 1
	)
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' || text[i] == '\t' || text[i] == '\n' {
			words++
		}
		if text[i] == '\n' {
			lines++
		}
		chars++
	}
	fmt.Printf("Text: '%s'\n", text)
	fmt.Printf("Chars: %d, Words: %d, Lines: %d\n", chars, words, lines)
	return chars, words, lines
}

// Exercise #2

func printPattern(centerLine int) {
	if centerLine%2 == 0 {
		fmt.Println("Please provide an odd number for the center line.")
		return
	}

	n := (centerLine + 1) / 2

	// Print upper part
	for i := 0; i < n; i++ {
		// Print spaces
		for j := 0; j < n-i-1; j++ {
			fmt.Print("  ")
		}
		// Print stars
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	// Print lower part
	for i := n - 2; i >= 0; i-- {
		// Print spaces
		for j := 0; j < n-i-1; j++ {
			fmt.Print("  ")
		}
		// Print stars
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

// Exercise #3

func rotateElems(slice *[]int, moves int, direction bool) []int {
	lastElem := len(*slice) - 1
	temp := 0
	for i := 0; i < moves; i++ {
		if direction == false {
			temp = (*slice)[lastElem]
			for j := lastElem; j > 0; j-- {
				(*slice)[j] = (*slice)[j-1]
			}
			(*slice)[0] = temp
		}
		if direction == true {
			temp = (*slice)[0]
			for j := 0; j < lastElem; j++ {
				(*slice)[j] = (*slice)[j+1]
			}
			(*slice)[lastElem] = temp
		}
	}
	fmt.Println(*slice)
	return *slice
}

// Exercise #4

type Shoe struct {
	Brand string
	Price float32
	Size  float32
}

func NewShoe(brand string, price, size float32) *Shoe {
	if size < 34 || size > 44 {
		fmt.Println("invalid size")
		return nil
	}
	return &Shoe{
		Brand: brand,
		Price: price,
		Size:  size,
	}
}

func sellShoe(inventory map[*Shoe]int, shoe *Shoe, quantity int) {
	if stock, exists := inventory[shoe]; exists {
		if stock > 0 && quantity <= stock {
			inventory[shoe] -= quantity
			return
		}
		fmt.Printf("No %s in %.2f size stock available.\n", shoe.Brand, shoe.Size)
		return
	}
	fmt.Println("Shoe not found in inventory.")
}

func printInventory(inventory map[*Shoe]int) {
	fmt.Println("Inventory:")
	for shoe, stock := range inventory {
		fmt.Printf("Brand: %s, Size: %.2f, Stock: %d\n", shoe.Brand, shoe.Size, stock)
	}
}

// Exercise #5

// TokenType represents the type of token.
type TokenType int

const (
	Number      TokenType = iota // Numeric value
	Operator                     // Arithmetic operator (+, -, *, /)
	Parenthesis                  // Parenthesis symbol (, )
)

// Token represents a lexical token.
type Token struct {
	Type  TokenType
	Value string
}

// tokenize converts an expression string into a list of tokens.
func tokenize(expression string) []Token {
	var tokens []Token
	for i := 0; i < len(expression); i++ {
		switch expression[i] {
		case ' ':
			// Skip spaces
			continue
		case '+', '-', '*', '/':
			// Operators are single characters
			tokens = append(tokens, Token{Type: Operator, Value: string(expression[i])})
		case '(', ')':
			// Parentheses are single characters
			tokens = append(tokens, Token{Type: Parenthesis, Value: string(expression[i])})
		default:
			// Parse numbers and decimal points
			j := i
			for j < len(expression) && (expression[j] >= '0' && expression[j] <= '9' || expression[j] == '.') {
				j++
			}
			tokens = append(tokens, Token{Type: Number, Value: expression[i:j]})
			i = j - 1
		}
	}
	return tokens
}

// isOperator checks if a string is a supported operator.
func isOperator(c string) bool {
	return c == "+" || c == "-" || c == "*" || c == "/"
}

// precedence returns the precedence level of an operator.
func precedence(op string) int {
	if op == "+" || op == "-" {
		return 1
	} else if op == "*" || op == "/" {
		return 2
	}
	return 0
}

// applyOperator applies an operator to two values and stores the result.
func applyOperator(operators *[]string, values *[]float64) {
	operator := (*operators)[len(*operators)-1]
	*operators = (*operators)[:len(*operators)-1]

	b := (*values)[len(*values)-1]
	*values = (*values)[:len(*values)-1]

	a := (*values)[len(*values)-1]
	*values = (*values)[:len(*values)-1]

	var result float64
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	}

	*values = append(*values, result)
}

// evaluateExpression evaluates an expression using the Shunting Yard algorithm.
func evaluateExpression(tokens []Token) (float64, error) {
	var operators []string
	var values []float64

	for _, token := range tokens {
		if token.Type == Number {
			num, err := strconv.ParseFloat(token.Value, 64)
			if err != nil {
				return 0, err
			}
			values = append(values, num)
		} else if token.Type == Operator {
			for len(operators) > 0 && isOperator(operators[len(operators)-1]) && precedence(operators[len(operators)-1]) >= precedence(token.Value) {
				applyOperator(&operators, &values)
			}
			operators = append(operators, token.Value)
		} else if token.Type == Parenthesis {
			if token.Value == "(" {
				operators = append(operators, token.Value)
			} else if token.Value == ")" {
				for len(operators) > 0 && operators[len(operators)-1] != "(" {
					applyOperator(&operators, &values)
				}
				if len(operators) > 0 && operators[len(operators)-1] == "(" {
					operators = operators[:len(operators)-1]
				}
			}
		}
	}

	for len(operators) > 0 {
		applyOperator(&operators, &values)
	}

	return values[0], nil
}

func main() {
	//vExercise #1
	fmt.Println("Exercise #1:")
	getCharsWordsLines("Hello World")
	fmt.Println("###########################################")

	// Exercise #2
	fmt.Println("Exercise #2:")
	printPattern(5)
	fmt.Println("##########################################")

	// Exercise #3
	fmt.Println("Exercise #3:")
	slice := []int{1, 2, 3, 4, 5}
	rotateElems(&slice, 3, true)
	rotateElems(&slice, 6, false)
	fmt.Println("##########################################")

	// Exercise #4
	fmt.Println("Exercise #4:")
	nike := NewShoe("Nike", 100000, 40)
	adidas := NewShoe("Adidas", 80000, 42)
	puma := NewShoe("Puma", 75000, 38)
	reebok := NewShoe("Reebok", 90000, 39)
	vans := NewShoe("Vans", 60000, 41)
	converse := NewShoe("Converse", 85000, 37)
	newBalance := NewShoe("New Balance", 95000, 43)
	jordan := NewShoe("Jordan", 120000, 42)
	underArmour := NewShoe("Under Armour", 70000, 44)
	sketchers := NewShoe("Sketchers", 50000, 35)

	inventory := map[*Shoe]int{
		nike:        10,
		adidas:      5,
		puma:        6,
		reebok:      4,
		vans:        3,
		converse:    9,
		newBalance:  10,
		jordan:      1,
		underArmour: 2,
		sketchers:   10,
	}

	sellShoe(inventory, nike, 2)
	sellShoe(inventory, adidas, 3)
	sellShoe(inventory, puma, 7)
	sellShoe(inventory, jordan, 1)
	sellShoe(inventory, newBalance, 10)
	sellShoe(inventory, sketchers, 5)

	printInventory(inventory)
	fmt.Println("##########################################")

	// Exercise #5
	fmt.Println("Exercise #5:")
	expressions := []string{"3 + (4 * 2) - (5 / 2)", "4 + 2", "( 8 + 9 ) * 8"}
	for _, expression := range expressions {
		tokens := tokenize(expression)
		result, _ := evaluateExpression(tokens)
		fmt.Printf("Expression: %s\n", expression)
		fmt.Printf("Result: %.2f\n", result)
	}
	fmt.Println("##########################################")
}
