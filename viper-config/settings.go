package main

import "time"

var Appf Settings

type Settings struct {
	Database
	AppSetting
}

type AppSetting struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Database struct {
	UserName     string
	Password     string
	Host         string
	Port         int
	DBName       string
	MaxIdleConns int
	MaxOpenConns int
}
