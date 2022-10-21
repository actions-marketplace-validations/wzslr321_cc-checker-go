package config

type Config struct {
	ExcludePatterns []string
	MaxComplexity   int
	WorkDir         string
}

func GetConfig() *Config {
	return &Config{
		ExcludePatterns: []string{"vendor"},
		MaxComplexity:   10,
		WorkDir:         ".",
	}
}
