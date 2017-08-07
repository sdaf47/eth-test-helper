package helper

type settings struct {
	GenesisPath string
	ChainName string
	GethParams params
}

func SettingsConstructor() settings {
	s := settings{}
	s.GethParams = params{}

	return s
}
