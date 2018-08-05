PWD = $(shell pwd)

build:
	qtc templates
	GOARCH=wasm GOOS=js go build -o app.wasm .

serve:
	docker build -t nginx .
	docker run --rm -p 3000:80 -v $(PWD):/usr/share/nginx/html nginx
