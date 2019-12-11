package main

import (
	"flag"
	"fmt"
	"github.com/dalianzhu/cmdhelper/controller"
)

func main() {
	tp := flag.String("t", "", "传入的go templates,按templates来进行渲染")
	js := flag.Bool("j", false, "stdin是一个json文件，只能使用-t来进行渲染")

	sp := flag.String("s", "", "传入分割符，"+
		"使用这个分割符来进行文字的切割，并配合-t渲染")
	regx := flag.String("r", "", "传入正则表达式，"+
		"可以使用正则来进行文字的切割，并配合-t渲染")

	flag.Parse()
	if *tp == "" {
		fmt.Println("-t 不能为空")
		return
	}
	if *js {
		if *sp != "" {
			fmt.Println("-j 不能与-s -r连用")
			return
		}
		if *regx != "" {
			fmt.Println("-j 不能与-s -r连用")
			return
		}
		controller.ReadJsFile(*tp)
	} else {
		if *sp == "" && *regx == "" {
			fmt.Println("-s -r至少传入一个")
			return
		}
		controller.Lines(*tp, *sp, *regx)
	}

}
