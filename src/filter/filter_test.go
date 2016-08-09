package filter

import (
	"testing"
)

var possibleVulns = []string{
	"name'",
	"name' OR 'a'='a",
	"name' OR 'b'='b",
	"name'); DELETE FROM items; --",
	"name'); DELETE FROM users; --",
	"name'); INSERT INTO items; --",
}

// Simple start to unit testing, I dislike using unit testing here till
// a more solid spec has been laid out for this, however this is the
//most important function, a little testing would be good
func TestFilters(t *testing.T) {
	LoadFilters("../")
	for _, val := range possibleVulns {
		if RunFilter(val) {
			t.Error(
				"For", val,
				"expected", false,
				"got", true,
			)
		}
	}
}
