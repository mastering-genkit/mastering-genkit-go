# Book Writing Template

A comprehensive setup for writing books in Markdown with multi-format publishing support.

## Formats Supported

- ✅ **Leanpub** - Direct Markdown publishing
- ✅ **Amazon Kindle** - EPUB/MOBI via Pandoc  
- ✅ **Web** - Static site with GitHub Pages
- ✅ **PDF** - Via Pandoc (optional)

## Directory Structure

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

## Quick Start

1. **Install dependencies:**
   ```bash
   npm install
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

## Writing Guidelines

### Chapter Files
- Name with numbers: `01-introduction.md`, `02-getting-started.md`
- Start with level 1 header: `# Chapter Title`
- Use levels 2-4 for sections: `## Section`, `### Subsection`

### Images
- Store in `src/images/`
- Reference with: `![Description](../images/filename.png)`
- Include a `cover.jpg` for book cover

### Cross-references
- Link chapters: `[See Chapter 2](02-getting-started.md)`
- Reference sections: `[Advanced Topics](#advanced-topics)`

### Multiple Authors
Support for single or multiple authors:

**Single author:**
```yaml
author: "Your Name"
```

**Multiple authors:**
```yaml
authors:
  - name: "First Author"
    bio: "Brief bio..."
    email: "first@example.com"
    website: "https://firstauthor.com"
  - name: "Second Author"
    bio: "Brief bio..."
    twitter: "@secondauthor"
```

Features:
- ✅ Auto-formatted author displays
- ✅ Enhanced "About the Authors" sections  
- ✅ Individual author contact info
- ✅ Backward compatibility with single author

## Publishing

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

## Development

```bash
npm run dev          # Start dev server with auto-rebuild
npm run validate     # Check book structure
npm run word-count   # Count words and estimate reading time
npm run lint         # Check Markdown formatting
```

## Requirements

- **Node.js** (for build scripts)
- **Pandoc** (for Kindle/PDF formats)
- **Calibre** (optional, for MOBI generation)

Install Pandoc: https://pandoc.org/installing.html
Install Calibre: https://calibre-ebook.com/

## Advanced Features

- 📊 Word count and reading time estimation
- 🔍 Content validation and structure checking
- 🎨 Customizable web themes via templates
- 📱 Mobile-responsive web version
- 🔍 SEO optimization for web version
- 📝 Markdown linting for consistency

## Troubleshooting

**Build errors?** Run `npm run validate` to check structure
**Missing Pandoc?** Install from official website
**Images not showing?** Check file paths in `src/images/`
**Web version not updating?** Clear browser cache

---

Happy writing! 📚✍️
