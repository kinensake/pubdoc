package epub

import (
	"archive/zip"
	"encoding/xml"
	"errors"
	"os"
)

const (
	PropertyCover = "cover-image"
	PropertyNav   = "nav"
	MetaNameCover = "cover"
	ContainerPath = "META-INF/container.xml"
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
	Spine    Spine    `xml:"spine"`
}

type Metadata struct {
	Identifier  string `xml:"identifier"`
	Title       string `xml:"title"`
	Language    string `xml:"language"`
	Publisher   string `xml:"publisher"`
	Date        string `xml:"date"`
	Description string `xml:"description"`
	Creator     string `xml:"creator"`
	CoverPath   string
	NavPath     string

	Meta []struct {
		Name    string `xml:"name,attr"`
		Content string `xml:"content,attr"`
	} `xml:"meta"`
}

type Manifest struct {
	Items []ManifestItem `xml:"item"`
	IDMap map[string]ManifestItem
}

type ManifestItem struct {
	ID         string `xml:"id,attr"`
	Href       string `xml:"href,attr"`
	MediaType  string `xml:"media-type,attr"`
	Properties string `xml:"properties,attr"`
}

type Spine struct {
	Items []SpintItem `xml:"itemref"`
}

type SpintItem struct {
	IDref string `xml:"idref,attr"`
}

func New(filePath string) (*Epub, error) {
	r, err := zip.OpenReader(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, ErrNotFound
		}
		if errors.Is(err, zip.ErrFormat) {
			return nil, ErrInvalid
		}
	}

	con, err := parseContainer(r)
	if err != nil {
		return nil, ErrInvalid
	}

	pack, err := parsePackage(r, con.RootFile.FullPath)
	if err != nil {
		return nil, ErrInvalid
	}

	return &Epub{
		Container: con,
		Package:   pack,
		reader:    r,
	}, nil
}

func parseContainer(r *zip.ReadCloser) (*Container, error) {
	var container Container

	file, err := r.Open(ContainerPath)
	if err != nil {
		return nil, err
	}

	decoder := xml.NewDecoder(file)
	if err := decoder.Decode(&container); err != nil {
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
	if err := decoder.Decode(&pack); err != nil {
		return nil, err
	}

	pack.Manifest.IDMap = make(map[string]ManifestItem)

	for _, v := range pack.Manifest.Items {
		pack.Manifest.IDMap[v.ID] = v

		switch v.Properties {
		case PropertyCover:
			pack.Metadata.CoverPath = v.Href
		case PropertyNav:
			pack.Metadata.NavPath = v.Href
		}
	}

	if pack.Metadata.CoverPath == "" {
		for _, v := range pack.Metadata.Meta {
			if v.Name == MetaNameCover {
				pack.Metadata.CoverPath = pack.Manifest.IDMap[v.Content].Href
			}
		}
	}

	return &pack, nil
}
