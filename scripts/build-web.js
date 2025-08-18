#!/usr/bin/env node

const fs = require('fs-extra');
const path = require('path');
const yaml = require('yaml');
const MarkdownIt = require('markdown-it');
const markdownItAnchor = require('markdown-it-anchor');
const markdownItTOC = require('markdown-it-table-of-contents');
const markdownItPrism = require('markdown-it-prism');
const Handlebars = require('handlebars');
const { getAuthorString } = require('./author-utils');

function registerHandlebarsHelpers() {
    // Helper for equality comparison
    Handlebars.registerHelper('eq', function(a, b) {
        return a === b;
    });
    
    // Helper for greater than comparison
    Handlebars.registerHelper('gt', function(a, b) {
        return a > b;
    });
    
    // Helper for less than comparison
    Handlebars.registerHelper('lt', function(a, b) {
        return a < b;
    });
    
    // Helper for addition
    Handlebars.registerHelper('add', function(a, b) {
        return parseInt(a) + parseInt(b);
    });
}

async function buildWeb() {
    console.log('🌐 Building web format...');
    
    try {
        // Ensure docs directory exists (GitHub Pages)
        await fs.ensureDir('docs');
        await fs.ensureDir('docs/chapters');
        await fs.ensureDir('docs/assets');
        
        // Read book metadata
        const bookMetadata = yaml.parse(
            await fs.readFile('src/metadata/book.yaml', 'utf8')
        );
        
        // Process authors for display
        const { getAuthorObjects } = require('./author-utils');
        bookMetadata.author = getAuthorString(bookMetadata);
        bookMetadata.authorObjects = getAuthorObjects(bookMetadata);
        
        // Setup Markdown processor
        const md = new MarkdownIt({
            html: true,
            linkify: true,
            typographer: true
        });
        
        md.use(markdownItAnchor, {
            permalink: markdownItAnchor.permalink.headerLink()
        });
        
        md.use(markdownItTOC, {
            includeLevel: [2, 3, 4]
        });
        
        // Add syntax highlighting with Prism
        md.use(markdownItPrism, {
            defaultLanguageForUnknown: 'js',
            defaultLanguageForUnspecified: 'js'
        });
        
        // Register Handlebars helpers
        registerHandlebarsHelpers();
        
        // Load templates
        const templates = await loadTemplates();
        
        // Get all chapter files
        const chaptersDir = 'src/chapters';
        const chapterFiles = await fs.readdir(chaptersDir);
        const markdownFiles = chapterFiles
            .filter(file => file.endsWith('.md'))
            .sort();
        
        // Process chapters
        const chapters = [];
        for (const file of markdownFiles) {
            const chapter = await processChapter(file, md, templates, bookMetadata);
            chapters.push(chapter);
        }
        
        // Generate index page
        await generateIndexPage(chapters, templates, bookMetadata);
        
        // Generate individual chapter pages
        for (const chapter of chapters) {
            await generateChapterPage(chapter, chapters, templates, bookMetadata);
        }
        
        // Copy assets
        await copyAssets();
        
        // Generate CSS
        await generateCSS();
        
        // Generate sitemap and other SEO files
        await generateSEOFiles(chapters, bookMetadata);
        
        console.log('✅ Web format built successfully!');
        console.log('📁 Files created in docs/ directory');
        console.log('🌐 Ready for GitHub Pages deployment');
        
    } catch (error) {
        console.error('❌ Error building web format:', error.message);
        process.exit(1);
    }
}

async function loadTemplates() {
    const templateDir = 'config/templates';
    await fs.ensureDir(templateDir);
    
    // Create default templates if they don't exist
    const defaultLayouts = {
        'layout.hbs': await createDefaultLayout(),
        'index.hbs': await createDefaultIndex(),
        'chapter.hbs': await createDefaultChapter()
    };
    
    const templates = {};
    
    for (const [filename, defaultContent] of Object.entries(defaultLayouts)) {
        const templatePath = path.join(templateDir, filename);
        
        if (!await fs.pathExists(templatePath)) {
            await fs.writeFile(templatePath, defaultContent);
        }
        
        const templateContent = await fs.readFile(templatePath, 'utf8');
        const templateName = filename.replace('.hbs', '');
        templates[templateName] = Handlebars.compile(templateContent);
    }
    
    return templates;
}

