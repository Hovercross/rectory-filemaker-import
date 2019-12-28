package filemaker_test

import (
	"encoding/xml"
	"testing"

	"github.com/go-test/deep"
	"github.com/hovercross/rectory-filemaker-import/pkg/filemaker"
)

func Test_SuccessfulRead(t *testing.T) {
	data := getSampleData()

	expected := &filemaker.File{
		XMLName: xml.Name{
			Space: "http://www.filemaker.com/fmpxmlresult",
			Local: "FMPXMLRESULT",
		},
		ErrorCode: 0,
		Product: filemaker.Product{
			XMLName: xml.Name{
				Space: "http://www.filemaker.com/fmpxmlresult",
				Local: "PRODUCT",
			},
			Build:   "10-27-2016",
			Name:    "FileMaker",
			Version: "Server 15.0.3",
		},
		Database: filemaker.Database{
			XMLName: xml.Name{
				Space: "http://www.filemaker.com/fmpxmlresult",
				Local: "DATABASE",
			},
			DateFormat: "M/d/yyyy",
			Layout:     "",
			Name:       "ksATTENDANCE.fmp12",
			Records:    "21689",
			TimeFormat: "h:mm:ss a",
		},
		Metadata: filemaker.Metadata{
			XMLName: xml.Name{
				Space: "http://www.filemaker.com/fmpxmlresult",
				Local: "METADATA",
			},
			Fileds: []filemaker.Field{
				filemaker.Field{
					XMLName: xml.Name{
						Space: "http://www.filemaker.com/fmpxmlresult",
						Local: "FIELD",
					},
					EmptyOK: "NO", MaxRepeat: "1", Name: "IDINCIDENT", Type: "TEXT"},
				filemaker.Field{
					XMLName: xml.Name{
						Space: "http://www.filemaker.com/fmpxmlresult",
						Local: "FIELD",
					},
					EmptyOK: "YES", MaxRepeat: "1", Name: "Det Date", Type: "DATE"},
			},
		},
		ResultSet: filemaker.ResultSet{
			XMLName: xml.Name{
				Space: "http://www.filemaker.com/fmpxmlresult",
				Local: "RESULTSET",
			},
			Found: "21689",
			Rows: []filemaker.Row{
				filemaker.Row{
					XMLName: xml.Name{
						Space: "http://www.filemaker.com/fmpxmlresult",
						Local: "ROW",
					},
					ModID:    "8",
					RecordID: "138",
					Cols: []filemaker.Col{
						filemaker.Col{
							XMLName: xml.Name{
								Space: "http://www.filemaker.com/fmpxmlresult",
								Local: "COL",
							},
							Data: []string{"134"},
						},
						filemaker.Col{
							XMLName: xml.Name{
								Space: "http://www.filemaker.com/fmpxmlresult",
								Local: "COL",
							},
							Data: []string{"9/13/2011"},
						},
					},
				},
			},
		},
	}

	parsed, err := filemaker.ReadData(data)

	if err != nil {
		t.Error(err)
		return
	}

	if err := deep.Equal(expected, parsed); err != nil {
		t.Error(err)
	}
}

func getSampleData() []byte {
	return []byte(`<?xml version="1.0" ?>
	<FMPXMLRESULT xmlns="http://www.filemaker.com/fmpxmlresult">
		<ERRORCODE>0</ERRORCODE>
		<PRODUCT BUILD="10-27-2016" NAME="FileMaker" VERSION="Server 15.0.3"/>
		<DATABASE DATEFORMAT="M/d/yyyy" LAYOUT="" NAME="ksATTENDANCE.fmp12" RECORDS="21689" TIMEFORMAT="h:mm:ss a"/>
		<METADATA>
			<FIELD EMPTYOK="NO" MAXREPEAT="1" NAME="IDINCIDENT" TYPE="TEXT"/>
			<FIELD EMPTYOK="YES" MAXREPEAT="1" NAME="Det Date" TYPE="DATE"/>
		</METADATA>
		<RESULTSET FOUND="21689">
			<ROW MODID="8" RECORDID="138">
				<COL>
					<DATA>134</DATA>
				</COL>
				<COL>
					<DATA>9/13/2011</DATA>
				</COL>
			</ROW>
		</RESULTSET>
	</FMPXMLRESULT>`)
}
