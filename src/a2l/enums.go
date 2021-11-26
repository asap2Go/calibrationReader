package a2l

import "errors"

//enums used in several A2L Datatypes:

type encodingEnum string

const (
	undefinedEncoding encodingEnum = emptyToken
	UTF8              encodingEnum = "UTF8"
	UTF16             encodingEnum = "UTF16"
	UTF32             encodingEnum = "UTF32"
)

func parseEncodingEnum(tok *tokenGenerator) (encodingEnum, error) {
	e := undefinedEncoding
	var err error
	switch tok.current() {
	case "UTF8":
		e = UTF8
	case "UTF16":
		e = UTF16
	case "UTF32":
		e = UTF32
	default:
		err = errors.New("incorrect value " + tok.current() + " for enum encoding")
	}
	return e, err
}

type dataTypeEnum string

const (
	undefinedDatatype dataTypeEnum = emptyToken
	UBYTE             dataTypeEnum = ubyteToken
	SBYTE             dataTypeEnum = sbyteToken
	UWORD             dataTypeEnum = uwordToken
	SWORD             dataTypeEnum = swordToken
	ULONG             dataTypeEnum = ulongToken
	SLONG             dataTypeEnum = slongToken
	aUint64           dataTypeEnum = aUint64Token
	aInt64            dataTypeEnum = aInt64Token
	float32Ieee       dataTypeEnum = float32IeeeToken
	float64Ieee       dataTypeEnum = float64IeeeToken
)

func parseDataTypeEnum(tok *tokenGenerator) (dataTypeEnum, error) {
	d := undefinedDatatype
	var err error
	switch tok.current() {
	case ubyteToken:
		d = UBYTE
	case sbyteToken:
		d = SBYTE
	case uwordToken:
		d = UWORD
	case swordToken:
		d = SWORD
	case ulongToken:
		d = ULONG
	case slongToken:
		d = SLONG
	case aUint64Token:
		d = aUint64
	case aInt64Token:
		d = aInt64
	case float32IeeeToken:
		d = float32Ieee
	case float64IeeeToken:
		d = float64Ieee
	default:
		err = errors.New("incorrect value " + tok.current() + " for enum datatype")
	}
	return d, err
}

type dataSizeEnum string

const (
	undefinedDatasize dataSizeEnum = emptyToken
	BYTE              dataSizeEnum = byteToken
	WORD              dataSizeEnum = wordToken
	LONG              dataSizeEnum = longToken
)

func parseDataSizeEnum(tok *tokenGenerator) (dataSizeEnum, error) {
	d := undefinedDatasize
	var err error
	switch tok.current() {
	case byteToken:
		d = BYTE
	case wordToken:
		d = WORD
	case longToken:
		d = LONG
	default:
		err = errors.New("incorrect value " + tok.current() + " for enum datasize")
	}
	return d, err
}

//AddrTypeEnum defines which address width is necessary.
type AddrTypeEnum string

const (
	undefinedAddrtype AddrTypeEnum = emptyToken
	PBYTE             AddrTypeEnum = pbyteToken
	PWORD             AddrTypeEnum = pwordToken
	PLONG             AddrTypeEnum = plongToken
	/*DIRECT: If an adjustable or measurable object is defined with indirect addressing
	(ADDRESS_TYPE is not DIRECT) and if the used interface does not support indirect
	addressing, it is the responsibility of the MC-System to dereference the object’s
	address before accessing the data or configuring measurement lists.*/
	DIRECT AddrTypeEnum = directToken
)

func parseAddrTypeEnum(tok *tokenGenerator) (AddrTypeEnum, error) {
	a := undefinedAddrtype
	var err error
	switch tok.current() {
	case pbyteToken:
		a = PBYTE
	case pwordToken:
		a = PWORD
	case plongToken:
		a = PLONG
	case directToken:
		a = DIRECT
	default:
		err = errors.New("incorrect value " + tok.current() + " for enum addrtype")
	}
	return a, err
}

type ByteOrderEnum string

const (
	undefinedByteorder ByteOrderEnum = emptyToken
	littleEndian       ByteOrderEnum = littleEndianToken
	bigEndian          ByteOrderEnum = bigEndianToken
	msbLast            ByteOrderEnum = msbLastToken
	msbFirst           ByteOrderEnum = msbFirstToken
)

