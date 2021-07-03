
dev: main.go
	go build ./
	./mylang

build: main.go
	go build ./

run: mylang
	./mylang