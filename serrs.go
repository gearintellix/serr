package serr

// SErrs type mutiple standard error
type SErrs []SErr

// CaptureSErr to capture standard error
func (ox *SErrs) CaptureSErr(errx SErr) {
	if errx != nil {
		*ox = append(*ox, errx)
	}
}
