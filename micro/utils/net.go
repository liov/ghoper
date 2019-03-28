package utils

import (
	"github.com/kataras/golog"
	"net/url"
)

// RelativeURLToAbsoluteURL 相对URL转绝对URL
func RelativeURLToAbsoluteURL(curURL string, baseURL string) (string, error) {
	curURLData, err := url.Parse(curURL)
	if err != nil {
		golog.Error(err)
		return "", err
	}
	baseURLData, err := url.Parse(baseURL)
	if err != nil {
		golog.Error(err)
		return "", err
	}
	curURLData = baseURLData.ResolveReference(curURLData)
	return curURLData.String(), nil
}