func parseByteOrderEnum(tok *tokenGenerator) (ByteOrderEnum, error) {
	b := undefinedByteorder
	var err error
	switch tok.current() {
	case littleEndianToken:
		b = littleEndian
	case bigEndianToken:
		b = bigEndian
	case msbLastToken:
		b = msbLast
	case msbFirstToken:
		b = msbFirst
	default:
		err = errors.New("incorrect value " + tok.current() + " for enum byteorder")
	}
	return b, err
}

type IndexOrderEnum string

const (
	undefinedIndexorder IndexOrderEnum = emptyToken
	indexIncr           IndexOrderEnum = indexIncrToken
	indexDecr           IndexOrderEnum = indexDecrToken
)

func parseIndexOrderEnum(tok *tokenGenerator) (IndexOrderEnum, error) {
	i := undefinedIndexorder
	var err error
	switch tok.current() {
	case indexIncrToken:
		i = indexIncr
	case indexDecrToken:
		i = indexDecr
	default:
		err = errors.New("incorrect value " + tok.current() + " for enum byteorder")
	}
	return i, err
}

type attributeEnum string

const (
	undefinedAttribute attributeEnum = emptyToken
	/*curveAxis type uses a separate CURVE CHARACTERISTIC to rescale the axis.
	The referenced CURVE is used to	lookup an axis index, and the index value is
	used by the controller to determine the	operating point in the CURVE or MAP.*/
	curveAxis attributeEnum = curveAxisToken
	/*comAxis: Group axis points or description of the axis
	points for deposit. For this variant of the	axis points the axis point values are
	separated from the table values of the curve or map in the emulation memory and
	must be described by a special AXIS_PTS	data record.
	The reference to this record occurs with the keyword 'AXIS_PTS_REF'.*/
	comAxis attributeEnum = comAxisToken
	/*fixAxis is a curve or a map with virtual axis
	points that are not deposited at EPROM.
	The axis points can be calculated from parameters defined with keywords:
	FIX_AXIS_PAR, FIX_AXIS_PAR_DIST	and FIX_AXIS_PAR_LIST.
	The axis points	cannot be modified.*/
	fixAxis attributeEnum = fixAxisToken
	/*Rescale axis. For this variant of the axis
	points the axis point values are separated from the table values of the curve or map in
	the emulation memory and must be described by a special AXIS_PTS data
	record. The reference to this record occurs	with the keyword 'AXIS_PTS_REF'.*/
	resAxis attributeEnum = resAxisToken
	stdAxis attributeEnum = stdAxisToken
	INTERN  attributeEnum = internToken
	EXTERN  attributeEnum = externToken
)

func parseAttributeEnum(tok *tokenGenerator) (attributeEnum, error) {
	a := undefinedAttribute
	var err error
	switch tok.current() {
	case curveAxisToken:
		a = curveAxis
	case comAxisToken:
		a = comAxis
	case fixAxisToken:
		a = fixAxis
	case resAxisToken:
		a = resAxis
	case stdAxisToken:
		a = stdAxis
	case internToken:
		a = INTERN
	case externToken:
		a = EXTERN
	default:
		err = errors.New("incorrect value " + tok.current() + " for enum attribute")
	}
	return a, err
}

type CalibrationAccessEnum string

const (
	undefinedCalibrationAccess CalibrationAccessEnum = emptyToken
	CALIBRATION                CalibrationAccessEnum = calibrationToken
	noCalibration              CalibrationAccessEnum = noCalibrationToken
	notInMcdSystem             CalibrationAccessEnum = notInMcdSystemToken
	offlineCalibration         CalibrationAccessEnum = offlineCalibrationToken
)

func parseCalibrationAccessEnum(tok *tokenGenerator) (CalibrationAccessEnum, error) {
	ca := undefinedCalibrationAccess
	var err error
	switch tok.next() {
	case calibrationToken:
		ca = CALIBRATION
	case noCalibrationToken:
		ca = noCalibration
	case notInMcdSystemToken:
		ca = notInMcdSystem
	case offlineCalibrationToken:
		ca = offlineCalibration
	default:
		err = errors.New("incorrect value " + tok.current() + " for enum calibrationAccess")
	}
	return ca, err
}

type typeEnum string

