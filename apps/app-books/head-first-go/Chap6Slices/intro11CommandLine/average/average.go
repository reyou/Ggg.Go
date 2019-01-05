package average

// GetAverage return average
func GetAverage(numbers []float64) float64 {

	var sum float64
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	return sum / sampleCount

}
