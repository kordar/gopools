package gopools

import (
	"sync"
)

// SingleUnsafePool 基于map寻址的连接池独立句柄，句柄需保证线程安全
type SingleUnsafePool struct {
	Files    []*File
	SyncOnce []sync.Once
}

func NewSingleUnsafePool(bucketSize int) *SingleUnsafePool {
	return &SingleUnsafePool{
		Files:    make([]*File, bucketSize),
		SyncOnce: make([]sync.Once, bucketSize),
	}
}

func (pool SingleUnsafePool) GetHandler(number uint32, fun func(number uint32) (*File, error)) (*File, error) {
	var err error
	pool.SyncOnce[number].Do(func() {
		pool.Files[number], err = fun(number)
	})
	return pool.Files[number], err
}
