package opsutilsgo

import "regexp"

var reg *regexp.Regexp = regexp.MustCompile(`(?:\w+://)?([^:/]+)(?::\d+)?`)

// GetDomain extract domain from url
func GetDomain(urlStr string) string {
	result := reg.FindStringSubmatch(urlStr)
	if len(result) == 0 {
		return ""
	}
	return result[1]
}
