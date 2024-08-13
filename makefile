include .env

.PHONY: run
.DEFAULT_GOAL:= run

run:
	go build -o /tmp/build . && /tmp/build