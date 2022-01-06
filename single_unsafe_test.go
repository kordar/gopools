package gopools

import (
	"testing"
)

func TestNewSingleUnsafePool(t *testing.T) {
	/*pool := NewSingleUnsafePool(100)
	dir := "data"
	handler, err := pool.GetHandler(1, func(number uint32) (*File, error) {
		if file, err := leveldb.OpenFile(dir, nil); err != nil {
			// 是否corrupt异常，是重新创建
			if errors.IsCorrupted(err) {
				if file, err = leveldb.RecoverFile(dir, nil); err == nil {
					return &File{File: file, Rw: sync.RWMutex{}}, nil
				}
			}
			return nil, fmt.Errorf("failed to open leveldb at %s: %s", dir, err)
		} else {
			return &File{File: file, Rw: sync.RWMutex{}}, nil
		}
	})
	log.Println(err)
	db := handler.File.(*leveldb.DB)
	_ = db.Put([]byte("AAA"), []byte("DDDDDD"), nil)
	get, _ := db.Get([]byte("AAA"), nil)
	log.Println(string(get))*/
}
