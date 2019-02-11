package upload

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path"
	"service/controller/common/logging"
	"service/initialize"
	"service/utils"
	"strings"
)

func GetImageFullUrl(name string) string {
	return initialize.Config.Server.ImagePrefixUrl + "/" + GetImagePath() + name
}

func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = utils.EncodeMD5(fileName)

	return fileName + ext
}

func GetImagePath() string {
	return initialize.Config.Server.ImagePath
}

func GetImageFullPath() string {
	return initialize.Config.Server.RuntimeRootPath + GetImagePath()
}

func CheckImageExt(fileName string) bool {
	ext := utils.GetExt(fileName)
	for _, allowExt := range initialize.Config.Server.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

func CheckImageSize(f multipart.File) bool {
	size, err := utils.GetSize(f)
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}

	return size <= initialize.Config.Server.ImageMaxSize
}

func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = utils.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := utils.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
