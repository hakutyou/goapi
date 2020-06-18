package Bang

import (
	"context"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/hakutyou/goapi/core/utils/cosfs"
	"github.com/hakutyou/goapi/moonlight/database"
	"github.com/hakutyou/goapi/moonlight/excel"
	uuid "github.com/satori/go.uuid"
	"strconv"
)

type Args struct {
}

type Reply struct {
	Url string
}

func (Bang) GenerateExcel(_ context.Context,
	_ Args, reply *Reply) (err error) {
	var (
		skills       []Skill
		skillDetails []SkillDetail
		streamWriter *excelize.StreamWriter

		e *excel.Writer
	)

	e = &excel.Writer{
		Title:  "Bang Skill",
		Author: "sirat",
	}
	if err = e.NewFile(); err != nil {
		return
	}

	database.DBCfg.DB.Find(&skills)
	for _, each := range skills {
		// 创建工作表
		if err = e.NewSheet(each.Name); err != nil {
			fmt.Println(err)
			return
		}
		// 设置列宽
		if err = e.F.SetColWidth(each.Name, "B", "D", 15); err != nil {
			fmt.Println(err)
			return
		}
		if err = e.F.SetColWidth(each.Name, "E", "E", 45); err != nil {
			fmt.Println(err)
			return
		}
		// 合并单元格
		if err = e.F.MergeCell(each.Name, "B3", "D3"); err != nil {
			fmt.Println(err)
			return
		}

		// 流式写入器
		if streamWriter, err = e.F.NewStreamWriter(each.Name); err != nil {
			fmt.Println(err)
			return
		}
		// 表头
		if err = streamWriter.SetRow("A1", []interface{}{
			"ID",
			each.ID,
			"名称",
			each.Name,
		}); err != nil {
			fmt.Println(err)
			return
		}
		if err = streamWriter.SetRow("A2", []interface{}{
			"类型",
			each.TypeId,
			// each.MaxLevel,
		}); err != nil {
			fmt.Println(err)
			return
		}
		if err = streamWriter.SetRow("A3", []interface{}{
			"描述",
			each.MainDes,
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
		database.DBCfg.DB.Where("skill_id = ?", each.ID).Find(&skillDetails)
		for i, eachDetail := range skillDetails {
			var pos string
			if pos, err = excelize.CoordinatesToCellName(1, 6+i); err != nil {
				fmt.Println(err)
			}
			if err = streamWriter.SetRow(pos, []interface{}{
				i,
				eachDetail.XiuWei,
				eachDetail.BangGong,
				eachDetail.SuiYin,
				eachDetail.Des,
				eachDetail.Props,
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
		lastPos := len(skillDetails) + 5
		lastPosStr := strconv.Itoa(lastPos)
		totalPosStr := strconv.Itoa(lastPos + 1)

		if err = e.F.SetCellFormula(each.Name, "B"+totalPosStr, "SUM(B6:BA"+lastPosStr+")"); err != nil {
			fmt.Println(err)
			return
		}
		if err = e.F.SetCellFormula(each.Name, "C"+totalPosStr, "SUM(C6:CA"+lastPosStr+")"); err != nil {
			fmt.Println(err)
			return
		}
		if err = e.F.SetCellFormula(each.Name, "D"+totalPosStr, "SUM(D6:DA"+lastPosStr+")"); err != nil {
			fmt.Println(err)
			return
		}
		// 创建单元格样式
		var boldStyle int
		if boldStyle, err = e.F.NewStyle(`{"font":{"bold":true}}`); err != nil {
			fmt.Println(err)
			return
		}
		// 设置单元格样式
		if err = e.F.SetCellStyle(each.Name, "A1", "A3", boldStyle); err != nil {
			fmt.Println(err)
			return
		}
		if err = e.F.SetCellStyle(each.Name, "C1", "C1", boldStyle); err != nil {
			fmt.Println(err)
			return
		}
		if err = e.F.SetCellStyle(each.Name, "A5", "F5", boldStyle); err != nil {
			fmt.Println(err)
			return
		}
		// 迷你图
		if err = e.MiniMap(each.Name,
			[]string{"B4", "C4", "D4"},
			[]string{
				each.Name + "!B6:B" + lastPosStr,
				each.Name + "!C6:C" + lastPosStr,
				each.Name + "!D6:D" + lastPosStr,
			}); err != nil {
			return
		}
	}
	// 删除默认的工作表
	e.F.DeleteSheet("Sheet1")
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