async function processChapter(filename, md, templates, bookMetadata) {
    const filePath = path.join('src/chapters', filename);
    const content = await fs.readFile(filePath, 'utf8');
    
    // Extract title from first heading
    const titleMatch = content.match(/^#\s+(.+)$/m);
    const title = titleMatch ? titleMatch[1] : filename.replace('.md', '');
    
    // Generate HTML
    let html = md.render(content);
    
    // Post-process HTML to add data-language attributes to code blocks
    html = addLanguageLabels(html);
    
    // Create chapter object
    const chapter = {
        filename: filename.replace('.md', '.html'),
        slug: filename.replace('.md', ''),
        title,
        content: html,
        rawContent: content
    };
    
    return chapter;
}

function addLanguageLabels(html) {
    // Add data-language attributes to pre elements with language classes
    return html.replace(/<pre([^>]*class="[^"]*language-([^"\s]+)[^"]*"[^>]*)>/g, (match, attributes, language) => {
        // Only add data-language if it doesn't already exist
        if (match.includes('data-language=')) {
            return match;
        }
        
        // Capitalize the language name for display
        const displayLanguage = language.charAt(0).toUpperCase() + language.slice(1);
        
        // Insert data-language attribute before the closing >
        return `<pre${attributes} data-language="${displayLanguage}">`;
    });
}

async function generateIndexPage(chapters, templates, bookMetadata) {
    const html = templates.index({
        book: bookMetadata,
        chapters,
        title: bookMetadata.title
    });
    
    const finalHtml = templates.layout({
        title: bookMetadata.title,
        content: html,
        book: bookMetadata,
        chapters,
        isHome: true
    });
    
    await fs.writeFile('docs/index.html', finalHtml);
}

async function generateChapterPage(chapter, allChapters, templates, bookMetadata) {
    // Find previous and next chapters
    const currentIndex = allChapters.findIndex(c => c.slug === chapter.slug);
    const prevChapter = currentIndex > 0 ? allChapters[currentIndex - 1] : null;
    const nextChapter = currentIndex < allChapters.length - 1 ? allChapters[currentIndex + 1] : null;
    
    const html = templates.chapter({
        chapter,
        prevChapter,
        nextChapter,
        book: bookMetadata
    });
    
    const finalHtml = templates.layout({
        title: `${chapter.title} - ${bookMetadata.title}`,
        content: html,
        book: bookMetadata,
        chapters: allChapters,
        chapter
    });
    
    await fs.writeFile(`docs/chapters/${chapter.filename}`, finalHtml);
}

async function copyAssets() {
    // Copy images
    if (await fs.pathExists('src/images')) {
        await fs.copy('src/images', 'docs/images');
    }
    
    // Copy any additional assets
    if (await fs.pathExists('config/assets')) {
        await fs.copy('config/assets', 'docs/assets');
    }
    
    // Copy enhanced JavaScript features
    const enhancedJSPath = 'config/enhanced-features.js';
    if (await fs.pathExists(enhancedJSPath)) {
        await fs.copy(enhancedJSPath, 'docs/assets/features.js');
    }
}

async function generateCSS() {
    // Copy the enhanced CSS file
    const enhancedCSSPath = 'config/enhanced-style.css';
    
    if (await fs.pathExists(enhancedCSSPath)) {
        await fs.copy(enhancedCSSPath, 'docs/assets/style.css');
    } else {
        // Fallback to basic CSS if enhanced version doesn't exist
        const basicCSS = `/* Basic fallback styles */
        body { font-family: system-ui, sans-serif; margin: 0; padding: 20px; }
        .container { max-width: 800px; margin: 0 auto; }
        header { background: #333; color: white; padding: 1rem; }
        nav { background: #666; padding: 0.5rem; }
        main { padding: 2rem 0; }`;
        
        await fs.writeFile('docs/assets/style.css', basicCSS);
    }
}

async function generateSEOFiles(chapters, bookMetadata) {
    // Generate sitemap.xml
    const baseUrl = bookMetadata.web?.base_url || 'https://example.com';
    const sitemap = `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
    <url>
        <loc>${baseUrl}/</loc>
        <changefreq>weekly</changefreq>
        <priority>1.0</priority>
    </url>
    ${chapters.map(chapter => `
    <url>
        <loc>${baseUrl}/chapters/${chapter.filename}</loc>
        <changefreq>monthly</changefreq>
        <priority>0.8</priority>
    </url>`).join('')}
</urlset>`;
    
    await fs.writeFile('docs/sitemap.xml', sitemap);
    
    // Generate robots.txt
    const robots = `User-agent: *
Allow: /

Sitemap: ${baseUrl}/sitemap.xml`;
    
    await fs.writeFile('docs/robots.txt', robots);
}

async function createDefaultLayout() {
    return `<!DOCTYPE html>
<html lang="en" data-theme="light">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{title}}</title>
    <meta name="description" content="{{book.description}}">
    <meta name="author" content="{{book.author}}">
    <meta name="keywords" content="book, {{book.title}}, {{book.author}}">
    <meta property="og:title" content="{{title}}">
    <meta property="og:description" content="{{book.description}}">
    <meta property="og:type" content="book">
    <link rel="stylesheet" href="/assets/style.css">
    <link rel="canonical" href="{{#if chapter}}{{book.web.base_url}}/chapters/{{chapter.filename}}{{else}}{{book.web.base_url}}/{{/if}}">
</head>
<body>
    <div class="reading-progress"></div>
    
    <header>
        <div class="container">
            <h1><a href="/">{{book.title}}</a></h1>
            {{#if book.subtitle}}<p>{{book.subtitle}}</p>{{/if}}
        </div>
    </header>
    
    <nav>
        <div class="container">
            <a href="/" class="nav-brand">📚 {{book.title}}</a>
            <ul>
                {{#each chapters}}
                <li><a href="/chapters/{{filename}}" {{#if (eq ../chapter.slug slug)}}class="active"{{/if}}>{{title}}</a></li>
                {{/each}}
            </ul>
            <button class="theme-toggle" onclick="toggleTheme()" aria-label="Toggle dark mode"></button>
        </div>
    </nav>
    
    <main>
        <div class="container fade-in">
            {{{content}}}
        </div>
    </main>
    
    <footer>
        <div class="container">
            <p>&copy; {{book.date}} {{book.author}}. All rights reserved.</p>
            <div class="social-links">
                {{#if book.social}}
                    {{#if book.social.github}}<a href="{{book.social.github}}" title="GitHub Repository"><i class="icon">📦</i> GitHub</a>{{/if}}
                    {{#if book.social.twitter}}<a href="https://twitter.com/{{book.social.twitter}}" title="Twitter"><i class="icon">🐦</i> Twitter</a>{{/if}}
                    {{#if book.social.linkedin}}<a href="{{book.social.linkedin}}" title="LinkedIn"><i class="icon">👔</i> LinkedIn</a>{{/if}}
                    {{#if book.social.website}}<a href="{{book.social.website}}" title="Website"><i class="icon">🌐</i> Website</a>{{/if}}
                {{/if}}
                
                {{#each book.authorObjects}}
                    {{#if this.email}}<a href="mailto:{{this.email}}" title="Email {{this.name}}"><i class="icon">📧</i> Email</a>{{/if}}
                    {{#if this.twitter}}<a href="https://twitter.com/{{this.twitter}}" title="Twitter {{this.name}}"><i class="icon">🐦</i> Twitter</a>{{/if}}
                    {{#if this.website}}<a href="{{this.website}}" title="Website {{this.name}}"><i class="icon">🌐</i> Website</a>{{/if}}
                {{/each}}
            </div>
        </div>
    </footer>
    
    <script src="/assets/features.js"></script>
</body>
</html>`;
}

async function createDefaultIndex() {
    return `<div class="home-content">
    <h1>Welcome to {{book.title}}</h1>
    
    {{#if book.description}}
    <div class="book-description">
        <p>{{book.description}}</p>
    </div>
    {{/if}}
    
    <h2>Table of Contents</h2>
    <ol class="chapter-list">
        {{#each chapters}}
        <li><a href="/chapters/{{filename}}">{{title}}</a></li>
        {{/each}}
    </ol>
    
    <div class="book-info">
        <p><strong>Author{{#if book.authorObjects}}{{#gt book.authorObjects.length 1}}s{{/gt}}{{/if}}:</strong> {{book.author}}</p>
        <p><strong>Last Updated:</strong> {{book.date}}</p>
        
        {{#if book.social}}
        <div class="book-social-links">
            <h3>Book Resources</h3>
            <div class="social-links book-links">
                {{#if book.social.github}}<a href="{{book.social.github}}" title="GitHub Repository"><i class="icon">📦</i> GitHub</a>{{/if}}
                {{#if book.social.twitter}}<a href="https://twitter.com/{{book.social.twitter}}" title="Twitter"><i class="icon">🐦</i> Twitter</a>{{/if}}
                {{#if book.social.linkedin}}<a href="{{book.social.linkedin}}" title="LinkedIn"><i class="icon">👔</i> LinkedIn</a>{{/if}}
                {{#if book.social.website}}<a href="{{book.social.website}}" title="Website"><i class="icon">🌐</i> Website</a>{{/if}}
            </div>
        </div>
        {{/if}}
        
        {{#if book.authorObjects}}
        <div class="author-social-links">
            {{#each book.authorObjects}}
            <div class="author-block">
                <h3>{{this.name}}</h3>
                <div class="social-links author-links">
                    {{#if this.email}}<a href="mailto:{{this.email}}" title="Email {{this.name}}"><i class="icon">📧</i> Email</a>{{/if}}
                    {{#if this.twitter}}<a href="https://twitter.com/{{this.twitter}}" title="Twitter {{this.name}}"><i class="icon">🐦</i> Twitter</a>{{/if}}
                    {{#if this.website}}<a href="{{this.website}}" title="Website {{this.name}}"><i class="icon">🌐</i> Website</a>{{/if}}
                </div>
            </div>
            {{/each}}
        </div>
        {{/if}}
    </div>
</div>`;
}

async function createDefaultChapter() {
    return `<div class="chapter-content">
    {{{chapter.content}}}
    
    <div class="chapter-nav">
        {{#if prevChapter}}
        <a href="/chapters/{{prevChapter.filename}}" class="chapter-nav-item chapter-nav-prev">
            <span>← Previous</span>
            <span>{{prevChapter.title}}</span>
        </a>
        {{else}}
        <div></div>
        {{/if}}
        
        <div class="chapter-nav-center">
            <div class="chapter-indicator active"></div>
        </div>
        
        {{#if nextChapter}}
        <a href="/chapters/{{nextChapter.filename}}" class="chapter-nav-item chapter-nav-next">
            <span>Next →</span>
            <span>{{nextChapter.title}}</span>
        </a>
        {{else}}
        <div></div>
        {{/if}}
    </div>
</div>`;
}

// Run the script
if (require.main === module) {
    buildWeb();
}

module.exports = buildWeb;
