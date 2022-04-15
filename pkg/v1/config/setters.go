package config

func (inst *Configuration) IsProduction(isProduction bool) *Configuration {
	inst.Production = isProduction
	return inst
}

func (inst *Configuration) SetPath(path string) *Configuration {
	inst.Path = path
	return inst
}
