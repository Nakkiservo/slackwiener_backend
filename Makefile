SHELL:=/bin/bash


SRC := $(shell find . -type f -name '*.go') 

.PHONY: all clean

all : build/ build/server build/app_config.toml

clean :
	rm -rf build
	rm -f front/dist/*
	rm -f front/cache/*

build/:
	mkdir -p build
	mkdir -p build/db

build/server: $(SRC)
	go build -buildmode=exe -o build/server

build/app_config.toml: app_config.toml
	@cp app_config.toml ./build/

