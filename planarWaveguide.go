package planarWaveguide 

import (
	"fmt"
	"math"
)

type PlanarWaveguide struct {
	coreHalfThicknessMeters float64
	wavelengthMeters	float64
	coreRefractiveIndex	float64
	claddingRefractiveIndex float64
}

func NewPlanarWaveguide(coreHalfThicknessMeters, 
			wavelengthMeters, 
			coreRefractiveIndex,
			claddingRefractiveIndex float64) (*PlanarWaveguide, error) {

	if coreHalfThicknessMeters <= 0 {
		return nil, fmt.Errorf("Core half-thickness must be > 0")
	}
	if wavelengthMeters <= 0 {
		return nil, fmt.Errorf("Wavelength must be > 0")
	}
	if coreRefractiveIndex <= claddingRefractiveIndex {
		return nil, fmt.Errorf("Core refractive index must be greater than cladding refractive index")
	}
	
	pwg := &PlanarWaveguide{
		coreHalfThicknessMeters: coreHalfThicknessMeters,
		wavelengthMeters: wavelengthMeters,
		coreRefractiveIndex: coreRefractiveIndex,
		claddingRefractiveIndex: claddingRefractiveIndex}
	return pwg, nil
}

func (pwg *PlanarWaveguide) WaveNumber() float64 {
	k0 := 2.0 * math.Pi / pwg.wavelengthMeters // rad/m
	return k0
}

func (pwg *PlanarWaveguide) NormalizedFrequency () float64 {
	k0 := pwg.WaveNumber()
	numericalAperture := math.Pow(pwg.coreRefractiveIndex, 2) - math.Pow(pwg.claddingRefractiveIndex, 2)
	V := k0 * pwg.coreHalfThicknessMeters * math.Sqrt(numericalAperture)
	return V // waveguide parameter
}



