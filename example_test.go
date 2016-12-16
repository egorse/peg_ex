package peg_ex

import (
	"fmt"
	"log"
	"testing"
)

type exec interface {
	Execute()
}

func (p *PegEx) Push(str string) {
	p.Acc = p.Acc + str
}
func (p *PegEx) Flush() {
	p.Words = append(p.Words, p.Acc)
	p.Acc = ""
}

func ExamplePegEx() {
	input := "v=T;v=H;v=E;|v=M;v=A;v=G;v=I;v=C;"

	p := PegEx{Buffer: input}
	p.Init()
	if err := p.Parse(); err != nil {
		log.Panicf("%v", err)
	}

	var pi interface{}
	pi = &p
	if i, ok := pi.(exec); ok {
		i.Execute()
	}
	fmt.Printf("%v", p.Words)

	// Output: [THE MAGIC]
}

//
//
//
func Benchmark100(b *testing.B) {
	benchmark(b, 100)
}
func Benchmark200(b *testing.B) {
	benchmark(b, 200)
}
func Benchmark500(b *testing.B) {
	benchmark(b, 500)
}
func Benchmark1k(b *testing.B) {
	benchmark(b, 1*1024)
}
func Benchmark10k(b *testing.B) {
	benchmark(b, 10*1024)
}
func Benchmark20k(b *testing.B) {
	benchmark(b, 10*1024)
}
func Benchmark50k(b *testing.B) {
	benchmark(b, 50*1024)
}
func Benchmark100k(b *testing.B) {
	benchmark(b, 100*1024)
}
func Benchmark200k(b *testing.B) {
	benchmark(b, 200*1024)
}
func Benchmark500k(b *testing.B) {
	benchmark(b, 500*1024)
}
func Benchmark1M(b *testing.B) {
	benchmark(b, 1*1024*1024)
}
func Benchmark2M(b *testing.B) {
	benchmark(b, 2*1024*1024)
}
func Benchmark4M(b *testing.B) {
	benchmark(b, 4*1024*1024)
}

//
func benchmark(b *testing.B, size int) {
	input := "v=T;v=H;v=E;|v=M;v=A;v=G;v=I;v=C;"
	for len(input) < size {
		input = input + "|v=T;v=H;v=E;|v=M;v=A;v=G;v=I;v=C;"
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p := PegEx{Buffer: input}
		p.Init()
		if err := p.Parse(); err != nil {
			log.Panicf("%v", err)
		}

		var pi interface{}
		pi = &p
		if i, ok := pi.(exec); ok {
			i.Execute()
		}
	}
}
