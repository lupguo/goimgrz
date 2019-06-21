package girls

// check str whether in str list
func inlist(str string, strList []string) bool {
	for _, s := range strList {
		if str == s {
			return true
		}
	}
	return false
}

//
func walkDir()  {

}