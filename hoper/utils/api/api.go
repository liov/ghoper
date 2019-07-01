package main

import (
	"encoding/json"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/swag"
	"gopkg.in/yaml.v2"
	"hoper/utils"
	"hoper/utils/ulog"
	"io/ioutil"
	"os"
	"path/filepath"
)

/*func ApiMiddle(ctx iris.Context) {
	currentRouteName:= ctx.GetCurrentRoute().Name()
	params:=ctx.Params().Store
	for i:= range params{
		key:=params[i].Key
		val:=params[i].ValueRaw
	}
}*/

func Generate(args ...string) {
	targetPath := "."
	if len(args) > 0 {
		targetPath = args[0]
	}
	realPath, err := filepath.Abs(targetPath)
	if err != nil {
		ulog.Error(err)
	}

	apiType := "json"
	if len(args) > 1 {
		apiType = args[1]
	}

	realPath = filepath.Join(realPath, "swagger."+apiType)

	if utils.CheckExist(realPath) {
		os.Remove(realPath)

	}
	_, err = os.Create(realPath)
	if err != nil {
		ulog.Error(err)
	}
	var b []byte
	var doc spec.Swagger

	info := new(spec.Info)
	doc.Info = info

	doc.Swagger = "2.0"
	doc.Paths = new(spec.Paths)
	doc.Definitions = make(spec.Definitions)

	info.Title = "Title"
	info.Description = "Description"
	info.Version = "0.01"
	info.TermsOfService = "TermsOfService"

	var contact spec.ContactInfo
	contact.Name = "Contact Name"
	contact.Email = "Contact Email"
	contact.URL = "Contact URL"
	info.Contact = &contact

	var license spec.License
	license.Name = "License Name"
	license.URL = "License URL"
	info.License = &license

	doc.Host = "localhost:80"
	doc.BasePath = "/"
	doc.Schemes = []string{"http","https"}
	doc.Consumes = []string{"application/json"}
	doc.Produces = []string{"application/json"}

	if apiType == "json" {
		b, err = json.MarshalIndent(doc, "", "  ")
	} else {
		// marshals as YAML
		b, err = json.Marshal(doc)
		if err == nil {
			d, ery := swag.BytesToYAMLDoc(b)
			if ery != nil {
				ulog.Error(ery)
			}
			b, err = yaml.Marshal(d)
		}
	}
	err = ioutil.WriteFile(realPath, b, 0644)

}

func main() {
	Generate("../config")
}
