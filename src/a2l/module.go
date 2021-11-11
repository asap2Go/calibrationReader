package a2l

import (
	"errors"
	"fmt"
	"sync"

	"github.com/rs/zerolog/log"
)

type module struct {
	name              string
	nameSet           bool
	longIdentifier    string
	longIdentifierSet bool
	a2ml              a2ml
	AxisPts           map[string]axisPts
	Characteristics   map[string]characteristic
	CompuMethods      map[string]compuMethod
	CompuTabs         map[string]compuTab
	CompuVTabs        map[string]compuVTab
	CompuVTabRanges   map[string]compuVTabRange
	frame             FRAME
	Functions         map[string]function
	Groups            map[string]group
	ifData            map[string]IfData
	Measurements      map[string]MEASUREMENT
	ModCommon         modCommon
	ModPar            modPar
	RecordLayouts     map[string]recordLayout
	Units             map[string]unit
	userRights        map[string]userRights
	variantCoding     variantCoding
}

func parseModule(tok *tokenGenerator) (module, error) {
	//Bulk init of an average number of objects contained in a modern a2l-file.
	myModule := module{}
	myModule.AxisPts = make(map[string]axisPts, 1000)
	myModule.Characteristics = make(map[string]characteristic, 10000)
	myModule.CompuMethods = make(map[string]compuMethod, 1000)
	myModule.CompuTabs = make(map[string]compuTab, 1000)
	myModule.CompuVTabs = make(map[string]compuVTab, 1000)
	myModule.CompuVTabRanges = make(map[string]compuVTabRange, 1000)
	myModule.Functions = make(map[string]function, 10000)
	myModule.Groups = make(map[string]group, 1000)
	myModule.ifData = make(map[string]IfData, 1000)
	myModule.Measurements = make(map[string]MEASUREMENT, 10000)
	myModule.RecordLayouts = make(map[string]recordLayout, 1000)
	myModule.Units = make(map[string]unit, 1000)
	myModule.userRights = make(map[string]userRights, 1000)
	var err error
	var bufAxisPts axisPts
	var bufCharacteristic characteristic
	var bufCompuMethod compuMethod
	var bufCompuTab compuTab
	var bufCompuVtab compuVTab
	var bufCompuVtabRange compuVTabRange
	var bufFunction function
	var bufGroup group
	var bufIfData IfData
	var bufMeasurement MEASUREMENT
	var bufRecordLayout recordLayout
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
		case beginMeasurementToken:
			bufMeasurement, err = parseMeasurement(tok)
			if err != nil {
				log.Err(err).Msg("module measurement could not be parsed")
				break forLoop
			}
			myModule.Measurements[bufMeasurement.name] = bufMeasurement
			log.Info().Msg("module measurement[bufMeasurement name] successfully parsed")
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
	myModule.Characteristics = make(map[string]characteristic, 10000)
	myModule.CompuMethods = make(map[string]compuMethod, 1000)
	myModule.CompuTabs = make(map[string]compuTab, 1000)
	myModule.CompuVTabs = make(map[string]compuVTab, 1000)
	myModule.CompuVTabRanges = make(map[string]compuVTabRange, 1000)
	myModule.Functions = make(map[string]function, 10000)
	myModule.Groups = make(map[string]group, 1000)
	myModule.ifData = make(map[string]IfData, 1000)
	myModule.Measurements = make(map[string]MEASUREMENT, 10000)
	myModule.RecordLayouts = make(map[string]recordLayout, 1000)
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
	cA2ml := make(chan a2ml, 1)
	cAxisPts := make(chan axisPts, 100)
	cCharacteristic := make(chan characteristic, 1000)
	cCompuMethod := make(chan compuMethod, 100)
	cCompuTab := make(chan compuTab, 100)
	cCompuVtab := make(chan compuVTab, 100)
	cCompuVtabRange := make(chan compuVTabRange, 100)
	cFrame := make(chan FRAME, 1)
	cFunction := make(chan function, 1000)
	cGroup := make(chan group, 100)
	cIfData := make(chan IfData, 100)
	cMeasurement := make(chan MEASUREMENT, 1000)
	cModCommon := make(chan modCommon, 1)
	cModPar := make(chan modPar, 1)
	cRecordLayout := make(chan recordLayout, 100)
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
		go parseModuleMainLoop(wgParsers, minIndex, maxIndex, cA2ml, cAxisPts, cCharacteristic, cCompuMethod,
			cCompuTab, cCompuVtab, cCompuVtabRange, cFrame, cFunction, cGroup, cIfData, cMeasurement, cModCommon,
			cModPar, cRecordLayout, cUnit, cUserRights, cVariantCoding)
	}
	//Start Go Routine that monitors when the parsers are done and then closes the channels.
	//this way the collectorroutines know when they're done.
	go closeChannelsAfterParsing(wgParsers, cA2ml, cAxisPts, cCharacteristic, cCompuMethod,
		cCompuTab, cCompuVtab, cCompuVtabRange, cFrame, cFunction, cGroup, cIfData, cMeasurement, cModCommon,
		cModPar, cRecordLayout, cUnit, cUserRights, cVariantCoding)

	//Select collector:
	myModule = collectChannelsSelect(myModule, cA2ml, cAxisPts, cCharacteristic, cCompuMethod,
		cCompuTab, cCompuVtab, cCompuVtabRange, cFrame, cFunction, cGroup, cIfData, cMeasurement, cModCommon,
		cModPar, cRecordLayout, cUnit, cUserRights, cVariantCoding)
	//Multithreaded collector:
	/*myModule = collectChannelsMultithreaded(myModule, cA2ml, cAxisPts, cCharacteristic, cCompuMethod,
	cCompuTab, cCompuVtab, cCompuVtabRange, cFrame, cFunction, cGroup, cIfData, cMeasurement, cModCommon,
	cModPar, cRecordLayout, cUnit, cUserRights, cVariantCoding)*/
	tok.index = endIndex
	return myModule, err

}

