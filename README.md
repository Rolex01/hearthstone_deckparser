# hearthstone_deckparser

## Getting Started

First, download the deckparser library:

```shell
go get github.com/rolex01/hearthstone_deckparser
```

Start using the library:

```go
deckCode := "AAECAf0EBMKhA4y2A8W4A97EAw1xuwKrBJYFn5sDoJsDv6QD9KsD8a8DwbgDwrgDjLkDgb8DAA=="
deck, _ := parseDeckString(deckCode)
log.Println("Name:", deck.Name)
log.Println("Heroes:", deck.Heroes)
log.Println("Format:", deck.Format)
log.Println("Version:", deck.Version)
log.Println("Cards:", deck.Cards)
```
