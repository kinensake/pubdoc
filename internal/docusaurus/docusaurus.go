package docusaurus

import (
	"bytes"
	"embed"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/PuerkitoBio/goquery"

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
			subDstDir := filepath.Join(dstDir, v.Name())
			subSrcDir := filepath.Join(srcDir, v.Name())

			if err := copyDir(subDstDir, subSrcDir); err != nil {
				return err
			}

			continue
		}

		// File
		dstFile := filepath.Join(dstDir, v.Name())
		srcFile := filepath.Join(srcDir, v.Name())

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

	docDir := filepath.Join("docs", slug.Make(e.Package.Metadata.Title))
	if err := os.Mkdir(docDir, os.ModePerm); err != nil {
		return fmt.Errorf("Mkdir: %v", err)
	}

	assetDir := filepath.Join(docDir, "assets")
	if err := os.Mkdir(assetDir, os.ModePerm); err != nil {
		return fmt.Errorf("Mkdir: %v", err)
	}

	refs := e.GetSpineIDRefs()

	for i, v := range refs {
		html, htmlDir, filename, err := getHTML(e, v)
		if err != nil {
			cobra.CompErrorln(err.Error())
			continue
		}

		html, err = copyImageToProject(e, html, htmlDir, assetDir)
		if err != nil {
			cobra.CompErrorln(err.Error())
			continue
		}

		html, err = replaceDocHref(html)
		if err != nil {
			cobra.CompErrorln(err.Error())
			continue
		}

		md, err := html2md.ConvertString(html)
		if err != nil {
			cobra.CompErrorln(err.Error())
			continue
		}

		filenameMD := strings.TrimSuffix(filename, ".html") + ".md"
		if err := writeToProject(md, docDir, i, filenameMD); err != nil {
			cobra.CompErrorln(err.Error())
			continue
		}
	}

	return nil
}

func getHTML(e *epub.Epub, idRef string) (string, string, string, error) {
	f := e.GetFile(idRef)
	if f == nil {
		return "", "", "", fmt.Errorf("GetFile: %v", idRef)
	}
	defer f.Close()

	return sanitizeHTML(f), e.GetDir(idRef), e.GetFilename(idRef), nil
}

func writeToProject(md string, docDir string, pos int, filename string) error {
	tmpl, err := template.New("doc").Parse(docTmpl)
	if err != nil {
		return fmt.Errorf("Parse: %v", err)
	}

	fp := filepath.Join(docDir, filename)

	f, err := os.Create(fp)
	if err != nil {
		return fmt.Errorf("Create: %v", err)
	}
	defer f.Close()

	if err := tmpl.Execute(f, map[string]interface{}{
		"Position": pos,
		"Content":  md,
	}); err != nil {
		return fmt.Errorf("Execute: %v", err)
	}

	return nil
}

func copyImageToProject(e *epub.Epub, html string, htmlDir string, assetDir string) (string, error) {
	r := bytes.NewBufferString(html)

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return "", fmt.Errorf("goquery.NewDocumentFromReader: %v", err)
	}

	s := doc.Find("img")
	for i := range s.Nodes {
		item := s.Eq(i)

		src, ok := item.Attr("src")
		if !ok {
			continue
		}

		filename := filepath.Base(src)
		srcImgPath := filepath.Join(htmlDir, src)
		dstImgPath := filepath.Join(assetDir, filename)

		f := e.GetFileFromPath(srcImgPath)
		if f == nil {
			return "", fmt.Errorf("epub.GetFileFromPath: %v", err)
		}

		content, err := io.ReadAll(f)
		if err != nil {
			return "", fmt.Errorf("io.ReadAll: %v", err)
		}

		if err := os.WriteFile(dstImgPath, content, 0o666); err != nil {
			return "", fmt.Errorf("os.WriteFile: %v", err)
		}

		item.SetAttr("src", filepath.Join("assets", filename))
	}

	modified, err := doc.Html()
	if err != nil {
		return "", fmt.Errorf("doc.Html: %v", err)
	}

	return modified, nil
}

func replaceDocHref(html string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBufferString(html))
	if err != nil {
		return "", fmt.Errorf("goquery.NewDocumentFromReader: %v", err)
	}

	s := doc.Find("a")
	for i := range s.Nodes {
		item := s.Eq(i)

		href, ok := item.Attr("href")
		if !ok {
			continue
		}

		filename := filepath.Base(href)
		if strings.HasSuffix(filename, ".html") {
			item.SetAttr("href", strings.TrimSuffix(filename, ".html")+".md")
		}
	}

	modified, err := doc.Html()
	if err != nil {
		return "", fmt.Errorf("doc.Html: %v", err)
	}

	return modified, nil
}
