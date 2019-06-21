package model

type Config struct {
	User struct{
		Account string `yaml:"account"`
		Passwd string `yaml:"passwd"`
	}
}

func init()  {
	print(11)
}