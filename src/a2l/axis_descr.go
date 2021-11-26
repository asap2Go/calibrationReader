package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

/*axisDescr is the Axis description within an adjustable object
Note:
With the 'input quantity' parameter a reference is made to a measurement object
(MEASUREMENT). The MEASUREMENT keyword also specifies the
'conversion', 'lower limit' and 'upper limit' parameters.
It is expected that both conversions are equivalent, i.e. they must lead to the
same result. The ‘upper limit’ and ‘lower limit’ parameters may be different.
Note: The keywords FIX_AXIS_PAR, FIX_AXIS_PAR_DIST, DEPOSIT and
FIX_AXIS_PAR_LIST are mutually exclusive, i.e. at most one of these keywords
is allowed to be used at the same AXIS_DESCR record.
Note: For the axis types COM_AXIS, RES_AXIS and CURVE_AXIS some attributes
are defined twice: both at the AXIS_DESCR record and at the referenced
AXIS_PTS resp. CHARACTERISTIC record. These redundant attributes are
InputQuantity, Conversion, MaxAxisPoints, LowerLimit, UpperLimit and some
optional parameters (e.g.: PHYS_UNIT). To support existing use cases where
one common axis is used with different input quantities (e.g. multiple cylinders) it
is recommended to ignore the redundant attributes defined at AXIS_PTS and use
the values of the AXIS_DESCR record instead. Exeptions are MaxAxisPoints and
MONOTONY which are used from AXIS_PTS.
AXIS_DESCR and AXIS_PTS partially support the same keyword parameters. If
a CHARACTERISTC (Curve, Map …) has an AXIS_DESCR of type COM_AXIS
and refers to an AXIS_PTS, then the parameters from AXIS_DESCR shall take
precedence. Exceptions to this rule are the parameters MaxAxisPoints,
DEPOSIT, BYTE_ORDER and MONOTONY, which shall be taken from
AXIS_PTS.
Note:
For the axis type COM_AXIS, it is necessary to have the same dimension for the
axis object and for the referencing curve / map object. Theoretically it is possible
to define the dimension both at the axis and at the curve resp. map – e.g. the
corresponding record layouts can both contain the NO_AXIS_PTS_X component.
It is recommended not to do so, but if the dimension is defined twice and if it is
not equal, then application systems shall always use the dimension of the
AXIS_PTS object.*/
type axisDescr struct {
	attribute    attributeEnum
	attributeSet bool
	/*inputQuantity references the data record for description of the input quantity (see MEASUREMENT).
	If there is no input quantity assigned, parameter 'InputQuantity' should be set to "NO_INPUT_QUANTITY"
	(measurement and calibration systems must be capable to treat this case).
	Note: If the referenced input quantity is an element of an array or a component of a structure,
	the identifier to be used for the reference shall be the name built according to rules described at INSTANCE.*/
	inputQuantity    string
	inputQuantitySet bool
	/*conversion references the relevant record of the description of the conversion method (see COMPU_METHOD).
	If there is no conversion method, as in the case of CURVE_AXIS, the parameter ‘Conversion’
	should be set to "NO_COMPU_METHOD" (measurement and calibration systems must be able to handle this case).*/
	conversion    string
	conversionSet bool
	//maximum number of axis points
	maxAxisPoints    uint16
	maxAxisPointsSet bool
	//lowerLimit quantifies the plausible range of axis point values, lower limit
	lowerLimit    float64
	lowerLimitSet bool
	//upperLimit quantifies the plausible range of axis point values, upper limit
	upperLimit     float64
	upperLimitSet  bool
	annotation     []annotation
	axisPtsRef     axisPtsRef
	byteOrder      byteOrder
	curveAxisRef   curveAxisRef
	deposit        deposit
	extendedLimits extendedLimits
	fixAxisPar     fixAxisPar
	fixAxisParDist fixAxisParDist
	fixAxisParList []fixAxisParList
	format         format
	maxGrad        MaxGrad
	monotony       Monotony
	physUnit       physUnit
	readOnly       readOnlyKeyword
	stepSize       StepSize
}

