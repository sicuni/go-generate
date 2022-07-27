package config

type InitConfig func(*Config)

type Config struct {
	ModelPath   string // 生成model代码输出路径
	ModelPkg    string // 生成model层包名
	DaoPath     string // 生成dao代码输出路径
	DaoPkg      string // 生成dao层包名
	ServicePath string // 生成service代码输出路径
	HandlePath  string // 生成handle代码输出路径
}

// New 初始化Config信息
func New(opts ...InitConfig) *Config {
	defaultConfig := &Config{
		ModelPath:   "./",
		DaoPath:     "./",
		ServicePath: "./",
		HandlePath:  "./",
		ModelPkg:    "model",
		DaoPkg:      "model",
	}
	for _, opt := range opts {
		opt(defaultConfig)
	}
	return defaultConfig
}

func WithModelPath(modelPath string) InitConfig {
	return func(config *Config) {
		config.ModelPath = modelPath
	}
}

func WithModelPkg(modelPkg string) InitConfig {
	return func(config *Config) {
		config.ModelPkg = modelPkg
	}
}

func WithDaoPath(daoPath string) InitConfig {
	return func(config *Config) {
		config.DaoPath = daoPath
	}
}

func WithDaoPkg(daoPkg string) InitConfig {
	return func(config *Config) {
		config.DaoPkg = daoPkg
	}
}

func WithServicePath(servicePath string) InitConfig {
	return func(config *Config) {
		config.ServicePath = servicePath
	}
}

func WithHandlePath(handlePath string) InitConfig {
	return func(config *Config) {
		config.HandlePath = handlePath
	}
}
