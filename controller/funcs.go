package controller

import (
	"fmt"
	"reflect"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
)

var masterMap = sprig.TxtFuncMap()

func init() {
	tp := make(template.FuncMap)
	for key, item := range masterMap {
		tp[key] = item
	}
	for key, item := range MapFuncs {
		tp[key] = item
	}
	MapFuncs = tp
}

var MapFuncs = template.FuncMap{
	"format": func(f string, i ...interface{}) string {
		return fmt.Sprintf(f, i...)
	},
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
