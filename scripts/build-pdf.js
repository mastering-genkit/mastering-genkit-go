#!/usr/bin/env node

const fs = require('fs-extra');
const path = require('path');
const yaml = require('yaml');
const { execSync } = require('child_process');
const { getAuthorString, getAuthorArray } = require('./author-utils');

async function buildPDF() {
    console.log('📄 Building PDF format...');
    
    try {
        // Check if pandoc is installed
        try {
            execSync('pandoc --version', { stdio: 'ignore' });
        } catch (error) {
            console.error('❌ Pandoc is not installed. Please install Pandoc first:');
            console.error('   macOS: brew install pandoc');
            console.error('   Ubuntu: sudo apt-get install pandoc');
            console.error('   Windows: Download from https://pandoc.org/installing.html');
            process.exit(1);
        }

        // Ensure output directory exists
        await fs.ensureDir('dist');
        
        // Read book metadata
        const bookMetadata = yaml.parse(
            await fs.readFile('src/metadata/book.yaml', 'utf8')
        );
        
        // Get all chapter files
        const chaptersDir = 'src/chapters';
        const chapterFiles = await fs.readdir(chaptersDir);
        const markdownFiles = chapterFiles
            .filter(file => file.endsWith('.md'))
            .sort();
        
        // Create a temporary combined markdown file
        const tempDir = 'temp';
        await fs.ensureDir(tempDir);
        
        // Copy images to temp directory for pandoc
        if (await fs.pathExists('src/images')) {
            await fs.copy('src/images', path.join(tempDir, 'images'));
            console.log('📷 Images copied to temp directory');
        }
        
        // Combine all chapters into one file
        let combinedContent = '';
        
        // Add title page
        combinedContent += createTitlePage(bookMetadata);
        
        // Add table of contents
        combinedContent += '\n\\newpage\n\\tableofcontents\n\\newpage\n\n';
        
        // Process each chapter
        for (const file of markdownFiles) {
            const filePath = path.join(chaptersDir, file);
            let content = await fs.readFile(filePath, 'utf8');
            
            // Process image paths to work with temp directory structure
            content = processImagePaths(content);
            
            // Add page break before each new chapter (except the first)
            if (combinedContent.includes('# ')) {
                combinedContent += '\n\\newpage\n\n';
            }
            
            combinedContent += content + '\n\n';
        }
        
        // Write combined file
        const combinedFilePath = path.join(tempDir, 'combined.md');
        await fs.writeFile(combinedFilePath, combinedContent);
        
        // Create LaTeX template for better formatting (if using LaTeX engine)
        // await createLatexTemplate(bookMetadata);
        
        // Build PDF using Pandoc - try different engines
        const outputPath = `dist/book.pdf`;
        
        console.log('🔄 Converting to PDF...');
        console.log('⏳ This may take a moment...');
        
        // Try different PDF engines in order of preference
        const engines = ['weasyprint', 'wkhtmltopdf', 'puppeteer', 'prince'];
        let success = false;
        
        for (const engine of engines) {
            try {
                console.log(`🔧 Trying with ${engine}...`);
                
                if (engine === 'weasyprint') {
                    // Convert to HTML first, then to PDF with WeasyPrint
                    const htmlPath = `temp/book.html`;
                    const htmlCommand = [
                        'pandoc',
                        combinedFilePath,
                        '-o', htmlPath,
                        '--standalone',
                        '--toc',
                        '--toc-depth=3',
                        '--number-sections',
                        '--highlight-style=tango',
                        '--css=../config/pdf-style.css',
                        '--resource-path=temp:src'
                    ].join(' ');
                    
                    execSync(htmlCommand, { stdio: 'inherit' });
                    
                    // Create CSS for PDF
                    await createPDFCSS();
                    
                    // Convert HTML to PDF
                    execSync(`weasyprint ${htmlPath} ${outputPath}`, { stdio: 'inherit' });
                    success = true;
                    break;
                    
                } else if (engine === 'puppeteer') {
                    // Convert to HTML first, then to PDF with Puppeteer
                    const htmlPath = `temp/book.html`;
                    const htmlCommand = [
                        'pandoc',
                        combinedFilePath,
                        '-o', htmlPath,
                        '--standalone',
                        '--toc',
                        '--toc-depth=3',
                        '--number-sections',
                        '--highlight-style=tango',
                        '--css=config/pdf-style.css',
                        '--resource-path=temp:src'
                    ].join(' ');
                    
                    execSync(htmlCommand, { stdio: 'inherit' });
                    
                    // Check if Puppeteer is available
                    try {
                        require('puppeteer');
                    } catch (e) {
                        throw new Error('Puppeteer not installed');
                    }
                    
                    // Use Puppeteer to convert HTML to PDF
                    const puppeteerScript = `
                        const puppeteer = require('puppeteer');
                        const path = require('path');
                        
                        (async () => {
                            const browser = await puppeteer.launch();
                            const page = await browser.newPage();
                            await page.goto('file://' + path.resolve('${htmlPath}'), { waitUntil: 'networkidle0' });
                            await page.pdf({
                                path: '${outputPath}',
                                format: 'A4',
                                printBackground: true,
                                margin: {
                                    top: '1in',
                                    right: '1in',
                                    bottom: '1in',
                                    left: '1in'
                                }
                            });
                            await browser.close();
                        })();
                    `;
                    
                    require('fs').writeFileSync('temp/puppeteer-pdf.js', puppeteerScript);
                    execSync('node temp/puppeteer-pdf.js', { stdio: 'inherit' });
                    require('fs').unlinkSync('temp/puppeteer-pdf.js');
                    
                    success = true;
                    break;
                    
                } else {
                    // Direct conversion with pandoc
                    const pandocCommand = [
                        'pandoc',
                        combinedFilePath,
                        '-o', outputPath,
                        `--pdf-engine=${engine}`,
                        '--toc',
                        '--toc-depth=3',
                        '--number-sections',
                        '--highlight-style=tango',
                        '--variable=geometry:margin=1in',
                        '--variable=fontsize:11pt',
                        '--variable=linestretch:1.2',
                        '--resource-path=temp:src'
                    ].join(' ');
                    
                    execSync(pandocCommand, { stdio: 'inherit' });
                    success = true;
                    break;
                }
                
            } catch (error) {
                console.log(`⚠️  ${engine} not available, trying next option...`);
                continue;
            }
        }
        
        if (!success) {
            // Fallback: Create HTML version and inform user
            console.log('⚠️  No PDF engine found. Creating HTML version instead...');
            const htmlPath = `dist/${bookMetadata.title.replace(/\s+/g, '-').toLowerCase()}.html`;
            
            // Create and copy CSS file to dist directory
            await createPDFCSS();
            await fs.copy('config/pdf-style.css', 'dist/pdf-style.css');
            
            // Copy images to dist directory as well
            if (await fs.pathExists('src/images')) {
                await fs.copy('src/images', 'dist/images');
            }
            
            const htmlCommand = [
                'pandoc',
                combinedFilePath,
                '-o', htmlPath,
                '--standalone',
                '--toc',
                '--toc-depth=3',
                '--number-sections',
                '--highlight-style=tango',
                '--css=pdf-style.css',
                '--resource-path=temp:src'
            ].join(' ');
            
            await createPDFCSS();
            execSync(htmlCommand, { stdio: 'inherit' });
            
            console.log('✅ HTML version created successfully!');
            console.log(`📄 HTML file created: ${htmlPath}`);
            console.log('💡 To convert to PDF:');
            console.log('   1. Install WeasyPrint: pip install weasyprint');
            console.log('   2. Install Puppeteer: npm install puppeteer');
            console.log('   3. Run: weasyprint ' + htmlPath + ' ' + outputPath);
            console.log('   Or open the HTML file in a browser and print to PDF');
            console.log('   Note: wkhtmltopdf has been discontinued on macOS Homebrew');
            
            return;
        }
        
        // Clean up temporary files
        await fs.remove(tempDir);
        
        console.log('✅ PDF format built successfully!');
        console.log(`📄 PDF created: ${outputPath}`);
        
        // Display file info
        const stats = await fs.stat(outputPath);
        const fileSizeInKB = Math.round(stats.size / 1024);
        console.log(`📊 File size: ${fileSizeInKB} KB`);
        
    } catch (error) {
        console.error('❌ Error building PDF format:', error.message);
        
        // Clean up on error
        if (await fs.pathExists('temp')) {
            await fs.remove('temp');
        }
        
        process.exit(1);
    }
}

