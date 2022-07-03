package global

import "blog-service/pkg/setting"

// Associate configuration information with applications

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
)