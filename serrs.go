package serr

// SErrs type mutiple standard error
type SErrs []SErr

// CaptureSErr to capture standard error
func (ox *SErrs) CaptureSErr(errx SErr) {
	if errx != nil {
		*ox = append(*ox, errx)
	}
}

func (ox *SErrs) String() string {
	var message string
	for _, errx := range *ox {
		if message != "" {
			message += "; "
		}

		message += errx.String()
	}

	return message
}

func (ox *SErrs) SimpleString() string {
	var message string
	for _, errx := range *ox {
		if message != "" {
			message += "; "
		}

		message += errx.SimpleString()
	}

	return message
}
