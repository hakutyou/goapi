package Excel

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/hakutyou/goapi/rpcx/utils/cosfs"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	uuid "github.com/satori/go.uuid"
)

type Excel int

type Args struct {
	Title string
}

type Reply struct {
	Url string
}

type WxSkillMain struct {
	SkillId   int            `json:"skillId"`
	SkillType int            `json:"skillType"`
	SkillName string         `json:"skillName"`
	MaxLevel  int            `json:"maxLevel"`
	MainDes   string         `json:"mainDes"`
	Levels    []WxSkillLevel `json:"levels"`
}

type WxSkillLevel struct {
	Xiuwei   int      `json:"xiuwei"`
	Banggong int      `json:"banggong"`
	Suiyin   int      `json:"suiyin"`
	Des      string   `json:"des"`
	Props    []string `json:"props"`
}

func (t *Excel) GenerateExcel(ctx context.Context, args *Args, reply *Reply) (err error) {
	var (
		data         []byte
		wxSkillMain  []WxSkillMain
		streamWriter *excelize.StreamWriter

		e *excelWrite
	)
	// 解析原始 json
	if data, err = ioutil.ReadFile("data/wx_skill.json"); err != nil {
		fmt.Println(err)
		return
	}
	if err = json.Unmarshal(data, &wxSkillMain); err != nil {
		fmt.Println(err)
		return
	}
	// 创建文件
	e = &excelWrite{
		args.Title,
		"sirat",
		nil,
	}
	if err = e.NewFile(); err != nil {
		return
	}

	for _, vMain := range wxSkillMain {
		// 创建工作表
		if err = e.NewSheet(vMain.SkillName); err != nil {
			fmt.Println(err)
			return
		}
		// 设置列宽
		if err = e.f.SetColWidth(vMain.SkillName, "B", "D", 15); err != nil {
			fmt.Println(err)
			return
		}
		if err = e.f.SetColWidth(vMain.SkillName, "E", "E", 45); err != nil {
			fmt.Println(err)
			return
		}
		// 合并单元格
		if err = e.f.MergeCell(vMain.SkillName, "B3", "D3"); err != nil {
			fmt.Println(err)
			return
		}

		// 流式写入器
		if streamWriter, err = e.f.NewStreamWriter(vMain.SkillName); err != nil {
			fmt.Println(err)
			return
		}
		// 表头
		if err = streamWriter.SetRow("A1", []interface{}{
			"ID",
			vMain.SkillId,
			"名称",
			vMain.SkillName,
		}); err != nil {
			fmt.Println(err)
			return
		}
		if err = streamWriter.SetRow("A2", []interface{}{
			"类型",
			vMain.SkillType,
			// vMain.MaxLevel,
		}); err != nil {
			fmt.Println(err)
			return
		}
		if err = streamWriter.SetRow("A3", []interface{}{
			"描述",
			vMain.MainDes,
		}); err != nil {
			fmt.Println(err)
			return
		}
		if err = streamWriter.SetRow("A5", []interface{}{
			"序号",
			"2",
			"3",
			"4",
			"描述",
		}); err != nil {
			fmt.Println(err)
			return
		}
		// 循环填充数据
		for i, vLevel := range vMain.Levels {
			var pos string
			if pos, err = excelize.CoordinatesToCellName(1, 6+i); err != nil {
				fmt.Println(err)
			}
			if err = streamWriter.SetRow(pos, []interface{}{
				i,
				vLevel.Xiuwei,
				vLevel.Banggong,
				vLevel.Suiyin,
				vLevel.Des,
				vLevel.Props,
			}); err != nil {
				fmt.Println(err)
				return
			}
		}
		// Flush 后写入
		if err = streamWriter.Flush(); err != nil {
			fmt.Println(err)
			return
		}
		// 写入总计
		lastPos := len(vMain.Levels) + 5
		lastPosStr := strconv.Itoa(lastPos)
		totalPosStr := strconv.Itoa(lastPos + 1)

		if err = e.f.SetCellFormula(vMain.SkillName, "B"+totalPosStr, "SUM(B6:BA"+lastPosStr+")"); err != nil {
			fmt.Println(err)
			return
		}
		if err = e.f.SetCellFormula(vMain.SkillName, "C"+totalPosStr, "SUM(C6:CA"+lastPosStr+")"); err != nil {
			fmt.Println(err)
			return
		}
		if err = e.f.SetCellFormula(vMain.SkillName, "D"+totalPosStr, "SUM(D6:DA"+lastPosStr+")"); err != nil {
			fmt.Println(err)
			return
		}
		// 创建单元格格式
		var boldStyle int

		if boldStyle, err = e.f.NewStyle(`{"font":{"bold":true}}`); err != nil {
			fmt.Println(err)
			return
		}
		// 设置单元格样式
		if err = e.f.SetCellStyle(vMain.SkillName, "A1", "A3", boldStyle); err != nil {
			fmt.Println(err)
			return
		}
		if err = e.f.SetCellStyle(vMain.SkillName, "C1", "C1", boldStyle); err != nil {
			fmt.Println(err)
			return
		}
		if err = e.f.SetCellStyle(vMain.SkillName, "A5", "F5", boldStyle); err != nil {
			fmt.Println(err)
			return
		}
		// 迷你图
		if err = e.MiniMap(vMain.SkillName,
			[]string{"B4", "C4", "D4"},
			[]string{
				vMain.SkillName + "!B6:B" + lastPosStr,
				vMain.SkillName + "!C6:C" + lastPosStr,
				vMain.SkillName + "!D6:D" + lastPosStr,
			}); err != nil {
			return
		}
	}
	// 删除默认的工作表
	e.f.DeleteSheet("Sheet1")
	// 设置工作簿的默认工作表
	// f.SetActiveSheet(1)

	// 保存文件
	var filename string
	filename = uuid.NewV4().String() + ".xlsx"

	if err = e.Save(filename, false); err != nil {
		return
	}
	reply.Url = fmt.Sprintf("https://%s.cos.%s.myqcloud.com/%s",
		cosfs.CosApi.Bucket, cosfs.CosApi.Region, filename)
	return
}
