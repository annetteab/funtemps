package conv

import (
	"reflect"
	"testing"
	"math"
	
)

/*
*

	Mal for testfunksjoner
	Du skal skrive alle funksjonene basert på denne malen
	For alle konverteringsfunksjonene (tilsammen 6)
	kan du bruke malen som den er (du må selvsagt endre
	funksjonsnavn og testverdier)
*/
func TestFahrenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 56.67},
	}

	for _, tc := range tests {
		got := FahrenheitToCelsius(tc.input)
		if !reflect.DeepEqual(tc.want, math.Round(got * 100)/100) {
			t.Errorf("expected: %v, got: %v", tc.want, math.Round(got * 100)/100)
		}
	}
}

// De andre testfunksjonene implementeres her
// ...
func TestFahrenheitToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 329.82},
	}

	for _, tc := range tests {
		got := FahrenheitToKelvin(tc.input)
		if !reflect.DeepEqual(tc.want, math.Round(got * 100)/100) {
			t.Errorf("expected: %v, got: %v", tc.want, math.Round(got * 100)/100)
		}
	}
}

func TestKelvinToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: -139.15},
	}

	for _, tc := range tests {
		got := KelvinToCelsius(tc.input)
		if !reflect.DeepEqual(tc.want, math.Round(got * 100)/100) {
			t.Errorf("expected: %v, got: %v", tc.want, math.Round(got * 100)/100)
		}
	}
}
func TestCelsiusToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 407.15},
	}

	for _, tc := range tests {
		got := CelsiusToKelvin(tc.input)
		if !reflect.DeepEqual(tc.want, math.Round(got * 100)/100) {
			t.Errorf("expected: %v, got: %v", tc.want, math.Round(got * 100)/100)
		}
	}
}
func TestCelsiusToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 273.2},
	}

	for _, tc := range tests {
		got := CelsiusToFahrenheit(tc.input)
		if !reflect.DeepEqual(tc.want, math.Round(got * 100)/100) {
			t.Errorf("expected: %v, got: %v", tc.want, math.Round(got * 100)/100)
		}
	}
}
func TestKelvinToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: -218.47},
	}

	for _, tc := range tests {
		got := KelvinToFahrenheit(tc.input)
		if !reflect.DeepEqual(tc.want, math.Round(got * 100)/100) {
			t.Errorf("expected: %v, got: %v", tc.want, math.Round(got * 100)/100)
		}
	}
}