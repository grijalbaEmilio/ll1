package asdf

// Función auxiliar que retorna true si un elemento está presente en un arreglo de strings
func Contains(arr []string, elem string) bool {
	for _, e := range arr {
		if e == elem {
			return true
		}
	}
	return false
}

// Función auxiliar que retorna un arreglo de strings sin elementos repetidos
func Unique(arr []string) []string {
	uniqueArr := []string{}
	for _, elem := range arr {
		if !Contains(uniqueArr, elem) {
			uniqueArr = append(uniqueArr, elem)
		}
	}
	return uniqueArr
}

func RemoveElement(slice []string, element string) []string {
	for i, v := range slice {
		if v == element {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
