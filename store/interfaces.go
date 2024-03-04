package store

// Store is an interface for getting and setting key-value pairs
type Store interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}