const (
	undefinedType typeEnum = emptyToken
	ASCII         typeEnum = asciiToken
	CURVE         typeEnum = curveToken
	MAP           typeEnum = mapToken
	CUBOID        typeEnum = cuboidToken
	cube4         typeEnum = cube4Token
	cube5         typeEnum = cube5Token
	valBlk        typeEnum = valBlkToken
	VALUE         typeEnum = valueToken
	DERIVED       typeEnum = derivedToken
	extendedSi    typeEnum = extendedSiToken
)

func parseTypeEnum(tok *tokenGenerator) (typeEnum, error) {
	t := undefinedType
	var err error
	switch tok.current() {
	case asciiToken:
		t = ASCII
	case curveToken:
		t = CURVE
	case mapToken:
		t = MAP
	case cuboidToken:
		t = CUBOID
	case cube4Token:
		t = cube4
	case cube5Token:
		t = cube5
	case valBlkToken:
		t = valBlk
	case valueToken:
		t = VALUE
	case derivedToken:
		t = DERIVED
	case extendedSiToken:
		t = extendedSi
	default:
		err = errors.New("incorrect value " + tok.current() + " for enum type")
	}
	return t, err
}

type ConversionTypeEnum string

const (
	undefinedConversionType ConversionTypeEnum = emptyToken
	IDENTICAL               ConversionTypeEnum = identicalToken
	FORM                    ConversionTypeEnum = formToken
	LINEAR                  ConversionTypeEnum = linearToken
	ratFunc                 ConversionTypeEnum = ratFuncToken
	tabIntp                 ConversionTypeEnum = tabIntpToken
	tabNointp               ConversionTypeEnum = tabNointpToken
	tabVerb                 ConversionTypeEnum = tabVerbToken
)

func parseConversionTypeEnum(tok *tokenGenerator) (ConversionTypeEnum, error) {
	ct := undefinedConversionType
	var err error
	switch tok.current() {
	case identicalToken:
		ct = IDENTICAL
	case formToken:
		ct = FORM
	case linearToken:
		ct = LINEAR
	case ratFuncToken:
		ct = ratFunc
	case tabIntpToken:
		ct = tabIntp
	case tabNointpToken:
		ct = tabNointp
	case tabVerbToken:
		ct = tabVerb
	default:
		err = errors.New("incorrect value " + tok.current() + " for enum conversionType")
	}
	return ct, err
}

type IndexModeEnum string

const (
	undefinedIndexMode IndexModeEnum = emptyToken
	alternateCurves    IndexModeEnum = alternateCurvesToken
	alternateWithX     IndexModeEnum = alternateWithXToken
	alternateWithY     IndexModeEnum = alternateWithYToken
	columnDir          IndexModeEnum = columnDirToken
	rowDir             IndexModeEnum = rowDirToken
)

func parseIndexModeEnum(tok *tokenGenerator) (IndexModeEnum, error) {
	im := undefinedIndexMode
	var err error
	switch tok.current() {
	case alternateCurvesToken:
		im = alternateCurves
	case alternateWithXToken:
		im = alternateWithX
	case alternateWithYToken:
		im = alternateWithY
	case columnDirToken:
		im = columnDir
	case rowDirToken:
		im = rowDir
	default:
		err = errors.New("incorrect value " + tok.current() + " for enum indexMode")
	}
	return im, err
}

type ModeEnum string

const (
	undefinedMode ModeEnum = emptyToken
	ABSOLUTE      ModeEnum = absoluteToken
	DIFFERENCE    ModeEnum = differenceToken
)

func parseModeEnum(tok *tokenGenerator) (ModeEnum, error) {
	m := undefinedMode
	var err error
	switch tok.current() {
	case absoluteToken:
		m = ABSOLUTE
	case differenceToken:
		m = DIFFERENCE
	default:
		err = errors.New("incorrect value " + tok.current() + " for enum mode")
	}
	return m, err
}

type prgTypeEnum string

const (
	undefinedPrgType     prgTypeEnum = emptyToken
	prgCode              prgTypeEnum = prgCodeToken
	prgData              prgTypeEnum = prgDataToken
	prgReserved          prgTypeEnum = prgReservedToken
	calibrationVariables prgTypeEnum = calibrationVariablesToken
	CODE                 prgTypeEnum = codeToken
	DATA                 prgTypeEnum = dataToken
	excludeFromFlash     prgTypeEnum = excludeFromFlashToken
	offlineData          prgTypeEnum = offlineDataToken
	RESERVED2            prgTypeEnum = reservedToken
	SERAM                prgTypeEnum = seramToken
	VARIABLES            prgTypeEnum = variablesToken
)

