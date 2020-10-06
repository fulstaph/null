package null

type Nullable interface {
	Null() bool
	Nullify()
	String() string
}
