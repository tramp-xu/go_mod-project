package main

import (
	"go_mod-project/cmd/reptile/engine"
	"go_mod-project/cmd/reptile/zhenai/parser"
)

func main() {

	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
