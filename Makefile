all: build run

build:
	docker build -t zyfdedh/crushon .

run:
	docker run -e NICKNAME=${NICKNAME} zyfdedh/crushon
