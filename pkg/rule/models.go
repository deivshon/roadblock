package rule

type Rule struct {
	AllOf       *[]Rule `yaml:"allOf"`
	OneOf       *[]Rule `yaml:"oneOf"`
	Not         *Rule   `yaml:"not"`
	ConditionId *string `yaml:"conditionId"`
	Message     *string `yaml:"message"`
}
