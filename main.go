package main

import (
	"github.com/eyasuyuki/yaml2excel/data"
	"github.com/xuri/excelize/v2"
	"gopkg.in/yaml.v2"
	"image"
	"io/ioutil"
	"log"
	"os"
	"github.com/urfave/cli/v2"
	"strconv"
	_ "image/jpeg"
	_ "image/gif"
	_ "image/png"
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
	rowHeight := int64(1)
	if col.Img != "" {
		xlsx.AddPicture(sheet.Name, getColumnName(rowNum, colNum), col.Img, "")
		if config.UseImageHeight {
			// TODO read img
			image, err := readImage(col.Img)
			if err != nil {
				log.Fatal(err)
			}
			h := float64(image.Bounds().Dy()) / config.HorizontalResolution * 25.4 / config.RowHeight
			rowHeight = int64(h)
		}
	} else {
		xlsx.SetCellStr(sheet.Name, getColumnName(rowNum, colNum), col.Text)
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

func getColumnName(rowNum int64, colNum int64) string {
	colName := ""
	n := colNum
	for n > 0 {
		rem := n % 26
		if rem == 0 {
			colName += "Z"
			n = (n / 26) - 1
		} else {
			colName += num2Col(rem - 1)
			n = n / 26
		}
	}
	result := colName + strconv.FormatInt(rowNum, 10)
	//fmt.Println(result)
	return result
}

func num2Col(n int64) string {
	cols := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	return cols[n];
}