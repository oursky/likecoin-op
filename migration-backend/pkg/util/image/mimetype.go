package image

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/gabriel-vasile/mimetype"
)

type SupportedMimeType string

const (
	WEBP SupportedMimeType = "image/webp"
	GIF  SupportedMimeType = "image/gif"
)

var (
	ErrMimeTypeNotSupported = errors.New("error mime type not supported")
)

type supportedMimeTypeMapType map[string]SupportedMimeType

func (m supportedMimeTypeMapType) Get(str string) (*SupportedMimeType, error) {
	s, ok := m[str]
	if !ok {
		return nil, ErrMimeTypeNotSupported
	}
	return &s, nil
}

var supportedMimeTypeMap = supportedMimeTypeMapType{
	string(WEBP): WEBP,
	string(GIF):  GIF,
}

type defaultMimeTypeConversion map[SupportedMimeType]SupportedMimeType

func (m defaultMimeTypeConversion) GetDefaultConvertedPath(srcPath string) (string, error) {
	srcMimeType, err := mimetype.DetectFile(srcPath)
	if err != nil {
		return "", err
	}
	supportedType, err := supportedMimeTypeMap.Get(srcMimeType.String())
	if err != nil {
		return "", err
	}
	defaultConvertedType, ok := m[*supportedType]
	if !ok {
		return "", ErrMimeTypeNotSupported
	}
	defaultConvertedMimeType := mimetype.Lookup(string(defaultConvertedType))
	fromExt := srcMimeType.Extension()
	toExt := defaultConvertedMimeType.Extension()
	regex := regexp.MustCompile(fmt.Sprintf("(%s)?$", fromExt))
	return regex.ReplaceAllString(srcPath, toExt), nil
}

var DefaultMimeTypeConversion = defaultMimeTypeConversion{
	WEBP: GIF,
}
