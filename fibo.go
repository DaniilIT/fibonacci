package main

import (
	"fmt"
	"math"
	"reflect"
)

func f(n int, cache []uint64) (num uint64) {
	if val := cache[n]; val != 0 {
		return val
	}

	f1 := f(n-1, cache)
	f2 := f(n-2, cache)
	if math.MaxUint64-f1 < f2 {
		panic(fmt.Sprintf("overflow of %s at step %d", reflect.TypeOf(num), n))
	}
	num = f1 + f2
	cache[n] = num
	return num
}

func CalcRecursive(n int) (res uint64, err error) {
	if n < 0 {
		err = fmt.Errorf("number must be >= 0")
		return
	}
	if n <= 1 {
		res = uint64(n)
		return
	}

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
		}
	}()

	cache := make([]uint64, n+1)
	cache[1], cache[2] = 1, 1
	res = f(n, cache)
	return res, nil
}

func CalcIterative(n int) (res uint64, err error) {
	if n < 0 {
		err = fmt.Errorf("number must be >= 0")
		return
	}
	if n <= 1 {
		res = uint64(n)
		return
	}

	var f1, f2 uint64 = 0, 1
	for i := 2; i <= n; i++ {
		if math.MaxUint64-f1 < f2 {
			return 0, fmt.Errorf("overflow of %s at step %d", reflect.TypeOf(res), i)
		}
		f1, f2 = f2, f1+f2
	}
	res = f2
	return res, nil
}

func CalcBinet(n int) (res uint64, err error) {
	if n < 0 {
		err = fmt.Errorf("number must be >= 0")
		return
	}

	phi := (1 + math.Sqrt(5)) / 2
	nf := float64(n)
	ff := (math.Pow(phi, nf) - (math.Pow(-phi, -nf))) / (2*phi - 1)

	ff = math.Round(ff)
	if ff > math.MaxUint64 {
		return 0, fmt.Errorf("overflow of %s", reflect.TypeOf(res))
	}

	res = uint64(ff)
	return res, nil
}
