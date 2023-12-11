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
		regexp.MustCompile(`\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}\b`),
		regexp.MustCompile(`\b(CR|DE|ME|RS)[0-9]{20}\b`),
		regexp.MustCompile(`\b(CZ|ES|SE|SK|TN)[0-9]{22}\b`),
		regexp.MustCompile(`\bSE45 5000 0000 0583 9825 7466\b`),
	}
)
