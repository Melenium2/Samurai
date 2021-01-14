package models

import (
	"fmt"
	"strings"
)

type Category string

func (c Category) Split() (string, string) {
	splited := strings.Split(string(c), "|")
	if len(splited) > 2 {
		panic("invalid category")
	}
	return splited[0], splited[1]
}

func NewCategory(cat, subcat string) Category {
	return Category(strings.ToLower(fmt.Sprintf("%s|%s", cat, subcat)))
}
