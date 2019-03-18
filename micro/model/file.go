package model

import "time"

type File struct {
	ID           uint   `gorm:"primary_key" json:"id"`
	FileName     string `gorm:"type:varchar(100);not null" json:"file_name"`
	OriginalName string `gorm:"type:varchar(100);not null" json:"original_name"`
	URL          string `json:"url"`
	MD5          string `gorm:"type:varchar(32)" json:"md5"`
	Mime         string `json:"mime"`
	Size         uint   `json:"size"`
}

type FileUploadInfo struct {
	File
	UUID           string    `gorm:"type:varchar(100);unique;not null" json:"uuid"`
	UploadUser     User      `json:"upload_user"`
	UploadUserID   uint      `json:"upload_user_id"`
	UploadAt       time.Time `json:"upload_at"`
	UploadFilePath string    `gorm:"type:varchar(100);not null" json:"upload_file_path"`
	Status         uint8     `gorm:"type:smallint;default:0" json:"status"`
}
