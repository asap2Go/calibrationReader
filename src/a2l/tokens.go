package a2l

//general tokens:
const emptyToken = ""
const spaceToken = " "
const beginMultilineCommentToken = "/*"
const endMultilineCommentToken = "*/"
const beginLineCommentToken = "//"
const quotationMarkToken = "\""
const slashToken = "/"

TypeDefAxisTokens
TypeDefBlobTokens
TypeDefStructureTokens
TypeDefCharacteristicTokens
TypeDefMeasurementTokens
//a2l tokens:
const modelLinkToken = "MODEL_LINK"
const beginBlobToken = "/begin BLOB"
const endBlobToken = "/end BLOB"
const addressTypeToken = "ADDRESS_TYPE"
const arPrototypeOfToken = "AR_PROTOTYPE_OF"
const endArComponentToken = "/end AR_COMPONENT"
const beginArComponentToken = "/begin AR_COMPONENT"
const beginVirtualCharacteristicToken = "/begin VIRTUAL_CHARACTERISTIC"
const endVirtualCharacteristicToken = "/end VIRTUAL_CHARACTERISTIC"
const calibrationHandleTextToken = "CALIBRATION_HANDLE_TEXT"
const beginCalibrationHandleToken = "/begin CALIBRATION_HANDLE"
const endCalibrationHandleToken = "/end CALIBRATION_HANDLE"
const beginA2mlToken = "/begin A2ML"
const endA2mlToken = "/end A2ML"
const beginAxisPtsToken = "/begin AXIS_PTS"
const endAxisPtsToken = "/end AXIS_PTS"
const beginCharacteristicToken = "/begin CHARACTERISTIC"
const endCharacteristicToken = "/end CHARACTERISTIC"
const beginCompuMethodToken = "/begin COMPU_METHOD"
const endCompuMethodToken = "/end COMPU_METHOD"
const beginCompuTabToken = "/begin COMPU_TAB"
const endCompuTabToken = "/end COMPU_TAB"
const beginCompuVtabToken = "/begin COMPU_VTAB"
const endCompuVtabToken = "/end COMPU_VTAB"
const beginCompuVtabRangeToken = "/begin COMPU_VTAB_RANGE"
const endCompuVtabRangeToken = "/end COMPU_VTAB_RANGE"
const beginFrameToken = "/begin FRAME"
const endFrameToken = "/end FRAME"
const beginFunctionToken = "/begin FUNCTION"
const endFunctionToken = "/end FUNCTION"
const beginGroupToken = "/begin GROUP"
const endGroupToken = "/end GROUP"
const beginIfDataToken = "/begin IF_DATA"
const endIfDataToken = "/end IF_DATA"
const beginMeasurementToken = "/begin MEASUREMENT"
const endMeasurementToken = "/end MEASUREMENT"
const beginModCommonToken = "/begin MOD_COMMON"
const endModCommonToken = "/end MOD_COMMON"
const beginModParToken = "/begin MOD_PAR"
const endModParToken = "/end MOD_PAR"
const beginRecordLayoutToken = "/begin RECORD_LAYOUT"
const endRecordLayoutToken = "/end RECORD_LAYOUT"
const beginUnitToken = "/begin UNIT"
const endUnitToken = "/end UNIT"
const beginUserRightsToken = "/begin USER_RIGHTS"
const endUserRightsToken = "/end USER_RIGHTS"
const beginVariantCodingToken = "/begin VARIANT_CODING"
const endVariantCodingToken = "/end VARIANT_CODING"
const beginModuleToken = "/begin MODULE"
const endModuleToken = "/end MODULE"
const projectNoToken = "PROJECT_NO"
const versionToken = "VERSION"
const beginHeaderToken = "/begin HEADER"
const endHeaderToken = "/end HEADER"
const beginVarCharacteristicToken = "/begin VAR_CHARACTERISTIC"
const endVarCharacteristicToken = "/end VAR_CHARACTERISTIC"
const beginVarCriterionToken = "/begin VAR_CRITERION"
const endVarCriterionToken = "/end VAR_CRITERION"
const beginVarForbiddenCombToken = "/begin VAR_FORBIDDEN_COMB"
const endVarForbiddenCombToken = "/end VAR_FORBIDDEN_COMB"
const varNamingToken = "VAR_NAMING"
const beginAnnotationToken = "/begin ANNOTATION"
const endAnnotationToken = "/end ANNOTATION"
const beginAxisDescrToken = "/begin AXIS_DESCR"
const endAxisDescrToken = "/end AXIS_DESCR"
const bitMaskToken = "BIT_MASK"
const byteOrderToken = "BYTE_ORDER"
const calibrationAccessToken = "CALIBRATION_ACCESS"
const comparisonQuantityToken = "COMPARISON_QUANTITY"
const beginDependentCharacteristicToken = "/begin DEPENDENT_CHARACTERISTIC"
const endDependentCharacteristicToken = "/end DEPENDENT_CHARACTERISTIC"
const discreteToken = "DISCRETE"
const displayIdentifierToken = "DISPLAY_IDENTIFIER"
const ecuAddressExtensionToken = "ECU_ADDRESS_EXTENSION"
const extendedLimitsToken = "EXTENDED_LIMITS"
const formatToken = "FORMAT"
const beginFunctionListToken = "/begin FUNCTION_LIST"
const endFunctionListToken = "/end FUNCTION_LIST"
const guardRailsToken = "GUARD_RAILS"
const beginMapListToken = "/begin MAP_LIST"
const endMapListToken = "/end MAP_LIST"
const matrixDimToken = "MATRIX_DIM"
const maxRefreshToken = "MAX_REFRESH"
const numberToken = "NUMBER"
const physUnitToken = "PHYS_UNIT"
const readOnlyToken = "READ_ONLY"
const refMemorySegmentToken = "REF_MEMORY_SEGMENT"
const stepSizeToken = "STEP_SIZE"
const symbolLinkToken = "SYMBOL_LINK"
const defaultValueToken = "DEFAULT_VALUE"
const beginOutMeasurementToken = "/begin OUT_MEASUREMENT"
const endOutMeasurementToken = "/end OUT_MEASUREMENT"
const beginVarAddressToken = "/begin VAR_ADDRESS"
const endVarAddressToken = "/end VAR_ADDRESS"
const beginRefCharacteristicToken = "/begin REF_CHARACTERISTIC"
const endRefCharacteristicToken = "/end REF_CHARACTERISTIC"
const beginLocMeasurementToken = "/begin LOC_MEASUREMENT"
const endLocMeasurementToken = "/end LOC_MEASUREMENT"
const alignmentByteToken = "ALIGNMENT_BYTE"
const alignmentFloat32IeeeToken = "ALIGNMENT_FLOAT32_IEEE"
const alignmentFloat64IeeeToken = "ALIGNMENT_FLOAT64_IEEE"
const alignmentInt64Token = "ALIGNMENT_INT64"
const alignmentLongToken = "ALIGNMENT_LONG"
const alignmentWordToken = "ALIGNMENT_WORD"
const axisPtsXToken = "AXIS_PTS_X"
const axisPtsYToken = "AXIS_PTS_Y"
const axisPtsZToken = "AXIS_PTS_Z"
const axisPts4Token = "AXIS_PTS_4"
const axisPts5Token = "AXIS_PTS_5"
const axisRescaleXToken = "AXIS_RESCALE_X"
const distOpXToken = "DIST_OP_X"
const distOpYToken = "DIST_OP_Y"
const distOpZToken = "DIST_OP_Z"
const distOp4Token = "DIST_OP_4"
const distOp5Token = "DIST_OP_5"
const fixNoAxisPtsXToken = "FIX_NO_AXIS_PTS_X"
const fixNoAxisPtsYToken = "FIX_NO_AXIS_PTS_Y"
const fixNoAxisPtsZToken = "FIX_NO_AXIS_PTS_Z"
const fixNoAxisPts4Token = "FIX_NO_AXIS_PTS_4"
const fixNoAxisPts5Token = "FIX_NO_AXIS_PTS_5"
const fncValuesToken = "FNC_VALUES"
const identificationToken = "IDENTIFICATION"
const noAxisPtsXToken = "NO_AXIS_PTS_X"
const noAxisPtsYToken = "NO_AXIS_PTS_Y"
const noAxisPtsZToken = "NO_AXIS_PTS_Z"
const noAxisPts4Token = "NO_AXIS_PTS_4"
const noAxisPts5Token = "NO_AXIS_PTS_5"
const noRescaleXToken = "NO_RESCALE_X"
const offsetXToken = "OFFSET_X"
const offsetYToken = "OFFSET_Y"
const offsetZToken = "OFFSET_Z"
const offset4Token = "OFFSET_4"
const offset5Token = "OFFSET_5"
const ripAddrWToken = "RIP_ADDR_W"
const ripAddrXToken = "RIP_ADDR_X"
const ripAddrYToken = "RIP_ADDR_Y"
const ripAddrZToken = "RIP_ADDR_Z"
const ripAddr4Token = "RIP_ADDR_4"
const ripAddr5Token = "RIP_ADDR_5"
const srcAddrXToken = "SRC_ADDR_X"
const srcAddrYToken = "SRC_ADDR_Y"
const srcAddrZToken = "SRC_ADDR_Z"
const srcAddr4Token = "SRC_ADDR_4"
const srcAddr5Token = "SRC_ADDR_5"
const shiftOpXToken = "SHIFT_OP_X"
const shiftOpYToken = "SHIFT_OP_Y"
const shiftOpZToken = "SHIFT_OP_Z"
const shiftOp4Token = "SHIFT_OP_4"
const shiftOp5Token = "SHIFT_OP_5"
const reservedToken = "RESERVED"
const staticRecordLayoutToken = "STATIC_RECORD_LAYOUT"
const beginDefCharacteristicToken = "/begin DEF_CHARACTERISTIC"
const endDefCharacteristicToken = "/end DEF_CHARACTERISTIC"
const annotationLabelToken = "ANNOTATION_LABEL"
const annotationOriginToken = "ANNOTATION_ORIGIN"
const beginAnnotationTextToken = "/begin ANNOTATION_TEXT"
const endAnnotationTextToken = "/end ANNOTATION_TEXT"
const leftShiftToken = "LEFT_SHIFT"
const rightShiftToken = "RIGHT_SHIFT"
const signExtendToken = "SIGN_EXTEND"
const beginBitOperationToken = "/begin BIT_OPERATION"
const endBitOperationToken = "/end BIT_OPERATION"
const beginMemoryLayoutToken = "/begin MEMORY_LAYOUT"
const endMemoryLayoutToken = "/end MEMORY_LAYOUT"
const depositToken = "DEPOSIT"
const monotonyToken = "MONOTONY"
const beginTypeToken = "/begin TYPE"
const endTypeToken = "/end TYPE"
const refUnitToken = "REF_UNIT"
const siExponentsToken = "SI_EXPONENTS"
const unitConversionToken = "UNIT_CONVERSION"
const beginProjectToken = "/begin PROJECT"
const endProjectToken = "/end PROJECT"
const defaultValueNumericToken = "DEFAULT_VALUE_NUMERIC"
const frameMeasurementToken = "FRAME_MEASUREMENT"
const beginRefMeasurementToken = "/begin REF_MEASUREMENT"
const endRefMeasurementToken = "/end REF_MEASUREMENT"
const rootToken = "ROOT"
const beginSubGroupToken = "/begin SUB_GROUP"
const endSubGroupToken = "/end SUB_GROUP"
const addrEpkToken = "ADDR_EPK"
const beginCalibrationMethodToken = "/begin CALIBRATION_METHOD"
const endCalibrationMethodToken = "/end CALIBRATION_METHOD"
const cpuTypeToken = "CPU_TYPE"
const customerToken = "CUSTOMER"
const customerNoToken = "CUSTOMER_NO"
const ecuToken = "ECU"
const ecuCalibrationOffsetToken = "ECU_CALIBRATION_OFFSET"
const epkToken = "EPK"
const beginMemorySegmentToken = "/begin MEMORY_SEGMENT"
const endMemorySegmentToken = "/end MEMORY_SEGMENT"
const noOfInterfacesToken = "NO_OF_INTERFACES"
const phoneNoToken = "PHONE_NO"
const supplierToken = "SUPPLIER"
const systemConstantToken = "SYSTEM_CONSTANT"
const beginUserToken = "/begin USER"
const endUserToken = "/end USER"
const beginInMeasurementToken = "/begin IN_MEASUREMENT"
const endInMeasurementToken = "/end IN_MEASUREMENT"
const formulaInvToken = "FORMULA_INV"
const beginFormulaToken = "/begin FORMULA"
const endFormulaToken = "/end FORMULA"
const beginRefGroupToken = "/begin REF_GROUP"
const endRefGroupToken = "/end REF_GROUP"
const arraySizeToken = "ARRAY_SIZE"
const ecuAddressToken = "ECU_ADDRESS"
const errorMaskToken = "ERROR_MASK"
const layoutToken = "LAYOUT"
const readWriteToken = "READ_WRITE"
const beginVirtualToken = "/begin VIRTUAL"
const endVirtualToken = "/end VIRTUAL"
const asap2VersionToken = "ASAP2_VERSION"
const a2mlVersionToken = "A2ML_VERSION"
const ubyteToken = "UBYTE"
const sbyteToken = "SBYTE"
const uwordToken = "UWORD"
const swordToken = "SWORD"
const ulongToken = "ULONG"
const slongToken = "SLONG"
const aUint64Token = "A_UINT64"
const aInt64Token = "A_INT64"
const float32IeeeToken = "FLOAT32_IEEE"
const float64IeeeToken = "FLOAT64_IEEE"
const byteToken = "BYTE"
const wordToken = "WORD"
const longToken = "LONG"
const pbyteToken = "PBYTE"
const pwordToken = "PWORD"
const plongToken = "PLONG"
const directToken = "DIRECT"
const littleEndianToken = "LITTLE_ENDIAN"
const bigEndianToken = "BIG_ENDIAN"
const msbLastToken = "MSB_LAST"
const msbFirstToken = "MSB_FIRST"
const indexIncrToken = "INDEX_INCR"
const indexDecrToken = "INDEX_DECR"
const curveAxisToken = "CURVE_AXIS"
const comAxisToken = "COM_AXIS"
const beginFixAxisToken = "/begin FIX_AXIS"
const endFixAxisToken = "/end FIX_AXIS"
const resAxisToken = "RES_AXIS"
const stdAxisToken = "STD_AXIS"
const internToken = "INTERN"
const externToken = "EXTERN"
const beginCalibrationToken = "/begin CALIBRATION"
const endCalibrationToken = "/end CALIBRATION"
const noCalibrationToken = "NO_CALIBRATION"
const notInMcdSystemToken = "NOT_IN_MCD_SYSTEM"
const offlineCalibrationToken = "OFFLINE_CALIBRATION"
const asciiToken = "ASCII"
const curveToken = "CURVE"
const beginMapToken = "/begin MAP"
const endMapToken = "/end MAP"
const cuboidToken = "CUBOID"
const cube4Token = "CUBE_4"
const cube5Token = "CUBE_5"
const valBlkToken = "VAL_BLK"
const valueToken = "VALUE"
const derivedToken = "DERIVED"
const extendedSiToken = "EXTENDED_SI"
const identicalToken = "IDENTICAL"
const beginFormToken = "/begin FORM"
const endFormToken = "/end FORM"
const linearToken = "LINEAR"
const ratFuncToken = "RAT_FUNC"
const tabIntpToken = "TAB_INTP"
const tabNointpToken = "TAB_NOINTP"
const tabVerbToken = "TAB_VERB"
const alternateCurvesToken = "ALTERNATE_CURVES"
const alternateWithXToken = "ALTERNATE_WITH_X"
const alternateWithYToken = "ALTERNATE_WITH_Y"
const columnDirToken = "COLUMN_DIR"
const rowDirToken = "ROW_DIR"
const absoluteToken = "ABSOLUTE"
const differenceToken = "DIFFERENCE"
const prgCodeToken = "PRG_CODE"
const prgDataToken = "PRG_DATA"
const prgReservedToken = "PRG_RESERVED"
const calibrationVariablesToken = "CALIBRATION_VARIABLES"
const codeToken = "CODE"
const dataToken = "DATA"
const excludeFromFlashToken = "EXCLUDE_FROM_FLASH"
const offlineDataToken = "OFFLINE_DATA"
const seramToken = "SERAM"
const variablesToken = "VARIABLES"
const eepromToken = "EEPROM"
const epromToken = "EPROM"
const flashToken = "FLASH"
const ramToken = "RAM"
const romToken = "ROM"
const registerToken = "REGISTER"
const monDecreaseToken = "MON_DECREASE"
const monIncreaseToken = "MON_INCREASE"
const strictDecreaseToken = "STRICT_DECREASE"
const strictIncreaseToken = "STRICT_INCREASE"
const monotonousToken = "MONOTONOUS"
const strictMonToken = "STRICT_MON"
const notMonToken = "NOT_MON"
const numericToken = "NUMERIC"
const varMeasurementToken = "VAR_MEASUREMENT"
const varSelectionCharacteristicToken = "VAR_SELECTION_CHARACTERISTIC"
const beginSubFunctionToken = "/begin SUB_FUNCTION"
const endSubFunctionToken = "/end SUB_FUNCTION"
const functionVersionToken = "FUNCTION_VERSION"
const axisPtsRefToken = "AXIS_PTS_REF"
const curveAxisRefToken = "CURVE_AXIS_REF"
const beginFixAxisParToken = "/begin FIX_AXIS_PAR"
const endFixAxisParToken = "/end FIX_AXIS_PAR"
const fixAxisParDistToken = "FIX_AXIS_PAR_DIST"
const beginFixAxisParListToken = "/begin FIX_AXIS_PAR_LIST"
const endFixAxisParListToken = "/end FIX_AXIS_PAR_LIST"
const maxGradToken = "MAX_GRAD"
const dataSizeToken = "DATA_SIZE"
const coeffsToken = "COEFFS"
const coeffsLinearToken = "COEFFS_LINEAR"
const compuTabRefToken = "COMPU_TAB_REF"
const statusStringRefToken = "STATUS_STRING_REF"
const fixAxisToken = "FIX_AXIS"
const notInEcuToken = "NOT_IN_ECU"
const calibrationToken = "CALIBRATION"
const mapToken = "MAP"
const formToken = "FORM"

