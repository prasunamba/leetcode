package module1

func convertTemperature(celsius float64) []float64 {
	var result []float64
	kelvin := celsius + 273.15
	fah := celsius*1.80 + 32.00
	result = append(result, kelvin, fah)
	return result

}
