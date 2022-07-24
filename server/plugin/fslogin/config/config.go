package config

type FsLogin struct {
	AppID       string `mapstructure:"appID" json:"appID" yaml:"appID"`                   // 飞书应用APPID
	AppSecret   string `mapstructure:"appSecret" json:"from" yaml:"appSecret"`            // 飞书应用AppSecret
	RedirectUri string `mapstructure:"redirectUri" json:"redirectUri" yaml:"redirectUri"` // 回调链接
	AuthorityID uint   `mapstructure:"authorityID" json:"authorityID" yaml:"authorityID"` // 角色ID
	ProjectID   uint   `mapstructure:"projectID" json:"projectID" yaml:"projectID"`       // 项目ID
}
