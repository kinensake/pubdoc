package epub

import (
	"archive/zip"
	"encoding/xml"
)

type Epub struct {
	Container *Container
	Package   *Package
	reader    *zip.ReadCloser
}

type Container struct {
	RootFile struct {
		FullPath string `xml:"full-path,attr"`
	} `xml:"rootfiles>rootfile"`
}

type Package struct {
	Metadata Metadata `xml:"metadata"`
	Manifest Manifest `xml:"manifest"`
}

type Metadata struct {
	Identifier  string `xml:"identifier"`
	Title       string `xml:"title"`
	Language    string `xml:"language"`
	Publisher   string `xml:"publisher"`
	Date        string `xml:"date"`
	Description string `xml:"description"`
	Creator     string `xml:"creator"`
}

type Manifest struct {
	Items []ManifestItem `xml:"item"`
}

type ManifestItem struct {
	ID         string `xml:"id,attr"`
	Href       string `xml:"href,attr"`
	MediaType  string `xml:"media-type,attr"`
	Properties string `xml:"properties,attr"`
}

func New(fileName string) *Epub {
	r, err := zip.OpenReader(fileName)
	if err != nil {
		panic(err)
	}

	con, err := parseContainer(r)
	if err != nil {
		panic(err)
	}

	pack, err := parsePackage(r, con.RootFile.FullPath)
	if err != nil {
		panic(err)
	}

	return &Epub{
		Container: con,
		Package:   pack,
		reader:    r,
	}
}

func parseContainer(r *zip.ReadCloser) (*Container, error) {
	var container Container

	file, err := r.Open("META-INF/container.xml")
	if err != nil {
		return nil, err
	}

	decoder := xml.NewDecoder(file)

	err = decoder.Decode(&container)
	if err != nil {
		return nil, err
	}

	return &container, nil
}

func parsePackage(r *zip.ReadCloser, path string) (*Package, error) {
	var pack Package

	file, err := r.Open(path)
	if err != nil {
		return nil, err
	}

	decoder := xml.NewDecoder(file)

	err = decoder.Decode(&pack)
	if err != nil {
		return nil, err
	}

	return &pack, nil
}
