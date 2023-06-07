package upload

import (
	"fmt"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"hoper/initialize"
	"hoper/utils"
)

func GetImageFullUrl(name string) string {
	return "/" + GetImagePath() + name
}

func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = utils.EncodeMD5(fileName)

	return fileName + ext
}

func GetImagePath() string {
	return initialize.Config.Server.UploadPath
}

func GetImageFullPath() string {
	return initialize.Config.Server.RuntimeRootPath + GetImagePath()
}

func CheckImageExt(fileName string) bool {
	ext := utils.GetExt(fileName)
	for _, allowExt := range initialize.Config.Server.UploadAllowExt {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

func CheckImageSize(f multipart.File) bool {
	size := utils.GetSize(f)
	if size == 0 {
		return false
	}

	return size <= initialize.Config.Server.UploadMaxSize
}

func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = utils.IsNotExistMkdir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkdir err: %v", err)
	}

	perm := utils.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
