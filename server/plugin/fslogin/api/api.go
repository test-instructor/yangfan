package api

import (
	"github.com/gin-gonic/gin"
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/model/common/response"
	"github.com/test-instructor/cheetah/server/model/system"
	systemReq "github.com/test-instructor/cheetah/server/model/system/request"
	"github.com/test-instructor/cheetah/server/plugin/fslogin/model"
	"github.com/test-instructor/cheetah/server/plugin/fslogin/service"
	"github.com/test-instructor/cheetah/server/utils"
	"go.uber.org/zap"
)

type FsLoginApi struct{}

// @Tags FsLogin
// @Summary 请手动填写接口功能
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /fsLogin/login [get]
func (p *FsLoginApi) Login(c *gin.Context) {
	code, _ := c.GetQuery("code")
	token, _ := c.GetQuery("state")
	j := utils.NewJWT()
	// parseToken 解析token包含的信息
	claims, err := j.ParseToken(token)
	FsUserInfo, err := fsLoginPassPort.FsLogin(code)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		c.Abort()
	}
	if claims != nil {
		// 绑定逻辑
		userId := claims.ID
		if err := service.ServiceGroupApp.BindFs(FsUserInfo, userId); err != nil {
			c.HTML(200, "fsError.html", model.LoginE{Err: err.Error()})
		} else {
			c.HTML(200, "fsBindSuccess.html", nil)
		}
	} else {
		// 登录逻辑
		if userInfo, err := service.ServiceGroupApp.LoginOrRegister(FsUserInfo); err != nil {
			global.GVA_LOG.Error("失败!", zap.Error(err))
			c.HTML(200, "fsError.html", model.LoginE{Err: err.Error()})
		} else {
			AuthO2Login(userInfo, c)
		}
	}

}

func AuthO2Login(user system.SysUser, c *gin.Context) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		c.HTML(200, "fsError.html", model.LoginE{Err: err.Error()})
	} else {
		u := model.LoginU{
			Test:      user.NickName,
			JWT:       token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}
		c.HTML(200, "fsLogin.html", u)
	}
}
