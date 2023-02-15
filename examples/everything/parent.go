package everything

import (
	"github.com/samber/mo"
	"time"
)

type Parent struct {
	ID ParentID `json:"id"`

	Number1 int              `json:"number1"`
	Number2 []int            `json:"number2"`
	Number3 *int             `json:"number3"`
	Number4 *[]int           `json:"number4"`
	Number5 []*int           `json:"number5"`
	Number6 mo.Option[int]   `json:"number6"`
	Number7 []mo.Option[int] `json:"number7"`

	Text1 string              `json:"text1"`
	Text2 []string            `json:"text2"`
	Text3 *string             `json:"text3"`
	Text4 *[]string           `json:"text4"`
	Text5 []*string           `json:"text5"`
	Text6 mo.Option[string]   `json:"text6"`
	Text7 []mo.Option[string] `json:"text7"`

	Date1 time.Time           `json:"date1"`
	Date2 []time.Time         `json:"date2"`
	Date3 *time.Time          `json:"date3"`
	Date4 *[]time.Time        `json:"date4"`
	Date5 []*time.Time        `json:"date5"`
	Date6 mo.Option[string]   `json:"date6"`
	Date7 []mo.Option[string] `json:"date7"`

	Child1 Child              `json:"child1"`
	Child2 []Child            `json:"child2"`
	Child3 *Child             `json:"child3"`
	Child4 *[]Child           `json:"child4"`
	Child5 []*Child           `json:"child5"`
	Child6 mo.Option[Child]   `json:"child6"`
	Child7 []mo.Option[Child] `json:"child7"`

	Map1 map[string]float64 `json:"map1"`
	Map2 map[ChildID]Child  `json:"map2_weird_name"`
}
