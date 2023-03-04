package atmosphere

import "math"

const REVISED_MAGNUS_BETA = 17.625
const REVISED_MAGNUS_LAMBDA = 243.04

// vaporPressure can be actual or saturated
func CalculateMixingRatio(stationPressure float64, vaporPressure float64) float64 {
	result := 621.97 * vaporPressure / (stationPressure - vaporPressure)
	return result
}

func CalculateDewPoint(temperatureInCelsius float64, relativeHumidityPercentage float64) float64 {

	a := math.Log(relativeHumidityPercentage / 100)
	b := (REVISED_MAGNUS_BETA * temperatureInCelsius) / (REVISED_MAGNUS_LAMBDA + temperatureInCelsius)

	humidityMultiplier := (a + b)

	numerator := REVISED_MAGNUS_LAMBDA * humidityMultiplier
	denominator := REVISED_MAGNUS_BETA - humidityMultiplier

	result := numerator / denominator
	return result
}

func CalculateRelativeHumidity(temperatureInCelsius float64, dewpointInCelsius float64) float64 {

	dewpointResult := math.Exp((REVISED_MAGNUS_BETA * dewpointInCelsius) /
		(REVISED_MAGNUS_LAMBDA + dewpointInCelsius))

	temperatureResult := math.Exp((REVISED_MAGNUS_BETA * temperatureInCelsius) /
		(REVISED_MAGNUS_LAMBDA + temperatureInCelsius))

	result := (dewpointResult / temperatureResult) * 100
	return result
}

// degreesInCelsius can be temperature or dewpoint
func CalculateVaporPressure(degreesInCelsius float64) float64 {
	result := 6.11 * math.Pow(10, 7.5*degreesInCelsius/(237.7+degreesInCelsius))
	return result
}
