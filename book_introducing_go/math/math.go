package math

func Average(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0.0
	}
	total := 0.0
	for _, number := range numbers {
		total += number
	}
	return total / float64(len(numbers))
}
