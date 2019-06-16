package _mysql

import "time"

type File struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	FileName    string `gorm:"type:varchar(100);not null" json:"file_name"`
	OrignalName string `gorm:"type:varchar(100);not null" json:"orignal_name"`
	URL         string `json:"url"`
	Mime        string `json:"mime"`
	Size        string `json:"size"`
}

type FileUploadInfo struct {
	File
	UploadAt       time.Time `json:"upload_at"`
	UploadDir      string    `gorm:"type:varchar(100);not null"`
	UploadFilePath string    `gorm:"type:varchar(100);not null"`
	UUIDName       string    `gorm:"type:varchar(100);unique;not null"`
	FileURL        string    `gorm:"type:varchar(100);unique;not null"`
	Status         uint8     `gorm:"type:tinyint(1) unsigned;default:0" json:"status"`
}
