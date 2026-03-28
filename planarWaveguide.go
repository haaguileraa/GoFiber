package planarWaveguide 

import (
	"fmt"
)

type PlanarWaveguide struct {
	coreHalfThicknessMeters float64
	wavelengthMeters		    float64
	coreRefractiveIndex			float64
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


