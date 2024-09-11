package storage

type Storage interface {
	Store(key string, data []byte) error
	Retrieve(key string) ([]byte, error)
	Delete(key string) error
}

type LocalStorage struct {
	// 实现本地存储的字段
}

// 实现Storage接口的方法