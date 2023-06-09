package pty

type IPTY interface {
	Read(*[][]rune)
	Write([]byte) (int, error)
	Close()
}
