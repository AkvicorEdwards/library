package config

import "library/mysql"

type ServerConfig struct {
	Title		string
	Addr		string
	Password	string
	Protocol	string
	CDN			string
	Key			string
}

type CertConfig struct {
	Cert string
	Key string
}

type PathConfig struct {
	Theme string
	Cover string
}

var MySQL = &mysql.MySQL{
	User:      "root",
	Password:  "password",
	Host:      "localhost",
	DBName:    "book",
	Charset:   "utf8mb4",
	ParseTime: "true",
	Loc:       "Local",
}

var Server = &ServerConfig{
	Title:    "Library",
	Addr:     ":8080",
	Password: "pass",
	Protocol: "https",
	CDN:      "/",
	Key:      "12345678123456781234567812345678",
}

var Path = PathConfig{
	Theme:"./tpl/ori/",
	Cover:"./cover/",
}

var Cert = &CertConfig{
	Cert: "./cert/cert.pem",
	Key:  "./cert/key.pem",
}