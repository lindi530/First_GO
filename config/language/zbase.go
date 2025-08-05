package language

type Language struct {
	Cpp    LanguageConfig `mapstructure:"cpp"`
	Go     LanguageConfig `mapstructure:"Go"`
	Python LanguageConfig `mapstructure:"Python"`
}

func (languageConfig *Language) Compile(language string) *LanguageConfig {
	switch language {
	case "cpp":
		return &languageConfig.Cpp
	case "go":
		return &languageConfig.Go
	case "python":
		return &languageConfig.Python
	default:
		return &LanguageConfig{}
	}
	return &LanguageConfig{}
}
