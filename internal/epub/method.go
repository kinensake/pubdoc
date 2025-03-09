package epub

import (
	"io/fs"
	"path/filepath"
)

func fullPath(path string) string {
	return filepath.Join("OEBPS", path)
}

func (e Epub) GetCoverFile() fs.File {
	file, err := e.reader.Open(fullPath(e.Package.Metadata.CoverPath))
	if err != nil {
		return nil
	}

	return file
}

func (e Epub) GetFileFromPath(path string) fs.File {
	file, err := e.reader.Open(path)
	if err != nil {
		return nil
	}

	return file
}

func (e Epub) GetFile(idRef string) fs.File {
	item, ok := e.Package.Manifest.IDMap[idRef]
	if !ok {
		return nil
	}

	return e.GetFileFromPath(fullPath(item.Href))
}

func (e Epub) GetSpineIDRefs() []string {
	var refs []string

	for _, v := range e.Package.Spine.Items {
		refs = append(refs, v.IDref)
	}

	return refs
}

func (e Epub) GetDir(idRef string) string {
	item, ok := e.Package.Manifest.IDMap[idRef]
	if !ok {
		return ""
	}

	return filepath.Dir(fullPath(item.Href))
}

func (e Epub) GetFilename(idRef string) string {
	item, ok := e.Package.Manifest.IDMap[idRef]
	if !ok {
		return ""
	}

	return filepath.Base(item.Href)
}
