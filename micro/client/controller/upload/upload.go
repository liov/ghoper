package upload

import (
	"github.com/kataras/iris"
	"github.com/satori/go.uuid"
	"hoper/client/controller/common"
	"hoper/initialize"
	"hoper/model"
	"hoper/utils"
	"io"
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

	fileUpload := model.FileUploadInfo{
		File:           model.File{FileName: filename},
		FileURL:        fileURL,
		UUIDName:       uuidName,
		UploadDir:      uploadDir,
		UploadFilePath: uploadFilePath,
	}
	return fileUpload
}

// Upload 文件上传
func Upload(ctx iris.Context) *model.FileUploadInfo {
	file, info, err := ctx.FormFile("file")

	if err != nil {
		common.Response(ctx, "参数无效")
		return nil
	}
	defer file.Close()
	var filename = info.Filename
	var index = strings.LastIndex(filename, ".")

	if index < 0 {
		common.Response(ctx, "无效的文件名")
		return nil
	}

	var ext = filename[index:]
	if len(ext) == 1 {
		common.Response(ctx, "无效的扩展名")
		return nil
	}
	var mimeType = mime.TypeByExtension(ext)

	if mimeType == "" && ext == ".jpeg" {
		mimeType = "f/jpeg"
	}
	if mimeType == "" {
		common.Response(ctx, "无效的图片类型")
		return nil
	}
	uploadedInfo := GenerateImgUploadedInfo(ext)

	if err := os.MkdirAll(uploadedInfo.UploadDir, 0777); err != nil {
		common.Response(ctx, err.Error())
		return nil
	}

	if err := SaveUploadedFile(info, uploadedInfo.UploadFilePath); err != nil {
		common.Response(ctx, err.Error())
		return nil
	}

	f := model.File{
		FileName:    uploadedInfo.FileName,
		OrignalName: filename,
		URL:         uploadedInfo.FileURL,
		Mime:        mimeType,
	}

	uploadedInfo.File = f

	if err := initialize.DB.Create(&uploadedInfo).Error; err != nil {

		common.Response(ctx, err.Error())
		return nil
	}

	return &uploadedInfo
}

// UploadHandler 文件上传
func UploadHandler(c iris.Context) {
	data := Upload(c)
	if data == nil {
		common.Response(c, 500, nil)
		return
	}
	common.Response(c, 200, data)
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
