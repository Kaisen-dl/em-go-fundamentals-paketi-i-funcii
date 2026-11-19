package main

import (
	"fmt"
	"github.com/Kaisen-dl/em-go-fundamentals-paketi-i-funcii.git/operations"
)

func main() {
	fmt.Println(operations.Add(3, 4))
	fmt.Println(operations.Mul(3, 4))
	// operations.isPositive(5) не скомпилируется
}