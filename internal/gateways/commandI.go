package gateways

type command interface {
	getCommand()
}

func Pega2() string {
	return "Hello, World!"
}
