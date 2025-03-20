# tspl-print-api
REST API written in Go to print using a label printer with TSPL support



### Installing

You can either download a binary from the [latest release](https://github.com/gergof/tspl-print-api/releases/latest) or use the [gergof/tspl-print-api](...) docker image.

When installing with docker you have to pass the required device to the container. For example: `docker run --device /dev/usb/lp0 gergof/tspl-print-api`.



### Usage

You can pass these command line arguments:

```
Usage of ./tspl-print-api:
  -addr string
    	Address to listen on (default "0.0.0.0:3000")
  -config string
    	Path of config file
```

You also have to provide a configuration file. You can see an example one in `dev/config.yml`.



### API

The REST API exposes these endpoints:

##### `GET /ping`

Always replies with JSON `{"message": "pong"}`.

##### `POST /label/{endpoint}/render`

Renders the TSPL code for the given endpoint.

The request body is parsed as JSON using the `args` mapping from the configuration.

##### `POST /label/{endpoint}/print`

Prints 1 label on the printer assigned to the endpoint.

The request body is parsed as JSON using the `args` mapping from the configuration.



### Configuration

Here's the sample config file with explanations:

```yaml
endpoints:
  ABC123: # the name of the endpoint. Think of this as an API key for public-facing apis
    printer:
      device: /dev/usb/lp0 # printer device
      direction: normal # can be normal or inverted
      label:
        width: 50 # with of label in mm
        height: 30 # height of label in mm
        gap: 2 # gap between labels in mm
        offset: 0 # offset of gap in mm
    args: # map arguments from query params / json body to template variables.
      ProductName: product_name # template vars should start with uppercase letter
      Id: product.id # you can use gjson getters (https://github.com/tidwall/gjson) here
      Expiration: expiration
      Manufacturer: manufacturer
    code: # the codes defined here are executed sequentially
      - type: text # TEXT tspl command
        x: 20 # x in pts
        y: 16 # y in pts
        font: "4" # font name
        align: center # alignment. Can be default, left, center, right
        content: "{{ .ProductName }}" # content with template
      - type: barcode # BARCODE tspl command
        x: 20 # x in pts
        y: 56 # y in pts
        height: 80 # height in pts
        codeType: "128" # code type
        humanReadable: center # position of human readable text: none, left, center, right
        align: center # alignment. Can be default, left, center, right
        content: "{{ .Id }}" # content of barcode
      - type: text
        x: 20
        y: 174
        font: "2"
        align: left
        content: "DD = {{ .Expiration }}"
      - type: text
        x: 20
        y: 202
        font: "1"
        align: "center"
        content: "MFR {{ .Manufacturer }}" # you can use templates with additional text
```

##### Supported Codes and arguments

- `text` - Print text
  - `x` - Horizontal starting coordinate (in dots)
  - `y` - Vertical starting coordinate (in dots)
  - `font` - Font to use
  - `align` - Text alignment (valid values: `default`, `left`, `right`, `center`)
  - `content` - The content to print (can use templating)
- `barcode` - Print 1D barcodes
  - `x` - Horizontal starting coordinate (in dots)
  - `y` - Vertical starting coordinate (in dots)
  - `height` - Height of barcode
  - `codeType` - Type of barcode
  - `humanReadable` - Position of human readable label (valid values: `none`, `left`, `center`, `right`)
  - `align` - Alignment of barcode
  - `content` - The content to print (can use templating)
- `pdf417` - Print PDF-417 2D codes
  - `x` - Horizontal starting coordinate (in dots)
  - `y` - Vertical starting coordinate (in dots)
  - `width` - The width of the barcode
  - `height` - The height of the barcode
  - `content` - The content to print (can use templating)
- `qr` - Print QR code
  - `x` - Horizontal starting coordinate (in dots)
  - `y` - Vertical starting coordinate (in dots)
  - `ecc` - ECC level to use for code (valid values: `L`, `M`, `Q`, `H`)
  - `cellWidth` - Width of a cell (valid values from 0 to 10)
  - `content` - The content to print (can use templating)
- `block` - Print text wrapped in a box
  - `x` - Horizontal starting coordinate (in dots)
  - `y` - Vertical starting coordinate (in dots)
  - `width` - The width of the text block
  - `height` - The height of the text block
  - `font` - Font to use
  - `space` - Space between the lines (default to 0)
  - `align` - Text alignment (valid values: `default`, `left`, `right`, `center`)
  - `content` - The content to print (can use templating)
