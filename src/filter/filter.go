package filter

// filter.go: A regex based filer engine for URI's
// filters.json contains an object array of Singatures
// Each signature contains a name and a regex string
// The array of signatures is read in, and then compiled
// into a list of regexes that will be applied

import (
	"encoding/json"
	"log"
	"os"
	"regexp"
)

// FilterObject contains two main arrays
// Signatures are matches for attacks
// When a signature matches it will
// prevent the request from continuing
// Shallows are matches for quick proccessing
// Shallows matches allow the packet to be passed on
// since it should be impossible/unlikely that there
// is any harm there
type FilterObject struct {
	Signatures []Signature
	Shallows   []Signature
}

// Signature struct is pretty simple, name and regex.
// In the future might want to have a description,
// and a way to enable/disable particular regexes
// (perhaps based on ports/HTTP Methods/etc)
// however for now, this should do
type Signature struct {
	Name  string
	Regex string
}

var filterObj = FilterObject{}
var signatureRegex []*regexp.Regexp
var shallowRegex []*regexp.Regexp

// LoadFilters: reads the filters.json file and then
// compiles the regexes.
// I would typically load them in from a DB (redis/mongo),
// however for this project it seemed simplicity and time
// constraints lead me to a standard json file.
// TODO: Enable a file update notification system
// that will reload the filters when the filters.json
// file changes.
func LoadFilters(projectRoot string) {
	var file *os.File
	if projectRoot != "" {
		file, _ = os.Open(projectRoot + "filters.json")
	} else {
		file, _ = os.Open("filters.json")
	}
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&filterObj)
	if err != nil {
		log.Println("Error loading filters")
	}
	compileRegex()
}

// compileRegex: run through the signatures and shallows list
// and grab out the regexes, compile them and save them
// for later
func compileRegex() {
	for _, value := range filterObj.Signatures {
		r, _ := regexp.Compile(value.Regex)
		signatureRegex = append(signatureRegex, r)
	}
	for _, value := range filterObj.Shallows {
		r, _ := regexp.Compile(value.Regex)
		shallowRegex = append(shallowRegex, r)
	}
}

// RunFilter: attempt to Match the regexs on the
// URI passed in. Given the scope of the assignment
// this was the simplist way to search for a SQL
// inject attack, however if we're assuming a lot of
// traffic, a two step search would be in order.
// I've included a very simple two step sort of search
// without too much preoptimization. Depending on traffic
// type and ammount, this would be changed.
func RunFilter(uri string) bool {
	// First check the easy ones that should be very fast
	for _, val := range shallowRegex {
		if val.MatchString(uri) {
			// Matched a shallow regex, let it go
			return true
		}
	}

	// Now check the actual signatures which could be much
	// more complex
	for _, val := range signatureRegex {
		if val.MatchString(uri) {
			// Matched an attack signature
			return false
		}
	}

	// Allow by default if nothing bad matches
	return true
}
