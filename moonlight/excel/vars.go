package excel

import (
	"bufio"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/hakutyou/goapi/core/utils/cosfs"
	"os"
	"time"
)

type Writer struct {
	Title  string
	Author string
	F      *excelize.File
}

func (e *Writer) NewFile() error {
	e.F = excelize.NewFile()
	e.F.SetDefaultFont("等距更纱黑体 SC")
	return e.F.SetDocProps(&excelize.DocProperties{
		Title:          e.Title,
		Subject:        e.Title,
		Creator:        e.Author,
		LastModifiedBy: e.Author,
		Created:        time.Now().Format(time.RFC3339),
		Identifier:     "xlsx",
		Language:       "zh_CN",
		Revision:       "0",
		Version:        "1.0.0",
	})
}

func (e *Writer) NewSheet(sheetName string) (err error) {
	var docProps *excelize.DocProperties

	e.F.NewSheet(sheetName)
	if docProps, err = e.F.GetDocProps(); err != nil {
		return
	}
	return e.F.SetHeaderFooter(sheetName, &excelize.FormatHeaderFooter{
		DifferentFirst:   true,
		DifferentOddEven: true,
		OddHeader:        "&R&P/&N",
		EvenHeader:       "&L&P/&N",
		FirstHeader:      `&R&P/&N&C` + docProps.Title + `&"-,Regular"`,
	})
}

func (e *Writer) MiniMap(sheetName string,
	location []string, _range []string) (err error) {
	return e.F.AddSparkline(sheetName, &excelize.SparklineOption{
		Location: location,
		Range:    _range,
		Markers:  true,
	})
}

func (e *Writer) Save(filename string, isUpload bool) (err error) {
	var filepath string

	filepath = "./temporary/" + filename
	if err = e.F.SaveAs(filepath); err != nil {
		return
	}

	if isUpload {
		var file *os.File
		if file, err = os.Open(filepath); err != nil {
			return
		}
		err = cosfs.CosApi.WriteFile(filename, bufio.NewReader(file))

		file.Close()
		os.Remove(filepath)
	}
	return
}
