package apierror

import (
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"rowing-registation-api/pkg/logger"
	"rowing-registation-api/pkg/translator"
)

type ApiError interface {
	GetError() JsonErr
}

type Err struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type JsonErr struct {
	Error Err `json:"error"`
}

func (e JsonErr) GetError() JsonErr {
	return e
}

func (e Err) Error() string {
	return e.Message
}

func CreateError(code int, msgKey string, lang string) JsonErr {
	message := GetTransErrorMsg(msgKey, lang)
	return JsonErr{Err{code, message}}
}

func GetTransErrorMsg(msgKey string, lang string) string {
	l := i18n.NewLocalizer(translator.Translator, lang, "en")
	m := i18n.LocalizeConfig{}
	m.MessageID = msgKey
	msg, err := l.Localize(&m)
	if err != nil {
		logger.Log(logger.WARNING, fmt.Sprintf("Translation not found : %s", err.Error()))
		return msgKey
	}
	return msg
}
