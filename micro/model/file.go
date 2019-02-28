package model

import "time"

type File struct {
	ID           uint   `gorm:"primary_key" json:"id"`
	FileName     string `gorm:"type:varchar(100);not null" json:"file_name"`
	OriginalName string `gorm:"type:varchar(100);not null" json:"original_name"`
	URL          string `json:"url"`
	Mime         string `json:"mime"`
	Size         uint   `json:"size"`
}

type FileUploadInfo struct {
	File
	UUID           string    `gorm:"type:varchar(100);unique;not null" json:"uuid"`
	UploadAt       time.Time `json:"upload_at"`
	UploadDir      string    `gorm:"type:varchar(100);not null" json:"upload_dir"`
	UploadFilePath string    `gorm:"type:varchar(100);not null" json:"upload_file_path"`
	Status         uint8     `gorm:"type:smallint;default:0" json:"status"`
}
