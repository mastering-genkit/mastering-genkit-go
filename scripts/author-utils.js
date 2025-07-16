/**
 * Author handling utilities for the book generation system
 * Supports both legacy single author string and new multiple authors array
 */

/**
 * Get authors as a formatted string for display
 * @param {Object} metadata - Book metadata
 * @returns {string} Formatted author string
 */
function getAuthorString(metadata) {
    // If new authors array exists, use it
    if (metadata.authors && Array.isArray(metadata.authors) && metadata.authors.length > 0) {
        const names = metadata.authors.map(author => author.name);
        
        if (names.length === 1) {
            return names[0];
        } else if (names.length === 2) {
            return `${names[0]} & ${names[1]}`;
        } else {
            return `${names.slice(0, -1).join(', ')} & ${names[names.length - 1]}`;
        }
    }
    
    // Fallback to legacy author field
    return metadata.author || '';
}

/**
 * Get authors as an array for Pandoc (supports multiple authors)
 * @param {Object} metadata - Book metadata
 * @returns {Array} Array of author names
 */
function getAuthorArray(metadata) {
    // If new authors array exists, use it
    if (metadata.authors && Array.isArray(metadata.authors) && metadata.authors.length > 0) {
        return metadata.authors.map(author => author.name);
    }
    
    // Fallback: split legacy author field by & or ,
    if (metadata.author) {
        return metadata.author
            .split(/\s*[&,]\s*/)
            .map(name => name.trim())
            .filter(name => name.length > 0);
    }
    
    return [];
}

/**
 * Get author objects with full metadata
 * @param {Object} metadata - Book metadata
 * @returns {Array} Array of author objects
 */
function getAuthorObjects(metadata) {
    // If new authors array exists, return it
    if (metadata.authors && Array.isArray(metadata.authors) && metadata.authors.length > 0) {
        return metadata.authors;
    }
    
    // Fallback: create basic author objects from legacy field
    if (metadata.author) {
        const names = metadata.author
            .split(/\s*[&,]\s*/)
            .map(name => name.trim())
            .filter(name => name.length > 0);
        
        return names.map(name => ({
            name: name,
            bio: '',
            email: '',
            twitter: '',
            website: ''
        }));
    }
    
    return [];
}

/**
 * Generate "About the Authors" section for multi-author books
 * @param {Object} metadata - Book metadata
 * @returns {string} Formatted about authors section
 */
function generateAboutAuthors(metadata) {
    const authors = getAuthorObjects(metadata);
    
    if (authors.length === 0) {
        return '# About the Author\n\nAdd your biography here.';
    }
    
    if (authors.length === 1) {
        const author = authors[0];
        return `# About the Author

${author.bio || `${author.name} is...`}

Add your biography here.`;
    }
    
    // Multiple authors
    let content = '# About the Authors\n\n';
    
    authors.forEach((author, index) => {
        content += `## ${author.name}\n\n`;
        content += `${author.bio || `${author.name} is...`}\n\n`;
        
        // Add contact info if available
        const contacts = [];
        if (author.website) contacts.push(`[Website](${author.website})`);
        if (author.twitter) contacts.push(`[Twitter](${author.twitter})`);
        if (author.email) contacts.push(`[Email](mailto:${author.email})`);
        
        if (contacts.length > 0) {
            content += `Contact: ${contacts.join(' • ')}\n\n`;
        }
        
        if (index < authors.length - 1) {
            content += '---\n\n';
        }
    });
    
    return content;
}

module.exports = {
    getAuthorString,
    getAuthorArray,
    getAuthorObjects,
    generateAboutAuthors
};
