package redis

type option func(*options)

type options struct {
}

func defaultConfig() *options {
	return &options{}
}

func NewOptions(opts ...option) *options {
	options := defaultConfig()
	for _, opt := range opts {
		opt(options)
	}
	return options
}
