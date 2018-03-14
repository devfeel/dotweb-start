package config

import (
	"encoding/xml"
)

//配置信息
type AppConfig struct {
	XMLName    xml.Name     `xml:"config"`
	AppSets    []*AppSet    `xml:"appsettings>add"`
	Redises    []*RedisInfo `xml:"redises>redis"`
	Databases    []*DataBaseInfo `xml:"databases>database"`
	AllowIps   []string     `xml:"allowips>ip"`
	ConfigPath string
}

//AppSetting配置
type AppSet struct {
	Key   string `xml:"key,attr"`
	Value string `xml:"value,attr"`
}

//Redis信息
type RedisInfo struct {
	ID        string `xml:"id,attr"`
	ServerUrl string `xml:"serverurl,attr"`
}

//DataBase信息
type DataBaseInfo struct {
	ID        string `xml:"id,attr"`
	ServerUrl string `xml:"serverurl,attr"`
}