var keywordList = []string{
	beginVirtualCharacteristicToken,
	endVirtualCharacteristicToken,
	calibrationHandleTextToken,
	beginCalibrationHandleToken,
	endCalibrationHandleToken,
	beginA2mlToken,
	endA2mlToken,
	beginAxisPtsToken,
	endAxisPtsToken,
	beginCharacteristicToken,
	endCharacteristicToken,
	beginCompuMethodToken,
	endCompuMethodToken,
	beginCompuTabToken,
	endCompuTabToken,
	beginCompuVtabToken,
	endCompuVtabToken,
	beginCompuVtabRangeToken,
	endCompuVtabRangeToken,
	beginFrameToken,
	endFrameToken,
	beginFunctionToken,
	endFunctionToken,
	beginGroupToken,
	endGroupToken,
	beginIfDataToken,
	endIfDataToken,
	beginMeasurementToken,
	endMeasurementToken,
	beginModCommonToken,
	endModCommonToken,
	beginModParToken,
	endModParToken,
	beginRecordLayoutToken,
	endRecordLayoutToken,
	beginUnitToken,
	endUnitToken,
	beginUserRightsToken,
	endUserRightsToken,
	beginVariantCodingToken,
	endVariantCodingToken,
	beginModuleToken,
	endModuleToken,
	projectNoToken,
	versionToken,
	beginHeaderToken,
	endHeaderToken,
	beginVarCharacteristicToken,
	endVarCharacteristicToken,
	beginVarCriterionToken,
	endVarCriterionToken,
	beginVarForbiddenCombToken,
	endVarForbiddenCombToken,
	varNamingToken,
	beginAnnotationToken,
	endAnnotationToken,
	beginAxisDescrToken,
	endAxisDescrToken,
	bitMaskToken,
	byteOrderToken,
	calibrationAccessToken,
	comparisonQuantityToken,
	beginDependentCharacteristicToken,
	endDependentCharacteristicToken,
	discreteToken,
	displayIdentifierToken,
	ecuAddressExtensionToken,
	extendedLimitsToken,
	formatToken,
	beginFunctionListToken,
	endFunctionListToken,
	guardRailsToken,
	beginMapListToken,
	endMapListToken,
	matrixDimToken,
	maxRefreshToken,
	numberToken,
	physUnitToken,
	readOnlyToken,
	refMemorySegmentToken,
	stepSizeToken,
	symbolLinkToken,
	defaultValueToken,
	beginOutMeasurementToken,
	endOutMeasurementToken,
	beginVarAddressToken,
	endVarAddressToken,
	beginRefCharacteristicToken,
	endRefCharacteristicToken,
	beginLocMeasurementToken,
	endLocMeasurementToken,
	alignmentByteToken,
	alignmentFloat32IeeeToken,
	alignmentFloat64IeeeToken,
	alignmentInt64Token,
	alignmentLongToken,
	alignmentWordToken,
	axisPtsXToken,
	axisRescaleXToken,
	distOpXToken,
	fixNoAxisPtsXToken,
	fncValuesToken,
	identificationToken,
	noAxisPtsXToken,
	noRescaleXToken,
	offsetXToken,
	reservedToken,
	ripAddrWToken,
	ripAddrXToken,
	srcAddrXToken,
	shiftOpXToken,
	staticRecordLayoutToken,
	beginDefCharacteristicToken,
	endDefCharacteristicToken,
	annotationLabelToken,
	annotationOriginToken,
	beginAnnotationTextToken,
	endAnnotationTextToken,
	leftShiftToken,
	rightShiftToken,
	signExtendToken,
	beginBitOperationToken,
	endBitOperationToken,
	beginMemoryLayoutToken,
	endMemoryLayoutToken,
	depositToken,
	monotonyToken,
	beginTypeToken,
	endTypeToken,
	refUnitToken,
	siExponentsToken,
	unitConversionToken,
	beginProjectToken,
	endProjectToken,
	defaultValueNumericToken,
	frameMeasurementToken,
	beginRefMeasurementToken,
	endRefMeasurementToken,
	rootToken,
	beginSubGroupToken,
	endSubGroupToken,
	addrEpkToken,
	beginCalibrationMethodToken,
	endCalibrationMethodToken,
	cpuTypeToken,
	customerToken,
	customerNoToken,
	ecuToken,
	ecuCalibrationOffsetToken,
	epkToken,
	beginMemorySegmentToken,
	endMemorySegmentToken,
	noOfInterfacesToken,
	phoneNoToken,
	supplierToken,
	systemConstantToken,
	beginUserToken,
	endUserToken,
	beginInMeasurementToken,
	endInMeasurementToken,
	formulaInvToken,
	beginFormulaToken,
	endFormulaToken,
	beginRefGroupToken,
	endRefGroupToken,
	arraySizeToken,
	ecuAddressToken,
	errorMaskToken,
	layoutToken,
	readWriteToken,
	beginVirtualToken,
	endVirtualToken,
	asap2VersionToken,
	a2mlVersionToken,
	ubyteToken,
	sbyteToken,
	uwordToken,
	swordToken,
	ulongToken,
	slongToken,
	aUint64Token,
	aInt64Token,
	float32IeeeToken,
	float64IeeeToken,
	byteToken,
	wordToken,
	longToken,
	pbyteToken,
	pwordToken,
	plongToken,
	directToken,
	littleEndianToken,
	bigEndianToken,
	msbLastToken,
	msbFirstToken,
	indexIncrToken,
	indexDecrToken,
	curveAxisToken,
	comAxisToken,
	beginFixAxisToken,
	endFixAxisToken,
	resAxisToken,
	stdAxisToken,
	internToken,
	externToken,
	beginCalibrationToken,
	endCalibrationToken,
	noCalibrationToken,
	notInMcdSystemToken,
	offlineCalibrationToken,
	asciiToken,
	curveToken,
	beginMapToken,
	endMapToken,
	cuboidToken,
	cube4Token,
	cube5Token,
	valBlkToken,
	valueToken,
	derivedToken,
	extendedSiToken,
	identicalToken,
	beginFormToken,
	endFormToken,
	linearToken,
	ratFuncToken,
	tabIntpToken,
	tabNointpToken,
	tabVerbToken,
	alternateCurvesToken,
	alternateWithXToken,
	alternateWithYToken,
	columnDirToken,
	rowDirToken,
	absoluteToken,
	differenceToken,
	prgCodeToken,
	prgDataToken,
	prgReservedToken,
	calibrationVariablesToken,
	codeToken,
	dataToken,
	excludeFromFlashToken,
	offlineDataToken,
	seramToken,
	variablesToken,
	eepromToken,
	epromToken,
	flashToken,
	ramToken,
	romToken,
	registerToken,
	monDecreaseToken,
	monIncreaseToken,
	strictDecreaseToken,
	strictIncreaseToken,
	monotonousToken,
	strictMonToken,
	notMonToken,
	numericToken,
	varMeasurementToken,
	varSelectionCharacteristicToken,
	beginSubFunctionToken,
	endSubFunctionToken,
	functionVersionToken,
	axisPtsRefToken,
	curveAxisRefToken,
	beginFixAxisParToken,
	endFixAxisParToken,
	fixAxisParDistToken,
	beginFixAxisParListToken,
	endFixAxisParListToken,
	maxGradToken,
	dataSizeToken,
	coeffsToken,
	coeffsLinearToken,
	compuTabRefToken,
	statusStringRefToken,
	fixAxisToken,
	notInEcuToken,
	calibrationToken,
	mapToken,
	formToken,
	axisPtsYToken,
	axisPtsZToken,
	axisPts4Token,
	axisPts5Token,
	distOpYToken,
	distOpZToken,
	distOp4Token,
	distOp5Token,
	fixNoAxisPtsYToken,
	fixNoAxisPtsZToken,
	fixNoAxisPts4Token,
	fixNoAxisPts5Token,
	offsetYToken,
	offsetZToken,
	offset4Token,
	offset5Token,
	ripAddrYToken,
	ripAddrZToken,
	ripAddr4Token,
	ripAddr5Token,
	srcAddrYToken,
	srcAddrZToken,
	srcAddr4Token,
	srcAddr5Token,
	shiftOpYToken,
	shiftOpZToken,
	shiftOp4Token,
	shiftOp5Token,
	arPrototypeOfToken,
	endArComponentToken,
	beginArComponentToken,
	addressTypeToken,
	beginBlobToken,
	endBlobToken,
	modelLinkToken,
}
