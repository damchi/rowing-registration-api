package validations

import (
	"context"
	"regexp"
	"rowing-registation-api/api/models"
	"rowing-registation-api/api/services"
	"rowing-registation-api/pkg/constants"
	"rowing-registation-api/pkg/translator"
	"strings"
)

const (
	OCCUPATION_CONSUMER = "consumer"
	OCCUPATION_OWNER    = "owner"
	OCCUPATION_MANAGER  = "manager"
	OCCUPATION_EMPLOYEE = "employee"
)

type CreateClubError struct {
	HasError       bool   `json:"-"`
	FirstName      string `json:"first_name,omitempty"`
	LastName       string `json:"last_name,omitempty"`
	Email          string `json:"email,omitempty"`
	Password       string `json:"password,omitempty"`
	ClubName       string `json:"club_name,omitempty"`
	Address1       string `json:"address1,omitempty"`
	Address2       string `json:"address2,omitempty"`
	Address3       string `json:"address3,omitempty"`
	AddressCity    string `json:"address_city,omitempty"`
	AddressCountry string `json:"address_country,omitempty"`
	AddressState   string `json:"address_state,omitempty"`
	PostCode       string `json:"post_code,omitempty"`
	Phone          string `json:"phone,omitempty"`
}

func CreateClubValidation(ctx context.Context, c models.ClubRegistrationParam, lang string) CreateClubError {
	err := CreateClubError{}

	if len(strings.Trim(c.FirstName, " ")) == 0 {
		err.HasError = true
		err.FirstName = translator.Trans("fieldIsMissing", lang, map[string]interface{}{"Field": "first_name"})
	} else {
		regxName := regexp.MustCompile(constants.RegexName)
		if !regxName.MatchString(c.FirstName) {
			err.HasError = true
			err.FirstName = translator.Trans("fieldIsNotValid", lang, map[string]interface{}{"Field": "first_name"})
		}
	}

	if len(strings.Trim(c.LastName, " ")) == 0 {
		err.HasError = true
		err.LastName = translator.Trans("fieldIsMissing", lang, map[string]interface{}{"Field": "last_name"})
	} else {
		regxName := regexp.MustCompile(constants.RegexName)
		if !regxName.MatchString(c.LastName) {
			err.HasError = true
			err.LastName = translator.Trans("fieldIsNotValid", lang, map[string]interface{}{"Field": "last_name"})
		}
	}

	if len(strings.Trim(c.ClubName, " ")) == 0 {
		err.HasError = true
		err.ClubName = translator.Trans("fieldIsMissing", lang, map[string]interface{}{"Field": "club_name"})
	} else {
		regxName := regexp.MustCompile(constants.RegexName)
		if !regxName.MatchString(c.LastName) {
			err.HasError = true
			err.ClubName = translator.Trans("fieldIsNotValid", lang, map[string]interface{}{"Field": "club_name"})
		}
	}

	if len(strings.Trim(c.Address1, " ")) == 0 {
		err.HasError = true
		err.Address1 = translator.Trans("fieldIsMissing", lang, map[string]interface{}{"Field": "address1"})
	} else {
		regxExclusion := regexp.MustCompile(constants.RegexExclusion)
		if !regxExclusion.MatchString(c.Address1) {
			err.HasError = true
			err.Address1 = translator.Trans("fieldIsNotValid", lang, map[string]interface{}{"Field": "address1"})
		}
	}

	if len(strings.Trim(c.Address2, " ")) > 0 {
		regxExclusion := regexp.MustCompile(constants.RegexExclusion)
		if !regxExclusion.MatchString(c.Address2) {
			err.HasError = true
			err.Address2 = translator.Trans("fieldIsNotValid", lang, map[string]interface{}{"Field": "address2"})
		}
	}
	if len(strings.Trim(c.Address3, " ")) > 0 {
		regxExclusion := regexp.MustCompile(constants.RegexExclusion)
		if !regxExclusion.MatchString(c.Address3) {
			err.HasError = true
			err.Address3 = translator.Trans("fieldIsNotValid", lang, map[string]interface{}{"Field": "address3"})
		}
	}

	if len(strings.Trim(c.AddressCity, " ")) == 0 {
		err.HasError = true
		err.AddressCity = translator.Trans("fieldIsMissing", lang, map[string]interface{}{"Field": "address_city"})
	} else {
		regxExclusion := regexp.MustCompile(constants.RegexExclusion)
		if !regxExclusion.MatchString(c.AddressCity) {
			err.HasError = true
			err.AddressCity = translator.Trans("fieldIsNotValid", lang, map[string]interface{}{"Field": "address_city"})
		}
	}

	if len(strings.Trim(c.AddressCountry, " ")) == 0 {
		err.HasError = true
		err.AddressCountry = translator.Trans("fieldIsMissing", lang, map[string]interface{}{"Field": "address_country"})
	} else {
		regxExclusion := regexp.MustCompile(constants.RegexExclusion)
		if !regxExclusion.MatchString(c.AddressCountry) {
			err.HasError = true
			err.AddressCountry = translator.Trans("fieldIsNotValid", lang, map[string]interface{}{"Field": "address_country"})
		}
	}
	if len(strings.Trim(c.AddressState, " ")) == 0 {
		err.HasError = true
		err.AddressState = translator.Trans("fieldIsMissing", lang, map[string]interface{}{"Field": "address_state"})
	} else {
		regxExclusion := regexp.MustCompile(constants.RegexExclusion)
		if !regxExclusion.MatchString(c.AddressState) {
			err.HasError = true
			err.AddressState = translator.Trans("fieldIsNotValid", lang, map[string]interface{}{"Field": "address_state"})
		}
	}

	if len(c.Email) > 0 {
		Re := regexp.MustCompile(constants.RegexEmail)
		if !Re.MatchString(c.Email) {
			err.HasError = true
			err.Email = translator.Trans("emailNotValid", lang, nil)
		} else {

			hasEmail := services.GetUserService(ctx, nil).FindByEmail(c.Email)
			if hasEmail {
				err.HasError = true
				err.Email = translator.Trans("emailAlreadyUsed", lang, nil)
			}

		}
	} else {
		err.HasError = true
		err.Email = translator.Trans("fieldIsMissing", lang, map[string]interface{}{"Field": "email"})
	}

	if len(strings.Trim(c.Phone, " ")) == 0 {
		err.HasError = true
		err.Phone = translator.Trans("fieldIsMissing", lang, map[string]interface{}{"Field": "phone"})
	} else {
		regxExclusion := regexp.MustCompile(constants.RegexExclusion)
		if !regxExclusion.MatchString(c.Phone) {
			err.HasError = true
			err.Phone = translator.Trans("fieldIsNotValid", lang, map[string]interface{}{"Field": "phone"})
		}
	}

	if len(strings.Trim(c.PostCode, " ")) == 0 {
		err.HasError = true
		err.PostCode = translator.Trans("fieldIsMissing", lang, map[string]interface{}{"Field": "post_code"})
	} else {
		regxExclusion := regexp.MustCompile(constants.RegexExclusion)
		if !regxExclusion.MatchString(c.PostCode) {
			err.HasError = true
			err.PostCode = translator.Trans("fieldIsNotValid", lang, map[string]interface{}{"Field": "post_code"})
		}
	}

	if !validatePassword(c.Password) {
		err.HasError = true
		err.Password = translator.Trans("weakPassword", lang, map[string]interface{}{"Field": "password"})
	}

	return err
}

func validatePassword(password string) bool {
	if len(password) >= 8 && len(password) <= 20 {
		hasSpecial := false
		hasLower := false
		hasUpper := false
		hasDigits := false
		hasErrorCharacter := false

		var chars []rune

		for i := 0; i < len(password); i++ {
			chars = append(chars, rune(password[i]))
		}

		for _, c := range chars {
			switch {
			case strings.ContainsRune(constants.SpecialChars, c):
				hasSpecial = true
			case strings.ContainsRune(constants.LowerChars, c):
				hasLower = true
			case strings.ContainsRune(constants.UpperChars, c):
				hasUpper = true
			case strings.ContainsRune(constants.DigitChars, c):
				hasDigits = true
			default:
				hasErrorCharacter = true
			}
		}
		if hasSpecial && hasLower && hasUpper && hasDigits && !hasErrorCharacter {
			return true
		}
	}
	return false
}
