package main

type Data_2021 []struct {
	Title string `json:"Title"`
	Date  string `json:"Date"`
}

func holiday_2021() Data_2021 {
	holidays_2021 := Data_2021{
		{
			Title: "元旦",
			Date:  "2021-01-01",
		},
		{
			Title: "成人の日",
			Date:  "2021-01-11",
		},
		{
			Title: "建国記念の日",
			Date:  "2021-02-11",
		},
		{
			Title: "天皇誕生日",
			Date:  "2021-02-23",
		},
		{
			Title: "春分の日",
			Date:  "2021-03-20",
		},
		{
			Title: "昭和の日",
			Date:  "2021-04-29",
		},
		{
			Title: "憲法記念日",
			Date:  "2021-05-03",
		},
		{
			Title: "みどりの日",
			Date:  "2021-05-04",
		},
		{
			Title: "こどもの日",
			Date:  "2021-05-05",
		},
		{
			Title: "海の日",
			Date:  "2021-07-22",
		},
		{
			Title: "スポーツの日",
			Date:  "2021-07-23",
		},
		{
			Title: "山の日",
			Date:  "2021-08-08",
		},
		{
			Title: "休日",
			Date:  "2021-08-09",
		},
		{
			Title: "敬老の日",
			Date:  "2021-09-20",
		},
		{
			Title: "秋分の日",
			Date:  "2021-09-23",
		},
		{
			Title: "文化の日",
			Date:  "2021-11-03",
		},
		{
			Title: "勤労感謝の日",
			Date:  "2021-11-23",
		},
	}

	return holidays_2021
}
