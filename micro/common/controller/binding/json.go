// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"bytes"
	"encoding/json"
	"github.com/kataras/iris"
	"io"
)

// EnableDecoderUseNumber is used to call the UseNumber method on the JSON
// Decoder instance. UseNumber causes the Decoder to unmarshal a number into an
// interface{} as a Number instead of as a float64.
var EnableDecoderUseNumber = false

type jsonBinding struct{}

func (jsonBinding) Name() string {
	return "json"
}

func (jsonBinding) Bind(ctx iris.Context, obj interface{}) error {
	return decodeJSON(ctx.Request().Body, obj)
}

func (jsonBinding) BindBody(body []byte, obj interface{}) error {
	return decodeJSON(bytes.NewReader(body), obj)
}

func decodeJSON(r io.Reader, obj interface{}) error {
	decoder := json.NewDecoder(r)
	if EnableDecoderUseNumber {
		decoder.UseNumber()
	}
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return validate(obj)
}
