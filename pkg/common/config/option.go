package config

import (
	"fmt"

	"dario.cat/mergo"
	"github.com/adnpa/fern/pkg/encoding"
	_ "github.com/adnpa/fern/pkg/encoding/json"
)

// Decoder Kv->map
type Decoder func(*KV, map[string]any) error

// Resolver resolve placeholder in config.
type Resolver func(map[string]any) error

// Merge is config merge func.
type Merge func(dst, src any) error

// NewOptions return default Options
func NewOptions() *Options {
	return &Options{
		decoder: defaultDecoder,
		merge:   defaultMerge,
	}
}

type Options struct {
	sources []Source
	decoder Decoder
	merge   Merge
}

// Option change val of Options
type Option func(*Options)

// WithSource with config source.
func WithSource(s ...Source) Option {
	return func(o *Options) {
		o.sources = s
	}
}

// WithDecoder with config decoder.
func WithDecoder(d Decoder) Option {
	return func(o *Options) {
		o.decoder = d
	}
}

// WithMergeFunc with config merge func.
func WithMergeFunc(m Merge) Option {
	return func(o *Options) {
		o.merge = m
	}
}

// 支持默认解码器
func defaultDecoder(cfg *KV, target map[string]any) error {
	fmt.Println("=====", cfg.Format)
	if codec := encoding.GetCodec(cfg.Format); codec != nil {
		return codec.Unmarshal(cfg.Value, &target)
	}
	return fmt.Errorf("unsupported key: %s format: %s", cfg.Key, cfg.Format)
}

// 合并不同的map，同名覆盖
func defaultMerge(dst, src any) error {
	return mergo.Map(dst, src, mergo.WithOverride)
}
