package mna

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const (
	tigoCommonName     = "Tigo"
	vodaCommonName     = "Vodacom"
	ttclCommonName     = "TTCL"
	airtelCommonName   = "Airtel"
	zantelCommonName   = "Zantel"
	smileCommonName    = "Smile"
	moCommonName       = "Mo Mobile"
	viettelCommonName  = "Halotel"
	mkulimaCommonName  = "Mkulima"
	wiAfricaCommonName = "Wiafrica"
	statusOperational  = "Operational"
)

var (

	ErrOperatorNotFound = errors.New("mobile operator not found")
	ErrInvalidFormat         = errors.New("invalid format, allowed formats are \"+255765XXXXXX\" and \"0765XXXXXX\"")
	ErrNumericOnly           = errors.New("phone numbers should contains numeric characters only")


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
			CommonName:   tigoCommonName,
		},
		{
			OperatorName: "Vodacom Tanzania PLC",
			CommonName:   vodaCommonName,
			Status:       statusOperational,
			Prefixes:     vodaPrefixes,
		},
		{
			OperatorName: "Tanzania Telecommunications Corporation",
			CommonName:   ttclCommonName,
			Status:       statusOperational,
			Prefixes:     ttclPrefixes,
		},
		{
			OperatorName: "Zanzibar Telecom PLC",
			CommonName:   zantelCommonName,
			Status:       statusOperational,
			Prefixes:     zantelPrefixes,
		},
		{
			OperatorName: "Airtel Tanzania PLC",
			CommonName:   airtelCommonName,
			Status:       statusOperational,
			Prefixes:     airtelPrefixes,
		},
		{
			OperatorName: "Smile Communications Tanzania Limited",
			CommonName:   smileCommonName,
			Status:       statusOperational,
			Prefixes:     smilePrefixes,
		},
		{
			OperatorName: "Viettel Tanzania PLC",
			CommonName:   viettelCommonName,
			Status:       statusOperational,
			Prefixes:     viettelPrefixes,
		},
		{
			OperatorName: "Mkulima African Telecommunication Company Limited",
			CommonName:   mkulimaCommonName,
			Status:       statusOperational,
			Prefixes:     mkulimaPrefixes,
		},
		{
			OperatorName: "Wiafrica Tanzania Limited",
			CommonName:   wiAfricaCommonName,
			Status:       statusOperational,
			Prefixes:     wiAfricaPrefixes,
		},
		{
			OperatorName: "MO Mobile Holding Limited",
			CommonName:   moCommonName,
			Status:       statusOperational,
			Prefixes:     moPrefixes,
		},
	}
)

type (
	Data struct {
		OperatorName string   `json:"operator_name"`
		CommonName   string   `json:"common_name"`
		Status       string   `json:"status"`
		Prefixes     []string `json:"prefixes"`
	}
)

func CheckNumber(phone string)  (Data,error){
	//sanitize
	prefix, err := sanitize(phone)
	if err != nil{
		return Data{}, err
	}

	return findUsingPrefix(prefix)
}

func mergePrefixes() map[string]string {
	var m map[string]string
	m = make(map[string]string)
	for _, prefix := range tigoPrefixes {
		m[prefix] = tigoCommonName
	}

	for _, prefix := range vodaPrefixes {
		m[prefix] = vodaCommonName
	}

	for _, prefix := range ttclPrefixes {
		m[prefix] = ttclCommonName
	}

	for _, prefix := range zantelPrefixes {
		m[prefix] = zantelCommonName
	}

	for _, prefix := range airtelPrefixes {
		m[prefix] = airtelCommonName
	}

	for _, prefix := range mkulimaPrefixes {
		m[prefix] = mkulimaCommonName
	}

	for _, prefix := range smilePrefixes {
		m[prefix] = smileCommonName
	}

	for _, prefix := range moPrefixes {
		m[prefix] = moCommonName
	}

	for _, prefix := range viettelPrefixes {
		m[prefix] = viettelCommonName
	}

	for _, prefix := range wiAfricaPrefixes {
		m[prefix] = wiAfricaCommonName
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
	if len(phoneNumber) == 3 && strings.HasPrefix(phoneNumber, "0") {
		return phoneNumber, nil
	}

	if len(phoneNumber) != 10 && len(phoneNumber) != 12 {
		return "", ErrInvalidFormat
	}

	startsWith255 := strings.HasPrefix(phoneNumber, "255") && len(phoneNumber) == 12
	startsWithZero := strings.HasPrefix(phoneNumber, "0") && len(phoneNumber) == 10

	if startsWithZero {
		chars := []rune(phoneNumber)
		prefix := string(chars[0:3])
		return prefix, err
	} else if startsWith255 {
		chars := []rune(phoneNumber)
		prefix := "0" + string(chars[3:5])
		return prefix, err
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
