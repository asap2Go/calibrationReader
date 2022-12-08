package calibrationReader

import (
	"errors"
	"math"

	"github.com/asap2Go/calibrationReader/a2l"
	"github.com/rs/zerolog/log"
)

func convDecToPhy(dec float64, cm *a2l.CompuMethod, cd *CalibrationData) (float64, error) {
	var err error
	switch cm.ConversionType {
	case a2l.Identical:
		return dec, nil
	case a2l.Form:
		//implement if found in prod data.
		err = errors.New("conversion method 'Form' not implemented for compuMethod " + cm.Name)
		log.Err(err).Msg("decimal value could not be converted")
		return 0, err
	case a2l.Linear:
		if !(cm.CoeffsLinear.ASet && cm.CoeffsLinear.BSet) {
			err = errors.New("CoeffsLinear not set in compuMethod: " + cm.Name)
			log.Err(err).Msg("decimal value could not be converted")
			return 0, err
		}
		return cm.CoeffsLinear.A*dec + cm.CoeffsLinear.B, err
	case a2l.RatFunc:
		if !(cm.Coeffs.ASet && cm.Coeffs.BSet && cm.Coeffs.CSet && cm.Coeffs.DSet && cm.Coeffs.ESet && cm.Coeffs.FSet) {
			err = errors.New("Coeffs not set in compuMethod: " + cm.Name)
			log.Err(err).Msg("decimal value could not be converted")
			return 0, err
		}
		phy, err := calcRatFunc(dec, cm)
		if err != nil {
			log.Err(err).Msg("decimal value could not be converted")
			return dec, err
		}
		return phy, err
	case a2l.TabIntp:
		return 0, err
	case a2l.TabNointp:
		phy, err := calcTabNoIntp(dec, cm, cd)
		if err != nil {
			log.Err(err).Msg("decimal value could not be converted")
			return dec, err
		}
		return phy, err
	case a2l.TabVerb:
		err = errors.New("conversion type tabVerb called by numeric function for compuMethod: " + cm.Name)
		log.Err(err).Msg("decimal value could not be converted")
		return 0, err
	default:
		err = errors.New("conversion Type undefined in compuMethod: " + cm.Name)
		log.Err(err).Msg("decimal value could not be converted")
		return 0, err
	}
}

func calcTabVerbRange(dec float64, cvr *a2l.CompuVTabRange, cd *CalibrationData) (string, error) {
	var err error
	tabRange, exists := cd.A2l.Project.Modules[cd.ModuleIndex].CompuVTabRanges[cvr.Name]
	if !exists {
		err = errors.New("conversion table " + cvr.Name)
		log.Err(err).Msg("decimal value could not be converted")
		return "", err
	}
	if tabRange.NumberValueTriplesSet {
		var i uint16
		for i = 0; i < tabRange.NumberValueTriples; i++ {
			if tabRange.InValMin[i] <= dec && dec < tabRange.InValMin[i] {
				return tabRange.OutVal[i], err
			}
		}
		if !tabRange.DefaultValue.DisplayStringSet {
			err = errors.New("no default output found for conversion table " + cvr.Name)
			log.Err(err).Msg("decimal value could not be converted")
			return "", err
		}
		return tabRange.DefaultValue.DisplayString, err
	} else {
		err = errors.New("no output at all found for conversion table " + cvr.Name)
		log.Err(err).Msg("decimal value could not be converted")
		return "", err
	}

}

func calcTabVerb(dec float64, cv *a2l.CompuVTab, cd *CalibrationData) (string, error) {
	var err error
	tab, exists := cd.A2l.Project.Modules[cd.ModuleIndex].CompuVTabs[cv.Name]
	if !exists {
		err = errors.New("conversion table " + cv.Name + " not found")
		log.Err(err).Msg("decimal value could not be converted")
		return "", err
	}
	if tab.InValSet && tab.OutValSet {
		for i := range tab.InVal {
			if dec == tab.InVal[i] {
				return tab.OutVal[i], err
			}
		}
		if !tab.DefaultValue.DisplayStringSet {
			err = errors.New("no default output found for conversion table " + cv.Name)
			log.Err(err).Msg("decimal value could not be converted")
			return "", err
		}
		return tab.DefaultValue.DisplayString, err
	} else {
		err = errors.New("no output at all found for conversion table " + cv.Name)
		log.Err(err).Msg("decimal value could not be converted")
		return "", err
	}
}

//ToDo: Hier weitermachen. Compu Method Tab_verb / Tab_Verb_Ranges / Status String Ref weiterimplementieren.

