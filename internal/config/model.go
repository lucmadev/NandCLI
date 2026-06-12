package config

type Config struct {
	Theme         string `yaml:"theme"`
	PreferredShell string `yaml:"preferredShell"`

	System SystemConfig `yaml:"system"`
	TUI    TUIConfig    `yaml:"tui"`

	Modules ModulesConfig `yaml:"modules"`
}

type SystemConfig struct {
	AutoUpdate bool `yaml:"autoUpdate"`
}

type TUIConfig struct {
	ShowLogo   bool `yaml:"showLogo"`
	Animations bool `yaml:"animations"`
	RefreshRate int `yaml:"refreshRate"`
}

type ModulesConfig struct {
	APT       ModuleConfig `yaml:"apt"`
	Docker    DockerConfig `yaml:"docker"`
	Git       ModuleConfig `yaml:"git"`
	Shells    ModuleConfig `yaml:"shells"`
	Stats     StatsConfig `yaml:"stats"`
	Systemctl ModuleConfig `yaml:"systemctl"`
}

type ModuleConfig struct {
	Enabled bool `yaml:"enabled"`
}

type DockerConfig struct {
	Enabled bool `yaml:"enabled"`
	Compose bool `yaml:"compose"`
}

type StatsConfig struct {
	Enabled bool `yaml:"enabled"`
	GpuType string `yaml:"gpuType"`
}