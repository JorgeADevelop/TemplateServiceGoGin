package utils

import (
	"encoding/json"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var Localizer *i18n.Localizer

func InitTranslate() {
	bundle := i18n.NewBundle(language.English)

	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustLoadMessageFile("translations/en.json")

	loc := i18n.NewLocalizer(bundle, language.English.String())

	Localizer = loc
}
