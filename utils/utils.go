package utils

import (
	"strings"
)

// ExtractUserIDFromMention takes in a mention (<@ID>) and spits out only the ID.
func ExtractUserIDFromMention(mention string) string {
	mention = strings.Replace(mention, "<", "", 1)
	mention = strings.Replace(mention, ">", "", 1)
	mention = strings.Replace(mention, "@", "", 1)
	return mention
}

// GetStringFromQuotes finds a "string which spans multiple spaces" in a split
// message. It then takes that and replaces the Quote string with a single
// string value of the quote contents.
func GetStringFromQuotes(parts []string) {
	var (
		// str is the string we're searching for in quotes.
		str string
		// startQuote holds the location of the quote
		startQuote int
	)

	length := len(parts)
	val := ""
	startQuote = -1
	for k := 0; k < length; k++ {
		val = parts[k]
		isQuoteStart := val[0] == '"'
		switch {
		case isQuoteStart && val[len(val)-1] == '"':
			parts[k] = val[:len(val)-1][1:]
		case isQuoteStart && startQuote == -1:
			startQuote = k
			str = val[1:] + " "
		case val[len(val)-1] == '"' && startQuote >= 0:
			parts = append(parts[:startQuote], str+val[:len(val)-1])
			parts = append(parts, parts[k+1:]...)
			k -= (length - len(parts)) + 1
			length = len(parts)
			startQuote = -1
		default:
			if startQuote >= 0 {
				str = str + val + " "
			}
		}
	}
}
