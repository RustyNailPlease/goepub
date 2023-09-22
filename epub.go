package goepub

import (
	"archive/zip"
	"encoding/xml"
	"errors"
	"io"
	"log"
)

type Epub struct {
	File      *zip.ReadCloser
	FilePaths map[string]*zip.File

	Container Container

	OPF OpenPackageFormat
}

func NewEpub(filePath string) (*Epub, error) {
	epub := &Epub{
		FilePaths: make(map[string]*zip.File),
	}

	zipFile, err := zip.OpenReader(filePath)
	if err != nil {
		return nil, err
	}

	epub.File = zipFile
	for _, file := range zipFile.File {
		//log.Println("---> ", file.Name)
		epub.FilePaths[file.Name] = file
	}

	// 1. read META-INF/container.xml
	err = epub.readContainer()
	if err != nil {
		return nil, err
	}

	// 2. read OPF(open package format)
	err = epub.readOPF()
	if err != nil {
		return nil, err
	}
	// read opd.guide.reference.cover
	//loadReferenceContent(epub)

	return epub, nil
}

//func loadReferenceContent(epub *Epub) {
//	if len(epub.OPF.Guide.Reference) > 0 {
//		for i := 0; i < len(epub.OPF.Guide.Reference); i++ {
//			reference := &epub.OPF.Guide.Reference[i]
//
//			if len(reference.Href) > 0 {
//				href, ok := epub.FilePaths["OEBPS/"+reference.Href]
//				if ok {
//					buf, err := readTextFromZipFile(href)
//					if err != nil {
//						log.Println("load reference failed: ", err.Error())
//						continue
//					}
//					reference.Content = string(buf)
//				} else {
//					log.Println("load reference failed2: ", reference.Href)
//				}
//			}
//		}
//	}
//}

func (epub *Epub) readContainer() error {
	containerXml, ok := epub.FilePaths["META-INF/container.xml"]
	if !ok {
		return errors.New("META-INF/container.xml not found")
	}

	bytes, err := readTextFromZipFile(containerXml)
	if err != nil {
		return err
	}

	var c Container
	err = xml.Unmarshal(bytes, &c)
	if err != nil {
		return err
	}

	epub.Container = c

	return nil
}

func (epub *Epub) readOPF() error {
	if len(epub.Container.RootFiles) == 0 {
		return errors.New("opf file not found")
	}
	opfPath := epub.FilePaths[epub.Container.RootFiles[0].FullPath]
	bytes, err := readTextFromZipFile(opfPath)
	if err != nil {
		return err
	}
	var opf OpenPackageFormat
	err = xml.Unmarshal(bytes, &opf)
	if err != nil {
		return err
	}
	epub.OPF = opf

	return nil
}

func readTextFromZipFile(containerXml *zip.File) ([]byte, error) {
	srcFile, err := containerXml.Open()
	if err != nil {
		panic(err.Error())
	}

	defer func(srcFile io.ReadCloser) {
		err := srcFile.Close()
		if err != nil {
			log.Println("close file[", containerXml.Name, "] error: ", err.Error())
			return
		}
	}(srcFile)

	bytes, err := io.ReadAll(srcFile)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
