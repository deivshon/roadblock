package condition

type Condition struct {
	Id        string             `yaml:"id"`
	Selector  ConditionSelector  `yaml:"select"`
	Evaluator ConditionEvaluator `yaml:"evaluate"`
}

type ConditionSelector struct {
	Command   *bool `yaml:"command"`
	WordIndex *int  `yaml:"wordIndex"`
	AnyWord   *bool `yaml:"anyWord"`
	EveryWord *bool `yaml:"everyWord"`
}

type ConditionEvaluator struct {
	Equals   *string `yaml:"equals"`
	Contains *string `yaml:"contains"`
	Regex    *string `yaml:"regex"`
}
