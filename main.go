package hearthstone_deckparser

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"log"
	"strings"
)

const (
	DECKSTRING_VERSION = 1
)

type Deck struct {
	Version int
	Format  int
	Name    string
	Heroes  []int
	Cards   []Card
}

type Card struct {
	Id    int
	Count int
}

func stringToLines(s string) (lines []string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}

func getVarInt(stream *bytes.Reader) (int, error) {
	var (
		shift  uint = 0
		result      = 0
	)

	for {
		c, err := stream.ReadByte()
		if err != nil {
			return -1, err
		}
		i := int(c)

		result |= (i & 0x7f) << shift
		shift += 7
		if (i & 0x80) == 0x0 {
			break
		}
	}

	return result, nil
}

func parseDeckString(deckString string) (deck Deck, err error) {
	var (
		line     string
		deckCode string
	)

	lines, err := stringToLines(deckString)
	for _, line = range lines {
		if len(line) > 3 && line[:3] == "###" {
			deck.Name = line[4:]
		} else if len(line) > 0 && line[0] == '#' {
			continue
		} else {
			if _, err := base64.StdEncoding.DecodeString(deckCode); err == nil {
				deckCode = line
			}
		}
	}

	decoded, err := base64.StdEncoding.DecodeString(deckCode)
	if err != nil {
		log.Println("decode error:", err)
		return
	}
	data := bytes.NewReader(decoded)

	if b, _ := getVarInt(data); b != 0x0 {
		log.Fatal("Invalid deckstring")
	}

	deck.Version, _ = getVarInt(data)
	if deck.Version != DECKSTRING_VERSION {
		log.Fatalf("Unsupported deckstring version %d", deck.Version)
	}
	deck.Format, _ = getVarInt(data)
	if deck.Format < 0 || deck.Format > 2 {
		log.Fatalf("Unsupported FormatType in deckstring %d", deck.Format)
	}

	heroesCount, _ := getVarInt(data)
	for i := 0; i < heroesCount; i++ {
		hero, _ := getVarInt(data)
		deck.Heroes = append(deck.Heroes, hero)
	}

	singleCardCount, _ := getVarInt(data)
	for i := 0; i < singleCardCount; i++ {
		cardId, _ := getVarInt(data)
		deck.Cards = append(deck.Cards, Card{cardId, 1})
	}

	pairCardCount, _ := getVarInt(data)
	for i := 0; i < pairCardCount; i++ {
		cardId, _ := getVarInt(data)
		deck.Cards = append(deck.Cards, Card{cardId, 2})
	}

	manyCardCount, _ := getVarInt(data)
	for i := 0; i < manyCardCount; i++ {
		cardId, _ := getVarInt(data)
		count, _ := getVarInt(data)
		deck.Cards = append(deck.Cards, Card{cardId, count})
	}

	return
}
