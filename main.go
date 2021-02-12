package main

import (
	"Samurai/appcmd"
)

// Почему то одинаковая выдача в US и UK в методе FLOW в appstore


// FIX:
// Выводить ошибку если backoff отработал полностью и реквест не починился
// разобраться с review в аппсторе Ошибка почеме то не срабатывает
// комитнуть изменения в Inhuman

func main() {
	appcmd.Execute()
}
