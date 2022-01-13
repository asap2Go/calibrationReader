package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

/*AXIS_PTS and AXIS_DESCR define the same parameters.
Which parameters are dominate is described at AXIS_DESCR.*/
type axisPts struct {
	/*The name has to be unique within all measure and adjustable objects of the MODULE,
	i.e. there must	not be another AXIS_PTS, MEASUREMENT, CHARACTERISTIC, BLOB or INSTANCE object
	with an equal identifier in the same MODULE. Furthermore it is not allowed to have a structure
	component with equal name in the MODULE. (Rules for building the full identifiers of structure
	components: see at INSTANCE).*/
	name              string
	nameSet           bool
	longIdentifier    string
	longIdentifierSet bool
	//address of the adjustable object in the emulation memory
	address    string //uint32
	addressSet bool
	/*inputQuantity references the data record for description of the input	quantity (see MEASUREMENT).
	If there is no input quantity assigned, parameter 'InputQuantity' should be set to "NO_INPUT_QUANTITY"
	(measurement and calibration systems must be capable to treat this case).*/
	inputQuantity    string
	inputQuantitySet bool
	//reference to the relevant data record for description of the record layout (see RECORD_LAYOUT)
	depositIdent    string
	depositIdentSet bool
	/*Maximum difference of physical value recommended for parameter change within one calibration step.
	If the difference in change exceeds this value, control	algorithms might fail.
	The value 0 describes that there is	no limit.*/
	maxDiff    float64
	maxDiffSet bool
	/*Reference to the relevant record of the description of the conversion method (see COMPU_METHOD).
	If there is no conversion method, as in the case of CURVE_AXIS,
	the parameter ‘Conversion’ should be set to "NO_COMPU_METHOD"
	(measurement and calibration systems must be able to handle this case).*/
	conversion    string
	conversionSet bool
	//maximum number of axis points
	maxAxisPoints    uint16
	maxAxisPointsSet bool
	//plausible range of axis point values, lower limit
	lowerLimit    float64
	lowerLimitSet bool
	//plausible range of axis point values, upper limit
	upperLimit          float64
	upperLimitSet       bool
	annotation          []annotation
	byteOrder           byteOrder
	calibrationAccess   calibrationAccessEnum
	deposit             deposit
	displayIdentifier   DisplayIdentifier
	ecuAddressExtension ecuAddressExtension
	extendedLimits      extendedLimits
	format              format
	functionList        []FunctionList
	guardRails          GuardRails
	ifData              []IfData
	monotony            Monotony
	physUnit            physUnit
	readOnly            readOnlyKeyword
	refMemorySegment    refMemorySegment
	stepSize            StepSize
	symbolLink          symbolLink
}

