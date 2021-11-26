package a2l

import (
	"errors"
	"fmt"
	"sync"

	"github.com/rs/zerolog/log"
)

/*ToDo:
[->TYPEDEF_MEASUREMENT]*
[->TYPEDEF_STRUCTURE]*
*/
type module struct {
	name                   string
	nameSet                bool
	longIdentifier         string
	longIdentifierSet      bool
	a2ml                   a2ml
	errors                 []error
	AxisPts                map[string]axisPts
	blobs                  map[string]blob
	Characteristics        map[string]characteristic
	CompuMethods           map[string]compuMethod
	CompuTabs              map[string]compuTab
	CompuVTabs             map[string]compuVTab
	CompuVTabRanges        map[string]compuVTabRange
	frame                  frame
	Functions              map[string]function
	Groups                 map[string]group
	ifData                 map[string]IfData
	instances              map[string]instance
	Measurements           map[string]measurement
	ModCommon              modCommon
	ModPar                 modPar
	RecordLayouts          map[string]recordLayout
	transformers           map[string]transformer
	typeDefAxis            map[string]typeDefAxis
	typeDefBlobs           map[string]typeDefBlob
	typeDefCharacteristics map[string]typeDefCharacteristic
	typeDefMeasurements    map[string]typeDefMeasurement
	typeDefStructures      map[string]typeDefStructure
	Units                  map[string]unit
	userRights             map[string]userRights
	variantCoding          variantCoding
}

