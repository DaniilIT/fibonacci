package main

import (
	"math/rand"
	"testing"
)

var strategies = [3]struct {
	name     string
	calcFibo func(int) (uint64, error)
}{
	{"Recursive", CalcRecursive},
	{"Iterative", CalcIterative},
	{"Binet", CalcBinet},
}

func TestCalc(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantRes uint64
		wantErr bool
	}{
		{
			name:    "#negativeNumber",
			args:    args{n: -1},
			wantErr: true,
		},
		{
			name:    "#0",
			args:    args{n: 0},
			wantRes: 0,
		},
		{
			name:    "#1",
			args:    args{n: 1},
			wantRes: 1,
		},
		{
			name:    "#2",
			args:    args{n: 2},
			wantRes: 1,
		},
		{
			name:    "#3",
			args:    args{n: 3},
			wantRes: 2,
		},
		{
			name:    "#10",
			args:    args{n: 10},
			wantRes: 55,
		},
		{
			name:    "#42",
			args:    args{n: 42},
			wantRes: 267914296,
		},
		{
			name:    "#70",
			args:    args{n: 70},
			wantRes: 190392490709135,
		},
		{
			name:    "#largeNumber",
			args:    args{n: 94},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, s := range strategies {
				gotRes, err := s.calcFibo(tt.args.n)
				if (err != nil) != tt.wantErr {
					t.Errorf("Calc%s(%d) error = %v, wantErr %v", s.name, tt.args.n, err, tt.wantErr)
					return
				}
				if gotRes != tt.wantRes {
					t.Errorf("Calc%s(%d) = %v, want %v", s.name, tt.args.n, gotRes, tt.wantRes)
				}
			}
		})
	}
}

const SEED = 42

func BenchmarkCalcRecursive(b *testing.B) {
	// 298.7 ns/op
	randGen := rand.New(rand.NewSource(SEED))
	for b.Loop() {
		n := randGen.Intn(100)
		CalcRecursive(n)
	}
}

func BenchmarkCalcIterative(b *testing.B) {
	// 26.78 ns/op
	randGen := rand.New(rand.NewSource(SEED))
	for b.Loop() {
		n := randGen.Intn(100)
		CalcIterative(n)
	}
}

func BenchmarkCalcBinet(b *testing.B) {
	// 51.33 ns/op
	randGen := rand.New(rand.NewSource(SEED))
	for b.Loop() {
		n := randGen.Intn(100)
		CalcBinet(n)
	}
}
