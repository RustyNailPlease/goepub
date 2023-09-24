# goepub

# Usage

#### Initialization

```go
epub, err := NewEpub("data/DomainDrivenDesign.epub")
//epub, err := NewEpub("data/TheMysteriousLord.epub")
if err != nil {
    t.Error(err.Error())
    return
}
```

#### Reading Data

```go
//  The src format may be text/xxxx.xhtml#xxxxxxx
src := strings.Split(epub.OPF.Guide.Reference[1].Href, "#")[0]
file := epub.FilePaths[epub.OEBPSPath+src]
openedFile, err := file.Open()
bytes, err := io.ReadAll(openedFile)
log.Println(string(bytes))
```