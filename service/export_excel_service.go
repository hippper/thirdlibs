package service

import (
	"fmt"
	"sync"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/luckyweiwei/base/utils"

	"golang.org/x/sync/semaphore"
)

type ExportExcelService struct {
	Sem *semaphore.Weighted
}

var exportExcelService *ExportExcelService
var exportExcelOnce sync.Once

func ExportExcelServiceInit() *ExportExcelService {
	exportExcelOnce.Do(func() {
		exportExcelService = &ExportExcelService{
			Sem: semaphore.NewWeighted(3), // 信号量 默认三个
		}
	})
	return exportExcelService
}

func ExportExcelServiceInstance() *ExportExcelService {
	utils.ASSERT(exportExcelService != nil)
	return exportExcelService
}

func (service *ExportExcelService) Export(xlsx *excelize.File, title []interface{}, in <-chan []interface{}, out chan bool) {
	defer service.Sem.Release(1)

	var (
		sheetname = "Sheet1"
		axis      = "A1"
		sheetNum  = 1
		rowNum    = 1
	)

	err := xlsx.SetSheetRow(sheetname, axis, &title)
	if err != nil {
		Log.Error(err)
		return
	}

	rowNum++
	for row := range in {
		if rowNum%50001 == 1 {
			sheetNum++
			rowNum = 1
			sheetname = fmt.Sprintf("Sheet%d", sheetNum)

			index := xlsx.NewSheet(sheetname)
			xlsx.SetActiveSheet(index)

			err := xlsx.SetSheetRow(sheetname, axis, &title)
			if err != nil {
				Log.Error(err)
				return
			}

			rowNum++
		}
		err := xlsx.SetSheetRow(sheetname, fmt.Sprintf("A%d", rowNum), &row)
		if err != nil {
			Log.Error(err)
			return
		}
		rowNum++
	}
	close(out)
	Log.Debug("export end!!!!!!")
}
