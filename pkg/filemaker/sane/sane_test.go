package sane_test

import (
	"testing"

	"github.com/hovercross/rectory-filemaker-import/pkg/filemaker/sane"
)

func Test_ImposeColString(t *testing.T) {
	col := &sane.Col{
		Data:  []string{"134"},
		Field: &sane.Field{EmptyOK: false, MaxRepeat: 1, Name: "IDINCIDENT", Type: "TEXT"},
	}

	type Record struct {
		Id string `filemaker:"IDINCIDENT"`
	}

	r := &Record{}

	if err := col.ImposeOnto(r); err != nil {
		t.Error(err)
	}

	if r.Id != "134" {
		t.Errorf("Id not as expected: %s", r.Id)
	}

}
