package epub

import (
	"io/fs"
	"path/filepath"
)

// fullPath constructs the complete path to a file within the EPUB package.
// It prepends the OEBPS directory to the provided path.
func fullPath(path string) string {
	return filepath.Join("OEBPS", path)
}

// GetCoverFile returns the cover file from the EPUB.
// Returns nil if the cover file cannot be opened.
func (e Epub) GetCoverFile() fs.File {
	file, err := e.reader.Open(fullPath(e.Package.Metadata.CoverPath))
	if err != nil {
		return nil
	}

	return file
}

// GetFileFromPath returns a file from the EPUB at the specified path.
// Returns nil if the file cannot be opened.
func (e Epub) GetFileFromPath(path string) fs.File {
	file, err := e.reader.Open(path)
	if err != nil {
		return nil
	}

	return file
}

// GetFile returns a file from the EPUB identified by the given ID reference.
// Returns nil if the ID reference is not found or the file cannot be opened.
func (e Epub) GetFile(idRef string) fs.File {
	item, ok := e.Package.Manifest.IDMap[idRef]
	if !ok {
		return nil
	}

	return e.GetFileFromPath(fullPath(item.Href))
}

// GetSpineIDRefs returns a slice of all ID references in the EPUB's spine.
// The spine defines the reading order of the EPUB content.
func (e Epub) GetSpineIDRefs() []string {
	var refs []string

	for _, v := range e.Package.Spine.Items {
		refs = append(refs, v.IDref)
	}

	return refs
}

// GetDir returns the directory path for a file identified by the given ID reference.
// Returns an empty string if the ID reference is not found.
func (e Epub) GetDir(idRef string) string {
	item, ok := e.Package.Manifest.IDMap[idRef]
	if !ok {
		return ""
	}

	return filepath.Dir(fullPath(item.Href))
}

// GetFilename returns the filename for a file identified by the given ID reference.
// Returns an empty string if the ID reference is not found.
func (e Epub) GetFilename(idRef string) string {
	item, ok := e.Package.Manifest.IDMap[idRef]
	if !ok {
		return ""
	}

	return filepath.Base(item.Href)
}
