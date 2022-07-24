package model

import model "github.com/test-instructor/cheetah/server/model/system"

type AccessReq struct {
	AccessToken      string `json:"access_token"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
}

type FsUserInfo struct {
	GvaUserId    uint   `json:"gva_user_id"`
	Sub          string `json:"sub"`
	Name         string `json:"name"`
	Picture      string `json:"picture"`
	OpenId       string `json:"open_id"`
	UnionId      string `json:"union_id"`
	EnName       string `json:"en_name"`
	TenantKey    string `json:"tenant_key"`
	AvatarUrl    string `json:"avatar_url"`
	AvatarThumb  string `json:"avatar_thumb"`
	AvatarMiddle string `json:"avatar_middle"`
	AvatarBig    string `json:"avatar_big"`
	Email        string `json:"email"`
	UserId       string `json:"user_id"`
	Mobile       string `json:"mobile"`
}

func (FsUserInfo) TableName() string {
	return "fs_user_info"
}

type LoginUserInfo struct {
	FsUserInfo
	model.SysUser
}

type LoginU struct {
	Test      string
	JWT       string
	ExpiresAt int64
}

type LoginE struct {
	Err string
}
