package helpers

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
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

func ValidateNIP(nip string, role string) error {
	if len(nip) != 13 {
		return errors.New("nip must be 13 digits")
	}

	if role == "it" && nip[:3] != "615" {
		return errors.New("nip should start with 615")
	}

	if role == "nurse" && nip[:3] != "303" {
		return errors.New("nip should start with 303")
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
		return errors.New("the fifth and eigth digit, must fill it with a year, starts from `2000` till current year")
	}

	randomDigits := nip[10:13]
	if matched, _ := regexp.MatchString(`^\d{3}$`, randomDigits); !matched {
		return errors.New("the eleventh and thirteenth, must fill it with three random digit, starts from `000` till `999`")
	}

	return nil
}

func IsItUser(nip int) bool {
	nipStr := strconv.Itoa(nip)
	fmt.Println(nipStr[:3] == "615")

	return nipStr[:3] == "615"
}
