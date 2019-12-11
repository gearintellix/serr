package serr

// GetStandardFormat function
func GetStandardFormat() string {
	return "In %s[%s:%d] %s.%s"
}

// GetStandardColorFormat function
func GetStandardColorFormat() string {
	frmt := ""
	frmt += applyForeColor("In", colorDarkGray) + " "
	frmt += applyForeColor("%s", colorLightYellow)
	frmt += applyForeColor("[", colorDarkGray)
	frmt += applyForeColor("%s:%d", colorMagenta)
	frmt += applyForeColor("]", colorDarkGray)
	frmt += " %s.%s"
	return frmt
}
