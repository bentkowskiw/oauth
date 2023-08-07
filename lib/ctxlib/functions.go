package ctxlib

import (
	"context"
	"errors"

	"github.com/oauth/i18n"
)

var (
	userKey contextKey = "user"
	langKey contextKey = "lang"
	tokKey  contextKey = "oauth"
)

type contextKey string

func GetUserId(ctx context.Context) (id *string, er error) {
	i := ctx.Value(userKey)
	if i == nil {
		er = errors.New("no user in context")
		return
	}
	idStr := i.(string)
	id = &idStr
	return
}

func SetUserId(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, userKey, id)
}
func SetLang(ctx context.Context, lang i18n.LangEnum) context.Context {
	return context.WithValue(ctx, langKey, lang)
}

func GetLang(ctx context.Context) i18n.LangEnum {
	if i := ctx.Value(userKey); i != nil {
		if lang, ok := i.(i18n.LangEnum); ok {
			return lang
		}
	}
	l := i18n.DefaultLanguage()
	return *l
}
