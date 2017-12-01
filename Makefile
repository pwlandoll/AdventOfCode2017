.PHONY: clean

clean:
	rm -f *.exe

.PHONY: ms1

ms1: clean ./ms1/p1.go ./ms1/input.txt
	go build ./ms1

