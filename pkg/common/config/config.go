package config

import (
	"encoding/json"
	"errors"
	"sync"
)

var ErrNotFound = errors.New("key not found") // ErrNotFound is key not found.

type Observer func(string, Value)

type Config interface {
	Load() error
	Scan(v any) error
	Value(key string) Value
}

func NewConfig(opts ...Option) Config {
	o := *NewOptions()
	for _, opt := range opts {
		opt(&o)
	}
	return &config{
		opts:   o,
		reader: NewReader(o),
	}
}

type config struct {
	opts   Options
	reader Reader
	cached sync.Map
}

func (c *config) Load() error {
	for _, src := range c.opts.sources {
		kvs, err := src.Load()
		if err != nil {
			return err
		}
		err = c.reader.Merge(kvs...)
		if err != nil {
			return err
		}
	}
	return nil
	// return c.reader.Resolve()
}

func (c *config) Scan(v any) error {
	data, err := c.reader.Source()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func (c *config) Value(key string) Value {
	if v, ok := c.cached.Load(key); ok {
		return v.(Value)
	}
	if v, ok := c.reader.Value(key); ok {
		c.cached.Store(key, v)
		return v
	}
	return &errValue{err: ErrNotFound}
}
