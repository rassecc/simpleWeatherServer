package models

import (
	"fmt"
	"strings"
)

type Weather struct {
	CityName   string      `json:"name"`
	Conditions []Condition `json:"weather"`
	Info       `json:"main"`
	Units      string
}

type Condition struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Info struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
}

type temperatures struct {
	current   int
	currMin   int
	currMax   int
	feelsLike int
}

const weatherExplanation = "In %s, the temperature today is a high of %d and a low of %d. It is currently %s at %d°%s."
const defaultConditionsMsg = "Conditions look clear out there."

// Explanation returns entire description of current weather.
func (w *Weather) Explanation() string {
	temps := w.getWeatherTemps()
	conditions := w.getWeatherConditions()

	return fmt.Sprintf(weatherExplanation+conditions, w.CityName, temps.currMax, temps.currMin, getTempFeel(temps.feelsLike), temps.current, w.getUnitSymbol())
}

// getWeatherTemps creates a temperatures obj with converted temps based on provided unit param to easily create weather description.
// If no unit param provided, default is kelvin
func (w *Weather) getWeatherTemps() temperatures {
	var convert ConvertFunc
	switch w.Units {
	case "imperial":
		convert = kelvinToFarenheit
	case "metric":
		convert = kelvinToCelcius
	default:
		convert = defaultKelvin
	}

	return temperatures{current: int(convert(w.Info.Temp)),
		currMin:   int(convert(w.Info.TempMin)),
		currMax:   int(convert(w.Info.TempMax)),
		feelsLike: int(convert(w.Info.FeelsLike))}
}

func (w *Weather) getWeatherConditions() string {
	if len(w.Conditions) < 1 {
		return defaultConditionsMsg
	}

	var sb strings.Builder
	sb.WriteString(" Current weather conditions include: ")
	for i, v := range w.Conditions {
		if i == len(w.Conditions)-1 {
			sb.WriteString(v.Description + ".")
			continue
		}
		sb.WriteString(v.Description + ", ")
	}

	return sb.String()
}

func (w *Weather) getUnitSymbol() string {
	switch w.Units {
	case "imperial":
		return "F"
	case "metric":
		return "C"
	default:
		return "K"
	}
}

func getTempFeel(temp int) string {
	switch {
	case temp > 82:
		return "hot"
	case temp < 68:
		return "cold"
	default:
		return "pretty nice outside"
	}
}

type ConvertFunc func(k float64) float64

func defaultKelvin(k float64) float64 {
	return k
}

func kelvinToFarenheit(k float64) float64 {
	return (k-273.15)*1.8 + 32
}

func kelvinToCelcius(k float64) float64 {
	return k - 273.15
}
