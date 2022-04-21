package Model

type IModel interface {
	Set(field string, s string)
	Get(field string) string
}
