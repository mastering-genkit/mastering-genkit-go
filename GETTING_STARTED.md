# Getting Started Guide

This is a quick reference guide for getting your book project up and running. For complete documentation, see the [README.md](README.md).

## Quick Setup

1. **Install dependencies:**
   ```bash
   npm install
   npm run setup  # Install all external tools
   ```

2. **Configure your book:**
   Edit `src/metadata/book.yaml` with your book details

3. **Start writing:**
   Create chapters in `src/chapters/` as numbered Markdown files

4. **Build and preview:**
   ```bash
   npm run dev  # Start development server
   ```

## Essential Commands

```bash
npm run build          # Build all formats
npm run dev            # Development server with auto-rebuild
npm run validate       # Check structure and content
npm run word-count     # Get writing statistics
```

## Writing Your First Chapter

1. Create `src/chapters/01-introduction.md`
2. Start with a level 1 header: `# Introduction`
3. Add your content using Markdown
4. Run `npm run dev` to preview

## Next Steps

- Read the complete [README.md](README.md) for detailed documentation
- Check [SETUP.md](SETUP.md) for troubleshooting installation issues
- Review the `src/metadata/book.yaml` file for configuration options

For questions or issues, refer to the troubleshooting section in the README.