func parseAxisPts(tok *tokenGenerator) (axisPts, error) {
	ap := axisPts{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case beginAnnotationToken:
			var buf annotation
			buf, err = parseAnnotation(tok)
			if err != nil {
				log.Err(err).Msg("axisPts annotation could not be parsed")
				break forLoop
			}
			ap.annotation = append(ap.annotation, buf)
			log.Info().Msg("axisPts annotation successfully parsed")
		case byteOrderToken:
			ap.byteOrder, err = parseByteOrder(tok)
			if err != nil {
				log.Err(err).Msg("axisPts byteOrder could not be parsed")
				break forLoop
			}
			log.Info().Msg("axisPts byteOrder successfully parsed")
		case calibrationAccessToken:
			ap.calibrationAccess, err = parseCalibrationAccessEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisPts calibrationAccess could not be parsed")
				break forLoop
			}
			log.Info().Msg("axisPts calibrationAccess successfully parsed")
		case depositToken:
			ap.deposit, err = parseDeposit(tok)
			if err != nil {
				log.Err(err).Msg("axisPts deposit could not be parsed")
				break forLoop
			}
			log.Info().Msg("axisPts deposit successfully parsed")
		case displayIdentifierToken:
			ap.displayIdentifier, err = parseDisplayIdentifier(tok)
			if err != nil {
				log.Err(err).Msg("axisPts displayIdentifier could not be parsed")
				break forLoop
			}
			log.Info().Msg("axisPts displayIdentifier successfully parsed")
		case ecuAddressExtensionToken:
			ap.ecuAddressExtension, err = parseECUAddressExtension(tok)
			if err != nil {
				log.Err(err).Msg("axisPts ecuAddressExtension could not be parsed")
				break forLoop
			}
			log.Info().Msg("axisPts ecuAddressExtension successfully parsed")
		case extendedLimitsToken:
			ap.extendedLimits, err = parseExtendedLimits(tok)
			if err != nil {
				log.Err(err).Msg("axisPts extendedLimits could not be parsed")
				break forLoop
			}
			log.Info().Msg("axisPts extendedLimits successfully parsed")
		case formatToken:
			ap.format, err = parseFormat(tok)
			if err != nil {
				log.Err(err).Msg("axisPts format could not be parsed")
				break forLoop
			}
			log.Info().Msg("axisPts format successfully parsed")
		case beginFunctionListToken:
			var buf FunctionList
			buf, err = parseFunctionList(tok)
			if err != nil {
				log.Err(err).Msg("axisPts functionList could not be parsed")
				break forLoop
			}
			ap.functionList = append(ap.functionList, buf)
			log.Info().Msg("axisPts functionList successfully parsed")
		case guardRailsToken:
			ap.guardRails, err = parseGuardRails(tok)
			if err != nil {
				log.Err(err).Msg("axisPts guardRails could not be parsed")
				break forLoop
			}
			log.Info().Msg("axisPts guardRails successfully parsed")
		case beginIfDataToken:
			var buf IfData
			buf, err = parseIfData(tok)
			if err != nil {
				log.Err(err).Msg("axisPts ifData could not be parsed")
				break forLoop
			}
			ap.ifData = append(ap.ifData, buf)
			log.Info().Msg("axisPts ifData successfully parsed")
		case monotonyToken:
			ap.monotony, err = parseMonotony(tok)
			if err != nil {
				log.Err(err).Msg("axisPts monotony could not be parsed")
				break forLoop
			}
			log.Info().Msg("axisPts monotony successfully parsed")
		case physUnitToken:
			ap.physUnit, err = parsePhysUnit(tok)
			if err != nil {
				log.Err(err).Msg("axisPts physUnit could not be parsed")
				break forLoop
			}
			log.Info().Msg("axisPts physUnit successfully parsed")
		case readOnlyToken:
			ap.readOnly, err = parseReadOnly(tok)
			if err != nil {
				log.Err(err).Msg("axisPts readOnly could not be parsed")
				break forLoop
			}
			log.Info().Msg("axisPts readOnly successfully parsed")
		case refMemorySegmentToken:
			ap.refMemorySegment, err = parseRefMemorySegment(tok)
			if err != nil {
				log.Err(err).Msg("axisPts refMemorySegment could not be parsed")
				break forLoop
			}
			log.Info().Msg("axisPts refMemorySegment successfully parsed")
		case stepSizeToken:
			ap.stepSize, err = parseStepSize(tok)
			if err != nil {
				log.Err(err).Msg("axisPts stepSize could not be parsed")
				break forLoop
			}
			log.Info().Msg("axisPts stepSize successfully parsed")
		case symbolLinkToken:
			ap.symbolLink, err = parseSymbolLink(tok)
			if err != nil {
				log.Err(err).Msg("axisPts symbolLink could not be parsed")
				break forLoop
			}
			log.Info().Msg("axisPts symbolLink successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("axisPts could not be parsed")
				break forLoop
			} else if tok.current() == endAxisPtsToken {
				break forLoop
			} else if !ap.nameSet {
				ap.name = tok.current()
				ap.nameSet = true
				log.Info().Msg("axisPts name successfully parsed")
			} else if !ap.longIdentifierSet {
				ap.longIdentifier = tok.current()
				ap.longIdentifierSet = true
				log.Info().Msg("axisPts longIdentifier successfully parsed")
			} else if !ap.addressSet {
				ap.address = tok.current()
				ap.addressSet = true
				log.Info().Msg("axisPts address successfully parsed")
			} else if !ap.inputQuantitySet {
				ap.inputQuantity = tok.current()
				ap.inputQuantitySet = true
				log.Info().Msg("axisPts inputQuantity successfully parsed")
			} else if !ap.depositIdentSet {
				ap.depositIdent = tok.current()
				ap.depositIdentSet = true
				log.Info().Msg("axisPts depositIdent successfully parsed")
			} else if !ap.maxDiffSet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("axisPts maxDiff could not be parsed")
					break forLoop
				}
				ap.maxDiff = buf
				ap.maxDiffSet = true
				log.Info().Msg("axisPts maxDiff successfully parsed")
			} else if !ap.conversionSet {
				ap.conversion = tok.current()
				ap.conversionSet = true
				log.Info().Msg("axisPts conversion successfully parsed")
			} else if !ap.maxAxisPointsSet {
				var buf uint64
				buf, err = strconv.ParseUint(tok.current(), 10, 16)
				if err != nil {
					log.Err(err).Msg("axisPts maxAxisPoints could not be parsed")
					break forLoop
				}
				ap.maxAxisPoints = uint16(buf)
				ap.maxAxisPointsSet = true
				log.Info().Msg("axisPts maxAxisPoints successfully parsed")
			} else if !ap.lowerLimitSet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("axisPts lowerLimit could not be parsed")
					break forLoop
				}
				ap.lowerLimit = buf
				ap.lowerLimitSet = true
				log.Info().Msg("axisPts lowerLimit successfully parsed")
			} else if !ap.upperLimitSet {
				var buf float64
				buf, err = strconv.ParseFloat(tok.current(), 64)
				if err != nil {
					log.Err(err).Msg("axisPts upperLimit could not be parsed")
					break forLoop
				}
				ap.upperLimit = buf
				ap.upperLimitSet = true
				log.Info().Msg("axisPts upperLimit successfully parsed")
			}
		}
	}
	return ap, err
}
