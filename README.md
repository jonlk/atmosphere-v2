# Atmosphere-v2

Simple Api for atmosphere calculations using Magnus Coefficients

Go Gin Gonic routing framework

To run:
- `go run .`

- http://localhost:3000/api/relativehumidity?temperatureInCelsius={t}&dewpointInCelsius={d}

- http://localhost:3000/api/dewpoint?temperatureInCelsius={t}&relativeHumidityPercentage={rh}

To test:
- `go test ./...`