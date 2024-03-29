package a2l

import (
	"errors"

	"github.com/x448/float16"
)

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

type DataTypeEnum string

const (
	undefinedDatatype DataTypeEnum = emptyToken
	//UBYTE is defined as an unsigned 8 bit integer
	UBYTE DataTypeEnum = ubyteToken
	//SBYTE is defined as an unsigned 8 bit integer
	SBYTE DataTypeEnum = sbyteToken
	//UWORD seems to be statically defined as a 32 bit unsigned integer
	UWORD DataTypeEnum = uwordToken
	//SWORD seems to be statically defined as a 32 bit signed integer
	SWORD DataTypeEnum = swordToken
	//ULONG is defined as a 64 bit unsigned integer
	ULONG DataTypeEnum = ulongToken
	//SLONG is defined as a 64 bit signed integer
	SLONG DataTypeEnum = slongToken
	//AUint64 is defined as a 64 bit unsigned integer
	AUint64 DataTypeEnum = aUint64Token
	//AInt64 is defined as a 64 bit signed integer
	AInt64 DataTypeEnum = aInt64Token
	//Float16Ieee is a standard 16 bit float
	Float16Ieee DataTypeEnum = float16IeeeToken
	//Float32Ieee is a standard 32 bit float
	Float32Ieee DataTypeEnum = float32IeeeToken
	//Float64Ieee is a standard 64 bit float
	Float64Ieee DataTypeEnum = float64IeeeToken
)

func (dte *DataTypeEnum) String() string {
	switch *dte {
	case undefinedDatatype:
		return emptyToken
	case UBYTE:
		return ubyteToken
	case SBYTE:
		return sbyteToken
	case UWORD:
		return uwordToken
	case SWORD:
		return swordToken
	case ULONG:
		return ulongToken
	case SLONG:
		return slongToken
	case AUint64:
		return aUint64Token
	case AInt64:
		return aInt64Token
	case Float16Ieee:
		return float16IeeeToken
	case Float32Ieee:
		return float32IeeeToken
	case Float64Ieee:
		return float64IeeeToken
	default:
		return emptyToken
	}

}

func (dte *DataTypeEnum) GetType() interface{} {
	switch *dte {
	case undefinedDatatype:
		return nil
	case UBYTE:
		var t uint8
		return t
	case SBYTE:
		var t int8
		return t
	case UWORD:
		var t uint32
		return t
	case SWORD:
		var t int32
		return t
	case ULONG:
		var t uint32
		return t
	case SLONG:
		var t int32
		return t
	case AUint64:
		var t uint64
		return t
	case AInt64:
		var t int64
		return t
	case Float16Ieee:
		var t float16.Float16
		return t
	case Float32Ieee:
		var t float32
		return t
	case Float64Ieee:
		var t float64
		return t
	default:
		return nil
	}
}

func (dte *DataTypeEnum) GetDatatypeLength() uint16 {
	switch *dte {
	case UBYTE:
		return 8
	case SBYTE:
		return 8
	case UWORD:
		return 32
	case SWORD:
		return 32
	case ULONG:
		return 32
	case SLONG:
		return 32
	case AUint64:
		return 64
	case AInt64:
		return 64
	case Float16Ieee:
		return 16
	case Float32Ieee:
		return 32
	case Float64Ieee:
		return 64
	default:
		return 0
	}

}

func parseDataTypeEnum(tok *tokenGenerator) (DataTypeEnum, error) {
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
		d = AUint64
	case aInt64Token:
		d = AInt64
	case float16IeeeToken:
		d = Float16Ieee
	case float32IeeeToken:
		d = Float32Ieee
	case float64IeeeToken:
		d = Float64Ieee
	default:
		err = errors.New("incorrect value " + tok.current() + " for enum datatype")
	}
	return d, err
}

type DataSizeEnum string

const (
	undefinedDatasize DataSizeEnum = emptyToken
	BYTE              DataSizeEnum = byteToken
	WORD              DataSizeEnum = wordToken
	LONG              DataSizeEnum = longToken
)

func (dse *DataSizeEnum) GetDataSizeLength() uint16 {
	switch *dse {
	case BYTE:
		return 8
	case WORD:
		return 32
	case LONG:
		return 64
	default:
		return 0
	}

}

