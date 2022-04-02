package Exception

type IException interface {
	GetCode() int
	GetMessage() string
}
