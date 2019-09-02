module go_mod-project

go 1.12

replace golang.org/x/text => github.com/golang/text v0.0.0-20190829152558-3d0f7978add9

replace golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190829043050-9756ffdc2472

replace golang.org/x/net => github.com/golang/net v0.0.0-20190827160401-ba9fcec4b297

replace golang.org/x/tools => github.com/golang/tools v0.0.0-20190830223141-573d9926052a

replace golang.org/x/sys => github.com/golang/sys v0.0.0-20190830142957-1e83adbbebd0

replace golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58

replace golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20190717185122-a985d3407aa7

require (
	github.com/astaxie/beego v1.12.0
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	golang.org/x/net v0.0.0-20190620200207-3b0461eec859
	golang.org/x/text v0.3.0
)
