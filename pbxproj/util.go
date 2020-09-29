package pbxproj

func isUUID(s string) bool {
	if 24 != len(s) {
		return false
	}
	match := false
	for _, c := range s {
		if (c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
			match = true
		} else {
			match = false
		}

		if !match {
			return false
		}
	}

	return true
}