func parseModule(tok *tokenGenerator) (module, error) {
	//Bulk init of an average number of objects contained in a modern a2l-file.
	myModule := module{}
	myModule.AxisPts = make(map[string]axisPts, 1000)
	myModule.blobs = make(map[string]blob, 5)
	myModule.Characteristics = make(map[string]characteristic, 10000)
	myModule.CompuMethods = make(map[string]compuMethod, 1000)
	myModule.CompuTabs = make(map[string]compuTab, 1000)
	myModule.CompuVTabs = make(map[string]compuVTab, 1000)
	myModule.CompuVTabRanges = make(map[string]compuVTabRange, 1000)
	myModule.Functions = make(map[string]function, 10000)
	myModule.Groups = make(map[string]group, 1000)
	myModule.ifData = make(map[string]IfData, 1000)
	myModule.instances = make(map[string]instance, 10)
	myModule.Measurements = make(map[string]measurement, 10000)
	myModule.RecordLayouts = make(map[string]recordLayout, 1000)
	myModule.transformers = make(map[string]transformer, 10)
	myModule.typeDefAxis = make(map[string]typeDefAxis, 10)
	myModule.typeDefBlobs = make(map[string]typeDefBlob, 10)
	myModule.typeDefCharacteristics = make(map[string]typeDefCharacteristic, 10)
	myModule.typeDefMeasurements = make(map[string]typeDefMeasurement, 10)
	myModule.typeDefStructures = make(map[string]typeDefStructure, 10)
	myModule.Units = make(map[string]unit, 1000)
	myModule.userRights = make(map[string]userRights, 1000)
	var err error
	var bufAxisPts axisPts
	var bufBlob blob
	var bufCharacteristic characteristic
	var bufCompuMethod compuMethod
	var bufCompuTab compuTab
	var bufCompuVtab compuVTab
	var bufCompuVtabRange compuVTabRange
	var bufFunction function
	var bufGroup group
	var bufIfData IfData
	var bufInstance instance
	var bufMeasurement measurement
	var bufRecordLayout recordLayout
	var bufTransformer transformer
	var bufTypeDefAxis typeDefAxis
	var bufTypeDefBlob typeDefBlob
	var bufTypeDefCharacteristic typeDefCharacteristic
	var bufTypeDefMeasurement typeDefMeasurement
	var bufTypeDefStructure typeDefStructure
	var bufUnit unit
	var bufUserRights userRights

forLoop:
	for {
		switch tok.next() {
		case beginA2mlToken:
			myModule.a2ml, err = parseA2ML(tok)
			if err != nil {
				log.Err(err).Msg("module a2ml could not be parsed")
				break forLoop
			}
			log.Info().Msg("module a2ml successfully parsed")
		case beginAxisPtsToken:
			bufAxisPts, err = parseAxisPts(tok)
			if err != nil {
				log.Err(err).Msg("module axisPts could not be parsed")
				break forLoop
			}
			myModule.AxisPts[bufAxisPts.name] = bufAxisPts
			log.Info().Msg("module axisPts successfully parsed")
		case beginBlobToken:
			bufBlob, err = parseBlob(tok)
			if err != nil {
				log.Err(err).Msg("module blob could not be parsed")
				break forLoop
			}
			myModule.blobs[bufBlob.name] = bufBlob
			log.Info().Msg("module blob successfully parsed")
		case beginCharacteristicToken:
			bufCharacteristic, err = parseCharacteristic(tok)
			if err != nil {
				log.Err(err).Msg("module characteristic could not be parsed")
				break forLoop
			}
			myModule.Characteristics[bufCharacteristic.Name] = bufCharacteristic
			log.Info().Msg("module characteristic successfully parsed")
		case beginCompuMethodToken:
			bufCompuMethod, err = parseCompuMethod(tok)
			if err != nil {
				log.Err(err).Msg("module compuMethod could not be parsed")
				break forLoop
			}
			myModule.CompuMethods[bufCompuMethod.name] = bufCompuMethod
			log.Info().Msg("module compuMethod successfully parsed")
		case beginCompuTabToken:
			bufCompuTab, err = parseCompuTab(tok)
			if err != nil {
				log.Err(err).Msg("module compuTab could not be parsed")
				break forLoop
			}
			myModule.CompuTabs[bufCompuTab.name] = bufCompuTab
			log.Info().Msg("module compuTab successfully parsed")
		case beginCompuVtabToken:
			bufCompuVtab, err = parseCompuVtab(tok)
			if err != nil {
				log.Err(err).Msg("module compuVtab could not be parsed")
				break forLoop
			}
			myModule.CompuVTabs[bufCompuVtab.name] = bufCompuVtab
			log.Info().Msg("module compuVtab successfully parsed")
		case beginCompuVtabRangeToken:
			bufCompuVtabRange, err = parseCompuVtabRange(tok)
			if err != nil {
				log.Err(err).Msg("module compuVtabRange could not be parsed")
				break forLoop
			}
			myModule.CompuVTabRanges[bufCompuVtabRange.name] = bufCompuVtabRange
			log.Info().Msg("module compuVtabRange successfully parsed")
		case beginFrameToken:
			myModule.frame, err = parseFrame(tok)
			if err != nil {
				log.Err(err).Msg("module frame could not be parsed")
				break forLoop
			}
			log.Info().Msg("module frame successfully parsed")
		case beginFunctionToken:
			bufFunction, err = parseFunction(tok)
			if err != nil {
				log.Err(err).Msg("module function could not be parsed")
				break forLoop
			}
			myModule.Functions[bufFunction.name] = bufFunction
			log.Info().Msg("module function successfully parsed")
		case beginGroupToken:
			bufGroup, err = parseGroup(tok)
			if err != nil {
				log.Err(err).Msg("module group could not be parsed")
				break forLoop
			}
			myModule.Groups[bufGroup.groupName] = bufGroup
			log.Info().Msg("module group successfully parsed")
		case beginIfDataToken:
			bufIfData, err = parseIfData(tok)
			if err != nil {
				log.Err(err).Msg("module ifData could not be parsed")
				break forLoop
			}
			myModule.ifData[bufIfData.name] = bufIfData
			log.Info().Msg("module ifData successfully parsed")
		case beginInstanceToken:
			bufInstance, err = parseInstance(tok)
			if err != nil {
				log.Err(err).Msg("module instance could not be parsed")
				break forLoop
			}
			myModule.instances[bufInstance.name] = bufInstance
			log.Info().Msg("module instance successfully parsed")
		case beginMeasurementToken:
			bufMeasurement, err = parseMeasurement(tok)
			if err != nil {
				log.Err(err).Msg("module measurement could not be parsed")
				break forLoop
			}
			myModule.Measurements[bufMeasurement.name] = bufMeasurement
			log.Info().Msg("module measurement successfully parsed")
		case beginModCommonToken:
			myModule.ModCommon, err = parseModCommon(tok)
			if err != nil {
				log.Err(err).Msg("module modCommon could not be parsed")
				break forLoop
			}
			log.Info().Msg("module modCommon successfully parsed")
		case beginModParToken:
			myModule.ModPar, err = parseModPar(tok)
			if err != nil {
				log.Err(err).Msg("module modPar could not be parsed")
				break forLoop
			}
			log.Info().Msg("module modPar successfully parsed")
		case beginRecordLayoutToken:
			bufRecordLayout, err = parseRecordLayout(tok)
			if err != nil {
				log.Err(err).Msg("module recordLayout could not be parsed")
				break forLoop
			}
			myModule.RecordLayouts[bufRecordLayout.name] = bufRecordLayout
			log.Info().Msg("module recordLayout successfully parsed")
		case beginTransformerToken:
			bufTransformer, err = parseTransformer(tok)
			if err != nil {
				log.Err(err).Msg("module transformer could not be parsed")
				break forLoop
			}
			myModule.transformers[bufTransformer.name] = bufTransformer
			log.Info().Msg("module transformer successfully parsed")
		case beginTypeDefAxisToken:
			bufTypeDefAxis, err = parseTypeDefAxis(tok)
			if err != nil {
				log.Err(err).Msg("module typeDefAxis could not be parsed")
				break forLoop
			}
			myModule.typeDefAxis[bufTypeDefAxis.name] = bufTypeDefAxis
			log.Info().Msg("module typeDefAxis successfully parsed")
		case beginTypeDefBlobToken:
			bufTypeDefBlob, err = parseTypeDefBlob(tok)
			if err != nil {
				log.Err(err).Msg("module typeDefBlob could not be parsed")
				break forLoop
			}
			myModule.typeDefBlobs[bufTypeDefBlob.name] = bufTypeDefBlob
			log.Info().Msg("module typeDefBlob successfully parsed")
		case beginTypeDefCharacteristicToken:
			bufTypeDefCharacteristic, err = parseTypeDefCharacteristic(tok)
			if err != nil {
				log.Err(err).Msg("module typeDefCharacteristic could not be parsed")
				break forLoop
			}
			myModule.typeDefCharacteristics[bufTypeDefCharacteristic.name] = bufTypeDefCharacteristic
			log.Info().Msg("module typeDefCharacteristic successfully parsed")
		case beginTypeDefMeasurementToken:
			bufTypeDefMeasurement, err = parseTypeDefMeasurement(tok)
			if err != nil {
				log.Err(err).Msg("module typeDefMeasurement could not be parsed")
				break forLoop
			}
			myModule.typeDefMeasurements[bufTypeDefMeasurement.name] = bufTypeDefMeasurement
			log.Info().Msg("module typeDefMeasurement successfully parsed")
		case beginTypeDefStructureToken:
			bufTypeDefStructure, err = parseTypeDefStructure(tok)
			if err != nil {
				log.Err(err).Msg("module typeDefStructure could not be parsed")
				break forLoop
			}
			myModule.typeDefStructures[bufTypeDefStructure.name] = bufTypeDefStructure
			log.Info().Msg("module typeDefStructure successfully parsed")
		case beginUnitToken:
			bufUnit, err = parseUnit(tok)
			if err != nil {
				log.Err(err).Msg("module unit could not be parsed")
				break forLoop
			}
			myModule.Units[bufUnit.name] = bufUnit
			log.Info().Msg("module unit successfully parsed")
		case beginUserRightsToken:
			bufUserRights, err = parseUserRights(tok)
			if err != nil {
				log.Err(err).Msg("module userRights could not be parsed")
				break forLoop
			}
			myModule.userRights[bufUserRights.userLevelId] = bufUserRights
			log.Info().Msg("module userRights successfully parsed")
		case beginVariantCodingToken:
			myModule.variantCoding, err = parseVariantCoding(tok)
			if err != nil {
				log.Err(err).Msg("module variantCoding could not be parsed")
				break forLoop
			}
			log.Info().Msg("module variantCoding successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("module could not be parsed")
				break forLoop
			} else if tok.current() == endModuleToken {
				break forLoop
			} else if !myModule.nameSet {
				myModule.name = tok.current()
				myModule.nameSet = true
				log.Info().Msg("module name successfully parsed")
			} else if !myModule.longIdentifierSet {
				myModule.longIdentifier = tok.current()
				myModule.longIdentifierSet = true
				log.Info().Msg("module longIdentifier successfully parsed")
			}
		}
	}
	return myModule, err
}

