package main

type Intl interface {
	Func()
}

type Config struct {
	Field1 string
}

func (c *Config) Func() {

}

// Intl 修改后需要同步修改Config的实现不然不能编译
var _ Intl = (*Config)(nil)

func main() {

}
