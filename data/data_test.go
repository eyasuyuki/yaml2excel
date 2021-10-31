package data_test

import (
	"github.com/eyasuyuki/yaml2excel/data"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"testing"
)

func TestRead(t *testing.T) {
	bytes, err := ioutil.ReadFile("../assets/Book1.yaml")
	if err != nil {
		log.Fatal(err)
	}

	data := data.CreateData()

	err = yaml.Unmarshal(bytes, &data);
	if err != nil {
		log.Fatal(err)
	}


	if data.Config == nil {
		t.Errorf("expect: nil, but got %v \n", data.Config)
	}

	if len(data.Books) != 1 {
		t.Errorf("expect len(data.Books): 1, but got %v\n", len(data.Books))
	}

	if data.Books[0].Name != "Book1" {
		t.Errorf("expect: Book1, but got %v \n", data.Books[0].Name)
	}

	if len(data.Books[0].Sheets) != 1 {
		t.Errorf("expect: 1, bat got %v \n", len(data.Books[0].Sheets))
	}

	if data.Books[0].Sheets[0].Name != "Sheet1" {
		t.Errorf("expect: Sheet1, but got %v \n", data.Books[0].Sheets[0].Name)
	}


	if len(data.Books[0].Sheets[0].Rows) != 4 {
		t.Errorf("expect: 4, but got %v \n", len(data.Books[0].Sheets[0].Rows))
	}

	if len(data.Books[0].Sheets[0].Rows[0].Cols) != 2 {
		t.Errorf("expect: 2, but got %v\n", len(data.Books[0].Sheets[0].Rows[0].Cols))
	}

	if data.Books[0].Sheets[0].Rows[0].Cols[0].Text != "This is test." {
		t.Errorf("expect: \"This is test.\", but got \"%v\" \n", data.Books[0].Sheets[0].Rows[0].Cols[0].Text)
	}

	if data.Books[0].Sheets[0].Rows[0].Cols[1].Text != "これがB1に入る予定" {
		t.Errorf("expect: \"これがB1に入る予定\", but got \"%v\" \n", data.Books[0].Sheets[0].Rows[0].Cols[1].Text)
	}

	if data.Books[0].Sheets[0].Rows[1].Cols[0].Img != "assets/pic1.png" {
		t.Errorf("expect: assets/pic1.png, but gut %v \n", data.Books[0].Sheets[0].Rows[1].Cols[0].Img)
	}

	if data.Books[0].Sheets[0].Rows[2].Cols[0].Img != "assets/pic2.jpg" {
		t.Errorf("expect: assets/pic2.jpg, but gut %v \n", data.Books[0].Sheets[0].Rows[2].Cols[0].Img)
	}

	if data.Books[0].Sheets[0].Rows[3].Cols[0].Text != "Test2" {
		t.Errorf("expect: Test2, but got %v \n", data.Books[0].Sheets[0].Rows[3].Cols[0].Text)
	}
}

