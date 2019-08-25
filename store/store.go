package store

type Closer interface {
	Close() error
}
