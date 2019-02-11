// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"github.com/golang/protobuf/proto"
	"github.com/kataras/iris"

	"io/ioutil"
)

type protobufBinding struct{}

func (protobufBinding) Name() string {
	return "protobuf"
}

func (b protobufBinding) Bind(ctx iris.Context, obj interface{}) error {
	buf, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return err
	}
	return b.BindBody(buf, obj)
}

func (protobufBinding) BindBody(body []byte, obj interface{}) error {
	if err := proto.Unmarshal(body, obj.(proto.Message)); err != nil {
		return err
	}
	// Here it's same to return validate(obj), but util now we cann't add
	// `binding:""` to the struct which automatically generate by gen-proto
	return nil
	// return validate(obj)
}