package main

import (
	"log"
	"os"
)

var paragraphStart = []byte{0xE5, 0x01, 0x00, 0xE6, 0x0A, 0x00}
var paragraphEnd = []byte{0xC4, 0x00}
var lastParagraphEnd = []byte{0xE6, 0x00, 0x00, 0xC4, 0x00}

var charset = map[int]string{
	0x1090: "А",
	0x1091: "Б",
	0x1092: "В",
	0x1093: "Г",
	0x1094: "Д",
	0x1095: "Е",
	0x1096: "Ж",
	0x1097: "З",
	0x1098: "И",
	0x1099: "Й",
	0x109A: "К",
	0x109B: "Л",
	0x109C: "М",
	0x109D: "Н",
	0x109E: "О",
	0x109F: "П",
	0x10A0: "Р",
	0x10A1: "С",
	0x10A2: "Т",
	0x10A3: "У",
	0x10A4: "Ф",
	0x10A5: "Х",
	0x10A6: "Ц",
	0x10A7: "Ч",
	0x10A8: "Ш",
	0x10A9: "Щ",
	0x10AA: "Ъ",
	0x10AB: "Ы",
	0x10AC: "Ь",
	0x10AD: "Э",
	0x10AE: "Ю",
	0x10AF: "Я",
	0x10B0: "а",
	0x10B1: "б",
	0x10B2: "в",
	0x10B3: "г",
	0x10B4: "д",
	0x10B5: "е",
	0x10B6: "ж",
	0x10B7: "з",
	0x10B8: "и",
	0x10B9: "й",
	0x10BA: "к",
	0x10BB: "л",
	0x10BC: "м",
	0x10BD: "н",
	0x10BE: "о",
	0x10BF: "п",
	0x1180: "р",
	0x1181: "с",
	0x1182: "т",
	0x1183: "у",
	0x1184: "ф",
	0x1185: "х",
	0x1186: "ц",
	0x1187: "ч",
	0x1188: "ш",
	0x1189: "щ",
	0x118A: "ъ",
	0x118B: "ы",
	0x118C: "ь",
	0x118D: "э",
	0x118E: "ю",
	0x118F: "я",
}

func main() {
	fileName := os.Args[1]

	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	outFile, err := os.Create(fileName + ".txt")
	defer outFile.Close()
	if err != nil {
		panic(err)
	}

	fileStat, err := file.Stat()

	log.Printf("Writing to %s.txt", fileName)

	data := make([]byte, fileStat.Size())
	dataSize, err := file.Read(data)
	if err != nil {
		panic(err)
	}

	paragraphCount := 0
	paragraphOpen := false
	var paragraph string

	var i int
	for i < dataSize {
		if checkToken(i, data, paragraphStart) {
			i = i + len(paragraphStart)
			paragraphOpen = true

			continue
		}

		if checkToken(i, data, paragraphEnd) {
			outFile.WriteString(paragraph + "\n")
			i = i + len(paragraphEnd)
			paragraphCount++
			paragraphOpen = false
			paragraph = ""

			continue
		}

		if checkToken(i, data, lastParagraphEnd) {
			outFile.WriteString(paragraph + "\n")
			i = i + len(lastParagraphEnd)
			paragraphCount++
			paragraphOpen = false
			paragraph = ""

			continue
		}

		if paragraphOpen {
			if data[i+1] == 0x10 || data[i+1] == 0x11 {
				// double-byte char
				ch := int(data[i+1])*256 + int(data[i])
				paragraph += charset[ch]
				i += 2
				continue
			} else {
				paragraph += string(data[i])
				i++
				continue
			}
		}

		i++
	}

	log.Printf("Written %d paragraphs", paragraphCount)
}

func checkToken(pos int, data []byte, token []byte) bool {
	var i int
	for i < len(token) {
		if data[pos+i] != token[i] {
			return false
		}
		i++
	}
	return true
}
