package ut

import (
	"path/filepath"
	"strings"
)

func GuessContentType(originalName string) string {
	var contentType string
	suffix := strings.ToLower(filepath.Ext(originalName))
	switch suffix {
	case ".txt":
		contentType = "text/plain"
	case ".gif":
		contentType = "image/gif"
	case ".png":
		contentType = "image/png"
	case ".jpeg", ".jpg":
		contentType = "image/jpeg"
	case ".pdf":
		contentType = "application/pdf"
	case ".doc", ".docx":
		contentType = "application/msword"
	case ".xls", ".xlsx":
		contentType = "application/vnd.ms-excel"
	case ".csv":
		contentType = "text/csv"
	case ".mp4":
		contentType = "video/mp4"
	case ".mp3":
		contentType = "audio/mpeg"
	case ".json":
		contentType = "application/json"
	case ".apk":
		contentType = "application/vnd.android.package-archive"
	default:
		contentType = "application/octet-stream"
	}

	return contentType
}
