// преобразует указанный в командной строке uuid или ref
// в ref или uuid соответственно
package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var uuid string
	var ref string
	var err error

	flag.StringVar(&uuid, "uuid", "", `uuid ссылочного объекта приложения 1С, например '9a25c95e-36df-11e9-80ed-0050569f2e9f'`)
	flag.StringVar(&ref, "ref", "", `ref ссылочного объекта приложения 1С из навигационной ссылки или из сообщения объект не найден, например '80ed0050569f2e9f11e936df9a25c95e'`)

	flag.Parse()

	// example: 9a25c95e-36df-11e9-80ed-0050569f2e9f
	uuidRegexp := regexp.MustCompile("(([0-9a-f]){8}-(([0-9a-f]){4}-){3}([0-9a-f]){12})")
	// example: 80ed0050569f2e9f11e936df9a25c95e
	refRegexp := regexp.MustCompile("([0-9a-f]){32}")

	if uuid != "" && uuidRegexp.MatchString(uuid) {
		ref, err = uuid2Ref(uuid)
		fmt.Println(ref)
	} else if ref != "" && refRegexp.MatchString(ref) {
		uuid, err = ref2UUID(ref)
		fmt.Println(uuid)
	} else {
		fmt.Println("Введены не корректные данные.")
		flag.PrintDefaults()
	}

	if err != nil {
		fmt.Println("Ошибка обработки данных.")
		fmt.Println(err)
	}
}

// преобразует uuid вида '9a25c95e-36df-11e9-80ed-0050569f2e9f'
// в ref вида '80ed0050569f2e9f11e936df9a25c95e',
// которую можно использовать для создания навигационной ссылки
func uuid2Ref(uuid string) (string, error) {
	var builder strings.Builder

	zeroUUID := uuid[0:8]
	firstUUID := uuid[9:13]
	secondUUID := uuid[14:18]
	thirdUUID := uuid[19:23]
	fourthUUID := uuid[24:]

	_, err := builder.WriteString(thirdUUID)
	_, err = builder.WriteString(fourthUUID)
	_, err = builder.WriteString(secondUUID)
	_, err = builder.WriteString(firstUUID)
	_, err = builder.WriteString(zeroUUID)

	return builder.String(), err
}

// преобразует ref вида '80ed0050569f2e9f11e936df9a25c95e'
// в  uuid вида '9a25c95e-36df-11e9-80ed-0050569f2e9f',
// который можно использовать для установки ссылки нового объекта
// или для поиска объекта по uuid
func ref2UUID(ref string) (string, error) {
	var builder strings.Builder

	zeroPart := ref[24:]
	firstPart := ref[20:24]
	secondPart := ref[16:20]
	thirdPart := ref[0:4]
	fourthPart := ref[4:16]

	_, err := builder.WriteString(zeroPart)
	_, err = builder.WriteString("-")
	_, err = builder.WriteString(firstPart)
	_, err = builder.WriteString("-")
	_, err = builder.WriteString(secondPart)
	_, err = builder.WriteString("-")
	_, err = builder.WriteString(thirdPart)
	_, err = builder.WriteString("-")
	_, err = builder.WriteString(fourthPart)

	return builder.String(), err
}
