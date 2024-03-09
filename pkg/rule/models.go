package rule

type RuleWithMeta struct {
	Name string `yaml:"name"`
	Rule Rule   `yaml:"rule"`
}

type Rule struct {
	AllOf       *[]Rule `yaml:"allOf"`
	OneOf       *[]Rule `yaml:"oneOf"`
	Not         *Rule   `yaml:"not"`
	ConditionId *string `yaml:"conditionId"`
}