func parsePrgTypeEnum(tok *tokenGenerator) (prgTypeEnum, error) {
	pt := undefinedPrgType
	var err error
	switch tok.current() {
	case prgCodeToken:
		pt = prgCode
	case prgDataToken:
		pt = prgData
	case prgReservedToken:
		pt = prgReserved
	case calibrationVariablesToken:
		pt = calibrationVariables
	case codeToken:
		pt = CODE
	case dataToken:
		pt = DATA
	case excludeFromFlashToken:
		pt = excludeFromFlash
	case offlineDataToken:
		pt = offlineData
	case reservedToken:
		pt = RESERVED2
	case seramToken:
		pt = SERAM
	case variablesToken:
		pt = VARIABLES
	default:
		err = errors.New("incorrect value " + tok.current() + " for enum prgType")
	}
	return pt, err
}

type memoryTypeEnum string

const (
	undefinedMemoryType memoryTypeEnum = emptyToken
	EEPROM              memoryTypeEnum = eepromToken
	EPROM               memoryTypeEnum = epromToken
	FLASH               memoryTypeEnum = flashToken
	RAM                 memoryTypeEnum = ramToken
	ROM                 memoryTypeEnum = romToken
	REGISTER            memoryTypeEnum = registerToken
	notInEcu            memoryTypeEnum = notInEcuToken
)

func parseMemoryTypeEnum(tok *tokenGenerator) (memoryTypeEnum, error) {
	mt := undefinedMemoryType
	var err error
	switch tok.current() {
	case eepromToken:
		mt = EEPROM
	case epromToken:
		mt = EPROM
	case flashToken:
		mt = FLASH
	case ramToken:
		mt = RAM
	case romToken:
		mt = ROM
	case registerToken:
		mt = REGISTER
	case notInEcuToken:
		mt = notInEcu
	default:
		err = errors.New("incorrect value " + tok.current() + " for enum memoryType")
	}
	return mt, err
}

type MonotonyTypeEnum string

const (
	undefinedMonotonyType MonotonyTypeEnum = emptyToken
	monDecrease           MonotonyTypeEnum = monDecreaseToken
	monIncrease           MonotonyTypeEnum = monIncreaseToken
	strictDecrease        MonotonyTypeEnum = strictDecreaseToken
	strictIncrease        MonotonyTypeEnum = strictIncreaseToken
	MONOTONOUS            MonotonyTypeEnum = monotonousToken
	strictMon             MonotonyTypeEnum = strictMonToken
	notMon                MonotonyTypeEnum = notMonToken
)

func parseMonotonyTypeEnum(tok *tokenGenerator) (MonotonyTypeEnum, error) {
	mt := undefinedMonotonyType
	var err error
	switch tok.current() {
	case monDecreaseToken:
		mt = monDecrease
	case monIncreaseToken:
		mt = monIncrease
	case strictDecreaseToken:
		mt = strictDecrease
	case strictIncreaseToken:
		mt = strictIncrease
	case monotonousToken:
		mt = MONOTONOUS
	case strictMonToken:
		mt = strictMon
	case notMonToken:
		mt = notMon
	default:
		err = errors.New("incorrect value " + tok.current() + " for enum monotonyType")
	}
	return mt, err
}

type tagEnum string

const (
	undefinedTag tagEnum = emptyToken
	NUMERIC      tagEnum = numericToken
)

func parseTagEnum(tok *tokenGenerator) (tagEnum, error) {
	mt := undefinedTag
	var err error
	switch tok.current() {
	case numericToken:
		mt = NUMERIC
	default:
		err = errors.New("incorrect value " + tok.current() + " for enum tag")
	}
	return mt, err
}

type triggerEnum string

const (
	undefinedTriggerEnum triggerEnum = emptyToken
	OnChange             triggerEnum = onChangeToken
	OnUserRequest        triggerEnum = onUserRequestToken
)

func parseTriggerEnum(tok *tokenGenerator) (triggerEnum, error) {
	mt := undefinedTriggerEnum
	var err error
	switch tok.current() {
	case onChangeToken:
		mt = OnChange
	case onUserRequestToken:
		mt = OnUserRequest
	default:
		err = errors.New("incorrect value " + tok.current() + " for enum trigger")
	}
	return mt, err
}
