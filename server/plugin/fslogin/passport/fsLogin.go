package passport

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/test-instructor/yangfan/server/plugin/fslogin/global"
	"github.com/test-instructor/yangfan/server/plugin/fslogin/model"
	"io"
	"net/http"
)

type FsLoginPassport struct {
}

func (f *FsLoginPassport) FsLogin(code string) (userInfo model.FsUserInfo, err error) {
	info := make(map[string]string)
	info["grant_type"] = "authorization_code"
	info["client_id"] = global.GlobalConfig.AppID
	info["client_secret"] = global.GlobalConfig.AppSecret
	info["code"] = code
	info["redirect_uri"] = global.GlobalConfig.RedirectUri

	bytesData, _ := json.Marshal(info)

	url := "https://passport.feishu.cn/suite/passport/oauth/token"
	reader := bytes.NewReader(bytesData)
	req, err := http.NewRequest("POST", url, reader)
	defer req.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	reqsBytes, err := io.ReadAll(resp.Body)
	var acReq model.AccessReq
	json.Unmarshal(reqsBytes, &acReq)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return f.GetUserInfo(acReq.AccessToken)
}

func (f *FsLoginPassport) GetUserInfo(accessToken string) (userInfo model.FsUserInfo, err error) {
	url := "https://passport.feishu.cn/suite/passport/oauth/userinfo"
	req, err := http.NewRequest("GET", url, nil)
	client := http.Client{}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	reqsBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	json.Unmarshal(reqsBytes, &userInfo)
	return userInfo, nil
}
