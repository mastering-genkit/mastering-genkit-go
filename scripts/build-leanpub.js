#!/usr/bin/env node

const fs = require('fs-extra');
const path = require('path');
const yaml = require('yaml');
const { generateAboutAuthors } = require('./author-utils');

async function buildLeanpub() {
    console.log('🚀 Building Leanpub format...');
    
    try {
        // Ensure manuscript directory exists
        await fs.ensureDir('manuscript');
        
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
        
        // Create Book.txt (defines chapter order for Leanpub)
        const bookTxt = markdownFiles.join('\n');
        await fs.writeFile('manuscript/Book.txt', bookTxt);
        
        // Create Sample.txt (sample chapters for preview)
        const sampleChapters = markdownFiles.slice(0, 2); // First 2 chapters
        const sampleTxt = sampleChapters.join('\n');
        await fs.writeFile('manuscript/Sample.txt', sampleTxt);
        
        // Copy all markdown files to manuscript directory
        for (const file of markdownFiles) {
            const sourcePath = path.join(chaptersDir, file);
            const destPath = path.join('manuscript', file);
            await fs.copy(sourcePath, destPath);
        }
        
        // Copy images directory if it exists
        if (await fs.pathExists('src/images')) {
            await fs.copy('src/images', 'manuscript/images');
        }
        
        // Create Leanpub-specific files
        await createLeanpubFiles(bookMetadata);
        
        console.log('✅ Leanpub format built successfully!');
        console.log('📁 Files created in manuscript/ directory');
        
    } catch (error) {
        console.error('❌ Error building Leanpub format:', error.message);
        process.exit(1);
    }
}

async function createLeanpubFiles(metadata) {
    // Create subtitle.txt if subtitle exists
    if (metadata.subtitle) {
        await fs.writeFile('manuscript/subtitle.txt', metadata.subtitle);
    }
    
    // Create dedication.txt template
    const dedication = `{frontmatter}

# Dedication

This book is dedicated to...

{mainmatter}`;
    await fs.writeFile('manuscript/dedication.txt', dedication);
    
    // Create about-author.txt template with support for multiple authors
    const aboutAuthor = `{backmatter}

${generateAboutAuthors(metadata)}`;
    await fs.writeFile('manuscript/about-author.txt', aboutAuthor);
    
    // Update Book.txt to include front and back matter
    const currentBookTxt = await fs.readFile('manuscript/Book.txt', 'utf8');
    const chaptersArray = currentBookTxt.split('\n').filter(line => line.trim());
    
    const fullBookTxt = [
        'dedication.txt',
        ...chaptersArray,
        'about-author.txt'
    ].join('\n');
    
    await fs.writeFile('manuscript/Book.txt', fullBookTxt);
}

// Run the script
if (require.main === module) {
    buildLeanpub();
}

module.exports = buildLeanpub;
