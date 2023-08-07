package i18n

import (
	"testing"
)

const (
	pl = `{
    "card": {
        "title": "Tytuł",
        "notification": "Powiadomienie"
    }
}`
	en = `{
    "card": {
        "title": "Title",
        "notification": "Notification"
    }
}`
)

func TestLang(t *testing.T) {

	ser := NewService(nil)
	l := ser.GetLang("pl_PL")
	if l.name != "pl_PL" {
		t.Errorf("expecting 'pl_PL', got %s ", l.name)
		return
	}
	l = ser.GetLang("en_US")
	if l.name != "en_US" {
		t.Errorf("expecting 'en_US', got %s ", l.name)
		return
	}
}

func TestTrans(t *testing.T) {
	ser := NewService(nil)
	ser.addTranslation("pl_PL", []byte(pl))
	ser.addTranslation("en_gb", []byte(en))

	v := ser.GetTranslation(ser.GetLang("pl_PL"), "card.title")
	if v != "Tytuł" {
		t.Errorf("got: %s", v)
	}
	v = ser.GetTranslation(ser.GetLang("de_CH"), "card.title")
	if v != "Title" {
		t.Errorf("got: %s", v)
	}
}
