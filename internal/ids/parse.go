package ids

import "strings"

func Prefix(value string) string {
	if i := strings.IndexByte(value, '-'); i > 0 {
		return value[:i]
	}
	return value
}
func IsGenerated(value string) bool { return strings.Count(value, "-") == 1 && len(value) > 10 }
func Join(parts ...string) string {
	clean := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.Trim(strings.TrimSpace(p), "-")
		if p != "" {
			clean = append(clean, p)
		}
	}
	return strings.Join(clean, "-")
}
