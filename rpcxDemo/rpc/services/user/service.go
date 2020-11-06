package user

import (
	"context"
	"fmt"
)

// Args 包含了请求的数据
type Args struct {
	UserID int
}

// Reply 回复
type Reply struct {
	U User
}

type User struct {
	Name string
	Age  int
}

func (t *User) GetUser(ctx context.Context, args *Args, reply *Reply) error {
	fmt.Printf("find user by userID: %d\n", args.UserID)
	reply.U = User{
		Name: "test8972",
		Age:  12,
	}
	return nil
}
