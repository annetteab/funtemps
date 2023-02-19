package conv
/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FahrenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/

// Konverterer Farhenheit til Celsius
func FahrenheitToCelsius( value float64) float64 {
		// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
  return ( (value - 32.0) * (5.0/9.0) )
}
// De andre konverteringsfunksjonene implementere her
// ...
func FahrenheitToKelvin (value float64) float64 {
  return ((value - 32.0)* (5.0/9.0) + 273.15)
}
func KelvinToCelsius (value float64) float64 {
  return (value - 273.15)
}
func CelsiusToKelvin (value float64) float64 {
  return (value + 273.15)
}
func CelsiusToFahrenheit (value float64) float64 {
  return (value *(9.0/5.0)+ 32.0)
}
func KelvinToFahrenheit (value float64) float64 {
  return ((value - 273.15) *(9.0/5.0)+ 32.0)
}