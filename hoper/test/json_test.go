package test

import (
	"github.com/json-iterator/go"
	"hoper/model"
	"testing"
)

var u = model.User{}

func Benchmark_One(b *testing.B) {

	for i := 0; i < b.N; i++ {
		jsoniter.ConfigCompatibleWithStandardLibrary.MarshalToString(u)
	}
}

func Benchmark_Two(b *testing.B) {

	for i := 0; i < b.N; i++ {
		jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(u)
	}
}
