package user

import (
	"context"
	"fmt"
	"log"
	"main/config"
	"main/services/auth"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/masterZSH/rservices/user"

	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/share"
)

func GetUser(c *gin.Context) {
	userIDString := c.Param("id")
	userID, _ := strconv.Atoi(userIDString)
	option := config.GetDefaultOptions()
	cfg := config.DefaultConfig()
	cg := config.AuthConfig()
	d := client.NewEtcdV3Discovery(cfg.BasePath, "User", []string{cfg.EtcdAddr}, cg)
	// 路由模式-随机选择
	xclient := client.NewXClient("User", client.Failtry, client.RandomSelect, d, option)
	xclient.Auth("")
	defer xclient.Close()
	args := &user.Args{
		UserID: userID,
	}

	reply := &user.Reply{}
	// 同步 异步使用Go
	ctx := context.WithValue(context.Background(), share.ReqMetaDataKey, make(map[string]string))
	err := xclient.Call(ctx, "GetUser", args, reply)
	// 广播模式 调用节点
	// err := xclient.Broadcast(context.Background(), "GetUser", args, reply)
	// token
	if err != nil && err.Error() == "invalid token" {
		fmt.Println("---------")
		// 获取token
		tk, err := auth.GetToken()
		if err != nil {
		}
		xclient.Auth(tk)
		err = xclient.Call(ctx, "GetUser", args, reply)
	}
	if err != nil {
		// debug.PrintStack()
		fmt.Printf("failed to call: %v", err)
	}
	log.Printf("get id : %d user %+v\n", args.UserID, reply.U)

	c.JSON(200, reply)
}