//parseModuleMultithreaded is the parallel parsing version of parseModule.
//it computes the start and the end of the module struct
//and splits it up among numProc number of goroutines
//which each execute a separate moduleMainLoop
func parseModuleMultithreaded(tok *tokenGenerator) (module, error) {
	//Bulk init of an average number of objects contained in a modern a2l-file.
	log.Info().Msg("creating maps for module subtypes")
	myModule := module{}
	myModule.AxisPts = make(map[string]axisPts, 1000)
	myModule.blobs = make(map[string]blob, 5)
	myModule.Characteristics = make(map[string]characteristic, 10000)
	myModule.CompuMethods = make(map[string]compuMethod, 1000)
	myModule.CompuTabs = make(map[string]compuTab, 1000)
	myModule.CompuVTabs = make(map[string]compuVTab, 1000)
	myModule.CompuVTabRanges = make(map[string]compuVTabRange, 1000)
	myModule.Functions = make(map[string]function, 10000)
	myModule.Groups = make(map[string]group, 1000)
	myModule.ifData = make(map[string]IfData, 1000)
	myModule.instances = make(map[string]instance, 10)
	myModule.Measurements = make(map[string]measurement, 10000)
	myModule.RecordLayouts = make(map[string]recordLayout, 1000)
	myModule.transformers = make(map[string]transformer, 10)
	myModule.typeDefAxis = make(map[string]typeDefAxis, 10)
	myModule.typeDefBlobs = make(map[string]typeDefBlob, 10)
	myModule.typeDefCharacteristics = make(map[string]typeDefCharacteristic, 10)
	myModule.typeDefMeasurements = make(map[string]typeDefMeasurement, 10)
	myModule.typeDefStructures = make(map[string]typeDefStructure, 10)
	myModule.Units = make(map[string]unit, 1000)
	myModule.userRights = make(map[string]userRights, 1000)
	var err error

forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("module could not be parsed")
			break forLoop
		} else if !myModule.nameSet {
			myModule.name = tok.current()
			myModule.nameSet = true
			log.Info().Msg("module name successfully parsed")
		} else if !myModule.longIdentifierSet {
			myModule.longIdentifier = tok.current()
			myModule.longIdentifierSet = true
			log.Info().Msg("module longIdentifier successfully parsed")
			break forLoop
		}
	}
	log.Info().Msg("creating channels")
	cError := make(chan error, numProc)
	cA2ml := make(chan a2ml, 1)
	cAxisPts := make(chan axisPts, 100)
	cBlob := make(chan blob, 5)
	cCharacteristic := make(chan characteristic, 1000)
	cCompuMethod := make(chan compuMethod, 100)
	cCompuTab := make(chan compuTab, 100)
	cCompuVtab := make(chan compuVTab, 100)
	cCompuVtabRange := make(chan compuVTabRange, 100)
	cFrame := make(chan frame, 1)
	cFunction := make(chan function, 1000)
	cGroup := make(chan group, 100)
	cIfData := make(chan IfData, 100)
	cInstance := make(chan instance, 10)
	cMeasurement := make(chan measurement, 1000)
	cModCommon := make(chan modCommon, 1)
	cModPar := make(chan modPar, 1)
	cRecordLayout := make(chan recordLayout, 100)
	cTransformer := make(chan transformer, 10)
	cTypeDefAxis := make(chan typeDefAxis, 10)
	cTypeDefBlob := make(chan typeDefBlob, 10)
	cTypeDefCharacteristic := make(chan typeDefCharacteristic, 10)
	cTypeDefMeasurement := make(chan typeDefMeasurement, 10)
	cTypeDefStructure := make(chan typeDefStructure, 10)
	cUnit := make(chan unit, 100)
	cUserRights := make(chan userRights, 10)
	cVariantCoding := make(chan variantCoding, 1)

	wgParsers := new(sync.WaitGroup)
	wgParsers.Add(numProc)

	var startIndex int
	var endIndex int
	startIndex = tok.index

	log.Info().Int("startIndex", startIndex).Msg("MODULE begins at index")
	//find /end MODULE token
	for i := len(tokenList) - 1; i >= 0; i-- {
		if tokenList[i] == endModuleToken {
			endIndex = i
		}
	}
	if endIndex <= startIndex {
		err = errors.New("no '/end module' token found")
		return myModule, err
	}
	log.Info().Int("endIndex", endIndex).Msg("MODULE ends at index")
	for i := 0; i < numProc; i++ {
		//Starte parser Threads
		minIndex := startIndex + ((endIndex-startIndex)/numProc)*i
		maxIndex := minIndex + ((endIndex - startIndex) / numProc) - 1
		if i+1 == numProc {
			maxIndex = endIndex
		}
		log.Info().Msg(("goroutine " + fmt.Sprint(i) + " starting index: " + fmt.Sprint(minIndex) + " until end at index: " + fmt.Sprint(maxIndex) + " of " + fmt.Sprint(endIndex)))
		go parseModuleMainLoop(wgParsers, minIndex, maxIndex, cA2ml, cAxisPts, cBlob, cCharacteristic, cCompuMethod,
			cCompuTab, cCompuVtab, cCompuVtabRange, cFrame, cFunction, cGroup, cIfData, cMeasurement, cModCommon,
			cModPar, cRecordLayout, cInstance, cTransformer, cTypeDefAxis, cTypeDefBlob, cTypeDefCharacteristic, cTypeDefMeasurement, cTypeDefStructure, cUnit, cUserRights, cVariantCoding, cError)
	}
	//Start Go Routine that monitors when the parsers are done and then closes the channels.
	//this way the collectorroutines know when they're done.
	go closeChannelsAfterParsing(wgParsers, cA2ml, cAxisPts, cBlob, cCharacteristic, cCompuMethod,
		cCompuTab, cCompuVtab, cCompuVtabRange, cFrame, cFunction, cGroup, cIfData, cMeasurement, cModCommon,
		cModPar, cRecordLayout, cInstance, cTransformer, cTypeDefAxis, cTypeDefBlob, cTypeDefCharacteristic, cTypeDefMeasurement, cTypeDefStructure, cUnit, cUserRights, cVariantCoding, cError)

	//Select collector:
	myModule = collectChannelsSelect(myModule, cA2ml, cAxisPts, cBlob, cCharacteristic, cCompuMethod,
		cCompuTab, cCompuVtab, cCompuVtabRange, cFrame, cFunction, cGroup, cIfData, cMeasurement, cModCommon,
		cModPar, cRecordLayout, cInstance, cTransformer, cTypeDefAxis, cTypeDefBlob, cTypeDefCharacteristic, cTypeDefMeasurement, cTypeDefStructure, cUnit, cUserRights, cVariantCoding, cError)
	//Multithreaded collector:
	/*myModule = collectChannelsMultithreaded(myModule, cA2ml, cAxisPts, cCharacteristic, cCompuMethod,
	cCompuTab, cCompuVtab, cCompuVtabRange, cFrame, cFunction, cGroup, cIfData, cMeasurement, cModCommon,
	cModPar, cRecordLayout, cUnit, cUserRights, cVariantCoding,cError)*/
	tok.index = endIndex
	if len(myModule.errors) > 0 {
		err = myModule.errors[0]
		log.Warn().Int("Number of errors", len(myModule.errors))
	}
	return myModule, err
}

