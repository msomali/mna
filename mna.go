package mna

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const (
	Tigo               = "Tigo"
	Vodacom            = "Vodacom"
	TTCL               = "TTCL"
	Airtel             = "Airtel"
	Zantel             = "Zantel"
	SmileCommonName    = "Smile"
	MoCommonName       = "Mo Mobile"
	Halotel            = "Halotel"
	MkulimaCommonName  = "Mkulima"
	WiAfricaCommonName = "Wiafrica"
	statusOperational  = "Operational"
)

var (
	ErrOperatorNotFound = errors.New("mobile operator not found")
	ErrInvalidFormat    = errors.New("invalid format, correct formats are +255[9-digits], 255[9-digits], 0[9-digits] or just last 9 digits")
	ErrNumericOnly      = errors.New("phone numbers should contains numeric characters only")

	tigoPrefixes     = []string{"071", "065", "067"}
	vodaPrefixes     = []string{"074", "075", "076"}
	ttclPrefixes     = []string{"073"}
	zantelPrefixes   = []string{"077"}
	airtelPrefixes   = []string{"078", "068", "069"}
	smilePrefixes    = []string{"066"}
	viettelPrefixes  = []string{"061", "062"}
	mkulimaPrefixes  = []string{"063"}
	wiAfricaPrefixes = []string{"064"}
	moPrefixes       = []string{"072"}

	repository = []Data{

		{
			OperatorName: "MIC Tanzania PLC",
			Status:       statusOperational,
			Prefixes:     tigoPrefixes,
			CommonName:   Tigo,
		},
		{
			OperatorName: "Vodacom Tanzania PLC",
			CommonName:   Vodacom,
			Status:       statusOperational,
			Prefixes:     vodaPrefixes,
		},
		{
			OperatorName: "Tanzania Telecommunications Corporation",
			CommonName:   TTCL,
			Status:       statusOperational,
			Prefixes:     ttclPrefixes,
		},
		{
			OperatorName: "Zanzibar Telecom PLC",
			CommonName:   Zantel,
			Status:       statusOperational,
			Prefixes:     zantelPrefixes,
		},
		{
			OperatorName: "Airtel Tanzania PLC",
			CommonName:   Airtel,
			Status:       statusOperational,
			Prefixes:     airtelPrefixes,
		},
		{
			OperatorName: "Smile Communications Tanzania Limited",
			CommonName:   SmileCommonName,
			Status:       statusOperational,
			Prefixes:     smilePrefixes,
		},
		{
			OperatorName: "Viettel Tanzania PLC",
			CommonName:   Halotel,
			Status:       statusOperational,
			Prefixes:     viettelPrefixes,
		},
		{
			OperatorName: "Mkulima African Telecommunication Company Limited",
			CommonName:   MkulimaCommonName,
			Status:       statusOperational,
			Prefixes:     mkulimaPrefixes,
		},
		{
			OperatorName: "Wiafrica Tanzania Limited",
			CommonName:   WiAfricaCommonName,
			Status:       statusOperational,
			Prefixes:     wiAfricaPrefixes,
		},
		{
			OperatorName: "MO Mobile Holding Limited",
			CommonName:   MoCommonName,
			Status:       statusOperational,
			Prefixes:     moPrefixes,
		},
	}
)

type (

	// Data contains basic details of a phone number include the mno
	Data struct {
		OperatorName string   `json:"operator"`
		CommonName   string   `json:"name"`
		Status       string   `json:"status"`
		Prefixes     []string `json:"prefixes"`
	}

	Prefixes []string
)