/*
calcRatFunc computes the following formula: f(Physical) = Decimal
y = (axx + bx + c) / (dxx + ex + f)
inverted fi(Decimal) = Physical
y = (e dec - b)/(2 (a - d dec)) ± sqrt((e dec - b)^2 - 4 (d dec - a) (f dec - c))/(2 (a - d dec))
*/
func calcRatFunc(dec float64, cm *a2l.CompuMethod) (float64, error) {
	var err error
	//following formula defines f(Physical) = Decimal
	//y = (axx + bx + c) / (dxx + ex + f)
	//inverted fi(Decimal) = Physical
	//y = (e dec - b)/(2 (a - d dec)) ± sqrt((e dec - b)^2 - 4 (d dec - a) (f dec - c))/(2 (a - d dec))
	firstDivisor := (2 * (cm.Coeffs.A - cm.Coeffs.D*dec))
	if firstDivisor == 0 {
		err = errors.New("rationality function cannot be computed(zero divisor) for compuMethod: " + cm.Name)
		log.Err(err).Msg("decimal value could not be converted")
		return 0, err
	}
	secondDivisorPositive := (2 * (cm.Coeffs.A - cm.Coeffs.D*dec)) +
		math.Sqrt(math.Pow((cm.Coeffs.E*dec-cm.Coeffs.B), 2)-4*(cm.Coeffs.D*dec-cm.Coeffs.A)*(cm.Coeffs.F*dec-cm.Coeffs.C))/firstDivisor
	secondDivisorNegative := (2 * (cm.Coeffs.A - cm.Coeffs.D*dec)) -
		math.Sqrt(math.Pow((cm.Coeffs.E*dec-cm.Coeffs.B), 2)-4*(cm.Coeffs.D*dec-cm.Coeffs.A)*(cm.Coeffs.F*dec-cm.Coeffs.C))/firstDivisor

	if secondDivisorPositive != 0 && secondDivisorNegative != 0 {
		plusVal := (cm.Coeffs.E*dec - cm.Coeffs.B) / secondDivisorPositive
		minusVal := (cm.Coeffs.E*dec - cm.Coeffs.B) / secondDivisorNegative
		testVal := (cm.Coeffs.A*plusVal*plusVal + cm.Coeffs.B*plusVal + cm.Coeffs.C) / (cm.Coeffs.D*plusVal*plusVal + cm.Coeffs.E*plusVal + cm.Coeffs.F)
		if testVal == dec {
			return plusVal, err
		} else {
			return minusVal, err
		}
	} else if secondDivisorPositive != 0 {
		plusVal := (cm.Coeffs.E*dec - cm.Coeffs.B) / secondDivisorPositive
		return plusVal, err
	} else if secondDivisorNegative != 0 {
		minusVal := (cm.Coeffs.E*dec - cm.Coeffs.B) / secondDivisorNegative
		return minusVal, err
	} else {
		err = errors.New("rationality function cannot be computed(zero divisor) for compuMethod: " + cm.Name)
		log.Err(err).Msg("decimal value could not be converted")
		return 0, err
	}

}

func calcTabNoIntp(dec float64, cm *a2l.CompuMethod, cd *CalibrationData) (float64, error) {
	var err error
	tab, exists := cd.A2l.Project.Modules[cd.ModuleIndex].CompuTabs[cm.CompuTabRef.ConversionTable]
	if !exists {
		if err != nil {
			err = errors.New("conversion table " + cm.CompuTabRef.ConversionTable + " not found for compuMethod: " + cm.Name)
			log.Err(err).Msg("decimal value could not be converted")
			return dec, err
		}
	}
	if tab.InValSet && tab.OutValSet {
		for i := range tab.InVal {
			if dec == tab.InVal[i] {
				return tab.OutVal[i], err
			}
		}
		if !tab.DefaultValueNumeric.DisplayValueSet {
			err = errors.New("no default output found for conversion table " + cm.CompuTabRef.ConversionTable + " in compu method " + cm.Name)
			log.Err(err).Msg("decimal value could not be converted")
			return dec, err
		}
		return tab.DefaultValueNumeric.DisplayValue, err
	} else if tab.DefaultValueNumeric.DisplayValueSet {
		return tab.DefaultValueNumeric.DisplayValue, err
	} else {
		err = errors.New("no output at all found for conversion table " + cm.CompuTabRef.ConversionTable + " in compu method " + cm.Name)
		log.Err(err).Msg("decimal value could not be converted")
		return dec, err
	}
}
