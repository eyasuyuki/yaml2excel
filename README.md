yaml2excel
====

Create Excel file from YAML description.

# Build

```
go get
go build
```

# Usage

```
yaml2exls <YAML file name>
```

# YAML structure

```
config:
books:
  - book:
    name:
    sheets:
      - sheet:
        name:
        rows:
          - row:
            cols:
              - col
                text:
                img:
```

# YAML example

```
config:
  img_scale: 0.5
books:
  - book:
    name: Book1
    sheets:
      - sheet:
        name: Sheet1
        rows:
          - row:
            cols:
              - text: This is test.
              - text: これがB1に入る予定
          - row:
            cols:
              - img: assets/pic1.png
          - row:
            cols:
              - img: assets/pic2.jpg
          - row:
            cols:
              - text: Test2
      - sheet:
        name: シート2だよ
        rows:
          - row:
            cols:
              - text: Some text
  - book:
    name: 日本語ブック名(仮)
    sheets:
      - sheet:
        name: あああ
        rows:
          - row:
            cols:
              - text: HogePiyo
      - sheet:
        name: 2.Test
        rows:
          - row:
            cols:
              - img: assets/pic2.jpg
```

# default config value

```
config:
  vertical_resolution: 96.0
  use_image_height: true
  img_scale: 1.0
  img_margine: 1.0
```

## ```vertital_resolution```

Excel's vertital DPI. Default value is ```96.0``` (96 DPI).

## ```use_image_height```

If true, It would be calculate image height to skip row number.

If false, Does not calculate image height and does not skip row number.

## ```img_scale```

Scaling ratio of image. Default value is ```1.0```.

## ```img_margin```

Vertical margin after image.

Default value is ```1.0```. It means: _one row height_.

If ```use_image_height``` to false, This value does not use.

# ```books```

Array of ```book```.

# ```book```

```
book:
  name:
  sheets:
```

## ```book``` ```name```

book name to be use .xlsx file name.

book name without ".xlsx".

## ```book``` ```sheets```

Array of sheet.

# ```sheet```

```
sheet:
  name:
  rows:
```

## ```sheet``` ```name```

It will be use sheet name.

## ```sheet``` ```rows```

Array of row.

# ```row```

```
row:
  cols:
```

## ```row``` ```cols```

Array of col.

# ```col```

```
col:
  text:
  img:
```

## ```col``` ```text```

Cell text.

## ```col``` ```img```

Image file path.

