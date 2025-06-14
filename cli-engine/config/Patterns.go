package config

type Patterns struct {
	Job    string `yaml:"job"`
	Exec   string `yaml:"exec"`
	DD     string `yaml:"dd"`
	Submit string `yaml:"submit"`
	ModTag string `yaml:"modtag"`
}
