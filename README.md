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

##### `GET/POST /label/{endpoint}/render`

Renders the TSPL code for the given endpoint.

When using a GET request, the query params are passed as arguments, when using a POST request, the request body is parsed as JSON.

##### `GET/POST /label/{endpoint}/print`

Prints 1 label on the printer assigned to the endpoint.

When using a GET request, the query params are passed as arguments, when using a POST request, the request body is parsed as JSON.



### Configuration

Here's the sample config file with explanations:

```yaml
endpoints:
  ABC123: # the name of the endpoint. Think of this as an API key for public-facing apis
    printer:
      device: /dev/usb/lp0 # printer device
      label:
        width: 50 # with of label in mm
        height: 30 # height of label in mm
        gap: 2 # gap between labels in mm
        offset: 0 # offset of gap in mm
    args: # map arguments from query params / json body to template variables.
      ProductName: product_name # template vars should start with uppercase letter
      Id: id
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

Currently only TEXT and BARCODE TSPL commands are supported. Feel free to open an issue or a PR if you need support for anything else. I only needed these two.
