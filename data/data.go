package data

type Data struct {
	Config *Config `yaml:"config"`
	Books  []*Book   `yaml:"books"`
}

func CreateData() *Data {
	return &Data{
		Config: CreateConfig(),
	}
}

func CreateConfig() *Config {
	return &Config{
		VerticalResolution: 96.0, // 96 dpi
		UseImageHeight:     true,
		ImgScale:           1.0,
		ImgMargin: 			1.0,
	}
}

type Config struct {
	VerticalResolution   float64 `yaml:"vertical_resolution"`
	UseImageHeight       bool    `yaml:"use_image_height"`
	ImgScale             float64 `yaml:"img_scale""`
	ImgMargin			float64 `yaml:"img_margin"`
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
}
