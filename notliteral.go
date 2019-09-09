package strlit

type NotLiteral interface {
	error
	NotLiteral()
}
