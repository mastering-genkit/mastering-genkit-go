#!/usr/bin/env node

const fs = require('fs-extra');
const path = require('path');

async function countWords() {
    console.log('📊 Counting words...');
    
    try {
        const chaptersDir = 'src/chapters';
        
        if (!await fs.pathExists(chaptersDir)) {
            console.log('No chapters directory found.');
            return;
        }
        
        const chapterFiles = await fs.readdir(chaptersDir);
        const markdownFiles = chapterFiles.filter(file => file.endsWith('.md'));
        
        let totalWords = 0;
        const chapterStats = [];
        
        for (const file of markdownFiles.sort()) {
            const filePath = path.join(chaptersDir, file);
            const content = await fs.readFile(filePath, 'utf8');
            
            // Remove markdown syntax and count words
            const plainText = content
                .replace(/^#{1,6}\s+/gm, '') // Remove headers
                .replace(/\*\*(.*?)\*\*/g, '$1') // Remove bold
                .replace(/\*(.*?)\*/g, '$1') // Remove italic
                .replace(/`(.*?)`/g, '$1') // Remove inline code
                .replace(/```[\s\S]*?```/g, '') // Remove code blocks
                .replace(/\[([^\]]+)\]\([^)]+\)/g, '$1') // Remove links
                .replace(/!\[([^\]]*)\]\([^)]+\)/g, '') // Remove images
                .replace(/^\s*[-*+]\s+/gm, '') // Remove list markers
                .replace(/^\s*\d+\.\s+/gm, '') // Remove numbered list markers
                .replace(/^\s*>\s+/gm, '') // Remove blockquotes
                .trim();
            
            const words = plainText
                .split(/\s+/)
                .filter(word => word.length > 0).length;
            
            totalWords += words;
            chapterStats.push({
                file,
                words,
                title: extractTitle(content)
            });
        }
        
        // Display results
        console.log('\n📈 Word Count Report');
        console.log('='.repeat(50));
        
        chapterStats.forEach((chapter, index) => {
            console.log(`${(index + 1).toString().padStart(2)}. ${chapter.title.padEnd(30)} ${chapter.words.toString().padStart(6)} words`);
        });
        
        console.log('='.repeat(50));
        console.log(`Total: ${totalWords.toLocaleString()} words`);
        
        // Estimate reading time (average 250 words per minute)
        const readingTimeMinutes = Math.ceil(totalWords / 250);
        const hours = Math.floor(readingTimeMinutes / 60);
        const minutes = readingTimeMinutes % 60;
        
        console.log(`Estimated reading time: ${hours > 0 ? `${hours}h ` : ''}${minutes}m`);
        
        // Estimate pages (average 250 words per page)
        const estimatedPages = Math.ceil(totalWords / 250);
        console.log(`Estimated pages: ${estimatedPages}`);
        
    } catch (error) {
        console.error('❌ Error counting words:', error.message);
        process.exit(1);
    }
}

function extractTitle(content) {
    const titleMatch = content.match(/^#\s+(.+)$/m);
    return titleMatch ? titleMatch[1] : 'Untitled';
}

// Run the script
if (require.main === module) {
    countWords();
}

module.exports = countWords;
