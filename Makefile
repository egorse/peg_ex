.PHONY:
all:
	# TODO Use -noast
	#peg -switch -inline ex.peg
	#go test
	rm *.bench

	cp ex.peg.go-vanilla ex.peg.go
	go test
	go test -benchmem -bench . | tee vanilla.bench
	cp ex.peg.go-mod128  ex.peg.go
	go test
	go test -benchmem -bench . | tee mod128.bench
	cp ex.peg.go-no-ast  ex.peg.go
	go test
	go test -benchmem -bench . | tee no-ast.bench

	benchcmp vanilla.bench no-ast.bench

