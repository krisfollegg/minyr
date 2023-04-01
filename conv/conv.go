package conv

//Konverteringer.
func FahrenheitToCelsius(value float64) float64 {
	return (value - 32.0) * (5.0/9.0)
}

func CelsiusToFahrenheit(value float64) float64 {
  return (value * 9 / 5) + 32
}

func KelvinToFahrenheit(value float64) float64 {
  return (value-273.15)*9/5 + 32
}

func FahrenheitToKelvin(value float64) float64 {
  return (value-32)*5/9 + 273.15
}

func CelciusToKelvin(value float64) float64 {
  return (value + 273.15)
}

func KelvinToCelcius(value float64) float64 {
  return (value - 273)
}


