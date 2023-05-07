package utils

import "strconv"

var (
	CountryCode = map[string]string{
		"AD": "Andorra",
		"AE": "United Arab Emirates",
		"AF": "AfghanIstan",
		"AG": "Antigua and Barbuda",
		"AI": "Anguilla",
		"AL": "Albania",
		"AM": "Armenia",
		"AO": "Angola",
		"AQ": "Antarctica",
		"AR": "Argentina",
		"AS": "American Samoa",
		"AT": "Austria",
		"AU": "Australia",
		"AW": "Aruba",
		"AX": "Åland Islands",
		"AZ": "Azerbaijan",
		"BA": "Bosnia and Herzegovina",
		"BB": "Barbados",
		"BD": "Bangladesh",
		"BE": "Belgium",
		"BF": "Burkina Faso",
		"BG": "Bulgaria",
		"BH": "Bahrain",
		"BI": "Burundi",
		"BJ": "Benin",
		"BL": "Saint Barthélemy",
		"BM": "Bermuda",
		"BN": "Brunei Darussalam",
		"BO": "Bolivia (Plurinational State of)",
		"BQ": "Bonaire, Sint Eustatius and Saba",
		"BR": "Brazil",
		"BS": "Bahamas",
		"BT": "Bhutan",
		"BV": "Bouvet Island",
		"BW": "Botswana",
		"BY": "Belarus",
		"BZ": "Belize",
		"CA": "Canada",
		"CC": "Cocos (Keeling) Islands",
		"CD": "Congo, Democratic Republic of the",
		"CF": "Central African Republic",
		"CG": "Congo",
		"CH": "Switzerland",
		"CI": "Côte d'Ivoire",
		"CK": "Cook Islands",
		"CL": "Chile",
		"CM": "Cameroon",
		"CN": "China",
		"CO": "Colombia",
		"CR": "Costa Rica",
		"CU": "Cuba",
		"CV": "Cabo Verde",
		"CW": "Curaçao",
		"CX": "ChrIstmas Island",
		"CY": "Cyprus",
		"CZ": "Czechia",
		"DE": "Germany",
		"DJ": "Djibouti",
		"DK": "Denmark",
		"DM": "Dominica",
		"DO": "Dominican Republic",
		"DZ": "Algeria",
		"EC": "Ecuador",
		"EE": "Estonia",
		"EG": "Egypt",
		"EH": "Western Sahara",
		"ER": "Eritrea",
		"ES": "Spain",
		"ET": "Ethiopia",
		"FI": "Finland",
		"FJ": "Fiji",
		"FK": "Falkland Islands (Malvinas)",
		"FM": "Micronesia (Federated States of)",
		"FO": "Faroe Islands",
		"FR": "France",
		"GA": "Gabon",
		"GB": "United Kingdom of Great Britain and Northern Ireland",
		"GD": "Grenada",
		"GE": "Georgia",
		"GF": "French Guiana",
		"GG": "Guernsey",
		"GH": "Ghana",
		"GI": "Gibraltar",
		"GL": "Greenland",
		"GM": "Gambia",
		"GN": "Guinea",
		"GP": "Guadeloupe",
		"GQ": "Equatorial Guinea",
		"GR": "Greece",
		"GS": "South Georgia and the South Sandwich Islands",
		"GT": "Guatemala",
		"GU": "Guam",
		"GW": "Guinea-BIssau",
		"GY": "Guyana",
		"HK": "Hong Kong",
		"HM": "Heard Island and McDonald Islands",
		"HN": "Honduras",
		"HR": "Croatia",
		"HT": "Haiti",
		"HU": "Hungary",
		"ID": "Indonesia",
		"IE": "Ireland",
		"IL": "Israel",
		"IM": "Isle of Man",
		"IN": "India",
		"IO": "BritIsh Indian Ocean Territory",
		"IQ": "Iraq",
		"IR": "Iran (Islamic Republic of)",
		"IS": "Iceland",
		"IT": "Italy",
		"JE": "Jersey",
		"JM": "Jamaica",
		"JO": "Jordan",
		"JP": "Japan",
		"KE": "Kenya",
		"KG": "Kyrgyzstan",
		"KH": "Cambodia",
		"KI": "Kiribati",
		"KM": "Comoros",
		"KN": "Saint Kitts and NevIs",
		"KP": "Korea (Democratic People's Republic of)",
		"KR": "Korea, Republic of",
		"KW": "Kuwait",
		"KY": "Cayman Islands",
		"KZ": "Kazakhstan",
		"LA": "Lao People's Democratic Republic",
		"LB": "Lebanon",
		"LC": "Saint Lucia",
		"LI": "Liechtenstein",
		"LK": "Sri Lanka",
		"LR": "Liberia",
		"LS": "Lesotho",
		"LT": "Lithuania",
		"LU": "Luxembourg",
		"LV": "Latvia",
		"LY": "Libya",
		"MA": "Morocco",
		"MC": "Monaco",
		"MD": "Moldova, Republic of",
		"ME": "Montenegro",
		"MF": "Saint Martin (French part)",
		"MG": "Madagascar",
		"MH": "Marshall Islands",
		"MK": "North Macedonia",
		"ML": "Mali",
		"MM": "Myanmar",
		"MN": "Mongolia",
		"MO": "Macao",
		"MP": "Northern Mariana Islands",
		"MQ": "Martinique",
		"MR": "Mauritania",
		"MS": "Montserrat",
		"MT": "Malta",
		"MU": "Mauritius",
		"MV": "Maldives",
		"MW": "Malawi",
		"MX": "Mexico",
		"MY": "Malaysia",
		"MZ": "Mozambique",
		"NA": "Namibia",
		"NC": "New Caledonia",
		"NE": "Niger",
		"NF": "Norfolk Island",
		"NG": "Nigeria",
		"NI": "Nicaragua",
		"NL": "Netherlands",
		"NO": "Norway",
		"NP": "Nepal",
		"NR": "Nauru",
		"NU": "Niue",
		"NZ": "New Zealand",
		"OM": "Oman",
		"PA": "Panama",
		"PE": "Peru",
		"PF": "French Polynesia",
		"PG": "Papua New Guinea",
		"PH": "Philippines",
		"PK": "PakIstan",
		"PL": "Poland",
		"PM": "Saint Pierre and Miquelon",
		"PN": "Pitcairn",
		"PR": "Puerto Rico",
		"PS": "Palestine, State of",
		"PT": "Portugal",
		"PW": "Palau",
		"PY": "Paraguay",
		"QA": "Qatar",
		"RE": "Réunion",
		"RO": "Romania",
		"RS": "Serbia",
		"RU": "Russian Federation",
		"RW": "Rwanda",
		"SA": "Saudi Arabia",
		"SB": "Solomon Islands",
		"SC": "Seychelles",
		"SD": "Sudan",
		"SE": "Sweden",
		"SG": "Singapore",
		"SH": "Saint Helena, Ascension and TrIstan da Cunha",
		"SI": "Slovenia",
		"SJ": "Svalbard and Jan Mayen",
		"SK": "Slovakia",
		"SL": "Sierra Leone",
		"SM": "San Marino",
		"SN": "Senegal",
		"SO": "Somalia",
		"SR": "Suriname",
		"SS": "South Sudan",
		"ST": "Sao Tome and Principe",
		"SV": "El Salvador",
		"SX": "Sint Maarten (Dutch part)",
		"SY": "Syrian Arab Republic",
		"SZ": "Eswatini",
		"TC": "Turks and Caicos Islands",
		"TD": "Chad",
		"TF": "French Southern Territories",
		"TG": "Togo",
		"TH": "Thailand",
		"TJ": "TajikIstan",
		"TK": "Tokelau",
		"TL": "Timor-Leste",
		"TM": "TurkmenIstan",
		"TN": "TunIsia",
		"TO": "Tonga",
		"TR": "Türkiye",
		"TT": "Trinidad and Tobago",
		"TV": "Tuvalu",
		"TW": "Taiwan, Province of China",
		"TZ": "Tanzania, United Republic of",
		"UA": "Ukraine",
		"UG": "Uganda",
		"UM": "United States Minor Outlying Islands",
		"US": "United States of America",
		"UY": "Uruguay",
		"UZ": "UzbekIstan",
		"VA": "Holy See",
		"VC": "Saint Vincent and the Grenadines",
		"VE": "Venezuela (Bolivarian Republic of)",
		"VG": "Virgin Islands (BritIsh)",
		"VI": "Virgin Islands (U.S.)",
		"VN": "Viet Nam",
		"VU": "Vanuatu",
		"WF": "WallIs and Futuna",
		"WS": "Samoa",
		"YE": "Yemen",
		"YT": "Mayotte",
		"ZA": "South Africa",
		"ZM": "Zambia",
		"ZW": "Zimbabwe",
	}

	ProviderSMSLIst = map[string]bool{
		"Topolo": true,
		"Rond":   true,
		"Kildy":  true,
	}

	ProviderVoiceLIst = map[string]bool{
		"TransparentCalls": true,
		"E-Voice":          true,
		"JustPhone":        true,
	}

	ProviderEmailLIst = map[string]bool{
		"Gmail":       true,
		"Yahoo":       true,
		"Hotmail":     true,
		"MSN":         true,
		"Orange":      true,
		"Comcast":     true,
		"AOL":         true,
		"Live":        true,
		"RediffMail":  true,
		"GMX":         true,
		"Proton Mail": true,
		"Yandex":      true,
		"Mail.ru":     true,
	}

	StatusCode = map[string]bool{
		"active": true,
		"closed": true,
	}
)

