package controller

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template"
)

func Trim(s string) string {
	return strings.Trim(s, "\n \t")
}

func Lines(tpStr string, sep, regxStr string) {
	var tmpl *template.Template
	var regx *regexp.Regexp
	var err error
	if tpStr != "" {
		tmpl = template.New("t1")
		tmpl = tmpl.Funcs(MapFuncs)
		tmpl = template.Must(tmpl.Parse(tpStr))
	}

	if regxStr != "" {
		//fmt.Println("init regx 1", regxStr)
		regx, err = regexp.Compile(regxStr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	br := bufio.NewReader(os.Stdin)
	for {
		line, err := br.ReadString('\n')
		if err != nil {
			//fmt.Println(err)
			break
		}
		if Trim(line) == "" {
			continue
		}

		if sep != "" {
			lineArr := strings.Split(line, sep)
			if tpStr != "" {
				var doc bytes.Buffer
				//fmt.Println("line arr", lineArr)
				trimedLineArr := make([]string, 0)
				for _, lineSepVal := range lineArr {
					trimVal := Trim(lineSepVal)
					if trimVal != "" {
						trimedLineArr = append(trimedLineArr, trimVal)
					}
				}
				tmpl.Execute(&doc, trimedLineArr)
				os.Stdout.Write(doc.Bytes())
			}
		} else if regxStr != "" {
			var doc bytes.Buffer
			groups := regx.FindStringSubmatch(line)
			//fmt.Println("groups", groups)
			tmpl.Execute(&doc, groups)
			os.Stdout.WriteString(doc.String())
		}
	}
}