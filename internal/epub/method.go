package epub

import (
	"fmt"
	"io/fs"
)

func fullPath(path string) string {
	return fmt.Sprintf("OEBPS/%s", path)
}

func (e Epub) GetCoverFile() fs.File {
	file, err := e.reader.Open(fullPath(e.Package.Metadata.CoverPath))
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

	file, err := e.reader.Open(fullPath(item.Href))
	if err != nil {
		return nil
	}

	return file
}

func (e Epub) GetSpineIDRefs() []string {
	var refs []string

	for _, v := range e.Package.Spine.Items {
		refs = append(refs, v.IDref)
	}

	return refs
}
