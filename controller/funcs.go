package controller

import (
	"fmt"
	"reflect"
	"strings"
	"text/template"
)

var MapFuncs = template.FuncMap{
	"eol": func() string {
		return "\n"
	},
	"cut": func(arr interface{}, start, end int) []interface{} {
		//fmt.Println("arr", arr)
		val := reflect.ValueOf(arr)
		ret := make([]interface{}, 0)
		switch val.Kind() {
		case reflect.Slice:
			if end <= 0 {
				end = val.Len() + end
			}
			for i := 0; i < val.Len(); i++ {
				e := val.Index(i)
				if i >= start && i < end {
					ret = append(ret, e.Interface())
				}
			}
		}
		return ret
	},
	"join": func(arr interface{}, sep string) string {
		val := reflect.ValueOf(arr)
		ret := make([]string, 0)
		switch val.Kind() {
		case reflect.Slice:
			for i := 0; i < val.Len(); i++ {
				e := val.Index(i)
				ret = append(ret, fmt.Sprint(e.Interface()))
			}
		}
		return strings.Join(ret, sep)
	},
}
