package disk_test

import (
	"encoding/xml"
	"testing"

	"github.com/go-test/deep"
	"github.com/hovercross/rectory-filemaker-import/pkg/filemaker/disk"
	"github.com/hovercross/rectory-filemaker-import/pkg/filemaker/sane"
)

func Test_Sane(t *testing.T) {
	raw := getTestParsedData()

	sanified, err := raw.Sane()

	if err != nil {
		t.Error(err)
		return
	}

	expected := &sane.Data{
		Product: &sane.Product{
			Build:   "10-27-2016",
			Name:    "FileMaker",
			Version: "Server 15.0.3",
		},
		Database: &sane.Database{
			DateFormat: "M/d/yyyy",
			Name:       "ksATTENDANCE.fmp12",
			Records:    21689,
			TimeFormat: "h:mm:ss a",
		},
		Metadata: &sane.Metadata{
			Fields: []*sane.Field{
				&sane.Field{EmptyOK: false, MaxRepeat: 1, Name: "IDINCIDENT", Type: "TEXT"},
				&sane.Field{EmptyOK: true, MaxRepeat: 1, Name: "Det Date", Type: "DATE"},
			},
		},
		ResultSet: &sane.ResultSet{
			Found: 21689,
			Rows: []*sane.Row{
				&sane.Row{
					ModID:    "8",
					RecordID: "138",
					Cols: []*sane.Col{
						&sane.Col{
							Data:  []string{"134"},
							Field: &sane.Field{EmptyOK: false, MaxRepeat: 1, Name: "IDINCIDENT", Type: "TEXT"},
						},
						&sane.Col{
							Data:  []string{"9/13/2011"},
							Field: &sane.Field{EmptyOK: true, MaxRepeat: 1, Name: "Det Date", Type: "DATE"},
						},
					},
				},
			},
		},
	}

	for _, diff := range deep.Equal(expected, sanified) {
		t.Error(diff)
	}
}

func getTestParsedData() *disk.File {
	data := getSampleData()

	parsed := &disk.File{}

	if err := xml.Unmarshal(data, &parsed); err != nil {
		panic(err)
	}

	return parsed
}
