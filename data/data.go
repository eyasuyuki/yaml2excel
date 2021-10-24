package data

type Data struct {
	Config *Config `yaml:"config"`
	Book   *Book   `yaml:"book"`
}

func CreateData() *Data {
	return &Data{
		Config: CreateConfig(),
	}
}

func CreateConfig() *Config {
	return &Config{
		HorizontalResolution: 96.0, // 96 dpi
		VerticalResolution:   96.0, // 96 dpi
		RowHeight:            6.61, // 18.75pt
		MaxWidth:             190,  // 210 - (10 x 2)
		UseImageHeight:       true,
		ImgScale:             1.0,
	}
}

type Config struct {
	HorizontalResolution float64 `yaml:"horizontal_resolution"`
	VerticalResolution   float64 `yaml:"vertical_resolution"`
	RowHeight            float64 `yaml:"row_height"`
	MaxWidth             float64 `yaml:"max_width"`
	UseImageHeight       bool    `yaml:"use_image_height"`
	ImgScale             float64 `yaml:"img_scale""`
}

type Book struct {
	Name   string   `yaml:"name"`
	Sheets []*Sheet `yaml:"sheets"`
}

type Sheet struct {
	Name string `yaml:"name"`
	Rows []*Row `yaml:"rows"`
}

type Row struct {
	LineNum int64  `yaml:"line_num"`
	Cols    []*Col `yaml:"cols"`
}

type Col struct {
	Text string `yaml:"text"`
	Img  string `yaml:"img"`
	Pdf  string `yaml:"pdf"`
}
