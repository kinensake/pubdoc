# PubDoc

PubDoc is a command-line tool that converts EPUB books into Docusaurus documentation sites. It extracts content from EPUB files, converts HTML to Markdown, and organizes the content in a Docusaurus project structure, making it easy to publish books as documentation websites.

## Installation

### Prerequisites
- Go 1.21 or higher

### Install from source
```bash
# Clone the repository
git clone https://github.com/kinensake/pubdoc.git
cd pubdoc

# Build the binary
go build -o pubdoc

# Move to a directory in your PATH (optional)
sudo mv pubdoc /usr/local/bin/
```

### Install using Go
```bash
go install github.com/kinensake/pubdoc@latest
```

## Usage

### Create a new Docusaurus project
```bash
pubdoc new my-book-site
cd my-book-site
```

### Add an EPUB book to your project
```bash
pubdoc add /path/to/your/book.epub
```

### Start the Docusaurus development server
```bash
npm start
```

### Build for production
```bash
npm run build
```

## Supporters

This project is enhanced with AI assistance for:
- Comprehensive code documentation with GoDoc style comments
- README generation and documentation
- Docusaurus template enhancement

---

## License

[MIT License](LICENSE)