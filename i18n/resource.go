package i18n

import (
	"io/fs"
	"github.com/oauth/lib/errlib"
)

type Service struct {
	languages    map[string]LangEnum
	translations map[LangEnum]Translation
}

func NewService(trans fs.FS) *Service {
	s := &Service{
		languages:    make(map[string]LangEnum),
		translations: make(map[LangEnum]Translation),
	}
	s.addLang(LangEnum{
		name:   "pl_PL",
		parent: DefaultLanguage(),
	})
	errlib.PanicOnErr(s.loadTranslations(trans))
	return s
}
