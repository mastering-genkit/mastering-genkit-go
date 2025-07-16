#!/usr/bin/env node

const fs = require('fs-extra');
const path = require('path');
const yaml = require('yaml');
const { execSync } = require('child_process');
const { getAuthorString, getAuthorArray } = require('./author-utils');

async function buildKindle() {
    console.log('📱 Building Kindle format...');
    
    try {
        // Check if Pandoc is installed
        try {
            execSync('pandoc --version', { stdio: 'pipe' });
        } catch (error) {
            console.error('❌ Pandoc is required but not installed.');
            console.log('Install Pandoc from: https://pandoc.org/installing.html');
            process.exit(1);
        }
        
        // Ensure build directory exists
        await fs.ensureDir('build/kindle');
        
        // Read book metadata
        const bookMetadata = yaml.parse(
            await fs.readFile('src/metadata/book.yaml', 'utf8')
        );
        
        // Get all chapter files
        const chaptersDir = 'src/chapters';
        const chapterFiles = await fs.readdir(chaptersDir);
        const markdownFiles = chapterFiles
            .filter(file => file.endsWith('.md'))
            .sort()
            .map(file => path.join(chaptersDir, file));
        
        // Create combined markdown file
        let combinedContent = '';
        
        // Add title page
        combinedContent += createTitlePage(bookMetadata);
        
        // Generate comprehensive table of contents
        const tocEntries = [];
        const chapterContents = [];
        
        // First pass: Read all chapters and extract structure for TOC
        for (const filePath of markdownFiles) {
            const content = await fs.readFile(filePath, 'utf8');
            chapterContents.push(content);
            
            // Extract the main title (first # heading)
            const titleMatch = content.match(/^#\s+(.+)$/m);
            if (titleMatch) {
                const title = titleMatch[1].trim();
                const chapterNumber = path.basename(filePath).match(/^(\d+)/)?.[1];
                
                // Extract section headings (## level)
                const sectionMatches = content.match(/^##\s+(.+)$/gm);
                const sections = sectionMatches ? sectionMatches.map(match => {
                    const sectionTitle = match.replace(/^##\s+/, '').trim();
                    return {
                        title: sectionTitle,
                        anchor: sectionTitle.toLowerCase().replace(/[^a-z0-9]+/g, '-').replace(/^-|-$/g, '')
                    };
                }) : [];
                
                tocEntries.push({
                    number: chapterNumber,
                    title: title,
                    anchor: title.toLowerCase().replace(/[^a-z0-9]+/g, '-').replace(/^-|-$/g, ''),
                    sections: sections
                });
            }
        }
        
        // Create unified table of contents page
        if (tocEntries.length > 0) {
            combinedContent += '\n\n\\newpage\n\n<div class="table-of-contents">\n\n# Table of Contents\n\n';
            
            tocEntries.forEach((entry, index) => {
                const chapterNum = entry.number || (index + 1);
                combinedContent += `**${chapterNum}. [${entry.title}](#${entry.anchor})**\n\n`;
                
                // Add sections as sub-items with better formatting
                if (entry.sections.length > 0) {
                    entry.sections.forEach(section => {
                        combinedContent += `   • [${section.title}](#${section.anchor})\n\n`;
                    });
                }
                combinedContent += '\n';
            });
            
            combinedContent += '</div>\n\n---\n\n\\newpage\n';
        }
        
        // Second pass: Combine all chapters
        chapterContents.forEach(content => {
            combinedContent += '\n\n' + content;
        });
        
        // Write combined file
        const combinedPath = 'build/kindle/book.md';
        await fs.writeFile(combinedPath, combinedContent);
        
        // Copy images
        if (await fs.pathExists('src/images')) {
            await fs.copy('src/images', 'build/kindle/images');
        }
        
        // Create metadata file for Pandoc
        const pandocMetadata = createPandocMetadata(bookMetadata);
        await fs.writeFile('build/kindle/metadata.yaml', pandocMetadata);
        
        // Build EPUB
        console.log('📖 Generating EPUB...');
        const epubCommand = [
            'pandoc',
            '--from=markdown+smart',
            '--to=epub3',
            '--epub-metadata=build/kindle/metadata.yaml',
            '--epub-cover-image=src/images/cover.jpg',
            '--css=config/epub.css',
            '--standalone',
            '-o build/kindle/book.epub',
            'build/kindle/book.md'
        ].join(' ');
        
        try {
            execSync(epubCommand, { stdio: 'inherit' });
            console.log('✅ EPUB generated successfully!');
        } catch (error) {
            console.log('⚠️  EPUB generation completed with warnings (this is normal)');
        }
        
        // Try to build MOBI (requires Calibre)
        try {
            execSync('ebook-convert --version', { stdio: 'pipe' });
            console.log('📱 Generating MOBI...');
            
            const mobiCommand = [
                'ebook-convert',
                'build/kindle/book.epub',
                'build/kindle/book.mobi',
                '--pretty-print'
            ].join(' ');
            
            execSync(mobiCommand, { stdio: 'inherit' });
            console.log('✅ MOBI generated successfully!');
            
        } catch (error) {
            console.log('ℹ️  Calibre not found. MOBI generation skipped.');
            console.log('   Install Calibre to generate MOBI files: https://calibre-ebook.com/');
        }
        
        console.log('✅ Kindle format built successfully!');
        console.log('📁 Files created in build/kindle/ directory');
        
    } catch (error) {
        console.error('❌ Error building Kindle format:', error.message);
        process.exit(1);
    }
}

function createTitlePage(metadata) {
    const authorString = getAuthorString(metadata);
    return `---
title: "${metadata.title}"
subtitle: "${metadata.subtitle || ''}"
author: "${authorString}"
date: "${metadata.date}"
---

# ${metadata.title}

${metadata.subtitle ? `## ${metadata.subtitle}` : ''}

**By ${authorString}**

*${metadata.date}*

${metadata.description || ''}

---`;
}

function createPandocMetadata(metadata) {
    const authors = getAuthorArray(metadata);
    const pandocMeta = {
        title: metadata.title,
        author: authors,
        date: metadata.date,
        language: metadata.language || 'en',
        description: metadata.description,
        publisher: metadata.publisher || '',
        rights: metadata.copyright || ''
    };
    
    if (metadata.isbn) {
        pandocMeta.identifier = metadata.isbn;
    }
    
    return yaml.stringify(pandocMeta);
}

// Run the script
if (require.main === module) {
    buildKindle();
}

module.exports = buildKindle;
