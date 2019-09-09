module go_mod-project

go 1.12

replace golang.org/x/text => github.com/golang/text v0.0.0-20190829152558-3d0f7978add9

replace golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190829043050-9756ffdc2472

replace golang.org/x/net => github.com/golang/net v0.0.0-20190827160401-ba9fcec4b297

replace golang.org/x/tools => github.com/golang/tools v0.0.0-20190830223141-573d9926052a

replace golang.org/x/sys => github.com/golang/sys v0.0.0-20190830142957-1e83adbbebd0

replace golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58

replace golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20190717185122-a985d3407aa7

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.2.1-0.20190905224727-74f33a446dd3

require (
	github.com/PuerkitoBio/goquery v1.5.0
	github.com/astaxie/beego v1.12.0
	github.com/boltdb/bolt v1.3.1
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gorilla/mux v1.7.3
	github.com/microcosm-cc/bluemonday v1.0.2
	github.com/russross/blackfriday v1.5.2
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/smartystreets/goconvey v0.0.0-20190731233626-505e41936337
	github.com/sourcegraph/annotate v0.0.0-20160123013949-f4cad6c6324d // indirect
	github.com/sourcegraph/syntaxhighlight v0.0.0-20170531221838-bd320f5d308e
	golang.org/x/net v0.0.0-20190620200207-3b0461eec859
	golang.org/x/text v0.3.0
)
