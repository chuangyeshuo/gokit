package str

func StringOrEmpty(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
