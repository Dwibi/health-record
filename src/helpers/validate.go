package helpers

import (
	"errors"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

/*
	- first until third digit, should start with `303`
	- the fourth digit, if it's male, fill it with `1`, else `2`
	- the fifth and eigth digit, fill it with a year, starts from `2000` till current year
	- the ninth and tenth, fill it with month, starts from `01` till `12`
	- the eleventh and thirteenth, fill it with three random digit, starts from `000` till `999`
*/

func ValidateURLWithDomain(u string) error {
	parsedURL, err := url.ParseRequestURI(u)
	if err != nil {
		return errors.New("invalid URL format")
	}

	// Regular expression to match a valid domain
	re := regexp.MustCompile(`\.[a-z]{2,}$`)
	if !re.MatchString(parsedURL.Host) {
		return errors.New("URL must contain a valid domain")
	}

	return nil
}

func ValidateNIP(nip string) error {
	if len(nip) < 13 || len(nip) > 15 {
		return errors.New("nip must be 13 - 15 digits")
	}

	if nip[3] != '1' && nip[3] != '2' {
		return errors.New("the fourth digit, if it's male, fill it with `1`, else `2`")
	}

	currentYear := time.Now().Year()

	year, err := strconv.Atoi(nip[4:8])
	if err != nil || year < 2000 || year > currentYear {
		return errors.New("the fifth and eigth digit, must fill it with a year, starts from `2000` till current year")
	}

	// Check the month part (9th to 10th digit)
	month, err := strconv.Atoi(nip[8:10])
	if err != nil || month < 1 || month > 12 {
		return errors.New("the ninth and tenth, fill it with month, starts from `01` till `12`")
	}

	randomDigits := nip[10:]
	if matched, _ := regexp.MatchString(`^\d{3}\d{0,4}$`, randomDigits); !matched {
		return errors.New("the eleventh and last, fill it with random digit, starts from `000` till `99999`")
	}

	return nil
}

func IsItUser(nip string) bool {
	nipStr := nip

	return nipStr[:3] == "615"
}

func IsItNurse(nip string) bool {
	nipStr := nip

	return nipStr[:3] == "303"
}

func ValidateIdentityNum(identityNum int) bool {
	identityStr := strconv.Itoa(identityNum)
	return len(identityStr) == 16
}

func ValidateInaPhoneNum(phoneNum string) bool {
	// identityStr := strconv.Itoa(identityNum)
	return strings.HasPrefix(phoneNum, "+62")
}

func ValidateGender(gender string) bool {
	// identityStr := strconv.Itoa(identityNum)
	return strings.ToLower(gender) == "male" || strings.ToLower(gender) == "female"
}

func ValidateDateFormat(date string) bool {
	formats := []string{
		time.RFC3339,              // "2006-01-02T15:04:05Z07:00"
		"2006-01-02",              // "2006-01-02"
		"2006-01-02T15:04:05",     // "2006-01-02T15:04:05"
		"2006-01-02T15:04:05.999", // "2006-01-02T15:04:05.999"
	}

	for _, format := range formats {
		if _, err := time.Parse(format, date); err == nil {
			return true
		}
	}
	return false
}
