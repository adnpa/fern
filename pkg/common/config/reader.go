package config

import (
	"encoding/json"
	"sync"

	"github.com/adnpa/fern/pkg/common/log"
	"github.com/jinzhu/copier"
)

// Reader convert row config from data to value
type Reader interface {
	// Merge different config
	Merge(...*KV) error
	// Get Single Value By Key
	Value(string) (Value, bool)
	// Get merged config raw content
	Source() ([]byte, error)
	//
	Resolve() error
}

func NewReader(opts Options) Reader {
	return &reader{
		opts:   opts,
		values: make(map[string]any),
		lock:   sync.Mutex{},
	}
}

type reader struct {
	opts   Options
	values map[string]any
	lock   sync.Mutex
}

func (r *reader) Merge(kvs ...*KV) error {
	// map并且需要遍历复杂操作 尽量减小临界区，所以深复制一份alue
	var merged map[string]any
	err := copier.Copy(&merged, r.values)
	if err != nil {
		return err
	}

	for _, kv := range kvs {
		next := make(map[string]any)
		if err := r.opts.decoder(kv, next); err != nil {
			log.Errorf("Failed to config decode error: %v key: %s value: %s", err, kv.Key, string(kv.Value))
			return err
		}
		if err := r.opts.merge(&merged, next); err != nil {
			log.Errorf("Failed to config merge error: %v key: %s value: %s", err, kv.Key, string(kv.Value))
			return err
		}
	}

	r.lock.Lock()
	r.values = merged
	r.lock.Unlock()
	return nil
}

func (r *reader) Value(k string) (Value, bool) {
	r.lock.Lock()
	defer r.lock.Unlock()
	v, ok := r.values[k]
	av := &atomicValue{}
	av.Store(v)
	return av, ok
}

func (r *reader) Source() ([]byte, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	return json.Marshal(r.values)
}

func (r *reader) Resolve() error {
	panic("not implemented") // TODO: Implement
}
