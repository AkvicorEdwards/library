package handler

import "strings"

func FilterStringBySep(author, sep string) []string {
	temp := strings.Split(author, sep)
	res := make([]string, 0, len(temp))
	for _, v := range temp {
		v = strings.TrimSpace(v)
		if len(v) > 0 {
			res = append(res, v)
		}
	}
	return res
}
