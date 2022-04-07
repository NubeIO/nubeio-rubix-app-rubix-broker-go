package config

const (
	DevPath        = "data/rubix-broker/config/config.json"
	ProductionPath = "data/rubix-broker/config/config.json"
)

//IsProduction if set to true then use path for production /data
func (inst *Configuration) IsProduction(isProduction bool) *Configuration {
	inst.Production = isProduction
	return inst
}

//SetPath if you want to override the dev or production path then pass in a new path
func (inst *Configuration) SetPath(path string) *Configuration {
	if path == "" {
		if inst.Production { //if path is nil and in production then use this path
			path = ProductionPath
		} else { // else use dev path
			path = DevPath
		} // or use user path
	}
	inst.Path = path
	return inst
}
