package gopools

import "sync"

type File struct {
	File interface{}
	Rw   sync.RWMutex  // 读写锁
}
