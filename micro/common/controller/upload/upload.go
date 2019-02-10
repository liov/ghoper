package upload

import (
	"errors"
	"fmt"
	"github.com/kataras/iris"
	"github.com/satori/go.uuid"
	"io"
	"micro/common/controller/common"
	"micro/common/initialize"
	"micro/common/model"
	"micro/common/utils"
	"mime"
	"mime/multipart"
	"os"
	"strings"
	"unicode/utf8"
)

// GenerateImgUploadedInfo 创建一个ImageUploadedInfo
func GenerateImgUploadedInfo(ext string) model.FileUploadInfo {

	sep := string(os.PathSeparator)
	uploadImgDir := initialize.Config.Server.UploadImgDir
	length := utf8.RuneCountInString(uploadImgDir)
	lastChar := uploadImgDir[length-1:]
	ymStr := utils.GetTodayYM(sep)

	var uploadDir string
	if lastChar != sep {
		uploadDir = uploadImgDir + sep + ymStr
	} else {
		uploadDir = uploadImgDir + ymStr
	}

	uuidName := uuid.NewV4().String()
	filename := uuidName + ext
	uploadFilePath := uploadDir + sep + filename
	fileURL := strings.Join([]string{
		"https://" + initialize.Config.Server.ImgHost + initialize.Config.Server.ImagePath,
		ymStr,
		filename,
	}, "/")
	var fileUpload model.FileUploadInfo

	fileUpload.FileName = filename
	fileUpload.FileURL = fileURL
	fileUpload.UUIDName = uuidName
	fileUpload.UploadDir = uploadDir
	fileUpload.UploadFilePath = uploadFilePath

	/*	fileUpload = model.FileUploadInfo{
		File:       model.File{FileName:filename,},
		FileURL:        fileURL,
		UUIDName:       uuidName,
		UploadDir:      uploadDir,
		UploadFilePath: uploadFilePath,
	}*/
	return fileUpload
}

// Upload 文件上传
func Upload(ctx iris.Context) (map[string]interface{}, error) {
	_, info, err := ctx.FormFile("upFile")

	if err != nil {
		return nil, errors.New("参数无效")
	}

	var filename = info.Filename
	var index = strings.LastIndex(filename, ".")

	if index < 0 {
		return nil, errors.New("无效的文件名")
	}

	var ext = filename[index:]
	if len(ext) == 1 {
		return nil, errors.New("无效的扩展名")
	}
	var mimeType = mime.TypeByExtension(ext)

	fmt.Printf("filename %s, index %d, ext %s, mimeType %s\n", filename, index, ext, mimeType)
	if mimeType == "" && ext == ".jpeg" {
		mimeType = "image/jpeg"
	}
	if mimeType == "" {
		return nil, errors.New("无效的图片类型")
	}

	imgUploadedInfo := GenerateImgUploadedInfo(ext)

	fmt.Println(imgUploadedInfo.UploadDir)

	if err := os.MkdirAll(imgUploadedInfo.UploadDir, 0777); err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("error")
	}

	if err := SaveUploadedFile(info, imgUploadedInfo.UploadFilePath); err != nil {
		return nil, errors.New("error1")
	}

	image := model.File{
		FileName:    imgUploadedInfo.FileName,
		OrignalName: filename,
		URL:         imgUploadedInfo.FileURL,
		Mime:        mimeType,
	}

	if err := initialize.DB.Create(&image).Error; err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("image error")
	}

	return map[string]interface{}{
		"id":       image.ID,
		"url":      imgUploadedInfo.FileURL,
		"title":    imgUploadedInfo.FileName, //新文件名
		"original": filename,                 //原始文件名
		"type":     mimeType,                 //文件类型
	}, nil
}

// UploadHandler 文件上传
func UploadHandler(ctx iris.Context) {
	data, err := Upload(ctx)
	if err != nil {
		common.Response(ctx, err, 500)
		return
	}
	common.Response(ctx, data, 200)
}

func SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	io.Copy(out, src)
	return nil
}
