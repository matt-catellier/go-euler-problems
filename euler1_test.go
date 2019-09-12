package main

import "testing"

// Unit Tests
func TestRecursive(t *testing.T) {
  actual := Recursive(1000)
  expected := 233168

  if actual != expected {
    t.Errorf("recursive failed -> actual: %v != expected: %v", actual, expected)
  }
}

func TestLoop(t *testing.T) {
  actual := Loop(1000)
  expected := 233168

  if actual != expected {
    t.Errorf("loop failed -> actual: %v != expected: %v", actual, expected)
  }
}

func TestEuler1_1000(t * testing.T) {
  actual := Euler1(1000)
  expected := 233168

  if actual != expected {
    t.Errorf("SumN failed: wanted %v got %v", actual, expected)
  }
}

// Benchmarks 
var result int // used to prevent compiler optimisations
func benchmarkRecursive(i int, b *testing.B) {
  var r int
  for n := 0; n < b.N; n++ {
    r = Recursive(i)
  }
  result = r
}

func BenchmarkRec10(b *testing.B) { benchmarkRecursive(10, b) }
func BenchmarkRec100(b *testing.B) { benchmarkRecursive(100, b) }
func BenchmarkRec1000(b *testing.B) { benchmarkRecursive(1000, b) }
func BenchmarkRec10000(b *testing.B) { benchmarkRecursive(10000, b) }

func benchmarkLoop(i int, b * testing.B) {
  for n := 0; n < b.N; n++ {
    Loop(i)
  }
}

func BenchmarkLoop10(b * testing.B) { benchmarkLoop(10, b) }
func BenchmarkLoop100(b * testing.B) { benchmarkLoop(100, b) }
func BenchmarkLoop1000(b * testing.B) { benchmarkLoop(1000, b) }
func BenchmarkLoop10000(b * testing.B) { benchmarkLoop(10, b) }

func benchmarkEuler1(i float64, b * testing.B) {
  var r int
  for n := 0; n < b.N; n++ {
    r = Euler1(i)
  }
  result = r
}

func BenchmarkEuler10(b *testing.B) { benchmarkEuler1(100, b) }
func BenchmarkEuler100(b *testing.B) { benchmarkEuler1(100, b) }
func BenchmarkEuler1000(b *testing.B) { benchmarkEuler1(10000, b) }
func BenchmarkEuler10000(b *testing.B) { benchmarkEuler1(10000, b) }

