package config

// KV is a map of namespace and content
type KV struct {
	Key    string
	Value  []byte
	Format string
}

// Source is the source of config, eg:file,env,remote...
type Source interface {
	Load() ([]*KV, error)
}
