package web

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jonlk/atmosphere-v2/atmosphere"
)

var dewpointRoute = func(ctx *gin.Context) {
	var errors []string

	temperatureInCelsius, err := strconv.ParseFloat(ctx.Query("temperatureInCelsius"), 64)
	if err != nil {
		errors = append(errors, err.Error())
	}

	relativeHumidityPercentage, err := strconv.ParseFloat(ctx.Query("relativeHumidityPercentage"), 64)
	if err != nil {
		errors = append(errors, err.Error())
	}

	if len(errors) > 0 {
		ctx.String(http.StatusBadRequest, strings.Join(errors, " | "))
	} else {
		result := atmosphere.CalculateDewPoint(temperatureInCelsius, relativeHumidityPercentage)
		ctx.String(http.StatusOK, fmt.Sprintf("%.2f", result))
	}
}

var relativeHumidityRoute = func(ctx *gin.Context) {

	var errors []string

	temperatureInCelsius, err := strconv.ParseFloat(ctx.Query("temperatureInCelsius"), 64)
	if err != nil {
		errors = append(errors, err.Error())
	}

	dewpointInCelsius, err := strconv.ParseFloat(ctx.Query("dewpointInCelsius"), 64)
	if err != nil {
		errors = append(errors, err.Error())
	}

	if len(errors) > 0 {
		ctx.String(http.StatusBadRequest, strings.Join(errors, " | "))
	} else {
		result := atmosphere.CalculateRelativeHumidity(temperatureInCelsius, dewpointInCelsius)
		ctx.String(http.StatusOK, fmt.Sprintf("%.2f", result))
	}
}
