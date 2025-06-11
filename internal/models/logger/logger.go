package logger

type Logger struct {
	Enabled   bool   `yaml:"enabled"`
	OutputDir string `yaml:"dir"`
	Format    string `yaml:"format"`
	Level     string `yaml:"level"`
}
