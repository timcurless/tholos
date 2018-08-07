package model

import "github.com/jinzhu/gorm"

type AWSCredential struct {
	gorm.Model
	AccessKey    string   `json:"ACCESS_KEY_ID"`
	AccessSecret string   `json:"ACCESS_SECRET_KEY"`
	Regions      []string `json:"regions"`
}
