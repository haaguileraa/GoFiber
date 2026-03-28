package planarWaveguide 

import (
	"fmt"
)

type PlanarWaveguide struct {
	CoreHalfThicknessMeters float64
	WavelengthMeters		    float64
	CoreRefractiveIndex			float64
	CladdingRefractiveIndex float64
}

func NewPlanarWaveguide(coreHalfThicknessMeters, 
												wavelengthMeters, 
												coreRefractiveIndex,
												claddingRefractiveIndex float64) (*PlanarWaveguide, error) {

	if CoreHalfThicknessMeters <= 0 {
		return nil, fmt.Errorf("Core half-thickness must be > 0")
	}
	if WavelengthMeters <= 0 {
		return nil, fmt.Errorf("Wavelength must be > 0")
	}
	if CoreRefractiveIndex <= CladdingRefractiveIndex {
		return nil, fmt.Errorf("Core refractive index must be greater than cladding refractive index")
	}
	
	pwg = &PlanarWaveguide{
		CoreHalfThicknessMeters: coreHalfThicknessMeters,
		WavelengthMeters: wavelengthMeters,
		CoreRefractiveIndex: coreRefractiveIndex,
		CladdingRefractiveIndex: claddingRefractiveIndex
	}
	return pwg, nil
}


