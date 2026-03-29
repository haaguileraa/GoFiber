package planarWaveguide 

import (
	"github.com/haaguileraa/GoFiber"
	"testing"
	"math"
)


const coreHalfThicknessMeters_t float64 = 1e-6
const wavelengthMeters_t float64 = 1.55e-6
const coreIndex_t float64 = 1.5
const claddIndex_t float64 = 1.0

const waveNumber_t float64 = 4053667.940116
const normalizedFrequency_t float64 = 4.53213853616
const float64Threshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64Threshold
}


func TestPlanarWaveGuide(t *testing.T){
	tests := []struct {
		name		string
		coreHalfThickness_m float64
		wavelength_m 	float64
		coreIndex 	float64
		cladIndex 	float64
		expectingError 	bool 
	}{
		{
			name: "zero core thickness",
			coreHalfThickness_m : 0,
			wavelength_m : wavelengthMeters_t,
			coreIndex : coreIndex_t,
			cladIndex : claddIndex_t,
			expectingError : true,
		},{
			name: "negative core thickness",
			coreHalfThickness_m : -1e-6,
			wavelength_m : wavelengthMeters_t,
			coreIndex : coreIndex_t,
			cladIndex : claddIndex_t,
			expectingError : true,
		},{
			name: "zero wavelength",
			coreHalfThickness_m : coreHalfThicknessMeters_t,
			wavelength_m : 0,
			coreIndex : coreIndex_t,
			cladIndex : claddIndex_t,
			expectingError : true,
		},{
			name: "equal core and cladding refractive indices",
			coreHalfThickness_m : coreHalfThicknessMeters_t,
			wavelength_m : wavelengthMeters_t,
			coreIndex : claddIndex_t,
			cladIndex : claddIndex_t,
			expectingError : true,
		},{
			name: "lower core than cladding refractive index",
			coreHalfThickness_m : coreHalfThicknessMeters_t,
			wavelength_m : wavelengthMeters_t,
			coreIndex : claddIndex_t,
			cladIndex : coreIndex_t,
			expectingError : true,
		},{
			name: "valid input",
			coreHalfThickness_m : coreHalfThicknessMeters_t,
			wavelength_m : wavelengthMeters_t,
			coreIndex : coreIndex_t,
			cladIndex : claddIndex_t,
			expectingError : false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T){
			pwg, err := planarWaveguide.NewPlanarWaveguide(
				test.coreHalfThickness_m,
				test.wavelength_m,
				test.coreIndex,
				test.cladIndex,
			)
			if test.expectingError{
				if err == nil {
					t.Fatalf("Expected error, got nil")
				}
				if pwg != nil {
					t.Fatalf("Expected nil planar waveguide struct on error")
				}
			} else {
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
				if pwg == nil {
					t.Fatalf("Expected valid planar waveguide struct, got nil")
				}
			}
		})
	}
}

func TestComputations(t *testing.T){
	pwg, _ := planarWaveguide.NewPlanarWaveguide(
		coreHalfThicknessMeters_t,
		wavelengthMeters_t,
		coreIndex_t,
		claddIndex_t)

	k0 := pwg.WaveNumber()
	if !almostEqual(k0*1e-6, waveNumber_t*1e-6) {
		t.Fatalf("Expected wave number, k0 = %f, got %f instead", waveNumber_t, k0)
	}

	V := pwg.NormalizedFrequency()
	if !almostEqual(V, normalizedFrequency_t) {
		t.Fatalf("Expected normalized frequency, V = %f, got %f instead", normalizedFrequency_t, V)
	}
}
