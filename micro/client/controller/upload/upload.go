package upload

import (
	"errors"
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
	"time"
)

func GetExt(file *multipart.FileHeader) (string, error) {
	var ext string
	var index = strings.LastIndex(file.Filename, ".")
	if index != -1 {
		return "", nil
	} else {
		ext = file.Filename[index:]
	}
	if len(ext) == 1 {
		return "", errors.New("无效的扩展名")
	}
	return ext, nil
}

func GetDirAndUrl(classify string, info *multipart.FileHeader) (string, string, error) {
	//sep := string(os.PathSeparator)
	sep := "/"
	var uploadDir, prefixUrl string
	ymdStr := utils.GetTodayYMD(sep)
	ext, err := GetExt(info)
	if err != nil {
		return "", "", err
	}

	if ext == "" {
		uploadDir = strings.Join([]string{initialize.Config.Server.UploadDir,
			"others",
			classify,
			ymdStr},
			"/")
		prefixUrl = strings.Join([]string{
			"https://" + initialize.Config.Server.UploadHost + initialize.Config.Server.UploadPath,
			"others",
			classify,
			ymdStr,
		}, "/")
		return uploadDir, prefixUrl, nil
	}

	var mimeType = mime.TypeByExtension(ext)
	if mimeType == "" && ext == ".jpeg" {
		mimeType = "image/jpeg"
	}

	dirType := strings.Split(mimeType, "/")

	uploadDir = strings.Join([]string{initialize.Config.Server.UploadDir,
		dirType[0] + "s",
		classify,
		ymdStr},
		"/")

	/*	length := utf8.RuneCountInString(uploadDir)
		lastChar := uploadDir[length-1:]

		if lastChar != sep {
			uploadDir = uploadDir + sep + ymdStr
		} else {
			uploadDir = uploadDir + ymdStr
		}
	*/

	prefixUrl = strings.Join([]string{
		"https://" + initialize.Config.Server.UploadHost + initialize.Config.Server.UploadPath,
		dirType[0] + "s",
		classify,
		ymdStr,
	}, "/")

	if err := os.MkdirAll(uploadDir, 0777); err != nil {
		return uploadDir, prefixUrl, err
	}
	return uploadDir, prefixUrl, nil
}

// Upload 文件上传
func Upload(ctx iris.Context) *model.FileUploadInfo {
	classify := ctx.Params().GetString("classify")
	file, info, err := ctx.FormFile("file")

	if err != nil {
		common.Response(ctx, "参数无效")
		return nil
	}
	defer file.Close()

	ext, err := GetExt(info)
	if ext == "" || err != nil {
		common.Response(ctx, "无效的图片类型")
		return nil
	}

	dir, url, err := GetDirAndUrl(classify, info)

	upInfo, err := SaveUploadedFile(info, dir, url)
	if err != nil {
		common.Response(ctx, err.Error())
		return nil
	}

	upInfo.File.Size = (uint)(utils.GetSize(file))

	if err := initialize.DB.Create(upInfo).Error; err != nil {
		common.Response(ctx, err.Error())
		return nil
	}
	common.Response(ctx, upInfo)
	return upInfo
}

func UploadMultiple(ctx iris.Context) {
	classify := ctx.Params().GetString("classify")
	//获取通过iris.WithPostMaxMemory获取的最大上传值大小。
	maxSize := ctx.Application().ConfigurationReadOnly().GetPostMaxMemory()
	err := ctx.Request().ParseMultipartForm(maxSize)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
		return
	}

	form := ctx.Request().MultipartForm
	files := form.File["files[]"]
	failures := 0
	dir, url, err := GetDirAndUrl(classify, files[0])
	for _, file := range files {
		upInfo, err := SaveUploadedFile(file, dir, url)
		if err != nil {
			failures++
			common.Response(ctx, file.Filename+"上传失败")
		} else {
			if err := initialize.DB.Create(&upInfo).Error; err != nil {
				common.Response(ctx, err.Error())
			}
		}
	}
	common.Response(ctx, len(files)-failures)
}

func SaveUploadedFile(file *multipart.FileHeader, dir string, url string) (*model.FileUploadInfo, error) {
	uuidName := uuid.NewV4().String()
	ext, err := GetExt(file)
	filename := uuidName + ext
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	out, err := os.Create(dir + filename)
	if err != nil {
		return nil, err
	}
	defer out.Close()

	fileUpload := model.FileUploadInfo{
		File: model.File{
			FileName:     filename,
			OriginalName: file.Filename,
			URL:          url + filename,
			Mime:         mime.TypeByExtension(ext),
		},
		UUID:           uuidName,
		UploadFilePath: dir + filename,
		UploadAt:       time.Now(),
	}
	io.Copy(out, src)
	return &fileUpload, nil
}
