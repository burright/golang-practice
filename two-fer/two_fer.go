package twofer

func ShareWith(name string) string {
	
	// If name is blank
	if name == "" {
		return "One for you, one for me."
	}
	// Else
	return "One for "+name+", one for me."
}
