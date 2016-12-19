.PHONY: all
all:
	# Generate parsers
	../peg/peg -switch -inline        ex.peg && mv ex.peg.go ex.peg.go-vanilla
	../peg/peg -switch -inline -noast ex.peg && mv ex.peg.go ex.peg.go-no-ast
	#go test
	rm -f *.bench

	cp ex.peg.go-vanilla ex.peg.go
	go test
	cp ex.peg.go-mod128  ex.peg.go
	go test
	cp ex.peg.go-no-ast  ex.peg.go
	go test

.PHONY: bench
bench:
	cp ex.peg.go-vanilla ex.peg.go
	go test -benchmem -bench . | tee vanilla.bench

	cp ex.peg.go-mod128  ex.peg.go
	go test -benchmem -bench . | tee mod128.bench

	cp ex.peg.go-no-ast  ex.peg.go
	go test -benchmem -bench . | tee no-ast.bench

	benchcmp vanilla.bench no-ast.bench
