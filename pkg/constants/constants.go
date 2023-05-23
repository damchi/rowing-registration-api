package constants

const (
	RegexBearer    = "^(Bearer ).*$"
	RegexBasic     = "^(Basic ).*$"
	RegexEmail     = `^[a-zA-Z0-9.!#$%&\'*+=?^_{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)+$`
	SpecialChars   = `!@#$%^&*(){}[]`
	LowerChars     = `abcdefghijklmnopqrstuvwxyz`
	UpperChars     = `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	DigitChars     = `0123456789`
	RegexName      = `[a-zA-ZàáâäãåąčćęèéêëėįìíîïłńòóôöõøùúûüųūÿýżźñçčšžÀÁÂÄÃÅĄĆČĖĘÈÉÊËÌÍÎÏĮŁŃÒÓÔÖÕØÙÚÛÜŲŪŸÝŻŹÑßÇŒÆČŠŽ∂ð ,.'-]+$`
	RegexExclusion = `^[^<>]+$`
)
