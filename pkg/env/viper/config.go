package envViper

var Env *envViper

type envViper struct {
	App app
}

type app struct {
	Name string `mapstructure:"name"`
}
