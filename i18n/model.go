package i18n

var defaultLanguage = &LangEnum{
	name: "en_US",
}

func DefaultLanguage() *LangEnum {
	return defaultLanguage
}

type LangEnum struct {
	name   string
	parent *LangEnum
}

type Translation struct {
	index int
	value map[string]interface{}
}