//collectChannelsMultithreaded uses anonymous function to collect the data sent by the goroutines running the moduleMainLoop.
//usually the Select Collector is to be prefered as it is mostly faster and always easier on memory
//as the additional goroutines spun up in collectChannelsMultithreaded seem to block the GC a lot
func collectChannelsMultithreaded(myModule module, cA2ml chan a2ml, cAxisPts chan axisPts, cBlob chan blob, cCharacteristic chan characteristic,
	cCompuMethod chan compuMethod, cCompuTab chan compuTab, cCompuVtab chan compuVTab,
	cCompuVtabRange chan compuVTabRange, cFrame chan frame, cFunction chan function,
	cGroup chan group, cIfData chan IfData, cMeasurement chan measurement,
	cModCommon chan modCommon, cModPar chan modPar, cRecordLayout chan recordLayout,
	cInstance chan instance, cTransformer chan transformer, cTypeDefAxis chan typeDefAxis,
	cTypeDefBlob chan typeDefBlob, cTypeDefCharacteristic chan typeDefCharacteristic,
	cTypeDefMeasurement chan typeDefMeasurement, cTypeDefStructure chan typeDefStructure,
	cUnit chan unit, cUserRights chan userRights, cVariantCoding chan variantCoding, cError chan error) module {

	log.Info().Msg("spinning up collector routines")
	wgCollectors := new(sync.WaitGroup)
	wgCollectors.Add(19)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for elem := range cError {
			myModule.errors = append(myModule.errors, elem)
		}
		log.Info().Msg("collected errors")
	}(wgCollectors)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for elem := range cA2ml {
			myModule.a2ml = elem
		}
		log.Info().Msg("collected a2ml")
	}(wgCollectors)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for elem := range cAxisPts {
			myModule.AxisPts[elem.name] = elem
		}
		log.Info().Msg("collected axisPts")
	}(wgCollectors)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for elem := range cBlob {
			myModule.blobs[elem.name] = elem
		}
		log.Info().Msg("collected blobs")
	}(wgCollectors)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for elem := range cCharacteristic {
			myModule.Characteristics[elem.Name] = elem
		}
		log.Info().Msg("collected characteristics")
	}(wgCollectors)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for elem := range cCompuMethod {
			myModule.CompuMethods[elem.name] = elem
		}
		log.Info().Msg("collected compuMethods")
	}(wgCollectors)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for elem := range cCompuTab {
			myModule.CompuTabs[elem.name] = elem
		}
		log.Info().Msg("collected compuTabs")
	}(wgCollectors)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for elem := range cCompuVtab {
			myModule.CompuVTabs[elem.name] = elem
		}
		log.Info().Msg("collected compuVtabs")
	}(wgCollectors)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for elem := range cCompuVtabRange {
			myModule.CompuVTabRanges[elem.name] = elem
		}
		log.Info().Msg("collected compuVtabRanges")
	}(wgCollectors)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for elem := range cFrame {
			myModule.frame = elem
		}
		log.Info().Msg("collected frame")
	}(wgCollectors)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for elem := range cFunction {
			myModule.Functions[elem.name] = elem
		}
		log.Info().Msg("collected functions")
	}(wgCollectors)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for elem := range cGroup {
			myModule.Groups[elem.groupName] = elem
		}
		log.Info().Msg("collected groups")
	}(wgCollectors)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for elem := range cIfData {
			myModule.ifData[elem.name] = elem
		}
		log.Info().Msg("collected ifDatas")
	}(wgCollectors)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for elem := range cMeasurement {
			myModule.Measurements[elem.name] = elem
		}
		log.Info().Msg("collected measurements")
	}(wgCollectors)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for elem := range cModCommon {
			myModule.ModCommon = elem
		}
		log.Info().Msg("collected modCommons")
	}(wgCollectors)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for elem := range cModPar {
			myModule.ModPar = elem
		}
		log.Info().Msg("collected modPars")
	}(wgCollectors)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for elem := range cRecordLayout {
			myModule.RecordLayouts[elem.name] = elem
		}
		log.Info().Msg("collected recordLayouts")
	}(wgCollectors)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for elem := range cInstance {
			myModule.instances[elem.name] = elem
		}
		log.Info().Msg("collected instances")
	}(wgCollectors)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for elem := range cTransformer {
			myModule.transformers[elem.name] = elem
		}
		log.Info().Msg("collected transformers")
	}(wgCollectors)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for elem := range cTypeDefAxis {
			myModule.typeDefAxis[elem.name] = elem
		}
		log.Info().Msg("collected typeDefAxis")
	}(wgCollectors)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for elem := range cTypeDefBlob {
			myModule.typeDefBlobs[elem.name] = elem
		}
		log.Info().Msg("collected typeDefBlobs")
	}(wgCollectors)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for elem := range cTypeDefCharacteristic {
			myModule.typeDefCharacteristics[elem.name] = elem
		}
		log.Info().Msg("collected typeDefCharacteristics")
	}(wgCollectors)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for elem := range cTypeDefMeasurement {
			myModule.typeDefMeasurements[elem.name] = elem
		}
		log.Info().Msg("collected typeDefMeasurements")
	}(wgCollectors)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for elem := range cTypeDefStructure {
			myModule.typeDefStructures[elem.name] = elem
		}
		log.Info().Msg("collected typeDefStructures")
	}(wgCollectors)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for elem := range cUnit {
			myModule.Units[elem.name] = elem
		}
		log.Info().Msg("collected units")
	}(wgCollectors)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for elem := range cUserRights {
			myModule.userRights[elem.userLevelId] = elem
		}
		log.Info().Msg("collected userRights")
	}(wgCollectors)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for elem := range cVariantCoding {
			myModule.variantCoding = elem
		}
		log.Info().Msg("collected variantCoding")
	}(wgCollectors)
	log.Info().Msg("waiting for collectors to finish")
	wgCollectors.Wait()
	log.Info().Msg("all collectors finished")
	return myModule
}

