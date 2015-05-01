#!/usr/bin/make -f

SHELL=/bin/bash

dev: 
	@cd app && go build
	@cd app && ./app ~/shareit/data 12345

deps:
	@go get github.com/nu7hatch/gouuid
	@go get github.com/gorilla/mux
	@go get gopkg.in/alecthomas/kingpin.v1
	@go get github.com/goji/httpauth