func parseAxisDescr(tok *tokenGenerator) (axisDescr, error) {
	ad := axisDescr{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case beginAnnotationToken:
			var buf annotation
			buf, err = parseAnnotation(tok)
			if err != nil {
				log.Err(err).Msg("axisDescr annotation could not be parsed")
				break forLoop
			}
			ad.annotation = append(ad.annotation, buf)
			log.Info().Msg("axisDescr annotation successfully parsed")
		case axisPtsRefToken:
			ad.axisPtsRef, err = parseAxisPtsRef(tok)
			if err != nil {
				log.Err(err).Msg("axisDescr axisPtsRef could not be parsed")
				break forLoop
			}
			log.Info().Msg("axisDescr axisPtsRef successfully parsed")
		case byteOrderToken:
			ad.byteOrder, err = parseByteOrder(tok)
			if err != nil {
				log.Err(err).Msg("axisDescr byteOrder could not be parsed")
				break forLoop
			}
			log.Info().Msg("axisDescr byteOrder successfully parsed")
		case curveAxisRefToken:
			ad.curveAxisRef, err = parseCurveAxisRef(tok)
			if err != nil {
				log.Err(err).Msg("axisDescr curveAxisRef could not be parsed")
				break forLoop
			}
			log.Info().Msg("axisDescr curveAxisRef successfully parsed")
		case depositToken:
			ad.deposit, err = parseDeposit(tok)
			if err != nil {
				log.Err(err).Msg("axisDescr deposit could not be parsed")
				break forLoop
			}
			log.Info().Msg("axisDescr deposit successfully parsed")
		case extendedLimitsToken:
			ad.extendedLimits, err = parseExtendedLimits(tok)
			if err != nil {
				log.Err(err).Msg("axisDescr extendedLimits could not be parsed")
				break forLoop
			}
			log.Info().Msg("axisDescr extendedLimits successfully parsed")
		case beginFixAxisParToken:
			ad.fixAxisPar, err = parseFixAxisPar(tok)
			if err != nil {
				log.Err(err).Msg("axisDescr fixAxisPar could not be parsed")
				break forLoop
			}
			log.Info().Msg("axisDescr fixAxisPar successfully parsed")
		case fixAxisParDistToken:
			ad.fixAxisParDist, err = parseFixAxisParDist(tok)
			if err != nil {
				log.Err(err).Msg("axisDescr fixAxisParDist could not be parsed")
				break forLoop
			}
			log.Info().Msg("axisDescr fixAxisParDist successfully parsed")
		case beginFixAxisParListToken:
			var buf fixAxisParList
			buf, err = parseFixAxisParList(tok)
			if err != nil {
				log.Err(err).Msg("axisDescr fixAxisParList could not be parsed")
				break forLoop
			}
			ad.fixAxisParList = append(ad.fixAxisParList, buf)
			log.Info().Msg("axisDescr fixAxisParList successfully parsed")
		case formatToken:
			ad.format, err = parseFormat(tok)
			if err != nil {
				log.Err(err).Msg("axisDescr format could not be parsed")
				break forLoop
			}
			log.Info().Msg("axisDescr format successfully parsed")
		case maxGradToken:
			ad.maxGrad, err = parseMaxGrad(tok)
			if err != nil {
				log.Err(err).Msg("axisDescr maxGrad could not be parsed")
				break forLoop
			}
			log.Info().Msg("axisDescr maxGrad successfully parsed")
		case monotonyToken:
			ad.monotony, err = parseMonotony(tok)
			if err != nil {
				log.Err(err).Msg("axisDescr monotony could not be parsed")
				break forLoop
			}
			log.Info().Msg("axisDescr monotony successfully parsed")
		case physUnitToken:
			ad.physUnit, err = parsePhysUnit(tok)
			if err != nil {
				log.Err(err).Msg("axisDescr physUnit could not be parsed")
				break forLoop
			}
			log.Info().Msg("axisDescr physUnit successfully parsed")
		case readOnlyToken:
			ad.readOnly, err = parseReadOnly(tok)
			if err != nil {
				log.Err(err).Msg("axisDescr readOnly could not be parsed")
				break forLoop
			}
			log.Info().Msg("axisDescr readOnly successfully parsed")
		case stepSizeToken:
			ad.stepSize, err = parseStepSize(tok)
			if err != nil {
				log.Err(err).Msg("axisDescr stepSize could not be parsed")
				break forLoop
			}
			log.Info().Msg("axisDescr stepSize successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("asap2Version could not be parsed")
				break forLoop
			} else if tok.current() == endAxisDescrToken {
				break forLoop
			} else if !ad.attributeSet {
				ad.attribute, err = parseAttributeEnum(tok)
				if err != nil {
					log.Err(err).Msg("axisDescr attribute could not be parsed")
					break forLoop
				}
				ad.attributeSet = true
				log.Info().Msg("axisDescr attribute successfully parsed")
			} else if !ad.inputQuantitySet {
				ad.inputQuantity = tok.current()
				ad.inputQuantitySet = true
				log.Info().Msg("axisDescr inputQuantity successfully parsed")
			} else if !ad.conversionSet {
				ad.conversion = tok.current()
				ad.inputQuantitySet = true
				log.Info().Msg("axisDescr inputQuantity successfully parsed")
			} else if !ad.maxAxisPointsSet {
				var buf uint64
				buf, err = strconv.ParseUint(tok.current(), 10, 16)
				if err != nil {
					log.Err(err).Msg("axisDescr maxAxisPoints could not be parsed")
					break forLoop
				}
				ad.maxAxisPoints = uint16(buf)
				ad.maxAxisPointsSet = true
				log.Info().Msg("axisDescr maxAxisPoints successfully parsed")
			} else if !ad.lowerLimitSet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("axisDescr lowerLimit could not be parsed")
					break forLoop
				}
				ad.lowerLimit = buf
				ad.lowerLimitSet = true
				log.Info().Msg("axisDescr lowerLimit successfully parsed")
			} else if !ad.upperLimitSet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("axisDescr upperLimit could not be parsed")
					break forLoop
				}
				ad.upperLimit = buf
				ad.upperLimitSet = true
				log.Info().Msg("axisDescr upperLimit successfully parsed")
			}
		}
	}
	return ad, err
}
