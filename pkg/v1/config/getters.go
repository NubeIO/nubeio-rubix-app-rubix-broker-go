package config

func (inst *Configuration) getPath() (path string) {
	return inst.Path
}

func (inst *Configuration) GetConfig() *Configuration {
	return inst
}
