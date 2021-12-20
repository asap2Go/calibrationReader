package main

import (
	"strings"
	"sync"

	"github.com/rs/zerolog/log"
)

type objectIdentifiers struct {
	isInitialized       bool
	axisPtsKeys         []string
	characteristicsKeys []string
	compuMethodsKeys    []string
	compuTabsKeys       []string
	compuVTabsKeys      []string
	compuVTabRangesKeys []string
	functionsKeys       []string
	groupsKeys          []string
	measurementsKeys    []string
	recordLayoutsKeys   []string
	unitsKeys           []string
}

func buildObjectKeys(cd *CalibrationData) objectIdentifiers {
	keys := objectIdentifiers{}
	var key string
	wgKeyCollectors := new(sync.WaitGroup)
	wgKeyCollectors.Add(11)
	for _, m := range cd.a2l.Project.Modules {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for key = range m.AxisPts {
				keys.axisPtsKeys = append(keys.axisPtsKeys, key)
				keys.axisPtsKeys = append(keys.axisPtsKeys, strings.ToLower(key))
			}
			log.Info().Msg("collected axisPtsKeys")
		}(wgKeyCollectors)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for key = range m.Characteristics {
				keys.characteristicsKeys = append(keys.characteristicsKeys, key)
				keys.characteristicsKeys = append(keys.characteristicsKeys, strings.ToLower(key))
			}
			log.Info().Msg("collected characteristicsKeys")
		}(wgKeyCollectors)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for key = range m.CompuMethods {
				keys.compuMethodsKeys = append(keys.compuMethodsKeys, key)
				keys.compuMethodsKeys = append(keys.compuMethodsKeys, strings.ToLower(key))
			}
			log.Info().Msg("collected compuMethodsKeys")
		}(wgKeyCollectors)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for key = range m.CompuTabs {
				keys.compuTabsKeys = append(keys.compuTabsKeys, key)
				keys.compuTabsKeys = append(keys.compuTabsKeys, strings.ToLower(key))
			}
			log.Info().Msg("collected compuTabsKeys")
		}(wgKeyCollectors)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for key = range m.CompuVTabs {
				keys.compuVTabsKeys = append(keys.compuVTabsKeys, key)
				keys.compuVTabsKeys = append(keys.compuVTabsKeys, strings.ToLower(key))
			}
			log.Info().Msg("collected compuVTabsKeys")
		}(wgKeyCollectors)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for key = range m.CompuVTabRanges {
				keys.compuVTabRangesKeys = append(keys.compuVTabRangesKeys, key)
				keys.compuVTabRangesKeys = append(keys.compuVTabRangesKeys, strings.ToLower(key))
			}
			log.Info().Msg("collected compuVTabRangesKeys")
		}(wgKeyCollectors)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for key = range m.Functions {
				keys.functionsKeys = append(keys.functionsKeys, key)
				keys.functionsKeys = append(keys.functionsKeys, strings.ToLower(key))
			}
			log.Info().Msg("collected functionsKeys")
		}(wgKeyCollectors)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for key = range m.Groups {
				keys.groupsKeys = append(keys.groupsKeys, key)
				keys.groupsKeys = append(keys.groupsKeys, strings.ToLower(key))
			}
			log.Info().Msg("collected groupsKeys")
		}(wgKeyCollectors)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for key = range m.Measurements {
				keys.measurementsKeys = append(keys.measurementsKeys, key)
				keys.measurementsKeys = append(keys.measurementsKeys, strings.ToLower(key))
			}
			log.Info().Msg("collected measurementsKeys")
		}(wgKeyCollectors)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for key = range m.RecordLayouts {
				keys.recordLayoutsKeys = append(keys.recordLayoutsKeys, key)
				keys.recordLayoutsKeys = append(keys.recordLayoutsKeys, strings.ToLower(key))
			}
			log.Info().Msg("collected recordLayoutsKeys")
		}(wgKeyCollectors)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for key = range m.Units {
				keys.unitsKeys = append(keys.unitsKeys, key)
				keys.unitsKeys = append(keys.unitsKeys, strings.ToLower(key))
			}
			log.Info().Msg("collected unitsKeys")
		}(wgKeyCollectors)
		log.Info().Str("module name", m.Name).Msg("waiting for collection of keys from module")
		//Wait until all keys have been collected from one module to avoid concurrent slice accesses.
		wgKeyCollectors.Wait()
		log.Info().Str("module name", m.Name).Msg("collected all keys from module")
	}
	keys.isInitialized = true
	return keys
}
