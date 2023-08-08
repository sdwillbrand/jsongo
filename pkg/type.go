package pkg

type JSType int64

const (
	String JSType = iota + 1
	Number
	Boolean
	Null
	Array
	Object
)
