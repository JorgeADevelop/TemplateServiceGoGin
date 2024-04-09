package utils

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func getLocalizer(languaje string) *i18n.Localizer {
	bundle := i18n.NewBundle(language.English)

	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustLoadMessageFile("translations/en.json")
	bundle.MustLoadMessageFile("translations/es.json")

	return i18n.NewLocalizer(bundle, languaje)
}

func getMessage(messageID string, meta map[string]interface{}, languaje string) string {
	localizer := getLocalizer(languaje)
	return localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: meta,
	})
}

func getLanguaje(ctx *gin.Context) string {
	languaje := ctx.GetHeader("Accept-Language")
	if languaje == "" {
		return "en"
	} else if languaje != "en" && languaje != "es" {
		return "en"
	}
	return languaje
}
