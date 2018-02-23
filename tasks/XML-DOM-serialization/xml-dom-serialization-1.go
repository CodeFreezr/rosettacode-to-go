package main

import (
    "fmt"
    dom "bitbucket.org/rj/xmldom-go"
)

func main() {
    d, err := dom.ParseStringXml(`
<?xml version="1.0" ?>
<root>
    <element>
        Some text here
    </element>
</root>`)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(d.ToXml()))
}

//\XML-DOM-serialization\xml-dom-serialization-1.go
