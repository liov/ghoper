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

var doc spec.Swagger

func GetDoc(stop bool,args ...string) {
	if doc.Swagger == "" {
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

		if utils.CheckNotExist(realPath) {
			generate()
		} else {
			file, err := os.Open(realPath)
			if err != nil {
				ulog.Error(err)
			}
			defer file.Close()
			data, err := ioutil.ReadAll(file)
			/*var buf bytes.Buffer
			err = json.Compact(&buf, data)
			if err != nil {
				ulog.Error(err)
			}*/
			if apiType == "json" {
				err = json.Unmarshal(data, &doc)
			} else {
				var v map[string]interface{}
				err = yaml.Unmarshal(data, &v)
				b, err := json.Marshal(data)
				if err != nil {
					ulog.Error(err)
				}
				json.Unmarshal(b, &doc)
			}
			if err != nil {
				ulog.Error(err)
			}
		}
	}
	if stop {
		defer WriteToFile(args...)
	}
}

func generate() {

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
	doc.Schemes = []string{"http", "https"}
	doc.Consumes = []string{"application/json"}
	doc.Produces = []string{"application/json"}
}

func WriteToFile(args ...string) {
	if doc.Swagger == "" {
		generate()
	}
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
	var file *os.File
	file, err = os.Create(realPath)
	if err != nil {
		ulog.Error(err)
	}
	defer file.Close()

	if apiType == "json" {
		enc := json.NewEncoder(file)
		enc.SetIndent("", "  ")
		enc.Encode(&doc)
	} else {
		b, err := yaml.Marshal(swag.ToDynamicJSON(&doc))
		if err != nil {
			ulog.Error(err)
		}
		if _, err := file.Write(b); err != nil {
			ulog.Error(err)
		}
	}
}

func main() {
	GetDoc(false,".", "yml")
}
