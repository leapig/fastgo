package utils

import (
	"encoding/base64"
	"errors"
	"github.com/leapig/fastgo/app/dal/model"
	"io"
	"net/http"
	"strings"
	"time"
)

// PutFile 上传文件通用接口
func PutFile(req *http.Request, paramName string) (object model.Oss, err error) {
	if object, err = PutFileWithFile(req, paramName); err == nil {
		return
	}
	if object, err = PutFileWithBase64(req, paramName); err == nil {
		return
	}
	if object, err = PutFileWithUrl(req, paramName); err == nil {
		return
	}
	object, err = PutFileWithBinary(req)
	return
}

// PutFileWithFile 上传表单文件
func PutFileWithFile(req *http.Request, paramName string) (object model.Oss, err error) {
	if file, header, err := req.FormFile(paramName); err == nil {
		if fileData, err := io.ReadAll(file); err == nil {
			object.Name = header.Filename
			nameArr := strings.Split(header.Filename, ".")
			pointer := len(nameArr) - 1
			suffix := nameArr[pointer]
			contentType := header.Header["Content-Type"][0]
			// 常见格式校验
			switch contentType {
			case "audio/aac": //音频
				if suffix != "aac" {
					return object, errors.New("file format error")
				}
				break
			case "audio/mp3": //音频
				if suffix != "mp3" {
					return object, errors.New("file format error")
				}
				break
			case "video/mp4": //视频
				if suffix != "mp4" {
					return object, errors.New("file format error")
				}
				break
			case "video/mpeg": //视频
				if suffix != "mpeg" {
					return object, errors.New("file format error")
				}
				break
			case "application/msword": // 微软word
				if suffix != "doc" {
					return object, errors.New("file format error")
				}
				break
			case "application/vnd.openxmlformats-officedocument.wordprocessingml.document": // 通用word
				if suffix != "doc" {
					return object, errors.New("file format error")
				}
				break
			case "application/vnd.ms-powerpoint": // 微软ppt
				if suffix != "ppt" {
					return object, errors.New("file format error")
				}
				break
			case "application/vnd.openxmlformats-officedocument.presentationml.presentation": // 通用ppt
				if suffix != "pptx" {
					return object, errors.New("file format error")
				}
				break
			case "application/vnd.ms-excel": // 微软Excel
				if suffix != "xls" {
					return object, errors.New("file format error")
				}
				break
			case "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet": // 通用Excel
				if suffix != "xlsx" {
					return object, errors.New("file format error")
				}
				break
			case "application/pdf": // 便携式文档
				if suffix != "pdf" {
					return object, errors.New("file format error")
				}
				break
			case "image/png":
				if suffix != "png" {
					return object, errors.New("file format error")
				}
				break
			case "image/jpeg":
				if suffix != "jpg" && suffix != "jpeg" {
					return object, errors.New("file format error")
				}
				break
			default:
				return object, errors.New("unsupported file format")
			}
			object.Suffix = suffix
			object.Extension = contentType
			object.Size = int64(len(fileData))
			object.Data = fileData
		}
	}
	return
}

// PutFileWithBinary 上传二进制文件
func PutFileWithBinary(req *http.Request) (object model.Oss, err error) {
	contentType := req.Header.Get("Content-Type")
	// 常见格式校验
	switch contentType {
	case "audio/aac": //音频
		object.Suffix = "aac"
		object.Extension = contentType
		break
	case "audio/mp3": //音频
		object.Suffix = "mp3"
		object.Extension = contentType
		break
	case "video/mp4": //视频
		object.Suffix = "mp4"
		object.Extension = contentType
		break
	case "video/mpeg": //视频
		object.Suffix = "mpeg"
		object.Extension = contentType
		break
	case "application/msword": // 微软word
		object.Suffix = "doc"
		object.Extension = contentType
		break
	case "application/vnd.openxmlformats-officedocument.wordprocessingml.document": // 通用word
		object.Suffix = "docx"
		object.Extension = contentType
		break
	case "application/vnd.ms-powerpoint": // 微软ppt
		object.Suffix = "ppt"
		object.Extension = contentType
		break
	case "application/vnd.openxmlformats-officedocument.presentationml.presentation": // 通用ppt
		object.Suffix = "pptx"
		object.Extension = contentType
		break
	case "application/vnd.ms-excel": // 微软Excel
		object.Suffix = "xls"
		object.Extension = contentType
		break
	case "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet": // 通用Excel
		object.Suffix = "xlsx"
		object.Extension = contentType
		break
	case "application/pdf": // 便携式文档
		object.Suffix = "pdf"
		object.Extension = contentType
		break
	case "image/png":
		object.Suffix = "png"
		object.Extension = contentType
		break
	case "image/jpeg":
		object.Suffix = "jpg"
		object.Extension = contentType
		break
	default:
		err = errors.New("unsupported file format")
		return
	}
	object.Data, err = io.ReadAll(req.Body)
	object.Size = int64(len(object.Data))
	object.Name = time.Now().Format("20060102150405") + "." + object.Suffix
	return
}

