package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	DbHost             string
	DbPort             string
	DbUser             string
	DbPassword         string
	DbName             string
	Charset            string
	EtcdHost           string
	EtcdPort           string
	UserServiceAddress string
	TestServiceAddress string
)

func Init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("配置文件错误，请检查配置文件路径", err)
	}
	LoadMySqlData(file)
	LoadEtcd(file)
	LoadServer(file)
}
func LoadEtcd(file *ini.File) {
	EtcdHost = file.Section("etcd").Key("EtcdHost").String()
	EtcdPort = file.Section("etcd").Key("EtcdPort").String()

}
func LoadServer(file *ini.File) {
	UserServiceAddress = file.Section("server").Key("UserServiceAddress").String()
	TestServiceAddress = file.Section("server").Key("TestServiceAddress").String()

}
func LoadMySqlData(file *ini.File) {
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassword = file.Section("mysql").Key("DbPassword").String()
	DbName = file.Section("mysql").Key("DbName").String()
	Charset = file.Section("mysql").Key("Charset").String()
}