// Format return a phone number starting with 255 or error
// if it can not be formatted.
// It tries it best to remove white spaces or hyphens put in between
// numbers and a plus sign at the beginning then replace 0 with 255
// if need be
func Format(phoneNumber string) (string,error) {
	phoneNumber = strings.TrimSpace(phoneNumber)
	replacer := strings.NewReplacer(" ", "", "-", "", "+", "")
	phoneNumber = replacer.Replace(phoneNumber)
	numericOnlyRegexStr := "^[0-9]+$"
	match, err := regexp.MatchString(numericOnlyRegexStr, phoneNumber)
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrNumericOnly, err)
	}
	if !match {
		return "", ErrNumericOnly
	}
	phoneNumberLen := len(phoneNumber)

	withoutZero := phoneNumberLen == 9 && !strings.HasPrefix(phoneNumber,"0")
	startsWith255 := strings.HasPrefix(phoneNumber, "255") && phoneNumberLen == 12
	startsWithZero := strings.HasPrefix(phoneNumber, "0") && phoneNumberLen == 10

	if withoutZero{
		return fmt.Sprintf("255%s",phoneNumber),nil
	}

	if startsWith255{
		return phoneNumber,nil
	}

	if startsWithZero{
		return fmt.Sprintf("255%s",phoneNumber[1:]),nil
	}

	return "", fmt.Errorf("pass the correct format")
}


// Details returns Data or err if the phone number inserted is not
// correct. after trying to sanitize it
func Details(phone string) (Data, error) {
	//sanitize
	prefix, err := sanitize(phone)
	if err != nil {
		return Data{}, err
	}

	return findUsingPrefix(prefix)
}

func mergePrefixes() map[string]string {
	var m map[string]string
	m = make(map[string]string)
	for _, prefix := range tigoPrefixes {
		m[prefix] = Tigo
	}

	for _, prefix := range vodaPrefixes {
		m[prefix] = Vodacom
	}

	for _, prefix := range ttclPrefixes {
		m[prefix] = TTCL
	}

	for _, prefix := range zantelPrefixes {
		m[prefix] = Zantel
	}

	for _, prefix := range airtelPrefixes {
		m[prefix] = Airtel
	}

	for _, prefix := range mkulimaPrefixes {
		m[prefix] = MkulimaCommonName
	}

	for _, prefix := range smilePrefixes {
		m[prefix] = SmileCommonName
	}

	for _, prefix := range moPrefixes {
		m[prefix] = MoCommonName
	}

	for _, prefix := range viettelPrefixes {
		m[prefix] = Halotel
	}

	for _, prefix := range wiAfricaPrefixes {
		m[prefix] = WiAfricaCommonName
	}

	return m
}

// sanitize takes a user input and tries to figure out if its in the
// right format. sanitize returns a 3 character prefix and nil error
// if the number is in the right format or an error if the number is
// bogus
func sanitize(phoneNumber string) (string, error) {
	phoneNumber = strings.TrimSpace(phoneNumber)
	replacer := strings.NewReplacer(" ", "", "-", "", "+", "")
	phoneNumber = replacer.Replace(phoneNumber)
	numericOnlyRegexStr := "^[0-9]+$"
	match, err := regexp.MatchString(numericOnlyRegexStr, phoneNumber)
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrNumericOnly, err)
	}
	if !match {
		return "", ErrNumericOnly
	}
	phoneNumberLen := len(phoneNumber)

	// the number len is not correct
	isWrongLen := phoneNumberLen !=9 && phoneNumberLen != 10 && phoneNumberLen != 12
	if isWrongLen{
		return "", ErrInvalidFormat
	}

	withoutZero := phoneNumberLen == 9 && !strings.HasPrefix(phoneNumber,"0")
	startsWith255 := strings.HasPrefix(phoneNumber, "255") && phoneNumberLen== 12
	startsWithZero := strings.HasPrefix(phoneNumber, "0") && phoneNumberLen == 10

	if startsWithZero {
		chars := []rune(phoneNumber)
		prefix := string(chars[0:3])
		return prefix, err
	} else if startsWith255 {
		chars := []rune(phoneNumber)
		prefix := "0" + string(chars[3:5])
		return prefix, err
	}else if withoutZero{
		chars := []rune(phoneNumber)
		prefix := string(chars[0:2])
		return fmt.Sprintf("0%s",prefix), err
	} else {
		return "", ErrInvalidFormat
	}

}

func findUsingPrefix(prefix string) (response Data, err error) {

	m := mergePrefixes()
	operator := m[prefix]
	var found bool

	for _, data := range repository {
		if data.CommonName == operator {
			found = true
			response = data
		}
	}

	if found {
		return response, nil
	} else {
		return Data{}, ErrOperatorNotFound
	}
}
