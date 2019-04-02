package main

import (
	"github.com/kataras/golog"
	"testing"
)

/**
 * @author     ：lbyi
 * @date       ：Created in 2019/4/2
 * @description：
 */

func Benchmark_Log(b *testing.B) {

	for i := 0; i < b.N; i++ {
		golog.Error("Hey, warning here")
	}
}
