# Setup Instructions

This repository includes a comprehensive setup script that installs all dependencies needed to build the book in multiple formats.

## Quick Setup

Run the automated setup script:

```bash
npm run setup
# or directly:
./scripts/setup-all.sh
```

This script will automatically detect your operating system and install:

- ✅ Node.js and npm dependencies
- ✅ Python 3 and pip
- ✅ Pandoc (for PDF and Kindle generation)
- ✅ WeasyPrint (for high-quality PDF generation)
- ✅ wkhtmltopdf (PDF fallback)
- ✅ Calibre (for MOBI generation)
- ✅ markdownlint-cli (for linting)

## Manual Setup

If you prefer to install dependencies manually:

### macOS (with Homebrew)
```bash
# Install Node.js and npm
brew install node

# Install Python 3
brew install python3

# Install Pandoc
brew install pandoc

# Install PDF generation tools
brew install cairo pango gdk-pixbuf libffi

# Try WeasyPrint via Homebrew first
brew install weasyprint

# If Homebrew fails due to externally-managed-environment, use pipx
if ! command -v weasyprint &> /dev/null; then
    brew install pipx
    pipx install weasyprint
fi

# Alternative: Virtual environment method
# python3 -m venv ~/.weasyprint-venv
# source ~/.weasyprint-venv/bin/activate
# pip install weasyprint
# deactivate

# Note: wkhtmltopdf has been discontinued on Homebrew as of 2024-12-16
# Alternative: Install Puppeteer for PDF generation
npm install --save-dev puppeteer

# Install Calibre (for MOBI)
brew install calibre

# Install project dependencies
npm install

# Install markdownlint globally
npm install -g markdownlint-cli
```

### Ubuntu/Debian
```bash
# Update package list
sudo apt-get update

# Install Node.js and npm
sudo apt-get install nodejs npm

# Install Python 3 and pip
sudo apt-get install python3 python3-pip

# Install Pandoc
sudo apt-get install pandoc

# Install WeasyPrint dependencies
sudo apt-get install build-essential python3-dev python3-setuptools python3-wheel python3-cffi libcairo2 libpango-1.0-0 libpangocairo-1.0-0 libgdk-pixbuf2.0-0 libffi-dev shared-mime-info
pip3 install weasyprint

# Install wkhtmltopdf
sudo apt-get install wkhtmltopdf

# Install Calibre
sudo apt-get install calibre

# Install project dependencies
npm install

# Install markdownlint globally
npm install -g markdownlint-cli
```

## Available Build Commands

After setup, you can use these commands:

```bash
npm run build          # Build all formats
npm run build:leanpub  # Build Leanpub format
npm run build:kindle   # Build Kindle format (EPUB + MOBI)
npm run build:web      # Build web format
npm run build:pdf      # Build PDF format
npm run dev            # Start development server
npm run lint           # Lint markdown files
npm run word-count     # Count words in all chapters
npm run validate       # Validate project structure
npm run clean          # Clean build artifacts
```

## Troubleshooting

### PDF Generation Issues
If PDF generation fails:
1. Try `npm run setup` again to reinstall dependencies
2. **Note**: wkhtmltopdf has been discontinued on macOS Homebrew (as of 2024-12-16)
3. Use Puppeteer as an alternative: `npm install puppeteer`
4. Use browser print-to-PDF with the generated HTML files in the `docs/` folder
5. Check the individual setup script: `./scripts/setup-pdf.sh`

### Python "externally-managed-environment" Error
If you get an externally-managed-environment error when installing WeasyPrint:

**Option 1: Use pipx (Recommended)**
```bash
brew install pipx
pipx install weasyprint
```

**Option 2: Use Homebrew**
```bash
brew install weasyprint
```

**Option 3: Use Virtual Environment**
```bash
python3 -m venv ~/.weasyprint-venv
source ~/.weasyprint-venv/bin/activate
pip install weasyprint
deactivate

# Create wrapper script
echo '#!/bin/bash
source "$HOME/.weasyprint-venv/bin/activate"
exec python -m weasyprint "$@"' | sudo tee /usr/local/bin/weasyprint
sudo chmod +x /usr/local/bin/weasyprint
```

**Option 4: Use Puppeteer instead**
```bash
npm install puppeteer
```

### MOBI Generation Issues
MOBI generation requires Calibre. If it fails:
1. Install Calibre manually from https://calibre-ebook.com/
2. Ensure `ebook-convert` is in your PATH
3. EPUB files will still be generated even if MOBI fails

### Permission Issues
If you get permission errors:
```bash
chmod +x scripts/*.sh
```

## Dependencies Overview

| Tool | Purpose | Required |
|------|---------|----------|
| Node.js | JavaScript runtime for build scripts | ✅ Yes |
| npm | Package manager | ✅ Yes |
| Python 3 | Runtime for WeasyPrint | ✅ Yes |
| Pandoc | Document converter | ✅ Yes |
| WeasyPrint | High-quality PDF generation | Recommended |
| Puppeteer | PDF generation via Chrome headless | Recommended |
| wkhtmltopdf | PDF generation (discontinued on macOS) | Optional |
| Calibre | MOBI generation for Kindle | Optional |
| markdownlint-cli | Markdown linting | Optional |

The setup script will attempt to install all dependencies and provide fallbacks where possible.
