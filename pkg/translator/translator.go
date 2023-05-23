package translator

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"io/ioutil"
	"log"
)

var Translator *i18n.Bundle

type Config struct {
	TranslationFolder string
}

const (
	LANGUAGE_EN = "en"
)

func InitTranslator(cfg Config) {
	Translator = i18n.NewBundle(language.English)
	Translator.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	var err error

	lstFiles, err := ioutil.ReadDir(cfg.TranslationFolder)

	if err != nil {
		log.Print(fmt.Sprintf("Failed to list the contain of the translation folder : %v", err))
	}

	for _, f := range lstFiles {
		filepath := cfg.TranslationFolder + "/" + f.Name()

		_, err = Translator.LoadMessageFile(filepath)

		if err != nil {
			log.Print(fmt.Sprintf("Failed to init translator with translation file %v", f.Name()))
		}
	}
}

func Trans(key string, lang string, transData map[string]interface{}) string {
	l := i18n.NewLocalizer(Translator, lang, LANGUAGE_EN)
	lc := i18n.LocalizeConfig{}
	lc.MessageID = key

	if transData != nil {
		lc.TemplateData = transData
	}

	msg, err := l.Localize(&lc)

	if err != nil {
		log.Print(fmt.Sprintf("translation key not found : %v", err))
		return key
	}

	return msg
}