//collectChannelsSelect uses a select statement to collect the data sent by the goroutines running the moduleMainLoop.
//preferred method compared to MultithreadedCollector
func collectChannelsSelect(myModule module, cA2ml chan a2ml, cAxisPts chan axisPts, cBlob chan blob, cCharacteristic chan characteristic,
	cCompuMethod chan compuMethod, cCompuTab chan compuTab, cCompuVtab chan compuVTab,
	cCompuVtabRange chan compuVTabRange, cFrame chan frame, cFunction chan function,
	cGroup chan group, cIfData chan IfData, cMeasurement chan measurement,
	cModCommon chan modCommon, cModPar chan modPar, cRecordLayout chan recordLayout,
	cInstance chan instance, cTransformer chan transformer, cTypeDefAxis chan typeDefAxis,
	cTypeDefBlob chan typeDefBlob, cTypeDefCharacteristic chan typeDefCharacteristic,
	cTypeDefMeasurement chan typeDefMeasurement, cTypeDefStructure chan typeDefStructure,
	cUnit chan unit, cUserRights chan userRights, cVariantCoding chan variantCoding, cError chan error) module {

	var c0Opn, c1Opn, c3Opn, c4Opn, c5Opn, c6Opn, c7Opn, c8Opn, c9Opn, c10Opn,
		c11Opn, c12Opn, c13Opn, c14Opn, c15Opn, c16Opn, c17Opn, c18Opn, c19Opn, c20Opn,
		c21Opn, c22Opn, c23Opn, c24Opn, c25Opn, c26Opn, c27Opn bool = true, true, true, true, true,
		true, true, true, true, true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true

forLoopSelect:
	select {
	case a, a2 := <-cA2ml:
		myModule.a2ml = a
		c0Opn = a2
	case b, b2 := <-cAxisPts:
		myModule.AxisPts[b.name] = b
		c1Opn = b2
	case c, c2 := <-cCharacteristic:
		myModule.Characteristics[c.Name] = c
		c3Opn = c2
	case d, d2 := <-cCompuMethod:
		myModule.CompuMethods[d.name] = d
		c4Opn = d2
	case e, e2 := <-cCompuTab:
		myModule.CompuTabs[e.name] = e
		c5Opn = e2
	case f, f2 := <-cCompuVtab:
		myModule.CompuVTabs[f.name] = f
		c6Opn = f2
	case g, g2 := <-cCompuVtabRange:
		myModule.CompuVTabRanges[g.name] = g
		c7Opn = g2
	case h, h2 := <-cFrame:
		myModule.frame = h
		c8Opn = h2
	case i, i2 := <-cFunction:
		myModule.Functions[i.name] = i
		c9Opn = i2
	case j, j2 := <-cGroup:
		myModule.Groups[j.groupName] = j
		c10Opn = j2
	case k, k2 := <-cIfData:
		myModule.ifData[k.name] = k
		c11Opn = k2
	case l, l2 := <-cMeasurement:
		myModule.Measurements[l.name] = l
		c12Opn = l2
	case m, m2 := <-cModCommon:
		myModule.ModCommon = m
		c13Opn = m2
	case n, n2 := <-cModPar:
		myModule.ModPar = n
		c14Opn = n2
	case o, o2 := <-cRecordLayout:
		myModule.RecordLayouts[o.name] = o
		c15Opn = o2
	case p, p2 := <-cUnit:
		myModule.Units[p.name] = p
		c16Opn = p2
	case q, q2 := <-cUserRights:
		myModule.userRights[q.userLevelId] = q
		c17Opn = q2
	case r, r2 := <-cVariantCoding:
		myModule.variantCoding = r
		c18Opn = r2
	case s, s2 := <-cBlob:
		myModule.blobs[s.name] = s
		c19Opn = s2
	case t, t2 := <-cInstance:
		myModule.instances[t.name] = t
		c20Opn = t2
	case u, u2 := <-cTransformer:
		myModule.transformers[u.name] = u
		c21Opn = u2
	case v, v2 := <-cTypeDefAxis:
		myModule.typeDefAxis[v.name] = v
		c22Opn = v2
	case w, w2 := <-cTypeDefBlob:
		myModule.typeDefBlobs[w.name] = w
		c23Opn = w2
	case x, x2 := <-cTypeDefCharacteristic:
		myModule.typeDefCharacteristics[x.name] = x
		c24Opn = x2
	case y, y2 := <-cTypeDefMeasurement:
		myModule.typeDefMeasurements[y.name] = y
		c25Opn = y2
	case z, z2 := <-cTypeDefStructure:
		myModule.typeDefStructures[z.name] = z
		c26Opn = z2
	case aa, aa2 := <-cError:
		myModule.errors = append(myModule.errors, aa)
		c27Opn = aa2
	default:
		if !(c0Opn || c1Opn || c3Opn || c4Opn || c5Opn || c6Opn || c7Opn || c8Opn || c9Opn || c10Opn ||
			c11Opn || c12Opn || c13Opn || c14Opn || c15Opn || c16Opn || c17Opn || c18Opn || c19Opn ||
			c20Opn || c21Opn || c22Opn || c23Opn || c24Opn || c25Opn || c26Opn || c27Opn) {
			break forLoopSelect
		}
	}
	log.Info().Msg("collected all channels")
	return myModule
}

