package gee

import (
	"strings"
)

func parsePattern(pattern string) []string{
	vs := strings.Split(pattern,"/")
	parts := make([]string,0)
	for _,item :=range vs{
		if item != ""{
			parts = append(parts,item)
			if item[0] == '*'{
				break
			}
		}
	}
	return parts
}