func parseDataSizeEnum(tok *tokenGenerator) (DataSizeEnum, error) {
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

// AddrTypeEnum defines which address width is necessary.
type AddrTypeEnum string

const (
	undefinedAddrtype AddrTypeEnum = emptyToken
	PBYTE             AddrTypeEnum = pbyteToken
	PWORD             AddrTypeEnum = pwordToken
	PLONG             AddrTypeEnum = plongToken
	PLONGLONG         AddrTypeEnum = plonLongToken
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
	LittleEndian       ByteOrderEnum = littleEndianToken
	BigEndian          ByteOrderEnum = bigEndianToken
	MsbLast            ByteOrderEnum = msbLastToken
	MsbFirst           ByteOrderEnum = msbFirstToken
	MsbFirstMswLast    ByteOrderEnum = msbFirstMswLastToken
	MsbLastMswFirst    ByteOrderEnum = msbLastMswFirstToken
)

func parseByteOrderEnum(tok *tokenGenerator) (ByteOrderEnum, error) {
	b := undefinedByteorder
	var err error
	switch tok.current() {
	case littleEndianToken:
		b = LittleEndian
	case bigEndianToken:
		b = BigEndian
	case msbLastToken:
		b = MsbLast
	case msbFirstToken:
		b = MsbFirst
	case msbFirstMswLastToken:
		b = MsbFirstMswLast
	case msbLastMswFirstToken:
		b = MsbLastMswFirst
	default:
		err = errors.New("incorrect value " + tok.current() + " for enum byteorder")
	}
	return b, err
}

func (boe *ByteOrderEnum) String() string {
	switch *boe {
	case LittleEndian:
		return littleEndianToken
	case BigEndian:
		return bigEndianToken
	case MsbLast:
		return msbLastToken
	case MsbFirst:
		return msbFirstToken
	case MsbFirstMswLast:
		return msbFirstMswLastToken
	case MsbLastMswFirst:
		return msbLastMswFirstToken
	default:
		return emptyToken
	}
}

type indexOrderEnum string

const (
	undefinedIndexorder indexOrderEnum = emptyToken
	indexIncr           indexOrderEnum = indexIncrToken
	indexDecr           indexOrderEnum = indexDecrToken
)

func parseIndexOrderEnum(tok *tokenGenerator) (indexOrderEnum, error) {
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

type calibrationAccessEnum string

const (
	undefinedCalibrationAccess calibrationAccessEnum = emptyToken
	calibration                calibrationAccessEnum = calibrationToken
	noCalibration              calibrationAccessEnum = noCalibrationToken
	notInMcdSystem             calibrationAccessEnum = notInMcdSystemToken
	offlineCalibration         calibrationAccessEnum = offlineCalibrationToken
)

func parseCalibrationAccessEnum(tok *tokenGenerator) (calibrationAccessEnum, error) {
	ca := undefinedCalibrationAccess
	var err error
	switch tok.next() {
	case calibrationToken:
		ca = calibration
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

// TypeEnum contains all types that an a2l characteristic might take
type TypeEnum string

const (
	undefinedType TypeEnum = emptyToken
	//String value in ASCII format.
	ASCII TypeEnum = asciiToken
	//Curve Datatype of a characteristic is like an array for the y axis plus an axis description on the x axis
	Curve TypeEnum = curveToken
	//Map defines a two dimensional array
	Map TypeEnum = mapToken
	//Cuboid defines a three dimensional array mostly represented as an array of Maps
	Cuboid TypeEnum = cuboidToken
	//The cuboid is stored as an array of maps with incremented or decremented Z coordinates. Rarely used
	Cube4 TypeEnum = cube4Token
	//The CUBE_5 is stored as an array of CUBE_4 with incremented or decremented Z5 coordinates. Rarely used
	Cube5 TypeEnum = cube5Token
	//Value Block is a a array with one, two or up to three dimensions.
	ValBlk TypeEnum = valBlkToken
	//Value is used in characteristics that only have one value like e.g. "1.0"
	Value      TypeEnum = valueToken
	Derived    TypeEnum = derivedToken
	ExtendedSi TypeEnum = extendedSiToken
)

func parseTypeEnum(tok *tokenGenerator) (TypeEnum, error) {
	t := undefinedType
	var err error
	switch tok.current() {
	case asciiToken:
		t = ASCII
	case curveToken:
		t = Curve
	case mapToken:
		t = Map
	case cuboidToken:
		t = Cuboid
	case cube4Token:
		t = Cube4
	case cube5Token:
		t = Cube5
	case valBlkToken:
		t = ValBlk
	case valueToken:
		t = Value
	case derivedToken:
		t = Derived
	case extendedSiToken:
		t = ExtendedSi
	default:
		err = errors.New("incorrect value " + tok.current() + " for enum type")
	}
	return t, err
}

type conversionTypeEnum string

const (
	undefinedConversionType conversionTypeEnum = emptyToken
	//Identical defines a OneToOne conversion from hex to decimal
	Identical conversionTypeEnum = identicalToken
	Form      conversionTypeEnum = formToken
	/*Linear function of the following type:
	f(x)=ax + b
	for which:
	PHYS=f(INT)
	The coefficients a and b have to be
	specified by the COEFFS_LINEAR
	keyword.*/
	Linear conversionTypeEnum = linearToken
	/*RatFunc is a fractional rational function of the following type:
	f(x)=(axx + bx + c)/(dxx + ex + f)
	for which:
	INT = f(PHYS)
	Coefficients a, b, c, d, e, f have to be specified by the COEFFS keyword.
	Note: For linear functions, use the	ConversionType LINEAR,
	for ident functions the ConversionType IDENT.
	For non linear functions it	must be possible to invert the formula within the limits of the
	AXIS_PTS, CHARACTERISTIC or	MEASUREMENT where it is used.
	Otherwise use the ConversionType FORM.*/
	RatFunc conversionTypeEnum = ratFuncToken
	//TabIntp defines a table with interpolation
	TabIntp conversionTypeEnum = tabIntpToken
	//TabIntp defines a table withOut interpolation
	TabNointp conversionTypeEnum = tabNointpToken
	//Tab Verb is a table to convert numeric values into strings. e.g.: 1 -> "True"
	TabVerb conversionTypeEnum = tabVerbToken
)

func parseConversionTypeEnum(tok *tokenGenerator) (conversionTypeEnum, error) {
	ct := undefinedConversionType
	var err error
	switch tok.current() {
	case identicalToken:
		ct = Identical
	case formToken:
		ct = Form
	case linearToken:
		ct = Linear
	case ratFuncToken:
		ct = RatFunc
	case tabIntpToken:
		ct = TabIntp
	case tabNointpToken:
		ct = TabNointp
	case tabVerbToken:
		ct = TabVerb
	default:
		err = errors.New("incorrect value " + tok.current() + " for enum conversionType")
	}
	return ct, err
}

type indexModeEnum string

const (
	undefinedIndexMode indexModeEnum = emptyToken
	/* curves which share a common axis are deposited in columns;
	each row of memory contains values for all the shared axis curves
	at a given axis breakpoint.
	Required in order to represent characteristics which correspond to
	arrays of structures in ECU program code.*/
	AlternateCurves indexModeEnum = alternateCurvesToken
	/*AlternateWithX defines that values of a map are
	stored in columns and the columns of table values alternate with the
	respective X-coordinates. A map of format
	 9 8 7
	[0 1 2
	 3 4 5
	 6 7 8]
	could be stored within the hex file as an array of format
	[9,0,3,6,8,1,4,7,7,2,5,8]
	The order of axis points and table values can be defined differently
	by the position statement in the FNC_VALUE
	In case of a curve the values of x-Axis and values alternate.*/
	AlternateWithX indexModeEnum = alternateWithXToken
	/*AlternateWithY defines that values of a map are
	deposited in rows, the rows of table values alternate with the
	respective Y-coordinates. A map of format
	9 [0 1 2
	8  3 4 5
	7  6 7 8]
	could be within the hex file as an array of format
	[9,0,1,2,8,3,4,5,7,6,7,8]
	The order of axis points and table values can be defined differently
	by the position statement in the FNC_VALUE
	Only applicable to maps*/
	AlternateWithY indexModeEnum = alternateWithYToken
	/*Column Direction defines that values of a map
	[0 1 2
	 3 4 5
	 6 7 8]
	are stored within the hex file as an array of format
	[0,3,6,1,4,7,2,5,8]	*/
	ColumnDir indexModeEnum = columnDirToken
	/*Row Direction defines that values of a map
	[0 1 2
	 3 4 5
	 6 7 8]
	are stored within the hex file as an array of format
	[0,1,2,3,4,5,6,7,8]	*/
	RowDir indexModeEnum = rowDirToken
)

func parseIndexModeEnum(tok *tokenGenerator) (indexModeEnum, error) {
	im := undefinedIndexMode
	var err error
	switch tok.current() {
	case alternateCurvesToken:
		im = AlternateCurves
	case alternateWithXToken:
		im = AlternateWithX
	case alternateWithYToken:
		im = AlternateWithY
	case columnDirToken:
		im = ColumnDir
	case rowDirToken:
		im = RowDir
	default:
		err = errors.New("incorrect value " + tok.current() + " for enum indexMode")
	}
	return im, err
}

type modeEnum string

const (
	undefinedMode modeEnum = emptyToken
	Absolute      modeEnum = absoluteToken
	Difference    modeEnum = differenceToken
)

func parseModeEnum(tok *tokenGenerator) (modeEnum, error) {
	m := undefinedMode
	var err error
	switch tok.current() {
	case absoluteToken:
		m = Absolute
	case differenceToken:
		m = Difference
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
	code                 prgTypeEnum = codeToken
	data                 prgTypeEnum = dataToken
	excludeFromFlash     prgTypeEnum = excludeFromFlashToken
	offlineData          prgTypeEnum = offlineDataToken
	reserved2            prgTypeEnum = reservedToken
	seram                prgTypeEnum = seramToken
	variables            prgTypeEnum = variablesToken
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
		pt = code
	case dataToken:
		pt = data
	case excludeFromFlashToken:
		pt = excludeFromFlash
	case offlineDataToken:
		pt = offlineData
	case reservedToken:
		pt = reserved2
	case seramToken:
		pt = seram
	case variablesToken:
		pt = variables
	default:
		err = errors.New("incorrect value " + tok.current() + " for enum prgType")
	}
	return pt, err
}

type memoryTypeEnum string

const (
	undefinedMemoryType memoryTypeEnum = emptyToken
	eeprom              memoryTypeEnum = eepromToken
	eprom               memoryTypeEnum = epromToken
	flash               memoryTypeEnum = flashToken
	ram                 memoryTypeEnum = ramToken
	rom                 memoryTypeEnum = romToken
	register            memoryTypeEnum = registerToken
	notInEcu            memoryTypeEnum = notInEcuToken
)

func parseMemoryTypeEnum(tok *tokenGenerator) (memoryTypeEnum, error) {
	mt := undefinedMemoryType
	var err error
	switch tok.current() {
	case eepromToken:
		mt = eeprom
	case epromToken:
		mt = eprom
	case flashToken:
		mt = flash
	case ramToken:
		mt = ram
	case romToken:
		mt = rom
	case registerToken:
		mt = register
	case notInEcuToken:
		mt = notInEcu
	default:
		err = errors.New("incorrect value " + tok.current() + " for enum memoryType")
	}
	return mt, err
}

type monotonyTypeEnum string

const (
	undefinedMonotonyType monotonyTypeEnum = emptyToken
	monDecrease           monotonyTypeEnum = monDecreaseToken
	monIncrease           monotonyTypeEnum = monIncreaseToken
	strictDecrease        monotonyTypeEnum = strictDecreaseToken
	strictIncrease        monotonyTypeEnum = strictIncreaseToken
	monotonous            monotonyTypeEnum = monotonousToken
	strictMon             monotonyTypeEnum = strictMonToken
	notMon                monotonyTypeEnum = notMonToken
)

func parseMonotonyTypeEnum(tok *tokenGenerator) (monotonyTypeEnum, error) {
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
		mt = monotonous
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
	numeric      tagEnum = numericToken
)

func parseTagEnum(tok *tokenGenerator) (tagEnum, error) {
	mt := undefinedTag
	var err error
	switch tok.current() {
	case numericToken:
		mt = numeric
	default:
		err = errors.New("incorrect value " + tok.current() + " for enum tag")
	}
	return mt, err
}

type triggerEnum string

const (
	undefinedTriggerEnum triggerEnum = emptyToken
	onChange             triggerEnum = onChangeToken
	onUserRequest        triggerEnum = onUserRequestToken
)

func parseTriggerEnum(tok *tokenGenerator) (triggerEnum, error) {
	mt := undefinedTriggerEnum
	var err error
	switch tok.current() {
	case onChangeToken:
		mt = onChange
	case onUserRequestToken:
		mt = onUserRequest
	default:
		err = errors.New("incorrect value " + tok.current() + " for enum trigger")
	}
	return mt, err
}