//closeChannelsAfterParsing obviously closes all channels when the parser routines have finished
//and wgParser.Wait() is over.
//channels have to be closed in order for the collector to recognize when it is done
//because no more data can be sent and all channels are empty
func closeChannelsAfterParsing(wg *sync.WaitGroup, cA2ml chan a2ml, cAxisPts chan axisPts, cBlob chan blob, cCharacteristic chan characteristic,
	cCompuMethod chan compuMethod, cCompuTab chan compuTab, cCompuVtab chan compuVTab,
	cCompuVtabRange chan compuVTabRange, cFrame chan frame, cFunction chan function,
	cGroup chan group, cIfData chan IfData, cMeasurement chan measurement,
	cModCommon chan modCommon, cModPar chan modPar, cRecordLayout chan recordLayout,
	cInstance chan instance, cTransformer chan transformer, cTypeDefAxis chan typeDefAxis,
	cTypeDefBlob chan typeDefBlob, cTypeDefCharacteristic chan typeDefCharacteristic,
	cTypeDefMeasurement chan typeDefMeasurement, cTypeDefStructure chan typeDefStructure,
	cUnit chan unit, cUserRights chan userRights, cVariantCoding chan variantCoding, cError chan error) {
	log.Info().Msg("waiting for the parsers to finish")
	wg.Wait()
	close(cError)
	close(cA2ml)
	close(cAxisPts)
	close(cBlob)
	close(cCharacteristic)
	close(cCompuMethod)
	close(cCompuTab)
	close(cCompuVtab)
	close(cCompuVtabRange)
	close(cFrame)
	close(cFunction)
	close(cGroup)
	close(cIfData)
	close(cInstance)
	close(cMeasurement)
	close(cModCommon)
	close(cModPar)
	close(cRecordLayout)
	close(cTransformer)
	close(cTypeDefAxis)
	close(cTypeDefBlob)
	close(cTypeDefCharacteristic)
	close(cTypeDefMeasurement)
	close(cTypeDefStructure)
	close(cUnit)
	close(cUserRights)
	close(cVariantCoding)
	log.Info().Msg("parsers finished, closed all channels")
}

