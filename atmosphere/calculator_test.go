package atmosphere

import "testing"

func TestCalculateDewPoint(t *testing.T) {

	var temperatureInCelsius float64 = 9
	var relativeHumidityPercentage float64 = 70

	min := 3.8
	max := 3.9

	result := CalculateDewPoint(temperatureInCelsius, relativeHumidityPercentage)

	if result < min || result > max {
		t.Fail()
	}
}

func TestCalculateRelativeHumidity(t *testing.T) {

	var temperatureInCelsius float64 = 9
	var dewpointInCelsius float64 = 4

	min := 70.0
	max := 71.0

	result := CalculateRelativeHumidity(temperatureInCelsius, dewpointInCelsius)

	if result < min || result > max {
		t.Fail()
	}
}
