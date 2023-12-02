package safelog

import (
	"regexp"
)

// https://docs.trellix.com/bundle/data-loss-prevention-11.10.x-classification-definitions-reference-guide/page/GUID-63579440-55DE-4D7B-BA96-3439CC27BDDF.html
var (
	keyRegexp = []*regexp.Regexp{
		regexp.MustCompile("IBAN"),
	}
	valueRegexp = []*regexp.Regexp{
		regexp.MustCompile(".*(CR|DE|ME|RS)[0-9]{20}.*"),
		regexp.MustCompile(".*(CZ|ES|SE|SK|TN)[0-9]{22}.*"),
	}
)
