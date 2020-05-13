package blockchain

import (
	"strconv"
	"strings"
)

const difficulty = 3

func getDataHash(data string) int {
	var hash int
	for i := 0; i < len(data); i++ {
		character := int([]rune(string(data[i]))[0])
		hash = ((hash << 5) - hash) + character
		hash = hash & hash
	}
	return hash;
}

// GetProvenHash will return a hash out of the Block.Data, Block.PrevHash and a increasing counter until it matches the prove of work rule
// the prove of work rule in this case is: lastHashDigits == 3
func GetProvenHash(block *Block) (hash string, nonce int) {
	var data string = block.Data + block.PrevHash
	for i := 0; ; i++ {
		hash := getDataHash(data + string(i))
		if rule, stringHash := doesHashProvesWork(strconv.Itoa(hash)); rule {
			return stringHash, i
		}
	}
}

func doesHashProvesWork(hash string) (bool, string) {
	splittedHash := strings.Split(hash, "")
	lastHashDigits := strings.Join(splittedHash[len(splittedHash)-difficulty:], "")
	repeatedDifficulty := strings.Repeat(strconv.Itoa(difficulty), difficulty)
	rule := lastHashDigits == repeatedDifficulty
	return rule, hash
}