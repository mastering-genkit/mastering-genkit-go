#!/usr/bin/env node

const fs = require('fs-extra');
const path = require('path');
const yaml = require('yaml');

async function validate() {
    console.log('🔍 Validating book structure...');
    
    let errors = [];
    let warnings = [];
    
    try {
        // Check required directories
        const requiredDirs = [
            'src',
            'src/chapters',
            'src/metadata',
            'src/images'
        ];
        
        for (const dir of requiredDirs) {
            if (!await fs.pathExists(dir)) {
                errors.push(`Missing required directory: ${dir}`);
            }
        }
        
        // Check metadata file
        const metadataPath = 'src/metadata/book.yaml';
        if (!await fs.pathExists(metadataPath)) {
            errors.push('Missing book metadata file: src/metadata/book.yaml');
        } else {
            await validateMetadata(metadataPath, errors, warnings);
        }
        
        // Check chapters
        await validateChapters(errors, warnings);
        
        // Check images
        await validateImages(errors, warnings);
        
        // Display results
        console.log('\n📋 Validation Report');
        console.log('='.repeat(50));
        
        if (errors.length === 0 && warnings.length === 0) {
            console.log('✅ All checks passed! Your book structure is valid.');
        } else {
            if (errors.length > 0) {
                console.log('❌ Errors found:');
                errors.forEach(error => console.log(`   • ${error}`));
            }
            
            if (warnings.length > 0) {
                console.log(warnings.length > 0 && errors.length > 0 ? '\n⚠️  Warnings:' : '⚠️  Warnings:');
                warnings.forEach(warning => console.log(`   • ${warning}`));
            }
        }
        
        if (errors.length > 0) {
            console.log('\nPlease fix the errors before building your book.');
            process.exit(1);
        }
        
    } catch (error) {
        console.error('❌ Validation failed:', error.message);
        process.exit(1);
    }
}

async function validateMetadata(metadataPath, errors, warnings) {
    try {
        const content = await fs.readFile(metadataPath, 'utf8');
        const metadata = yaml.parse(content);
        
        // Required fields
        const requiredFields = ['title', 'description', 'date'];
        for (const field of requiredFields) {
            if (!metadata[field]) {
                errors.push(`Missing required metadata field: ${field}`);
            }
        }
        
        // Author validation - must have either 'author' or 'authors'
        if (!metadata.author && (!metadata.authors || !Array.isArray(metadata.authors) || metadata.authors.length === 0)) {
            errors.push('Missing author information: must have either "author" field or "authors" array');
        }
        
        // Validate authors array if present
        if (metadata.authors && Array.isArray(metadata.authors)) {
            metadata.authors.forEach((author, index) => {
                if (!author.name) {
                    errors.push(`Author at index ${index} is missing "name" field`);
                }
                if (typeof author.name !== 'string') {
                    errors.push(`Author at index ${index} "name" field must be a string`);
                }
            });
        }
        
        // Optional but recommended fields
        const recommendedFields = ['subtitle', 'language', 'version'];
        for (const field of recommendedFields) {
            if (!metadata[field]) {
                warnings.push(`Missing recommended metadata field: ${field}`);
            }
        }
        
        // Check cover image if specified
        if (metadata.kindle?.cover_image) {
            const coverPath = path.resolve('src', metadata.kindle.cover_image.replace('../', ''));
            if (!await fs.pathExists(coverPath)) {
                errors.push(`Cover image not found: ${coverPath}`);
            }
        } else {
            warnings.push('No cover image specified in metadata');
        }
        
    } catch (parseError) {
        errors.push(`Invalid YAML in metadata file: ${parseError.message}`);
    }
}

async function validateChapters(errors, warnings) {
    const chaptersDir = 'src/chapters';
    
    if (!await fs.pathExists(chaptersDir)) {
        return; // Already reported as missing directory
    }
    
    const files = await fs.readdir(chaptersDir);
    const markdownFiles = files.filter(file => file.endsWith('.md')).sort();
    
    if (markdownFiles.length === 0) {
        errors.push('No chapter files found in src/chapters/');
        return;
    }
    
    // Check chapter naming convention
    const properNaming = markdownFiles.every(file => /^\d{2}-/.test(file));
    if (!properNaming) {
        warnings.push('Chapter files should follow naming convention: 01-chapter-name.md');
    }
    
    // Validate each chapter
    for (const file of markdownFiles) {
        const filePath = path.join(chaptersDir, file);
        const content = await fs.readFile(filePath, 'utf8');
        
        // Check for title
        if (!content.match(/^#\s+.+$/m)) {
            warnings.push(`Chapter ${file} missing main title (# heading)`);
        }
        
        // Check for very short chapters
        const wordCount = content.split(/\s+/).length;
        if (wordCount < 100) {
            warnings.push(`Chapter ${file} is very short (${wordCount} words)`);
        }
        
        // Check for broken image references
        const imageRefs = content.match(/!\[.*?\]\(([^)]+)\)/g);
        if (imageRefs) {
            for (const ref of imageRefs) {
                const imagePath = ref.match(/!\[.*?\]\(([^)]+)\)/)[1];
                if (imagePath.startsWith('../images/')) {
                    const fullImagePath = path.join('src', imagePath.replace('../', ''));
                    if (!await fs.pathExists(fullImagePath)) {
                        errors.push(`Broken image reference in ${file}: ${imagePath}`);
                    }
                }
            }
        }
    }
    
    console.log(`📚 Found ${markdownFiles.length} chapter(s)`);
}

async function validateImages(errors, warnings) {
    const imagesDir = 'src/images';
    
    if (!await fs.pathExists(imagesDir)) {
        return; // Already reported as missing directory
    }
    
    const files = await fs.readdir(imagesDir);
    const imageFiles = files.filter(file => 
        /\.(png|jpg|jpeg|gif|svg|webp)$/i.test(file)
    );
    
    if (imageFiles.length === 0) {
        warnings.push('No image files found in src/images/');
    } else {
        console.log(`🖼️  Found ${imageFiles.length} image(s)`);
        
        // Check for cover image
        const hasCover = imageFiles.some(file => 
            file.toLowerCase().includes('cover')
        );
        
        if (!hasCover) {
            warnings.push('No cover image found (should contain "cover" in filename)');
        }
    }
    
    // Check for large images
    for (const file of imageFiles) {
        const filePath = path.join(imagesDir, file);
        const stats = await fs.stat(filePath);
        const sizeMB = stats.size / (1024 * 1024);
        
        if (sizeMB > 5) {
            warnings.push(`Large image file: ${file} (${sizeMB.toFixed(1)}MB)`);
        }
    }
}

// Run the script
if (require.main === module) {
    validate();
}

module.exports = validate;
