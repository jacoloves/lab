package main

type Holiday_2022 struct {
	Year string    `json:"year"`
	Data Data_2022 `json: "data"`
}

type Data_2022 []struct {
	Title string `json:"Title"`
	Date  string `json:"Date"`
}

func holiday_2022() []Holiday_2022 {
	holidays_2022 := []Holiday_2022{{
		Year: "2022",
		Data: Data_2022{
			{
				Title: "元旦",
				Date:  "2022-01-01",
			},
			{
				Title: "成人の日",
				Date:  "2022-01-10",
			},
			{
				Title: "建国記念の日",
				Date:  "2022-02-11",
			},
		},
	}}

	return holidays_2022
}
