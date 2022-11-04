# calibrationReader
[![Go Report Card](https://goreportcard.com/badge/github.com/asap2Go/calibrationReader)](https://goreportcard.com/report/github.com/asap2Go/calibrationReader)  [![Codacy Badge](https://app.codacy.com/project/badge/Grade/e19560faf3484ccb88922ad3548b19ad)](https://www.codacy.com/gh/asap2Go/calibrationReader/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=asap2Go/calibrationReader&amp;utm_campaign=Badge_Grade) [![Codacy Badge](https://app.codacy.com/project/badge/Coverage/e19560faf3484ccb88922ad3548b19ad)](https://www.codacy.com/gh/asap2Go/calibrationReader/dashboard?utm_source=github.com&utm_medium=referral&utm_content=asap2Go/calibrationReader&utm_campaign=Badge_Coverage) [![Go](https://github.com/asap2Go/calibrationReader/actions/workflows/go-build-test-and-license.yml/badge.svg)](https://github.com/asap2Go/calibrationReader/actions/workflows/go-build-test-and-license.yml) [![CodeQL](https://github.com/asap2Go/calibrationReader/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/asap2Go/calibrationReader/actions/workflows/codeql-analysis.yml) [![Go Reference](https://pkg.go.dev/badge/github.com/asap2Go/calibrationReader.svg)](https://pkg.go.dev/github.com/asap2Go/calibrationReader)
## Scope
 The calibrationReader package reads characteristics, system constants, measurement definitions etc. 
 from an a2l file and correlates it with the data from a corresponding hex oder s19 file. 
 It is build to cover the ASAM MCD 2MC specification in its current version 1.7.1
 As of now the package only uses the metadata part of the ASAM MCD 2MC standard. 
 The interface descriptions which are used to communicate with an ECU 
 and are defined in the AML datastructure are not within the scope of this package.

## Current capablilites
 Right now the package can parse a2l-files as well as the corresponding IntelHex32 or Motorola S19 files. 
 And it is quite fast at that. Currently a real world A2L(80MB) with its corresponding Hex File(10MB) will be parsed in less than a second.
### and what is still left to do
 The package still lacks the last bit of work which is implementing the methods for axis_pts, axis_descr, record_layout and fnc_values 
 in order to understand the memory layout and position of a specific characteristic.
 This currently worked on as everything else seems to be quite stable now.
 Once completely implemented version 1.0 will be released.
 
##  Usage
 This is only a preliminary explanation on how to access the datastructures 
 as the API will be formalized when version 1.0 releases with the full scope of the package implemented.

 `calibrationData, err := ReadCalibration(a2lFilePath, hexFilePath)`

 parses an a2l and hex file into datastructures.
 All relevant information e.g. Record Layouts, Measurements, Characterstics, etc. 
 are part of a module which is in turn part of the Project whithin the a2l data structure.
 An a2l file can contain several modules but in most real world applications only contains one. 
 So indexOfModuleInProject (see below) can basically assumed to be 0.
 In the future this will be solved more elegantly with a cd.SetModule method. 
 So the user can define the module he is going to work on.

 `mod := calibrationData.A2l.Project.Modules[indexOfModuleInProject]`

 Most datastructures that are directly below a module 
 within the a2l data structure hierarchy are accessible through maps by their identifiers.

 To access System Constants:

 `sc, exists := mod.SystemConstants["NameOfSystemConstant"]`

 To access Characteristics:

 `c, exists := mod.Characteristics["NameOfCharacteristic"]`

 and so on. 

 To access a specific memory location in the hex-file (contains a single byte):

 `b, exists := calibrationData.hex[12345]`
 
##  Disclaimer
 I am a mechanical engineer, so any code you see might not be up to the standards of true/correct/modern/acceptable software development ;)
 Feedback is always appreciated.

## Dependencies and Licensing
 The only dependency outside the go standard library is currently zerolog.
 The package is - and will always be - released under MIT license.
 Feel free to do with it what you want :)
