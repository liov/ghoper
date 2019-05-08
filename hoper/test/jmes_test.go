package test

import (
	"fmt"
	"github.com/jmespath/go-jmespath"
	"testing"
)

func Test(t *testing.T) {
	user := struct {
		ID   int
		Name string
	}{
		ID:   1,
		Name: "a",
	}
	users := []interface{}{user}
	res, _ := jmespath.Search("[?ID!='1']", users)
	fmt.Println(res)
}
