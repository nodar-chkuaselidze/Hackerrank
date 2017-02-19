package main

import "bufio"
import "fmt"

import "os"
import "sort"
import "strings"

const Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type cipherCols []string

func (cc cipherCols) Len() int           { return len(cc) }
func (cc cipherCols) Swap(i, j int)      { cc[i], cc[j] = cc[j], cc[i] }
func (cc cipherCols) Less(i, j int) bool { return cc[i][0] < cc[j][0] }

func uniqueString(str string) (unique string) {
	usedMap := make(map[byte]struct{})

	for _, b := range []byte(str) {
		if _, ok := usedMap[b]; !ok {
			unique += string(b)
			usedMap[b] = struct{}{}
		}
	}

	return
}

type CipherText struct {
	secret     string
	ctext      string
	text       string
	decryptMap map[byte]byte
}

func (ct *CipherText) GenerateSubtitutions() {
	ct.decryptMap = make(map[byte]byte)
	secret := uniqueString(ct.secret)
	alphabet := uniqueString(secret + Alphabet)

	//generate unsorted Cols
	columns := make(cipherCols, len(secret))

	for i, v := range alphabet {
		columns[i%len(secret)] += string(v)
	}

	sort.Sort(columns)

	substitution := strings.Join(columns, "")

	for i, v := range []byte(substitution) {
		ct.decryptMap[v] = Alphabet[i]
	}
}

func (ct *CipherText) Decrypt() {
	for _, v := range []byte(ct.ctext) {
		if r, ok := ct.decryptMap[v]; ok {
			ct.text += string(r)
		} else {
			ct.text += string(v)
		}
	}
}

func main() {
	var count int
	var texts []CipherText

	fmt.Scan(&count)

	texts = make([]CipherText, count)
	scanner := bufio.NewScanner(os.Stdin)

	for i := range texts {
		scanner.Scan()
		texts[i].secret = scanner.Text()
		scanner.Scan()
		texts[i].ctext = scanner.Text()

		texts[i].GenerateSubtitutions()
		texts[i].Decrypt()
		fmt.Println(texts[i].text)
	}
}
