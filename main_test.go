package hearthstone_deckparser

import (
	"testing"
)

func parseDeckStringTest(t *testing.T) {
	deckCode := `### CastomMage
# Класс: Маг
# Формат: Стандартный
# Год Феникса
#
# 2x (1) Мурлок-налетчик
# 2x (1) Чародейские стрелы
# 2x (2) Инженер-новичок
# 2x (2) Речной кроколиск
# 2x (2) Чародейский взрыв
# 2x (2) Ящер Кровавой Топи
# 2x (3) Всадник на волке
# 2x (3) Интеллект чародея
# 2x (3) Лидер рейда
# 2x (4) Оазисный хрустогрыз
# 2x (4) Огненный шар
# 2x (4) Превращение
# 2x (4) Щитоносец Сен'джин
# 2x (5) Ночной Клинок
# 2x (6) Огр Тяжелого Кулака
# 
AAECAf0EAA9NvwHYAZwCoQK7Ar8DqwS0BPsEngXZCtoK+QqWDQA=
# 
# Чтобы использовать эту колоду, скопируйте ее в буфер обмена и создайте новую колоду в Hearthstone.
`

	deck, _ := parseDeckString(deckCode)
	t.Log("Name:", deck.Name)
	t.Log("Heroes:", deck.Heroes)
	t.Log("Format:", deck.Format)
	t.Log("Version:", deck.Version)
	t.Log("Cards:", deck.Cards)
}
