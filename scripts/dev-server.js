#!/usr/bin/env node

const express = require('express');
const chokidar = require('chokidar');
const path = require('path');
const buildWeb = require('./build-web');

const app = express();
const PORT = 3000;

// Serve static files from docs directory
app.use(express.static('docs'));

// Watch for changes in source files
const watcher = chokidar.watch(['src/**/*', 'config/**/*'], {
    ignored: /node_modules/,
    persistent: true
});

let buildTimeout;

watcher.on('change', (filePath) => {
    console.log(`📝 File changed: ${filePath}`);
    
    // Debounce builds to avoid rapid rebuilds
    clearTimeout(buildTimeout);
    buildTimeout = setTimeout(async () => {
        try {
            console.log('🔄 Rebuilding...');
            await buildWeb();
            console.log('✅ Rebuild complete!');
        } catch (error) {
            console.error('❌ Rebuild failed:', error.message);
        }
    }, 500);
});

// Initial build
buildWeb().then(() => {
    app.listen(PORT, () => {
        console.log(`🌐 Development server running at http://localhost:${PORT}`);
        console.log('👀 Watching for changes...');
        console.log('Press Ctrl+C to stop');
    });
}).catch(error => {
    console.error('❌ Failed to start development server:', error.message);
    process.exit(1);
});

// Graceful shutdown
process.on('SIGINT', () => {
    console.log('\n👋 Shutting down development server...');
    watcher.close();
    process.exit(0);
});
