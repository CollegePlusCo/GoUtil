package serializable

type Serializable interface {
	Serialize() []byte
}
