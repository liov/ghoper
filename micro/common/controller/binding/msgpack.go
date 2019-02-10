// Copyright 2017 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"bytes"
	"github.com/kataras/iris"
	"github.com/ugorji/go/codec"
	"io"
)

type msgpackBinding struct{}

func (msgpackBinding) Name() string {
	return "msgpack"
}

func (msgpackBinding) Bind(ctx iris.Context, obj interface{}) error {
	return decodeMsgPack(ctx.Request().Body, obj)
}

func (msgpackBinding) BindBody(body []byte, obj interface{}) error {
	return decodeMsgPack(bytes.NewReader(body), obj)
}

func decodeMsgPack(r io.Reader, obj interface{}) error {
	cdc := new(codec.MsgpackHandle)
	if err := codec.NewDecoder(r, cdc).Decode(&obj); err != nil {
		return err
	}
	return validate(obj)
}
