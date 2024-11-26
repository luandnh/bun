module github.com/uptrace/bun

go 1.22.0

require (
	github.com/jinzhu/inflection v1.0.0
	github.com/puzpuzpuz/xsync/v3 v3.4.0
	github.com/rs/zerolog v1.33.0
	github.com/stretchr/testify v1.8.1
	github.com/tmthrgd/go-hex v0.0.0-20190904060850-447a3041c3bc
	github.com/uptrace/bun/extra/buncli v0.0.0-00010101000000-000000000000
	github.com/vmihailenco/msgpack/v5 v5.4.1
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.5 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/urfave/cli/v2 v2.27.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/xrash/smetrics v0.0.0-20240521201337-686a1a2994c1 // indirect
	golang.org/x/sys v0.27.0 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/uptrace/bun/extra/buncli => ./extra/buncli
