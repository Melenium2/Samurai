package main

import (
	"Samurai/appcmd"
)

// Почему то одинаковая выдача в US и UK в методе FLOW в appstore


// FIX:
// Убрать контекст по времени
// Выводить ошибку если backoff отработал полностью и реквест не починился

func main() {
	appcmd.Execute()
}
