# Detailed Setup Instructions

This document provides comprehensive setup instructions for all dependencies needed to build the book in multiple formats. For a quick overview, see the [README.md](README.md).

## Automated Setup (Recommended)

Run the automated setup script that detects your operating system:

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

## Common Troubleshooting Issues

### PDF Generation Issues
If PDF generation fails:
1. **Missing WeasyPrint**: Try `npm run setup` again to reinstall dependencies
2. **Alternative options**: Use Puppeteer: `npm install puppeteer`
3. **Browser fallback**: Use browser print-to-PDF with generated HTML files in `docs/` folder
4. **Individual setup**: Check the setup script: `./scripts/setup-pdf.sh`

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
MOBI generation requires Calibre:
1. Install Calibre manually from https://calibre-ebook.com/
2. Ensure `ebook-convert` is in your PATH
3. EPUB files will still be generated even if MOBI fails

### Permission Issues
If you get permission errors:
```bash
chmod +x scripts/*.sh
```

### Node.js Version Issues
Ensure you're using a supported Node.js version:
```bash
node --version  # Should be v16 or higher
```

## Dependencies Reference

| Tool | Purpose | Installation | Required |
|------|---------|-------------|----------|
| Node.js | JavaScript runtime for build scripts | [nodejs.org](https://nodejs.org/) | ✅ Yes |
| npm | Package manager | Included with Node.js | ✅ Yes |
| Python 3 | Runtime for WeasyPrint | [python.org](https://python.org/) | ✅ Yes |
| Pandoc | Document converter | [pandoc.org](https://pandoc.org/) | ✅ Yes |
| WeasyPrint | High-quality PDF generation | `pip install weasyprint` | Recommended |
| Puppeteer | PDF generation via Chrome headless | `npm install puppeteer` | Recommended |
| Calibre | MOBI generation for Kindle | [calibre-ebook.com](https://calibre-ebook.com/) | Optional |
| markdownlint-cli | Markdown linting | `npm install -g markdownlint-cli` | Optional |

The setup script attempts to install all dependencies and provides fallbacks where possible.

## Getting Help

If you continue having setup issues:

1. Check the [README.md](README.md) troubleshooting section
2. Run `npm run validate` to check your project structure
3. Ensure all required tools are in your PATH
4. Try the manual installation steps for your platform
5. Check individual build scripts in the `scripts/` directory

For more information, see the main [README.md](README.md) documentation.