//collectChannelsMultithreaded uses anonymous function to collect the data sent by the goroutines running the moduleMainLoop.
//usually the Select Collector is to be prefered as it is mostly faster and always easier on memory
//as the additional goroutines spun up in collectChannelsMultithreaded seem to block the GC a lot
func collectChannelsMultithreaded(myModule module, cA2ml chan a2ml, cAxisPts chan axisPts, cCharacteristic chan characteristic,
	cCompuMethod chan compuMethod, cCompuTab chan compuTab, cCompuVtab chan compuVTab,
	cCompuVtabRange chan compuVTabRange, cFrame chan FRAME, cFunction chan function,
	cGroup chan group, cIfData chan IfData, cMeasurement chan MEASUREMENT,
	cModCommon chan modCommon, cModPar chan modPar, cRecordLayout chan recordLayout,
	cUnit chan unit, cUserRights chan userRights, cVariantCoding chan variantCoding) module {

	log.Info().Msg("spinning up collector routines")
	wgCollectors := new(sync.WaitGroup)
	wgCollectors.Add(18)
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
func collectChannelsSelect(myModule module, cA2ml chan a2ml, cAxisPts chan axisPts, cCharacteristic chan characteristic,
	cCompuMethod chan compuMethod, cCompuTab chan compuTab, cCompuVtab chan compuVTab,
	cCompuVtabRange chan compuVTabRange, cFrame chan FRAME, cFunction chan function,
	cGroup chan group, cIfData chan IfData, cMeasurement chan MEASUREMENT,
	cModCommon chan modCommon, cModPar chan modPar, cRecordLayout chan recordLayout,
	cUnit chan unit, cUserRights chan userRights, cVariantCoding chan variantCoding) module {

	var aOpn, bOpn, cOpn, dOpn, eOpn, fOpn, gOpn, hOpn, iOpn, jOpn, kOpn, lOpn, mOpn, nOpn, oOpn, pOpn, qOpn, rOpn bool

forLoopSelect:
	select {
	case a, a2 := <-cA2ml:
		myModule.a2ml = a
		aOpn = a2
	case b, b2 := <-cAxisPts:
		myModule.AxisPts[b.name] = b
		bOpn = b2
	case c, c2 := <-cCharacteristic:
		myModule.Characteristics[c.Name] = c
		cOpn = c2
	case d, d2 := <-cCompuMethod:
		myModule.CompuMethods[d.name] = d
		dOpn = d2
	case e, e2 := <-cCompuTab:
		myModule.CompuTabs[e.name] = e
		eOpn = e2
	case f, f2 := <-cCompuVtab:
		myModule.CompuVTabs[f.name] = f
		fOpn = f2
	case g, g2 := <-cCompuVtabRange:
		myModule.CompuVTabRanges[g.name] = g
		gOpn = g2
	case h, h2 := <-cFrame:
		myModule.frame = h
		hOpn = h2
	case i, i2 := <-cFunction:
		myModule.Functions[i.name] = i
		iOpn = i2
	case j, j2 := <-cGroup:
		myModule.Groups[j.groupName] = j
		jOpn = j2
	case k, k2 := <-cIfData:
		myModule.ifData[k.name] = k
		kOpn = k2
	case l, l2 := <-cMeasurement:
		myModule.Measurements[l.name] = l
		lOpn = l2
	case m, m2 := <-cModCommon:
		myModule.ModCommon = m
		mOpn = m2
	case n, n2 := <-cModPar:
		myModule.ModPar = n
		nOpn = n2
	case o, o2 := <-cRecordLayout:
		myModule.RecordLayouts[o.name] = o
		oOpn = o2
	case p, p2 := <-cUnit:
		myModule.Units[p.name] = p
		pOpn = p2
	case q, q2 := <-cUserRights:
		myModule.userRights[q.userLevelId] = q
		qOpn = q2
	case r, r2 := <-cVariantCoding:
		myModule.variantCoding = r
		rOpn = r2
	default:
		if !(aOpn || bOpn || cOpn || dOpn || eOpn || fOpn || gOpn || hOpn || iOpn || jOpn || kOpn || lOpn || mOpn || nOpn || oOpn || pOpn || qOpn || rOpn) {
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
func closeChannelsAfterParsing(wg *sync.WaitGroup, cA2ml chan a2ml, cAxisPts chan axisPts, cCharacteristic chan characteristic,
	cCompuMethod chan compuMethod, cCompuTab chan compuTab, cCompuVtab chan compuVTab,
	cCompuVtabRange chan compuVTabRange, cFrame chan FRAME, cFunction chan function,
	cGroup chan group, cIfData chan IfData, cMeasurement chan MEASUREMENT,
	cModCommon chan modCommon, cModPar chan modPar, cRecordLayout chan recordLayout,
	cUnit chan unit, cUserRights chan userRights, cVariantCoding chan variantCoding) {
	log.Info().Msg("waiting for the parsers to finish")
	wg.Wait()
	close(cA2ml)
	close(cAxisPts)
	close(cCharacteristic)
	close(cCompuMethod)
	close(cCompuTab)
	close(cCompuVtab)
	close(cCompuVtabRange)
	close(cFrame)
	close(cFunction)
	close(cGroup)
	close(cIfData)
	close(cMeasurement)
	close(cModCommon)
	close(cModPar)
	close(cRecordLayout)
	close(cUnit)
	close(cUserRights)
	close(cVariantCoding)
	log.Info().Msg("parsers finished, closed all channels")
}

//parseModuleMainLoop is used by the parseModuleMultithreaded function to run the module parser in individual goroutines
func parseModuleMainLoop(wg *sync.WaitGroup, minIndex int, maxIndex int,
	cA2ml chan a2ml, cAxisPts chan axisPts, cCharacteristic chan characteristic,
	cCompuMethod chan compuMethod, cCompuTab chan compuTab, cCompuVtab chan compuVTab,
	cCompuVtabRange chan compuVTabRange, cFrame chan FRAME, cFunction chan function,
	cGroup chan group, cIfData chan IfData, cMeasurement chan MEASUREMENT,
	cModCommon chan modCommon, cModPar chan modPar, cRecordLayout chan recordLayout,
	cUnit chan unit, cUserRights chan userRights, cVariantCoding chan variantCoding) {

	defer wg.Done()

	tg := tokenGenerator{}
	tg.index = minIndex
	var err error
	var bufA2ml a2ml
	var bufAxisPts axisPts
	var bufCharacteristic characteristic
	var bufCompuMethod compuMethod
	var bufCompuTab compuTab
	var bufCompuVtab compuVTab
	var bufCompuVtabRange compuVTabRange
	var bufFrame FRAME
	var bufFunction function
	var bufGroup group
	var bufIfData IfData
	var bufMeasurement MEASUREMENT
	var bufModCommon modCommon
	var bufModPar modPar
	var bufRecordLayout recordLayout
	var bufUnit unit
	var bufUserRights userRights
	var bufVariantCoding variantCoding

forLoop:
	for {
		if tg.index >= maxIndex {
			break forLoop
		}
		switch tg.next() {
		case beginA2mlToken:
			bufA2ml, err = parseA2ML(&tg)
			if err != nil {
				log.Err(err).Msg("module a2ml could not be parsed")
				break forLoop
			}
			cA2ml <- bufA2ml
			log.Info().Msg("module a2ml successfully parsed")
		case beginAxisPtsToken:
			bufAxisPts, err = parseAxisPts(&tg)
			if err != nil {
				log.Err(err).Msg("module axisPts could not be parsed")
				break forLoop
			}
			cAxisPts <- bufAxisPts
			log.Info().Msg("module axisPts[bufAxisPts name] successfully parsed")
		case beginCharacteristicToken:
			bufCharacteristic, err = parseCharacteristic(&tg)
			if err != nil {
				log.Err(err).Msg("module characteristic could not be parsed")
				break forLoop
			}
			cCharacteristic <- bufCharacteristic
			log.Info().Msg("module characteristic[bufCharacteristic name] successfully parsed")
		case beginCompuMethodToken:
			bufCompuMethod, err = parseCompuMethod(&tg)
			if err != nil {
				log.Err(err).Msg("module compuMethod could not be parsed")
				break forLoop
			}
			cCompuMethod <- bufCompuMethod
			log.Info().Msg("module compuMethod[bufCompuMethod name] successfully parsed")
		case beginCompuTabToken:
			bufCompuTab, err = parseCompuTab(&tg)
			if err != nil {
				log.Err(err).Msg("module compuTab could not be parsed")
				break forLoop
			}
			cCompuTab <- bufCompuTab
			log.Info().Msg("module compuTab[bufCompuTab name] successfully parsed")
		case beginCompuVtabToken:
			bufCompuVtab, err = parseCompuVtab(&tg)
			if err != nil {
				log.Err(err).Msg("module compuVtab could not be parsed")
				break forLoop
			}
			cCompuVtab <- bufCompuVtab
			log.Info().Msg("module compuVtab[bufCompuVtab name] successfully parsed")
		case beginCompuVtabRangeToken:
			bufCompuVtabRange, err = parseCompuVtabRange(&tg)
			if err != nil {
				log.Err(err).Msg("module compuVtabRange could not be parsed")
				break forLoop
			}
			cCompuVtabRange <- bufCompuVtabRange
			log.Info().Msg("module compuVtabRange[bufCompuVtabRange name] successfully parsed")
		case beginFrameToken:
			bufFrame, err = parseFrame(&tg)
			if err != nil {
				log.Err(err).Msg("module frame could not be parsed")
				break forLoop
			}
			cFrame <- bufFrame
			log.Info().Msg("module frame successfully parsed")
		case beginFunctionToken:
			bufFunction, err = parseFunction(&tg)
			if err != nil {
				log.Err(err).Msg("module function could not be parsed")
				break forLoop
			}
			cFunction <- bufFunction
			log.Info().Msg("module function[bufFunction name] successfully parsed")
		case beginGroupToken:
			bufGroup, err = parseGroup(&tg)
			if err != nil {
				log.Err(err).Msg("module group could not be parsed")
				break forLoop
			}
			cGroup <- bufGroup
			log.Info().Msg("module group[bufGroup groupName] successfully parsed")
		case beginIfDataToken:
			bufIfData, err = parseIfData(&tg)
			if err != nil {
				log.Err(err).Msg("module ifData could not be parsed")
				break forLoop
			}
			cIfData <- bufIfData
			log.Info().Msg("module ifData[bufIfData name] successfully parsed")
		case beginMeasurementToken:
			bufMeasurement, err = parseMeasurement(&tg)
			if err != nil {
				log.Err(err).Msg("module measurement could not be parsed")
				break forLoop
			}
			cMeasurement <- bufMeasurement
			log.Info().Msg("module measurement[bufMeasurement name] successfully parsed")
		case beginModCommonToken:
			bufModCommon, err = parseModCommon(&tg)
			if err != nil {
				log.Err(err).Msg("module modCommon could not be parsed")
				break forLoop
			}
			cModCommon <- bufModCommon
			log.Info().Msg("module modCommon successfully parsed")
		case beginModParToken:
			bufModPar, err = parseModPar(&tg)
			if err != nil {
				log.Err(err).Msg("module modPar could not be parsed")
				break forLoop
			}
			cModPar <- bufModPar
			log.Info().Msg("module modPar successfully parsed")
		case beginRecordLayoutToken:
			bufRecordLayout, err = parseRecordLayout(&tg)
			if err != nil {
				log.Err(err).Msg("module recordLayout could not be parsed")
				break forLoop
			}
			cRecordLayout <- bufRecordLayout
			log.Info().Msg("module recordLayout[bufRecordLayout name] successfully parsed")
		case beginUnitToken:
			bufUnit, err = parseUnit(&tg)
			if err != nil {
				log.Err(err).Msg("module unit could not be parsed")
				break forLoop
			}
			cUnit <- bufUnit
			log.Info().Msg("module unit[bufUnit name] successfully parsed")
		case beginUserRightsToken:
			bufUserRights, err = parseUserRights(&tg)
			if err != nil {
				log.Err(err).Msg("module userRights could not be parsed")
				break forLoop
			}
			cUserRights <- bufUserRights
			log.Info().Msg("module userRights[bufUserRights userLevelId] successfully parsed")
		case beginVariantCodingToken:
			bufVariantCoding, err = parseVariantCoding(&tg)
			if err != nil {
				log.Err(err).Msg("module variantCoding could not be parsed")
				break forLoop
			}
			cVariantCoding <- bufVariantCoding
			log.Info().Msg("module variantCoding successfully parsed")
		default:
			if tg.current() == emptyToken {
				fmt.Println("empty_token")
				err := errors.New("unexpected end of file")
				log.Err(err).Msg("module could not be parsed")
				break forLoop
			} else if tg.current() == endModuleToken {
				break forLoop
			} else if tg.index >= maxIndex {
				break forLoop
			}
		}
	}
}
