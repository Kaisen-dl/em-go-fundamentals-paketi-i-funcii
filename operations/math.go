package operations 

func Add(a, b float64) float64 { 
	return a + b 
}

func Mul(a, b float64) float64 { 
	return a * b 
}

func isPositive(n float64) bool {  // неэкспортируемая
	return n > 0 
} 