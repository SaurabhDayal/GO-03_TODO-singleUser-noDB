package models

var Todos []Todo

var Counter = 2

type Todo struct {
	Id          int    `json:"id"`
	UserId      int    `json:"userId"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func init() {
	Todos = []Todo{
		{
			Id:          Counter,
			UserId:      124,
			Title:       "Bank checks",
			Description: "Collect checks house and deposit in all the bank branches",
		},
		{
			Id:          Counter,
			UserId:      234,
			Title:       "Motor Cycle Service",
			Description: "Wash and Dry",
		},
	}

}
