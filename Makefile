tspl-print-api: *.go
	go build

run:
	go run . -config=dev/config.yml
