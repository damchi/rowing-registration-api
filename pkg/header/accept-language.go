package headers

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func GetAcceptLanguage(c *gin.Context) string {
	lang := c.GetHeader("Accept-Language")

	switch strings.ToLower(lang) {
	case "en", "fr":
		return strings.ToLower(lang)
	default:
		return "en"
	}
}
