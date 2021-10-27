package main

import (
	"github.com/eyasuyuki/yaml2excel/data"
	"github.com/urfave/cli/v2"
	"github.com/xuri/excelize/v2"
	"gopkg.in/yaml.v2"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	app := &cli.App{
		Name:  "yaml2excel",
		Usage: "yaml2excel <yaml filename>",
		Action: convert,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func convert(c *cli.Context) error {
	bytes, err := ioutil.ReadFile(c.Args().First())
	if err != nil {
		log.Fatal(err)
	}

	data := data.CreateData()
	err = yaml.Unmarshal(bytes, &data)
	if err != nil {
		log.Fatal(err)
	}

	var xlsx = excelize.NewFile()
	if err != nil {
		log.Fatal(err)	}

	writeBook(xlsx, data.Book, data.Config)

	err = xlsx.SaveAs(data.Book.Name + ".xlsx")
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func writeBook(xlsx *excelize.File, book *data.Book, config *data.Config) {
	for _, sheet := range book.Sheets {
		writeSheet(xlsx, sheet, config)
	}
}

func writeSheet(xlsx *excelize.File, sheet *data.Sheet, config *data.Config) {
	if xlsx.SheetCount == 1 {
		xlsx.SetSheetName("Sheet1", sheet.Name)
	} else {
		xlsx.NewSheet(sheet.Name)
	}
	xlsx.NewSheet(sheet.Name)
	rowNum := int64(0)
	for _, row := range sheet.Rows {
		rowNum++
		rowNum = writeRow(xlsx, sheet, rowNum, row, config)
	}
}

func writeRow(xlsx *excelize.File, sheet *data.Sheet, rowNum int64, row *data.Row, config *data.Config) int64 {
	colNum := int64(0)
	rowHeight := int64(1)
	for _, col := range row.Cols {
		colNum++
		h := writeCol(xlsx, sheet, rowNum, colNum, col, config)
		if h > rowHeight {
			rowHeight = h
		}
	}
	return rowNum + rowHeight
}

func writeCol(xlsx *excelize.File, sheet *data.Sheet, rowNum int64, colNum int64, col *data.Col, config *data.Config) int64 {
	rh, err := xlsx.GetRowHeight(sheet.Name, 1)
	if err != nil {
		log.Fatal(err)
	}
	rowHeight := int64(1)
	colName, err := excelize.CoordinatesToCellName(int(colNum), int(rowNum))
	if err != nil {
		log.Fatal(err)
	}
	if col.Img != "" {
		scale := strconv.FormatFloat(config.ImgScale, 'f', -1, 64)
		format := "{\"x_scale\": " + scale + ", \"y_scale\": " + scale + "}"
		err = xlsx.AddPicture(sheet.Name, colName, col.Img, format)
		if err != nil {
			log.Fatal(err)
		}
		if config.UseImageHeight {
			image, err := readImage(col.Img)
			if err != nil {
				log.Fatal(err)
			}
			h := float64(image.Bounds().Dy()) / config.VerticalResolution * config.ImgScale / float64(rh) * 72.0
			rowHeight = int64(h)
		}
	} else {
		xlsx.SetCellStr(sheet.Name, colName, col.Text)
	}
	return rowHeight
}

func readImage(file string) (image.Image, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}