async function createPDFCSS() {
    const css = `
@page {
    margin: 1in;
    size: letter;
}

body {
    font-family: 'Georgia', 'Times New Roman', serif;
    font-size: 11pt;
    line-height: 1.4;
    color: #333;
    max-width: none;
}

h1, h2, h3, h4, h5, h6 {
    font-family: 'Arial', 'Helvetica', sans-serif;
    color: #2c3e50;
    margin-top: 1.5em;
    margin-bottom: 0.5em;
    page-break-after: avoid;
}

h1 {
    font-size: 24pt;
    text-align: center;
    page-break-before: always;
    margin-top: 2em;
}

h2 {
    font-size: 18pt;
    border-bottom: 2px solid #3498db;
    padding-bottom: 0.2em;
}

h3 {
    font-size: 14pt;
    color: #34495e;
}

h4 {
    font-size: 12pt;
    font-style: italic;
}

p {
    margin: 0.8em 0;
    text-align: justify;
    orphans: 2;
    widows: 2;
}

code {
    font-family: 'Courier New', 'Monaco', monospace;
    font-size: 8pt;
    background-color: #f8f8f8;
    padding: 0.1em 0.2em;
    border-radius: 2px;
    word-break: break-all;
}

pre {
    font-family: 'Courier New', 'Monaco', monospace;
    font-size: 8pt;
    background-color: #f8f8f8;
    padding: 0.8em;
    border-radius: 4px;
    border-left: 4px solid #3498db;
    overflow-x: visible;
    white-space: pre-wrap;
    word-wrap: break-word;
    word-break: break-all;
    page-break-inside: auto;
    margin: 1em 0;
    max-width: 100%;
    box-sizing: border-box;
}

blockquote {
    border-left: 4px solid #bdc3c7;
    padding-left: 1em;
    margin: 1em 0;
    font-style: italic;
    color: #555;
}

ul, ol {
    margin: 1em 0;
    padding-left: 2em;
}

li {
    margin: 0.3em 0;
}

table {
    border-collapse: collapse;
    width: 100%;
    margin: 1em 0;
    page-break-inside: avoid;
}

th, td {
    border: 1px solid #ddd;
    padding: 0.5em;
    text-align: left;
}

th {
    background-color: #f2f2f2;
    font-weight: bold;
}

a {
    color: #3498db;
    text-decoration: none;
}

a:hover {
    text-decoration: underline;
}

/* Table of Contents */
#TOC {
    page-break-before: always;
    page-break-after: always;
    margin-bottom: 2em;
}

#TOC h1 {
    text-align: center;
    margin-bottom: 1em;
}

#TOC ul {
    list-style: none;
    padding-left: 0;
}

#TOC ul ul {
    padding-left: 1.5em;
}

#TOC a {
    text-decoration: none;
    color: #2c3e50;
}

#TOC a:hover {
    color: #3498db;
}

/* Print-specific styles */
@media print {
    body {
        font-size: 10pt;
        line-height: 1.3;
    }
    
    h1 {
        page-break-before: always;
        margin-top: 0;
    }
    
    h2, h3, h4, h5, h6 {
        page-break-after: avoid;
    }
    
    pre {
        font-size: 7pt;
        line-height: 1.2;
        padding: 0.6em;
        page-break-inside: auto;
        white-space: pre-wrap;
        word-wrap: break-word;
        word-break: break-all;
        overflow: visible;
        max-width: 100%;
    }
    
    code {
        font-size: 7pt;
        word-break: break-all;
    }
    
    blockquote, table {
        page-break-inside: avoid;
    }
    
    img {
        max-width: 100%;
        page-break-inside: avoid;
    }
}
`;

    await fs.writeFile('config/pdf-style.css', css);
}

