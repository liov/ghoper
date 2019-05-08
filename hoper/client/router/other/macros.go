package other

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/hero"
	"reflect"
	"sort"
	"strings"
)

func Macros(app *iris.Application) {
	app.Macros().Get("string").RegisterFunc("range", func(minLength, maxLength int) func(string) bool {
		return func(paramValue string) bool {
			return len(paramValue) >= minLength && len(paramValue) <= maxLength
		}
	})

	app.Get("/api/limitchar/{name:string range(1,200) else 400}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.Writef(`Hello %s | the name should be between 1 and 200 characters length
    otherwise this handler will not be executed`, name)
	})

	app.Macros().Get("string").RegisterFunc("has", func(validNames []string) func(string) bool {
		return func(paramValue string) bool {
			for _, validName := range validNames {
				if validName == paramValue {
					return true
				}
			}

			return false
		}
	})

	app.Macros().Register("slice", "", false, true, func(paramValue string) (interface{}, bool) {
		return strings.Split(paramValue, "/"), true
	}).RegisterFunc("contains", func(expectedItems []string) func(paramValue []string) bool {
		sort.Strings(expectedItems)
		return func(paramValue []string) bool {
			if len(paramValue) != len(expectedItems) {
				return false
			}

			sort.Strings(paramValue)
			for i := 0; i < len(paramValue); i++ {
				if paramValue[i] != expectedItems[i] {
					return false
				}
			}

			return true
		}
	})

	context.ParamResolvers[reflect.TypeOf([]string{})] = func(paramIndex int) interface{} {
		return func(ctx context.Context) []string {
			// When you want to retrieve a parameter with a value type that it is not supported by-default, such as ctx.Params().GetInt
			// then you can use the `GetEntry` or `GetEntryAt` and cast its underline `ValueRaw` to the desired type.
			// The type should be the same as the macro's evaluator function (last argument on the Macros#Register) return value.
			return ctx.Params().GetEntryAt(paramIndex).ValueRaw.([]string)
		}
	}

	app.Get("/api/test_slice_hero/{myparam:slice}", hero.Handler(func(myparam []string) string {
		return fmt.Sprintf("myparam's value (a trailing path parameter type) is: %#v\n", myparam)
	}))

	/*
		http://localhost:8080/test_slice_contains/notcontains1/value2 ->
		(404) Not Found
		http://localhost:8080/test_slice_contains/value1/value2 ->
		myparam's value (a trailing path parameter type) is: []string{"value1", "value2"}
	*/
	app.Get("/api/test_slice_contains/{myparam:slice contains([value1,value2])}", func(ctx context.Context) {
		// When it is not a built'n function available to retrieve your value with the type you want, such as ctx.Params().GetInt
		// then you can use the `GetEntry.ValueRaw` to get the real value, which is set-ed by your macro above.
		myparam := ctx.Params().GetEntry("myparam").ValueRaw.([]string)
		ctx.Writef("myparam's value (a trailing path parameter type) is: %#v\n", myparam)
	})
}
