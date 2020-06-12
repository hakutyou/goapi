package Excel

import (
	"bufio"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/hakutyou/goapi/rpcx/utils/cosfs"
	"os"
	"time"
)

type excelWrite struct {
	Title  string
	Author string
	f      *excelize.File
}

func (e *excelWrite) NewFile() error {
	e.f = excelize.NewFile()
	e.f.SetDefaultFont("等距更纱黑体 SC")
	return e.f.SetDocProps(&excelize.DocProperties{
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

func (e *excelWrite) NewSheet(sheetName string) (err error) {
	var docProps *excelize.DocProperties

	e.f.NewSheet(sheetName)
	if docProps, err = e.f.GetDocProps(); err != nil {
		return
	}
	return e.f.SetHeaderFooter(sheetName, &excelize.FormatHeaderFooter{
		DifferentFirst:   true,
		DifferentOddEven: true,
		OddHeader:        "&R&P/&N",
		EvenHeader:       "&L&P/&N",
		FirstHeader:      `&R&P/&N&C` + docProps.Title + `&"-,Regular"`,
	})
}

func (e *excelWrite) MiniMap(sheetName string,
	location []string, _range []string) (err error) {
	return e.f.AddSparkline(sheetName, &excelize.SparklineOption{
		Location: location,
		Range:    _range,
		Markers:  true,
	})
}

func (e *excelWrite) Save(filename string, isUpload bool) (err error) {
	var filepath string

	filepath = "./temporary/" + filename
	if err = e.f.SaveAs(filepath); err != nil {
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
