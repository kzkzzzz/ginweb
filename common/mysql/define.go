package mysql

// OpenClose 开启关闭
type OpenClose = uint8

const (
	OPEN  OpenClose = 1
	CLOSE OpenClose = 2
)
