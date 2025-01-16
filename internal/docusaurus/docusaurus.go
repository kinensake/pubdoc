package docusaurus

import (
	"embed"
	"fmt"
	"os"
	"path"
	"text/template"

	html2md "github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/gosimple/slug"
	"github.com/kinensake/pubdoc/internal/epub"
	"github.com/spf13/cobra"
)

//go:embed project/*
var em embed.FS

func NewProject(name string) error {
	return copyDir(name, "project")
}

func copyDir(dstDir string, srcDir string) error {
	_ = os.Mkdir(dstDir, os.ModePerm)

	entries, err := em.ReadDir(srcDir)
	if err != nil {
		return fmt.Errorf("ReadDir: %v", err)
	}

	for _, v := range entries {
		if v.IsDir() {
			subDstDir := path.Join(dstDir, v.Name())
			subSrcDir := path.Join(srcDir, v.Name())

			if err := copyDir(subDstDir, subSrcDir); err != nil {
				return err
			}

			continue
		}

		// File
		dstFile := path.Join(dstDir, v.Name())
		srcFile := path.Join(srcDir, v.Name())

		content, err := em.ReadFile(srcFile)
		if err != nil {
			return fmt.Errorf("ReadFile: %v", err)
		}

		if err := os.WriteFile(dstFile, content, 0o666); err != nil {
			return fmt.Errorf("WriteFile: %v", err)
		}
	}

	return nil
}

func AddEpub(epubPath string) error {
	e, err := epub.New(epubPath)
	if err != nil {
		return fmt.Errorf("epub.New: %v", err)
	}

	dir := path.Join("docs", slug.Make(e.Package.Metadata.Title))
	if err := os.Mkdir(dir, os.ModePerm); err != nil {
		return fmt.Errorf("Mkdir: %v", err)
	}

	refs := e.GetSpineIDRefs()
	mds := make([]string, 0, len(refs))

	for _, v := range refs {
		md, err := convertToMarkdown(e, v)
		if err != nil {
			cobra.CompErrorln(err.Error())
			continue
		}

		mds = append(mds, md)
	}

	if err := writeToProject(mds, dir); err != nil {
		return fmt.Errorf("writeToProject: %v", err)
	}

	return nil
}

func convertToMarkdown(e *epub.Epub, idRef string) (string, error) {
	f := e.GetFile(idRef)
	if f == nil {
		return "", fmt.Errorf("GetFile: %v", idRef)
	}
	defer f.Close()

	b, err := html2md.ConvertString(sanitizeHTML(f))
	if err != nil {
		return "", fmt.Errorf("ConvertReader: %v", err)
	}

	return b, nil
}

func writeToProject(mds []string, dir string) error {
	tmpl, err := template.New("doc").Parse(docTmpl)
	if err != nil {
		return fmt.Errorf("Parse: %v", err)
	}

	for i, v := range mds {
		fp := path.Join(dir, fmt.Sprintf("%d.md", i))

		f, err := os.Create(fp)
		if err != nil {
			return fmt.Errorf("Create: %v", err)
		}
		defer f.Close()

		if err := tmpl.Execute(f, map[string]interface{}{
			"Position": i,
			"Content":  string(v),
		}); err != nil {
			return fmt.Errorf("Execute: %v", err)
		}
	}

	return nil
}
