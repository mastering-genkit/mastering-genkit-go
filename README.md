# Book Writing Repository

This repository provides a complete setup for writing a book in Markdown with support for multiple output formats:

- **Leanpub** - Direct Markdown support
- **Amazon Kindle** - EPUB/MOBI conversion via Pandoc
- **Web** - Static site generation with GitHub Pages

## 📁 Project Structure

```
├── manuscript/           # Leanpub manuscript files
│   ├── Book.txt         # Chapter order for Leanpub
│   ├── Sample.txt       # Sample chapters for preview
│   └── *.md            # Individual chapter files
├── src/                 # Source markdown files
│   ├── chapters/       # Book chapters
│   ├── images/         # Book images
│   └── metadata/       # Book metadata
├── docs/               # GitHub Pages website
├── build/              # Generated output files
├── scripts/            # Build automation scripts
└── config/             # Configuration files
```

## 🚀 Quick Start

1. **Write your content** in `src/chapters/` as Markdown files
2. **Update metadata** in `src/metadata/book.yaml`
3. **Build outputs**:
   ```bash
   # Build all formats
   npm run build

   # Build specific formats
   npm run build:leanpub
   npm run build:kindle
   npm run build:web
   ```

## 📖 Writing Guidelines

### Chapter Files
- Name chapters with numbers: `01-introduction.md`, `02-getting-started.md`
- Use level 1 headers (`#`) for chapter titles
- Use level 2-4 headers for sections within chapters

### Images
- Place images in `src/images/`
- Reference with relative paths: `![Alt text](../images/diagram.png)`
- Supported formats: PNG, JPG, SVG

### Cross-references
- Use standard Markdown links: `[Chapter 2](02-getting-started.md)`
- For web output, links will be converted automatically

### Multiple Authors Support
The system supports both single and multiple authors:

**Single Author (legacy format):**
```yaml
author: "John Doe"
```

**Multiple Authors (recommended):**
```yaml
authors:
  - name: "John Doe"
    bio: "John is a software engineer..."
    email: "john@example.com"
    twitter: "@johndoe"
    website: "https://johndoe.com"
  - name: "Jane Smith"
    bio: "Jane is a technical writer..."
    email: "jane@example.com"
    website: "https://janesmith.com"
```

The system automatically:
- Formats author names for display (e.g., "John Doe & Jane Smith")
- Generates enhanced "About the Authors" sections
- Provides proper metadata for all output formats
- Maintains backward compatibility with single author format

## 🔧 Build System

The build system uses Pandoc and custom scripts to generate:

1. **Leanpub format** - Copies files to `manuscript/` directory
2. **Kindle format** - Generates EPUB and MOBI files
3. **Web format** - Creates a static website in `docs/`

## 📚 Publishing

### Leanpub
1. Connect your GitHub repository to Leanpub
2. Set the manuscript directory to `manuscript/`
3. Use the "Preview" and "Publish" buttons in Leanpub

### Amazon Kindle
1. Use the generated EPUB file in `build/kindle/`
2. Upload to Amazon KDP (Kindle Direct Publishing)

### Web
1. Enable GitHub Pages in repository settings
2. Set source to `docs/` folder
3. Your book will be available at `https://username.github.io/repository-name`

## 🛠️ Requirements

- Node.js (for build scripts)
- Pandoc (for format conversion)
- Git (for version control)

## 📄 License

This template is released under the MIT License. Your book content retains your chosen license.
