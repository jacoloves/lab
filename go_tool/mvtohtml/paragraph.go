package main

func is_Empty(s *[]string) bool {
	return len(*s) == 0
}

func slice_Delete(s *[]string) {
	res := []string{}

	*s = res
}

func slice_Join(s *[]string) string {
	var str string
	for _, v := range *s {
		str += v
	}
	return str
}

func last_Paragraph(slice_arr *[]string) string {
	html_line := ""
	if !(is_Empty(slice_arr)) {
		*slice_arr = append(*slice_arr, "</p>\n")
		html_line = slice_Join(slice_arr)
		slice_Delete(slice_arr)
	}
	return html_line
}

func paragraph(pattern string, html_line *string, slice_arr *[]string) {
	if pattern != "NONE" {
		if !(is_Empty(slice_arr)) {
			*slice_arr = append(*slice_arr, "</p>\n")
			*html_line = slice_Join(slice_arr)
			slice_Delete(slice_arr)
		}
	} else {
		if is_Empty(slice_arr) {
			*slice_arr = append(*slice_arr, "<p>")
			*slice_arr = append(*slice_arr, *html_line)
		} else {
			*slice_arr = append(*slice_arr, *html_line)
		}
		*html_line = ""
	}
}
