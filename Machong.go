package codewar

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var sortMap map[string]int32

// 定義麻將牌的類型
type MahjongTile struct {
	Suit  string
	Value int
}

// 創建一副麻將牌
func createMahjongTiles() []MahjongTile {
	suits := []string{"筒", "條", "萬"}
	var tiles []MahjongTile

	for _, suit := range suits {
		for value := 1; value <= 9; value++ {
			tile := MahjongTile{Suit: suit, Value: value}
			tiles = append(tiles, tile)
		}
	}

	// 加入風牌和箭牌
	winds := []string{"東", "南", "西", "北"}
	dragons := []string{"中", "發", "白"}

	for _, wind := range winds {
		tiles = append(tiles, MahjongTile{Suit: "風", Value: convertWindAndDragon(wind)})
	}

	for _, dragon := range dragons {
		tiles = append(tiles, MahjongTile{Suit: "箭", Value: convertWindAndDragon(dragon)})
	}
	var total []MahjongTile

	//複製四份丟回總排堆
	for i := 0; i < 4; i++ {
		total = append(total, tiles...)
	}
	//加入花牌
	for i := 1; i < 9; i++ {
		total = append(total, MahjongTile{Suit: "花", Value: 1})
	}
	//初始化各類排序
	sortMap = make(map[string]int32)
	sortMap["筒"] = 0
	sortMap["條"] = 1
	sortMap["萬"] = 2
	sortMap["風"] = 3
	sortMap["箭"] = 4
	sortMap["花"] = 5
	return total
}
func convertWindAndDragon(arg string) int {
	switch arg {
	case "東":
		return 1

	case "南":
		return 2

	case "西":
		return 3

	case "北":
		return 4

	case "中":
		return 5

	case "發":
		return 6

	case "白":
		return 7
	default:
		return 0
	}
}

// 洗牌
func shuffleTiles(tiles []MahjongTile) []MahjongTile {
	rand.Seed(time.Now().UnixNano())
	shuffledTiles := make([]MahjongTile, len(tiles))
	perm := rand.Perm(len(tiles))
	for i, v := range perm {
		shuffledTiles[i] = tiles[v]
	}
	return shuffledTiles
}

// 初始化麻將遊戲
func initializeMahjongGame() []MahjongTile {
	tiles := createMahjongTiles()
	shuffledTiles := shuffleTiles(tiles)
	return shuffledTiles
}

func ShowMahjong() {
	// 初始化麻將遊戲
	mahjongTiles := initializeMahjongGame()

	// 顯示初始化後的麻將牌
	fmt.Println("初始化後的麻將牌：")
	a := MahjongSlice(mahjongTiles[:16])
	sort.Sort(a)
	for i, tile := range a {
		fmt.Printf("%d: %s %v\n", i+1, tile.Suit, tile.Value)
	}
}

//麻將排序
type MahjongSlice []MahjongTile

func (s MahjongSlice) Len() int {
	return len(s)
}

func (s MahjongSlice) Less(i, j int) bool {
	if s[i].Suit == s[j].Suit {
		return s[i].Value < s[j].Value

	} else {
		return sortMap[s[i].Suit] < sortMap[s[j].Suit]
	}
}

func (s MahjongSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
