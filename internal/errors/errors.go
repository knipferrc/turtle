package errors

// CheckError checks if an error occurs
func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}
