package model

type Cache interface {
	Ping() error
	Set(key string, value string) error
	Get(key string) (string, error)
}
