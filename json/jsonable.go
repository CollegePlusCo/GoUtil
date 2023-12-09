package json

type Jsonable interface {
	ToJSON() []byte
}
