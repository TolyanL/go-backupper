package job

type Job struct {
	ContainerName string `yaml:"container_name"`
	UseCommand    string `yaml:"use_command"`
	WorkDirPath   string `yaml:"work_dir"`
	MediaPath     string `yaml:"media_path"`
}
