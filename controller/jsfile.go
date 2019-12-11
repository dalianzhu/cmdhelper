package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

func ReadJsFile(tpStr string) {
	in, _ := ioutil.ReadAll(os.Stdin)
	var strInData interface{}
	err := json.Unmarshal(in, &strInData)
	if err != nil {
		fmt.Println(err)
		return
	}

	var doc bytes.Buffer
	tmpl := template.New("t1")
	tmpl = tmpl.Funcs(MapFuncs)

	tmpl = template.Must(tmpl.Parse(tpStr))
	tmpl.Execute(&doc, strInData)
	os.Stdout.Write(doc.Bytes())
}