function createTitlePage(metadata) {
    const authorArray = getAuthorArray(metadata);
    
    // Get the description with proper indentation for YAML
    const descText = metadata.pdfDescription || metadata.description || '';
    const formattedDesc = descText
        .split('\n')
        .map(line => `  ${line}`)
        .join('\n');
    
    // Format authors as YAML list
    const authorList = authorArray.map(author => `- ${author}`).join('\n');
    
    return `---
title: "${metadata.title}"
subtitle: "${metadata.subtitle || ''}"
author:
${authorList}
abstract: |
${formattedDesc}
abstract-title: "**About this Book**"
description: |
${formattedDesc}
date: "${new Date().toLocaleDateString()}"
documentclass: book
classoption: [12pt, oneside]
geometry: [margin=1in]
fontfamily: libertinus
colorlinks: true
linkcolor: blue
urlcolor: blue
citecolor: blue
---

`;
}

async function createLatexTemplate(metadata) {
    const template = `\\documentclass[$if(fontsize)$$fontsize$,$endif$$if(lang)$$babel-lang$,$endif$$if(papersize)$$papersize$paper,$endif$$for(classoption)$$classoption$$sep$,$endfor$]{$documentclass$}

% Packages
\\usepackage{lmodern}
\\usepackage{amssymb,amsmath}
\\usepackage{ifxetex,ifluatex}
\\usepackage{fixltx2e}
\\usepackage[T1]{fontenc}
\\usepackage[utf8]{inputenc}
\\usepackage{microtype}
\\usepackage[unicode=true]{hyperref}
\\usepackage{graphicx}
\\usepackage{longtable,booktabs}
\\usepackage{listings}
\\usepackage{fancyhdr}
\\usepackage{geometry}
\\usepackage{setspace}
\\usepackage{xcolor}

% Custom styling
\\definecolor{chaptercolor}{RGB}{0, 100, 200}
\\definecolor{codebackground}{RGB}{248, 248, 248}

% Header and footer
\\pagestyle{fancy}
\\fancyhf{}
\\fancyhead[LE,RO]{\\thepage}
\\fancyhead[LO,RE]{$title$}
\\renewcommand{\\headrulewidth}{0.4pt}

% Chapter styling
\\usepackage{titlesec}
\\titleformat{\\chapter}[display]
{\\normalfont\\huge\\bfseries\\color{chaptercolor}}
{\\chaptertitlename\\ \\thechapter}{20pt}{\\Huge}

% Code block styling
\\lstset{
    backgroundcolor=\\color{codebackground},
    basicstyle=\\ttfamily\\small,
    breaklines=true,
    frame=single,
    framesep=5pt,
    rulecolor=\\color{gray},
    showspaces=false,
    showstringspaces=false
}

% Metadata
$if(title)$
\\title{$title$$if(thanks)$\\thanks{$thanks$}$endif$}
$endif$
$if(subtitle)$
\\providecommand{\\subtitle}[1]{}
\\subtitle{$subtitle$}
$endif$
$if(author)$
\\author{$for(author)$$author$$sep$ \\and $endfor$}
$endif$
$if(abstract)$
% Define abstract text with proper line breaks
\\def\\abstracttext{
  \\begin{minipage}{\\linewidth}
    \\setlength{\\parindent}{0pt}
    \\setlength{\\parskip}{0.8em}
    $for(abstract)$
    $abstract$
    
    $endfor$
  \\end{minipage}
}
$endif$
$if(date)$
\\date{$date$}
$endif$

% Geometry
$if(geometry)$
\\usepackage[$for(geometry)$$geometry$$sep$,$endfor$]{geometry}
$endif$

% Line spacing
$if(linestretch)$
\\usepackage{setspace}
\\setstretch{$linestretch$}
$endif$

\\begin{document}

% Title page
$if(title)$
\\maketitle
$endif$

% Abstract (after title page, no heading)
$if(abstract)$
\\vspace{1em}
\\noindent\\abstracttext
\\vspace{2em}
$endif$

% Table of contents
$if(toc)$
{
\\hypersetup{linkcolor=$if(toccolor)$$toccolor$$else$black$endif$}
\\setcounter{tocdepth}{$toc-depth$}
\\tableofcontents
}
$endif$

% Main content
$body$

\\end{document}`;

    await fs.writeFile('temp/template.latex', template);
}

function processImagePaths(content) {
    // Convert relative image paths to work with the temp directory structure
    // This handles paths like: ![alt](../images/image.jpg) or ![alt](images/image.jpg)
    return content.replace(
        /!\[([^\]]*)\]\((?:\.\.\/)?(?:src\/)?images\/([^)]+)\)/g,
        '![$1](images/$2)'
    );
}

// Run the script
if (require.main === module) {
    buildPDF();
}

module.exports = buildPDF;
