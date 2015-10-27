# HTML Meta

Golang library for grabbing some HTML meta tags from an io.Reader. Includes title, description, and all open graph tags. Useful for generating link previews

--------------------------------------------------------------------------------

## Installation

```bash
$ go get github.com/jonlaing/htmlmeta
```

## Usage

```golang

import (
	"net/http"
	"github.com/jonlaing/htmlmeta"
)

func main() {
	response, err := http.Get("https://www.github.com/")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		meta := htmlmeta.Extract(response.Body)
		fmt.Println(meta.OGTitle) // print the open graph title
		fmt.Println(meta.OGImage) // print the open graph image
	}
}
```

## Supported Meta Tags

```golang

type HTMLMeta struct {
	Title         string
	Description   string
	OGTitle       string
	OGDescription string
	OGImage       string
	OGAuthor      string
	OGPublisher   string
	OGSiteName    string
}

```