// PutFileWithBase64 上传base64文件
func PutFileWithBase64(req *http.Request, paramName string) (object model.Oss, err error) {
	if stream := req.PostFormValue(paramName); strings.HasPrefix(stream, "data:") && strings.Contains(stream, "base64,") {
		streamBak := stream
		name := strings.Split(streamBak, ";base64,")[0]
		contentType := name[5:]
		// 常见格式校验
		switch contentType {
		case "audio/aac": //音频
			object.Suffix = "aac"
			object.Extension = contentType
			break
		case "audio/mp3": //音频
			object.Suffix = "mp3"
			object.Extension = contentType
			break
		case "video/mp4": //视频
			object.Suffix = "mp4"
			object.Extension = contentType
			break
		case "video/mpeg": //视频
			object.Suffix = "mpeg"
			object.Extension = contentType
			break
		case "application/msword": // 微软word
			object.Suffix = "doc"
			object.Extension = contentType
			break
		case "application/vnd.openxmlformats-officedocument.wordprocessingml.document": // 通用word
			object.Suffix = "docx"
			object.Extension = contentType
			break
		case "application/vnd.ms-powerpoint": // 微软ppt
			object.Suffix = "ppt"
			object.Extension = contentType
			break
		case "application/vnd.openxmlformats-officedocument.presentationml.presentation": // 通用ppt
			object.Suffix = "pptx"
			object.Extension = contentType
			break
		case "application/vnd.ms-excel": // 微软Excel
			object.Suffix = "xls"
			object.Extension = contentType
			break
		case "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet": // 通用Excel
			object.Suffix = "xlsx"
			object.Extension = contentType
			break
		case "application/pdf": // 便携式文档
			object.Suffix = "pdf"
			object.Extension = contentType
			break
		case "image/png":
			object.Suffix = "png"
			object.Extension = contentType
			break
		case "image/jpeg":
			object.Suffix = "jpg"
			object.Extension = contentType
			break
		default:
			err = errors.New("unsupported file format")
			return
		}
		if object.Data, err = base64.StdEncoding.DecodeString(stream); err == nil {
			object.Size = int64(len(object.Data))
		}
		object.Name = time.Now().Format("20060102150405") + "." + object.Suffix
	} else {
		err = errors.New("unsupported file format")
	}
	return
}

// PutFileWithUrl 上传远程文件
func PutFileWithUrl(req *http.Request, paramName string) (object model.Oss, err error) {
	if url := req.PostFormValue(paramName); strings.HasPrefix(url, "http") {
		if res, err := http.Get(strings.TrimSpace(url)); err == nil {
			contentType := res.Header["Content-Type"][0]
			// 常见格式校验
			switch contentType {
			case "audio/aac": //音频
				object.Suffix = "aac"
				object.Extension = contentType
				break
			case "audio/mp3": //音频
				object.Suffix = "mp3"
				object.Extension = contentType
				break
			case "video/mp4": //视频
				object.Suffix = "mp4"
				object.Extension = contentType
				break
			case "video/mpeg": //视频
				object.Suffix = "mpeg"
				object.Extension = contentType
				break
			case "application/msword": // 微软word
				object.Suffix = "doc"
				object.Extension = contentType
				break
			case "application/vnd.openxmlformats-officedocument.wordprocessingml.document": // 通用word
				object.Suffix = "docx"
				object.Extension = contentType
				break
			case "application/vnd.ms-powerpoint": // 微软ppt
				object.Suffix = "ppt"
				object.Extension = contentType
				break
			case "application/vnd.openxmlformats-officedocument.presentationml.presentation": // 通用ppt
				object.Suffix = "pptx"
				object.Extension = contentType
				break
			case "application/vnd.ms-excel": // 微软Excel
				object.Suffix = "xls"
				object.Extension = contentType
				break
			case "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet": // 通用Excel
				object.Suffix = "xlsx"
				object.Extension = contentType
				break
			case "application/pdf": // 便携式文档
				object.Suffix = "pdf"
				object.Extension = contentType
				break
			case "image/png":
				object.Suffix = "png"
				object.Extension = contentType
				break
			case "image/jpeg":
				object.Suffix = "jpg"
				object.Extension = contentType
				break
			default:
				return object, errors.New("unsupported file format")
			}
			object.Data, _ = io.ReadAll(res.Body)
			object.Size = int64(len(object.Data))
			object.Name = time.Now().Format("20060102150405") + "." + object.Suffix
		}
	}
	return
}
