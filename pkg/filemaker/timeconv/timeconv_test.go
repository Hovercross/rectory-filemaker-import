package timeconv_test

import (
	"testing"

	"github.com/hovercross/rectory-filemaker-import/pkg/filemaker/timeconv"
)

func Test_DateParse(t *testing.T) {
	p := &timeconv.Parser{}

	p.Defaults()

	stringFunc, found := p.String("M/d/yyyy")

	if !found {
		t.Error("string func was not found")
	}

	formatted, err := stringFunc("3/14/1592")

	if err != nil {
		t.Error(err)
		return
	}

	if formatted != "1592-03-14" {
		t.Errorf("Unexpected formatted time: %s", formatted)
	}
}

func Test_DateParse_NotFound(t *testing.T) {
	p := &timeconv.Parser{}

	p.Defaults()

	stringFunc, found := p.String("Pie/d/yyyy")

	if found {
		t.Error("Something was found")
	}

	if stringFunc != nil {
		t.Error("string func was not nil")
	}
}
