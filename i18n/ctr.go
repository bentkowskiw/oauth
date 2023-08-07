package i18n

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

func (s *Service) GetLang(lang string) LangEnum {
	l, ok := s.languages[strings.ToLower(lang)]
	if ok {
		return l
	}
	split := strings.Split(lang, "_")
	if len(split) == 2 {
		return s.GetLang(split[0])
	}
	return *DefaultLanguage()
}

func (s *Service) GetTranslation(lang LangEnum, key string) string {
	split := strings.Split(key, ".")
	trns, ok := s.translations[lang]
	if !ok {
		return key
	}
	tr := trns.Get(split)
	if tr != nil {
		return *tr
	}
	return fmt.Sprintf("%v.%s", lang, key)
}

func (t *Translation) Get(keys []string) *string {
	if t.index < len(keys) {
		if v, ok := t.value[keys[t.index]]; ok {
			if str, ok := v.(string); ok {
				return &str
			}
			if m, ok := v.(map[string]interface{}); ok {
				t.value = m
				t.index++
				return t.Get(keys)
			}
		}
	}
	return nil
}

func (s *Service) addLang(lang LangEnum) {
	s.languages[strings.ToLower(lang.name)] = lang
	parent := lang.parent
	if parent != nil {
		name := strings.ToLower(parent.name)
		if _, ok := s.languages[name]; !ok {
			s.languages[name] = *parent
		}
	}
}

func (s *Service) loadTranslations(fsys fs.FS) error {
	if fsys == nil {
		return nil
	}

	fs.WalkDir(fsys, "resources/translation", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		f, err := fsys.Open(path)
		if err != nil {
			return err
		}
		b, err := io.ReadAll(f)
		if err != nil {
			return err
		}
		info, err := f.Stat()
		if err != nil {
			return err
		}
		name := info.Name()
		len := len(name)
		name = name[:len-5]
		return s.addTranslation(name, b)
	})
	return nil
}

func (s *Service) addTranslation(name string, b []byte) error {
	data := make(map[string]interface{})
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}
	lang := s.GetLang(name)
	tr := Translation{
		value: data,
	}
	s.translations[lang] = tr
	return nil
}
