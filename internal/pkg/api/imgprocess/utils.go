package imgprocess

import "path"

func RemoveExtension(str string) string {
	if len(str) == 0 {
		return str
	}

	var newstr string
	ext := path.Ext(str)
	if ext != "" {
		 newstr = str[:len(str) - len(ext)]
	}

	return newstr
}
