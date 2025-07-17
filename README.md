# Book Writing Template

A comprehensive setup for writing books in Markdown with multi-format publishing support.

> 📚 **New to this template?** Check out [DOCS.md](DOCS.md) for a complete documentation overview, or jump to [GETTING_STARTED.md](GETTING_STARTED.md) for a quick start guide.

## 📋 Formats Supported

- ✅ **Leanpub** - Direct Markdown publishing
- ✅ **Amazon Kindle** - EPUB/MOBI via Pandoc  
- ✅ **Web** - Static site with GitHub Pages
- ✅ **PDF** - Via Pandoc with WeasyPrint/Puppeteer

## 📁 Directory Structure

```
src/
├── chapters/           # Your book chapters
├── images/            # Book images and diagrams  
└── metadata/          # Book information

manuscript/            # Generated Leanpub files
build/                # Generated output files
docs/                 # Generated website
scripts/              # Build automation
config/               # Configuration files
```

## 🚀 Quick Start

1. **Install dependencies:**
   ```bash
   npm install
   # For complete setup with all tools:
   npm run setup
   ```

2. **Update book information:**
   Edit `src/metadata/book.yaml` with your book details

3. **Write your content:**
   Add chapters in `src/chapters/` as `01-chapter-name.md` files

4. **Build your book:**
   ```bash
   npm run build        # Build all formats
   npm run build:web    # Build website only
   npm run dev          # Start development server
   ```

## � Writing Guidelines

### Chapter Files
- Name with numbers: `01-introduction.md`, `02-getting-started.md`
- Start with level 1 header: `# Chapter Title`
- Use levels 2-4 for sections: `## Section`, `### Subsection`

### Images
- Store in `src/images/`
- Reference with: `![Description](../images/filename.png)`
- Include a `cover.jpg` for book cover
- Supported formats: PNG, JPG, SVG

### Cross-references
- Link chapters: `[See Chapter 2](02-getting-started.md)`
- Reference sections: `[Advanced Topics](#advanced-topics)`


## 🔧 Available Commands

```bash
npm run build          # Build all formats
npm run build:leanpub  # Build Leanpub format
npm run build:kindle   # Build Kindle format (EPUB + MOBI)
npm run build:web      # Build web format
npm run build:pdf      # Build PDF format
npm run dev            # Start development server with auto-rebuild
npm run validate       # Check book structure
npm run word-count     # Count words and estimate reading time
npm run lint           # Check Markdown formatting
npm run clean          # Clean build artifacts
npm run setup          # Install all dependencies
```

## 📚 Publishing

### Leanpub
1. Connect GitHub repo to Leanpub
2. Set manuscript directory to `manuscript/`
3. Use Leanpub's Preview/Publish buttons

### Kindle (Amazon KDP)
1. Use generated EPUB: `build/kindle/book.epub`
2. Upload to Amazon KDP
3. Or use MOBI if generated: `build/kindle/book.mobi`

### Web (GitHub Pages)
1. Enable Pages in repo settings
2. Set source to `docs/` folder  
3. Access at: `https://username.github.io/repo-name`

## 🛠️ Setup & Requirements

### Quick Setup
Run the automated setup script:
```bash
npm run setup
# or directly:
./scripts/setup-all.sh
```

This script automatically installs:
- ✅ Node.js and npm dependencies
- ✅ Python 3 and pip
- ✅ Pandoc (for PDF and Kindle generation)
- ✅ WeasyPrint (for high-quality PDF generation)
- ✅ Puppeteer (PDF generation via Chrome headless)
- ✅ Calibre (for MOBI generation)
- ✅ markdownlint-cli (for linting)

## 🎨 Advanced Features

- 📊 Word count and reading time estimation
- 🔍 Content validation and structure checking
- 🎨 Customizable web themes via templates
- 📱 Mobile-responsive web version
- 🔍 SEO optimization for web version
- 📝 Markdown linting for consistency


## Dependencies Overview

| Tool | Purpose | Required |
|------|---------|----------|
| Node.js | JavaScript runtime for build scripts | ✅ Yes |
| Pandoc | Document converter | ✅ Yes |
| Python 3 | Runtime for WeasyPrint | ✅ Yes |
| WeasyPrint | High-quality PDF generation | Recommended |
| Puppeteer | PDF generation via Chrome headless | Recommended |
| Calibre | MOBI generation for Kindle | Optional |
| markdownlint-cli | Markdown linting | Optional |

---

## 📄 License

This template is released under the MIT License. Your book content retains your chosen license.

---

Happy writing! 📚✍️
