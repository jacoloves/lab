package main

type Data_2022 []struct {
	Title string `json:"Title"`
	Date  string `json:"Date"`
}

func holiday_2022() Data_2022 {
	holidays_2022 := Data_2022{
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
		{
			Title: "天皇誕生日",
			Date:  "2022-02-23",
		},
		{
			Title: "春分の日",
			Date:  "2022-03-21",
		},
		{
			Title: "昭和の日",
			Date:  "2022-04-29",
		},
		{
			Title: "憲法記念日",
			Date:  "2022-05-03",
		},
		{
			Title: "みどりの日",
			Date:  "2022-05-04",
		},
		{
			Title: "こどもの日",
			Date:  "2022-05-05",
		},
		{
			Title: "海の日",
			Date:  "2022-07-18",
		},
		{
			Title: "山の日",
			Date:  "2022-08-11",
		},
		{
			Title: "敬老の日",
			Date:  "2022-09-19",
		},
		{
			Title: "秋分の日",
			Date:  "2022-09-23",
		},
		{
			Title: "スポーツの日",
			Date:  "2022-10-10",
		},
		{
			Title: "文化の日",
			Date:  "2022-11-03",
		},
		{
			Title: "勤労感謝の日",
			Date:  "2022-11-23",
		},
	}

	return holidays_2022
}