func CodeToCountry(code string) string {
	country, ok := CountryCode[code]
	if !ok {
		return ""
	}

	return country
}

func IsValidCountry(country string) bool {
	_, ok := CountryCode[country]

	return ok
}

func IsValidBandwidth(bandwidth string) bool {
	width, err := strconv.Atoi(bandwidth)
	if err != nil {
		return false
	}

	return width >= 0 && width <= 100
}

func IsValidResponseTime(responseTime string) bool {
	time, err := strconv.Atoi(responseTime)
	if err != nil {
		return false
	}

	return time >= 0
}

func IsValidSMSProvider(provider string) bool {
	_, ok := ProviderSMSLIst[provider]

	return ok
}

func IsValidVoiceProvider(provider string) bool {
	_, ok := ProviderVoiceLIst[provider]

	return ok
}

func IsValidEmailProvider(provider string) bool {
	_, ok := ProviderEmailLIst[provider]

	return ok
}

func IsValidLoad(load string) bool {
	currentLoad, err := strconv.Atoi(load)
	if err != nil {
		return false
	}

	return currentLoad >= 0 && currentLoad <= 100
}

func IsValidStability(stability string) bool {
	_, err := strconv.ParseFloat(stability, 32)

	return err == nil
}

func IsValidPurity(purity string) bool {
	_, err := strconv.Atoi(purity)

	return err == nil
}

func IsValidTTFB(ttfb string) bool {
	_, err := strconv.Atoi(ttfb)

	return err == nil
}

func IsMedianDuration(median string) bool {
	_, err := strconv.Atoi(median)

	return err == nil
}

func IsValidDeliveryTime(deliveryTime string) bool {
	time, err := strconv.Atoi(deliveryTime)
	if err != nil {
		return false
	}

	return time > 0
}

func IsValidStatus(status string) bool {
	_, ok := StatusCode[status]

	return ok
}
