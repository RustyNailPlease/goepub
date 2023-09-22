package goepub

import (
	"encoding/xml"
)

/**
*********************************************************
**
**  META-INF/container.xml
**
*********************************************************
**/

type Container struct {
	XMLName xml.Name `xml:"container"`

	Version   string     `xml:"version,attr"`
	RootFiles []RootFile `xml:"rootfiles>rootfile"`
}

type RootFile struct {
	FullPath  string `xml:"full-path,attr"`
	MediaType string `xml:"media-type,attr"`
}

/**
*********************************************************
**
**  content.opf
**
*********************************************************
**/

type OpenPackageFormat struct {
	XMLName xml.Name `xml:"package"`

	Version          string `xml:"version,attr"`
	UniqueIdentifier string `xml:"unique-identifier,attr"`

	MetaData OpenPackageFormatMetaData `xml:"metadata"`

	Guide Guide `xml:"guide"`

	Spine Spine `xml:"spine"`

	Manifest Manifest `xml:"manifest"`
}

type DCDate struct {
	Text  string `xml:",chardata"`
	Event string `xml:"event,attr"`
}

type OpenPackageFormatMetaData struct {
	Metas        []OpenPackageFormatMeta `xml:"meta"`
	DCTitle      string                  `xml:"title"`
	DCCreator    string                  `xml:"creator"`
	DCDate       []DCDate                `xml:"date"`
	DCRights     string                  `xml:"rights"`
	DCLanguage   string                  `xml:"language"`
	DCIdentifier string                  `xml:"identifier"`
}

type OpenPackageFormatMeta struct {
	Name    string `xml:"name,attr"`
	Content string `xml:"content,attr"`
}

type GuideReference struct {
	Type  string `xml:"type,attr"`
	Title string `xml:"title,attr"`
	Href  string `xml:"href,attr"`

	Content string `json:"Content"`
}
type Guide struct {
	Reference []GuideReference `xml:"reference"`
}

type Spine struct {
	Toc string `xml:"toc,attr"`

	ItemRefs []SpineItemRef `xml:"itemref"`
}

type SpineItemRef struct {
	IDRef      string `xml:"idref,attr"`
	Linear     string `xml:"linear,attr"`
	Properties string `xml:"properties,attr"`
}

type Manifest struct {
	Items []ManifestItem `xml:"item"`
}

type ManifestItem struct {
	ID        string `xml:"id,attr"`
	Href      string `xml:"href,attr"`
	MediaType string `xml:"media-type,attr"`
}

/**
*********************************************************
**
**  toc.ncx todo
**
*********************************************************
**/
type TocNcx struct {
}
