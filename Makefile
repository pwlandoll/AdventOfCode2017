GOPATH=$(shell pwd)

.PHONY: clean
clean: ./bin/
	rm -f ./bin/*

setup: clean ./src/

ms1: setup ./src/ms1_p1/p1.go ./src/ms1_p2/p2.go ./inputs/ms1.txt
	go install ms1_p1 && go install ms1_p2

ms2: setup ./src/ms2_p1/p1.go ./inputs/ms2.txt
	go install ms2_p1

.PHONY: all
all: ms1 ms2