//parseModuleMainLoop is used by the parseModuleMultithreaded function to run the module parser in individual goroutines
func parseModuleMainLoop(wg *sync.WaitGroup, minIndex int, maxIndex int,
	cA2ml chan a2ml, cAxisPts chan axisPts, cBlob chan blob, cCharacteristic chan characteristic,
	cCompuMethod chan compuMethod, cCompuTab chan compuTab, cCompuVtab chan compuVTab,
	cCompuVtabRange chan compuVTabRange, cFrame chan frame, cFunction chan function,
	cGroup chan group, cIfData chan IfData, cMeasurement chan measurement,
	cModCommon chan modCommon, cModPar chan modPar, cRecordLayout chan recordLayout,
	cInstance chan instance, cTransformer chan transformer, cTypeDefAxis chan typeDefAxis,
	cTypeDefBlob chan typeDefBlob, cTypeDefCharacteristic chan typeDefCharacteristic,
	cTypeDefMeasurement chan typeDefMeasurement, cTypeDefStructure chan typeDefStructure,
	cUnit chan unit, cUserRights chan userRights, cVariantCoding chan variantCoding, cError chan error) {

	defer wg.Done()

	tg := tokenGenerator{}
	tg.index = minIndex
	var err error
	var bufAxisPts axisPts
	var bufBlob blob
	var bufCharacteristic characteristic
	var bufCompuMethod compuMethod
	var bufCompuTab compuTab
	var bufCompuVtab compuVTab
	var bufCompuVtabRange compuVTabRange
	var bufFunction function
	var bufGroup group
	var bufIfData IfData
	var bufInstance instance
	var bufMeasurement measurement
	var bufRecordLayout recordLayout
	var bufTransformer transformer
	var bufTypeDefAxis typeDefAxis
	var bufTypeDefBlob typeDefBlob
	var bufTypeDefCharacteristic typeDefCharacteristic
	var bufTypeDefMeasurement typeDefMeasurement
	var bufTypeDefStructure typeDefStructure
	var bufUnit unit
	var bufUserRights userRights

forLoop:
	for {
		if tg.index >= maxIndex {
			break forLoop
		}
		switch tg.next() {
		case beginA2mlToken:
			var bufA2ml a2ml
			bufA2ml, err = parseA2ML(&tg)
			if err != nil {
				log.Err(err).Msg("module a2ml could not be parsed")
				cError <- err
				break forLoop
			}
			cA2ml <- bufA2ml
			log.Info().Msg("module a2ml successfully parsed")
		case beginAxisPtsToken:
			bufAxisPts, err = parseAxisPts(&tg)
			if err != nil {
				log.Err(err).Msg("module axisPts could not be parsed")
				cError <- err
				break forLoop
			}
			cAxisPts <- bufAxisPts
			log.Info().Msg("module axisPts successfully parsed")
		case beginBlobToken:
			bufBlob, err = parseBlob(&tg)
			if err != nil {
				log.Err(err).Msg("module blob could not be parsed")
				cError <- err
				break forLoop
			}
			cBlob <- bufBlob
			log.Info().Msg("module blob successfully parsed")
		case beginCharacteristicToken:
			bufCharacteristic, err = parseCharacteristic(&tg)
			if err != nil {
				log.Err(err).Msg("module characteristic could not be parsed")
				cError <- err
				break forLoop
			}
			cCharacteristic <- bufCharacteristic
			log.Info().Msg("module characteristic successfully parsed")
		case beginCompuMethodToken:
			bufCompuMethod, err = parseCompuMethod(&tg)
			if err != nil {
				log.Err(err).Msg("module compuMethod could not be parsed")
				cError <- err
				break forLoop
			}
			cCompuMethod <- bufCompuMethod
			log.Info().Msg("module compuMethod successfully parsed")
		case beginCompuTabToken:
			bufCompuTab, err = parseCompuTab(&tg)
			if err != nil {
				log.Err(err).Msg("module compuTab could not be parsed")
				cError <- err
				break forLoop
			}
			cCompuTab <- bufCompuTab
			log.Info().Msg("module compuTab successfully parsed")
		case beginCompuVtabToken:
			bufCompuVtab, err = parseCompuVtab(&tg)
			if err != nil {
				log.Err(err).Msg("module compuVtab could not be parsed")
				cError <- err
				break forLoop
			}
			cCompuVtab <- bufCompuVtab
			log.Info().Msg("module compuVtab successfully parsed")
		case beginCompuVtabRangeToken:
			bufCompuVtabRange, err = parseCompuVtabRange(&tg)
			if err != nil {
				log.Err(err).Msg("module compuVtabRange could not be parsed")
				cError <- err
				break forLoop
			}
			cCompuVtabRange <- bufCompuVtabRange
			log.Info().Msg("module compuVtabRange successfully parsed")
		case beginFrameToken:
			var bufFrame frame
			bufFrame, err = parseFrame(&tg)
			if err != nil {
				log.Err(err).Msg("module frame could not be parsed")
				cError <- err
				break forLoop
			}
			cFrame <- bufFrame
			log.Info().Msg("module frame successfully parsed")
		case beginFunctionToken:
			bufFunction, err = parseFunction(&tg)
			if err != nil {
				log.Err(err).Msg("module function could not be parsed")
				cError <- err
				break forLoop
			}
			cFunction <- bufFunction
			log.Info().Msg("module function successfully parsed")
		case beginGroupToken:
			bufGroup, err = parseGroup(&tg)
			if err != nil {
				log.Err(err).Msg("module group could not be parsed")
				cError <- err
				break forLoop
			}
			cGroup <- bufGroup
			log.Info().Msg("module group successfully parsed")
		case beginIfDataToken:
			bufIfData, err = parseIfData(&tg)
			if err != nil {
				log.Err(err).Msg("module ifData could not be parsed")
				cError <- err
				break forLoop
			}
			cIfData <- bufIfData
			log.Info().Msg("module ifData successfully parsed")
		case beginInstanceToken:
			bufInstance, err = parseInstance(&tg)
			if err != nil {
				log.Err(err).Msg("module instance could not be parsed")
				cError <- err
				break forLoop
			}
			cInstance <- bufInstance
			log.Info().Msg("module instance successfully parsed")
		case beginMeasurementToken:
			bufMeasurement, err = parseMeasurement(&tg)
			if err != nil {
				log.Err(err).Msg("module measurement could not be parsed")
				cError <- err
				break forLoop
			}
			cMeasurement <- bufMeasurement
			log.Info().Msg("module measurement successfully parsed")
		case beginModCommonToken:
			var bufModCommon modCommon
			bufModCommon, err = parseModCommon(&tg)
			if err != nil {
				log.Err(err).Msg("module modCommon could not be parsed")
				cError <- err
				break forLoop
			}
			cModCommon <- bufModCommon
			log.Info().Msg("module modCommon successfully parsed")
		case beginModParToken:
			var bufModPar modPar
			bufModPar, err = parseModPar(&tg)
			if err != nil {
				log.Err(err).Msg("module modPar could not be parsed")
				cError <- err
				break forLoop
			}
			cModPar <- bufModPar
			log.Info().Msg("module modPar successfully parsed")
		case beginRecordLayoutToken:
			bufRecordLayout, err = parseRecordLayout(&tg)
			if err != nil {
				log.Err(err).Msg("module recordLayout could not be parsed")
				cError <- err
				break forLoop
			}
			cRecordLayout <- bufRecordLayout
			log.Info().Msg("module recordLayout successfully parsed")
		case beginTransformerToken:
			bufTransformer, err = parseTransformer(&tg)
			if err != nil {
				log.Err(err).Msg("module transformer could not be parsed")
				cError <- err
				break forLoop
			}
			cTransformer <- bufTransformer
			log.Info().Msg("module transformer successfully parsed")
		case beginTypeDefAxisToken:
			bufTypeDefAxis, err = parseTypeDefAxis(&tg)
			if err != nil {
				log.Err(err).Msg("module typeDefAxis could not be parsed")
				cError <- err
				break forLoop
			}
			cTypeDefAxis <- bufTypeDefAxis
			log.Info().Msg("module typeDefAxis successfully parsed")
		case beginTypeDefBlobToken:
			bufTypeDefBlob, err = parseTypeDefBlob(&tg)
			if err != nil {
				log.Err(err).Msg("module typeDefBlob could not be parsed")
				cError <- err
				break forLoop
			}
			cTypeDefBlob <- bufTypeDefBlob
			log.Info().Msg("module typeDefBlob successfully parsed")
		case beginTypeDefCharacteristicToken:
			bufTypeDefCharacteristic, err = parseTypeDefCharacteristic(&tg)
			if err != nil {
				log.Err(err).Msg("module typeDefCharacteristic could not be parsed")
				cError <- err
				break forLoop
			}
			cTypeDefCharacteristic <- bufTypeDefCharacteristic
			log.Info().Msg("module typeDefCharacteristic successfully parsed")
		case beginTypeDefMeasurementToken:
			bufTypeDefMeasurement, err = parseTypeDefMeasurement(&tg)
			if err != nil {
				log.Err(err).Msg("module typeDefMeasurement could not be parsed")
				cError <- err
				break forLoop
			}
			cTypeDefMeasurement <- bufTypeDefMeasurement
			log.Info().Msg("module typeDefMeasurement successfully parsed")
		case beginTypeDefStructureToken:
			bufTypeDefStructure, err = parseTypeDefStructure(&tg)
			if err != nil {
				log.Err(err).Msg("module typeDefStructure could not be parsed")
				cError <- err
				break forLoop
			}
			cTypeDefStructure <- bufTypeDefStructure
			log.Info().Msg("module typeDefStructure successfully parsed")
		case beginUnitToken:
			bufUnit, err = parseUnit(&tg)
			if err != nil {
				log.Err(err).Msg("module unit could not be parsed")
				cError <- err
				break forLoop
			}
			cUnit <- bufUnit
			log.Info().Msg("module unit successfully parsed")
		case beginUserRightsToken:
			bufUserRights, err = parseUserRights(&tg)
			if err != nil {
				log.Err(err).Msg("module userRights could not be parsed")
				cError <- err
				break forLoop
			}
			cUserRights <- bufUserRights
			log.Info().Msg("module userRights successfully parsed")
		case beginVariantCodingToken:
			var bufVariantCoding variantCoding
			bufVariantCoding, err = parseVariantCoding(&tg)
			if err != nil {
				log.Err(err).Msg("module variantCoding could not be parsed")
				cError <- err
				break forLoop
			}
			cVariantCoding <- bufVariantCoding
			log.Info().Msg("module variantCoding successfully parsed")
		default:
			if tg.current() == emptyToken {
				fmt.Println("empty_token")
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("module could not be parsed")
				cError <- err
				break forLoop
			} else if tg.current() == endModuleToken {
				break forLoop
			} else if tg.index >= maxIndex {
				break forLoop
			}
		}
	}
}
