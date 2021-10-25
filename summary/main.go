package main

import "fmt"

// T type from json
type T struct {
	Status int `json:"status"`
	Code   int `json:"code"`
	Data   struct {
		List []struct {
			Name      string `json:"name"`
			CreatedAt string `json:"created_at"`
			Id        int    `json:"id"`
			Type      int    `json:"type"`
			Value     int    `json:"value"`
			Url       string `json:"url"`
		} `json:"list"`
		Total int `json:"total"`
	} `json:"data"`
	Message string `json:"message"`
}

type AwardRecord struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

/*type AwardItem struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
	Type uint   `json:"type"`
}*/

type AwardItem struct {
	AwardRecord
	Url  string `json:"url"`
	Type uint   `json:"type"`
}

type AwardItemFull struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
	Type uint   `json:"type"`
}

// AwardItemNest 嵌套结构体
type AwardItemNest struct {
	AwardRecord AwardRecord `json:"award_record"`
	Url         string      `json:"url"`
	Type        uint        `json:"type"`
}

type A struct {
	AwardItem
	A string `json:"a"`
}

func main() {
	fmt.Println("Hello world")
	structTest()
}

func structTest() {

	record := AwardRecord{
		Id:   0,
		Name: "aaa",
	}

	fmt.Printf("record:%v\n", record)

	// item := AwardItem{
	// 	Id:   0,
	// 	Name: "aaa",
	// 	Url:  "baidu.com",
	// 	Type: 1,
	// }

	// fmt.Printf("item:%v\n", item)

	/*item := AwardItem{
		Url:  "baidu",
		Type: 0,
	}
	item.Id = 1
	item.Name = "aaa"*/

	item := AwardItem{
		AwardRecord: AwardRecord{
			Id:   0,
			Name: "aaa",
		},
		Url:  "baidu.com",
		Type: 1,
	}

	itemFull := AwardItemFull{
		Id:   0,
		Name: "aaa",
		Url:  "baidu.com",
		Type: 1,
	}

	awardItemNest := AwardItemNest{
		AwardRecord: AwardRecord{
			Id:   0,
			Name: "aaa",
		},
		Url:  "baidu",
		Type: 1,
	}

	fmt.Printf("\titem:%v\n \titemFull:%v\n \tawardItemNest:%v\n", item, itemFull, awardItemNest)

}
