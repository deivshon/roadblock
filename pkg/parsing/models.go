package parsing

type CommandParsingOpts struct {
	AddDefaultWrappers *bool    `yaml:"addDefaultWrappers"`
	RemoveWrappers     *bool    `yaml:"removeWrappers"`
	Wrappers           []string `yaml:"wrappers"`
}
