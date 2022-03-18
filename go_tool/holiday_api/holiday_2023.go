package main

type Holiday_2023 struct {
	Year string    `json:"year"`
	Data Data_2023 `json: "data"`
}

type Data_2023 []struct {
	Title string `json:"Title"`
	Date  string `json:"Date"`
}

func holiday_2023() Holiday_2023 {
	holidays_2023 := Holiday_2023{
		Year: "2023",
		Data: Data_2023{
			{
				Title: "元旦",
				Date:  "2023-01-01",
			},
			{
				Title: "休日",
				Date:  "2023-01-02",
			},
			{
				Title: "成人の日",
				Date:  "2023-01-09",
			},
			{
				Title: "建国記念の日",
				Date:  "2023-02-11",
			},
			{
				Title: "天皇誕生日",
				Date:  "2023-02-23",
			},
			{
				Title: "春分の日",
				Date:  "2023-03-21",
			},
			{
				Title: "昭和の日",
				Date:  "2023-04-29",
			},
			{
				Title: "憲法記念日",
				Date:  "2023-05-03",
			},
			{
				Title: "みどりの日",
				Date:  "2023-05-04",
			},
			{
				Title: "こどもの日",
				Date:  "2023-05-05",
			},
			{
				Title: "海の日",
				Date:  "2023-07-17",
			},
			{
				Title: "山の日",
				Date:  "2023-08-11",
			},
			{
				Title: "敬老の日",
				Date:  "2023-09-18",
			},
			{
				Title: "秋分の日",
				Date:  "2023-09-23",
			},
			{
				Title: "スポーツの日",
				Date:  "2023-10-09",
			},
			{
				Title: "文化の日",
				Date:  "2023-11-03",
			},
			{
				Title: "勤労感謝の日",
				Date:  "2023-11-23",
			},
		},
	}

	return holidays_2023
}
