package model

import "time"

//系统配置
type SysInfoModel struct {
	Name string //系统名称
}

//用户相关
type UserModel struct {
	UserName    string
	PWD         string
	Token       string
	Head        string
	Email       string
	Description string
	Initialized bool
	Avatar      string
	NickName    string
	Public      string
}

//服务配置
type ServerModel struct {
	HttpPort     string
	RunMode      string
	ServerApi    string
	LockAccount  bool
	Handshake    string
	Token        string
	UDPPort      string
	USBAutoMount string
}

//服务配置
type APPModel struct {
	LogSavePath    string
	LogSaveName    string
	LogFileExt     string
	DateStrFormat  string
	DateTimeFormat string
	TimeFormat     string
	DateFormat     string
	ProjectPath    string
	RootPath       string
}

//公共返回模型
type Result struct {
	Success int         `json:"success" example:"200"`
	Message string      `json:"message" example:"ok"`
	Data    interface{} `json:"data" example:"返回结果"`
}

//redis配置文件
type RedisModel struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

type SystemConfig struct {
	ConfigStr  string `json:"config_str"`
	WidgetList string `json:"widget_list"`
	ConfigPath string `json:"config_path"`
	SyncPort   string `json:"sync_port"`
	SyncKey    string `json:"sync_key"`
	Analyse    string `json:"analyse"`
}

type CasaOSGlobalVariables struct {
	AppChange bool
}

type FileSetting struct {
	ShareDir    []string `json:"share_dir" delim:"|"`
	DownloadDir string   `json:"download_dir"`
}
