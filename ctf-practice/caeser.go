package main

import "fmt"

func main() {
	str := "EBG KVVV vf n fvzcyr yrggre fhofgvghgvba pvcure gung ercynprf n yrggre jvgu gur yrggre KVVV yrggref nsgre vg va gur nycunorg. EBG KVVV vf na rknzcyr bs gur Pnrfne pvcure, qrirybcrq va napvrag Ebzr. Synt vf SYNTFjmtkOWFNZdjkkNH. Vafreg na haqrefpber vzzrqvngryl nsgre SYNT."
	var transStr string
	for _, c := range str {
		if c >= 65 && c <= 90 {
			c += 13
			if c >= 91 {
				c -= 26
			}
		} else if c >= 97 && c <= 122 {
			c += 13
			if c >= 123 {
				c -= 26
			}
		}
		transStr += string(c)
	}

	fmt.Println(transStr)
}
