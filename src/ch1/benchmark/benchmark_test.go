package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo1()
	}
}

func Echo1() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%.2fs elapsed", time.Since(start).Seconds())
}

func Echo2() {
	start := time.Now()
	s, seq := "", ""
	for _, arg := range os.Args[1:] {
		s += arg + seq
		seq = " "
	}
	fmt.Println(s)
	fmt.Printf("%.2fs elapsed", time.Since(start).Seconds())
}

func TestEcho(t *testing.T) {
	Echo1()
	Echo2()
}
