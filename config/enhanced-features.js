// Enhanced Book Website JavaScript

document.addEventListener('DOMContentLoaded', function() {
    initializeFeatures();
});

function initializeFeatures() {
    // Initialize all interactive features
    initThemeToggle();
    initReadingProgress();
    initSmoothScrolling();
    initTableOfContents();
    initKeyboardNavigation();
    initNavigationDropdowns();
    initEnhancedKeyboardNavigation();
    initChapterGrouping();
    enhanceChapterNavigation();
    initCopyButtons(); // Add this line
}

// Enhanced Navigation Functions
function initNavigationDropdowns() {
    // Close dropdowns when clicking outside
    document.addEventListener('click', function(e) {
        if (!e.target.closest('.nav-dropdown')) {
            const activeDropdowns = document.querySelectorAll('.nav-dropdown.active');
            activeDropdowns.forEach(dropdown => {
                dropdown.classList.remove('active');
            });
        }
    });
    
    // Close dropdowns on escape key
    document.addEventListener('keydown', function(e) {
        if (e.key === 'Escape') {
            const activeDropdowns = document.querySelectorAll('.nav-dropdown.active');
            activeDropdowns.forEach(dropdown => {
                dropdown.classList.remove('active');
            });
            
            const activeMobileMenu = document.querySelector('.nav-menu.active');
            if (activeMobileMenu) {
                toggleMobileMenu();
            }
        }
    });
}

function toggleChaptersDropdown() {
    const dropdown = document.querySelector('.nav-dropdown');
    if (dropdown) {
        dropdown.classList.toggle('active');
    }
}

function toggleMobileMenu() {
    const toggle = document.querySelector('.mobile-menu-toggle');
    const menu = document.querySelector('.nav-menu');
    
    if (toggle && menu) {
        toggle.classList.toggle('active');
        menu.classList.toggle('active');
        
        // Prevent body scroll when mobile menu is open
        if (menu.classList.contains('active')) {
            document.body.style.overflow = 'hidden';
        } else {
            document.body.style.overflow = '';
        }
    }
}

// Enhanced chapter navigation with progress indication
function enhanceChapterNavigation() {
    const chapterDropdownItems = document.querySelectorAll('.chapter-dropdown-item');
    
    chapterDropdownItems.forEach((item, index) => {
        // Add progress indication
        const progressIndicator = document.createElement('div');
        progressIndicator.className = 'chapter-progress';
        
        // Check if chapter has been visited (you can implement this logic based on localStorage)
        const isVisited = localStorage.getItem(`chapter-${index}-visited`) === 'true';
        if (isVisited) {
            progressIndicator.classList.add('visited');
        }
        
        item.appendChild(progressIndicator);
        
        // Mark as visited when clicked
        item.addEventListener('click', () => {
            localStorage.setItem(`chapter-${index}-visited`, 'true');
        });
    });
}

// Filter chapters in dropdown
function filterChapters(searchTerm) {
    const chaptersGrid = document.getElementById('chapters-grid');
    const chapterItems = chaptersGrid.querySelectorAll('.chapter-dropdown-item');
    
    chapterItems.forEach(item => {
        const title = item.getAttribute('data-chapter-title').toLowerCase();
        const shouldShow = title.includes(searchTerm.toLowerCase());
        item.style.display = shouldShow ? 'flex' : 'none';
    });
}

// Enhanced keyboard navigation for chapters
function initEnhancedKeyboardNavigation() {
    document.addEventListener('keydown', function(e) {
        // Ctrl/Cmd + K: Open chapter search
        if ((e.ctrlKey || e.metaKey) && e.key === 'k') {
            e.preventDefault();
            const chaptersDropdown = document.querySelector('.nav-dropdown');
            if (chaptersDropdown) {
                chaptersDropdown.classList.add('active');
                const searchInput = document.getElementById('chapter-search');
                if (searchInput) {
                    setTimeout(() => searchInput.focus(), 100);
                }
            }
        }
        
        // Escape: Close all dropdowns
        if (e.key === 'Escape') {
            const activeDropdowns = document.querySelectorAll('.nav-dropdown.active');
            activeDropdowns.forEach(dropdown => {
                dropdown.classList.remove('active');
            });
        }
    });
}

// Progressive enhancement for large chapter lists
function initChapterGrouping() {
    const chaptersGrid = document.getElementById('chapters-grid');
    const chapterItems = chaptersGrid.querySelectorAll('.chapter-dropdown-item');
    
    // If more than 20 chapters, consider grouping
    if (chapterItems.length > 20) {
        groupChaptersByPrefix(chapterItems, chaptersGrid);
    }
}

function groupChaptersByPrefix(chapterItems, container) {
    const groups = new Map();
    
    chapterItems.forEach(item => {
        const title = item.querySelector('.chapter-title').textContent;
        const prefix = extractChapterPrefix(title);
        
        if (!groups.has(prefix)) {
            groups.set(prefix, []);
        }
        groups.get(prefix).push(item);
    });
    
    // Only group if we have meaningful groups (more than 1 group with multiple items)
    const meaningfulGroups = Array.from(groups.entries()).filter(([_, items]) => items.length > 1);
    
    if (meaningfulGroups.length > 1) {
        container.innerHTML = '';
        
        groups.forEach((items, groupName) => {
            if (items.length > 1) {
                const groupDiv = document.createElement('div');
                groupDiv.className = 'chapter-group';
                
                const groupHeader = document.createElement('div');
                groupHeader.className = 'chapter-group-header';
                groupHeader.textContent = groupName;
                
                const groupGrid = document.createElement('div');
                groupGrid.className = 'chapters-grid';
                
                items.forEach(item => {
                    groupGrid.appendChild(item);
                });
                
                groupDiv.appendChild(groupHeader);
                groupDiv.appendChild(groupGrid);
                container.appendChild(groupDiv);
            } else {
                // Single items go directly in container
                container.appendChild(items[0]);
            }
        });
    }
}

function extractChapterPrefix(title) {
    // Extract meaningful prefixes for grouping
    // Examples: "Part 1: Introduction" -> "Part 1"
    //          "Chapter 5: Advanced" -> "Chapters 1-10"
    //          "Getting Started" -> "Getting Started"
    
    const partMatch = title.match(/^(Part \d+)/i);
    if (partMatch) return partMatch[1];
    
    const chapterMatch = title.match(/^Chapter (\d+)/i);
    if (chapterMatch) {
        const num = parseInt(chapterMatch[1]);
        const groupSize = 5;
        const groupStart = Math.floor((num - 1) / groupSize) * groupSize + 1;
        const groupEnd = groupStart + groupSize - 1;
        return `Chapters ${groupStart}-${groupEnd}`;
    }
    
    return 'Chapters';
}

function toggleTheme() {
    const html = document.documentElement;
    const currentTheme = html.getAttribute('data-theme');
    const newTheme = currentTheme === 'dark' ? 'light' : 'dark';
    html.setAttribute('data-theme', newTheme);
    localStorage.setItem('theme', newTheme);
    
    // Add transition effect
    html.style.transition = 'background-color 0.3s ease';
    setTimeout(() => {
        html.style.transition = '';
    }, 300);
}

// Theme Management
function initThemeToggle() {
    // Load saved theme preference
    const savedTheme = localStorage.getItem('theme') || 
                      (window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light');
    document.documentElement.setAttribute('data-theme', savedTheme);
    
    // Listen for system theme changes
    window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', (e) => {
        if (!localStorage.getItem('theme')) {
            document.documentElement.setAttribute('data-theme', e.matches ? 'dark' : 'light');
        }
    });
}

// Reading Progress Indicator
function initReadingProgress() {
    const progressBar = document.querySelector('.reading-progress');
    if (!progressBar) return;
    
    function updateProgress() {
        const article = document.querySelector('.chapter-content');
        if (!article) return;
        
        const scrollTop = window.scrollY;
        const docHeight = document.documentElement.scrollHeight - window.innerHeight;
        const scrollPercent = (scrollTop / docHeight) * 100;
        
        progressBar.style.width = Math.min(100, Math.max(0, scrollPercent)) + '%';
    }
    
    window.addEventListener('scroll', updateProgress);
    window.addEventListener('resize', updateProgress);
    updateProgress();
}

// Smooth Scrolling
function initSmoothScrolling() {
    document.addEventListener('click', function(e) {
        if (e.target.tagName === 'A' && e.target.getAttribute('href')?.startsWith('#')) {
            e.preventDefault();
            const targetId = e.target.getAttribute('href').substring(1);
            const targetElement = document.getElementById(targetId);
            if (targetElement) {
                targetElement.scrollIntoView({ 
                    behavior: 'smooth',
                    block: 'start'
                });
            }
        }
    });
}

// Dynamic Table of Contents
function initTableOfContents() {
    const content = document.querySelector('.chapter-content');
    if (!content) return;
    
    const headings = content.querySelectorAll('h1, h2, h3, h4, h5, h6');
    if (headings.length < 2) return;
    
    // Create TOC sidebar structure
    const tocSidebar = document.createElement('div');
    tocSidebar.className = 'toc-sidebar';
    tocSidebar.innerHTML = `
        <div class="toc-header">
            <span>Contents</span>
            <button class="toc-collapse" aria-label="Collapse TOC">−</button>
        </div>
        <div class="toc-content">
            <ul id="toc-list"></ul>
        </div>
    `;

    // Create mobile toggle button
    const tocToggle = document.createElement('button');
    tocToggle.className = 'toc-toggle';
    tocToggle.setAttribute('aria-label', 'Toggle Table of Contents');
    tocToggle.title = 'Table of Contents';

    // Add mobile close button
    const mobileCloseBtn = document.createElement('button');
    mobileCloseBtn.className = 'mobile-toc-close';
    mobileCloseBtn.innerHTML = '×';
    mobileCloseBtn.setAttribute('aria-label', 'Close Table of Contents');
    tocSidebar.appendChild(mobileCloseBtn);

    // Add to page
    document.body.appendChild(tocSidebar);
    document.body.appendChild(tocToggle);

    // Adjust TOC position based on navigation height
    function adjustTocPosition() {
        const nav = document.querySelector('nav');
        const header = document.querySelector('header');
        
        if (nav && window.innerWidth > 768) {
            const navHeight = nav.offsetHeight;
            const headerHeight = header ? header.offsetHeight : 0;
            const scrollTop = window.scrollY;
            
            // Calculate the effective top position
            let topPosition;
            
            // If we're at the very top (header is visible)
            if (scrollTop <= headerHeight) {
                // Position below both header and nav
                topPosition = headerHeight + navHeight + 20;
            } else {
                // Header is scrolled out of view, position below sticky nav
                topPosition = navHeight + 20;
            }
            
            // Ensure minimum distance from top (never higher than nav + 20px)
            topPosition = Math.max(topPosition, navHeight + 20);
            // Also ensure it's never unreasonably high
            topPosition = Math.max(topPosition, 80);
            
            tocSidebar.style.top = topPosition + 'px';
            tocSidebar.style.maxHeight = `calc(100vh - ${topPosition + 40}px)`;
            
            // Update toc-content max-height as well
            const tocContent = tocSidebar.querySelector('.toc-content');
            if (tocContent) {
                tocContent.style.maxHeight = `calc(100vh - ${topPosition + 120}px)`;
            }
        }
    }

    // Call on scroll and resize with throttling for better performance
    let tocPositionTimeout;
    function throttledAdjustPosition() {
        if (tocPositionTimeout) clearTimeout(tocPositionTimeout);
        tocPositionTimeout = setTimeout(adjustTocPosition, 10);
    }
    
    window.addEventListener('scroll', throttledAdjustPosition);
    window.addEventListener('resize', adjustTocPosition);
    
    // Initial adjustment with multiple attempts to ensure proper positioning
    setTimeout(adjustTocPosition, 100);
    setTimeout(adjustTocPosition, 500);
    setTimeout(adjustTocPosition, 1000);

    // Build TOC list with proper numbering and nesting
    const tocList = document.getElementById('toc-list');
    let tocHTML = '';
    let previousLevel = 0;
    let counters = { h1: 0, h2: 0, h3: 0, h4: 0, h5: 0, h6: 0 };
    
    headings.forEach((heading, index) => {
        const level = parseInt(heading.tagName.charAt(1));
        const text = heading.textContent.trim();
        const id = heading.id || `heading-${heading.tagName.toLowerCase()}-${index}`;
        
        if (!heading.id) {
            heading.id = id;
        }

        // Reset counters for deeper levels when we encounter a higher level heading
        for (let i = level + 1; i <= 6; i++) {
            counters[`h${i}`] = 0;
        }
        counters[`h${level}`]++;

        // Calculate nesting
        const levelDiff = level - previousLevel;
        let levelClass = `toc-level-${level}`;
        let indent = Math.max(0, (level - 1) * 1); // rem units for indentation
        
        // Create number/bullet based on level
        let marker = '';
        if (level === 1) {
            marker = `<span class="toc-marker toc-marker-h1">${counters.h1}</span>`;
        } else if (level === 2) {
            marker = `<span class="toc-marker toc-marker-h2">${counters.h1}.${counters.h2}</span>`;
        } else if (level === 3) {
            marker = `<span class="toc-marker toc-marker-h3">•</span>`;
        } else {
            marker = `<span class="toc-marker toc-marker-h4">–</span>`;
        }

        tocHTML += `
            <li class="${levelClass}" style="margin-left: ${indent}rem;">
                <a href="#${id}" data-heading-id="${id}" class="toc-link">
                    ${marker}
                    <span class="toc-text">${text}</span>
                </a>
            </li>
        `;
        
        previousLevel = level;
    });
    
    tocList.innerHTML = tocHTML;

    // Handle collapse/expand
    const collapseBtn = tocSidebar.querySelector('.toc-collapse');
    const tocContent = tocSidebar.querySelector('.toc-content');
    let isCollapsed = false;
    
    collapseBtn.addEventListener('click', () => {
        isCollapsed = !isCollapsed;
        tocContent.style.display = isCollapsed ? 'none' : 'block';
        collapseBtn.textContent = isCollapsed ? '+' : '−';
        tocSidebar.classList.toggle('collapsed', isCollapsed);
    });

    // Handle mobile toggle
    tocToggle.addEventListener('click', () => {
        tocSidebar.classList.toggle('mobile-active');
        document.body.classList.toggle('toc-mobile-open');
    });

    // Handle mobile close
    mobileCloseBtn.addEventListener('click', () => {
        tocSidebar.classList.remove('mobile-active');
        document.body.classList.remove('toc-mobile-open');
    });

    // Close on mobile when clicking a link
    tocList.addEventListener('click', (e) => {
        if (e.target.closest('.toc-link') && window.innerWidth <= 768) {
            tocSidebar.classList.remove('mobile-active');
            document.body.classList.remove('toc-mobile-open');
        }
    });

    // Close TOC on escape key or clicking outside
    document.addEventListener('keydown', (e) => {
        if (e.key === 'Escape') {
            tocSidebar.classList.remove('mobile-active');
            document.body.classList.remove('toc-mobile-open');
        }
    });

    // Close when clicking outside on mobile
    document.addEventListener('click', (e) => {
        if (window.innerWidth <= 768 && 
            tocSidebar.classList.contains('mobile-active') &&
            !tocSidebar.contains(e.target) &&
            !tocToggle.contains(e.target)) {
            tocSidebar.classList.remove('mobile-active');
            document.body.classList.remove('toc-mobile-open');
        }
    });

    // Update active TOC item on scroll with enhanced logic
    function updateActiveTocItem() {
        const scrollPosition = window.scrollY + 120; // Reduced offset for better top detection
        let activeHeading = null;
        let closestDistance = Infinity;

        // Find the current active heading based on scroll position
        headings.forEach(heading => {
            const headingTop = heading.offsetTop;
            const distance = Math.abs(scrollPosition - headingTop);
            
            if (headingTop <= scrollPosition && distance < closestDistance) {
                activeHeading = heading;
                closestDistance = distance;
            }
        });

        // If no heading is above scroll position or we're near the top, use the first heading
        if (!activeHeading || scrollPosition < 200) {
            if (headings.length > 0) {
                activeHeading = headings[0];
            }
        }

        // Update TOC active states
        const tocLinks = document.querySelectorAll('.toc-content .toc-link');
        tocLinks.forEach(link => {
            link.classList.remove('active');
            if (activeHeading && link.getAttribute('data-heading-id') === activeHeading.id) {
                link.classList.add('active');
                
                // Auto-scroll the TOC to keep active item visible
                const tocContent = link.closest('.toc-content');
                if (tocContent && !tocSidebar.classList.contains('mobile-active')) {
                    const linkTop = link.offsetTop;
                    const linkHeight = link.offsetHeight;
                    const tocContentHeight = tocContent.clientHeight;
                    const tocContentScrollTop = tocContent.scrollTop;
                    
                    if (linkTop < tocContentScrollTop || linkTop + linkHeight > tocContentScrollTop + tocContentHeight) {
                        tocContent.scrollTo({
                            top: linkTop - tocContentHeight / 2 + linkHeight / 2,
                            behavior: 'smooth'
                        });
                    }
                }
            }
        });
    }
    
    // Throttle scroll events for better performance
    let scrollTimeout;
    window.addEventListener('scroll', () => {
        if (scrollTimeout) {
            clearTimeout(scrollTimeout);
        }
        scrollTimeout = setTimeout(updateActiveTocItem, 10);
    });
    
    // Update on resize
    window.addEventListener('resize', updateActiveTocItem);
    
    // Initial update with a small delay to ensure proper rendering
    setTimeout(() => {
        updateActiveTocItem();
        // Ensure TOC is visible on page load
        if (window.innerWidth > 768) {
            tocSidebar.style.display = 'block';
            tocSidebar.style.opacity = '1';
            tocSidebar.style.visibility = 'visible';
        }
    }, 200);

    // Also update immediately for faster response
    updateActiveTocItem();

    // Enhanced smooth scrolling for TOC links
    tocList.addEventListener('click', (e) => {
        const link = e.target.closest('.toc-link');
        if (link) {
            e.preventDefault();
            const targetId = link.getAttribute('data-heading-id');
            const targetElement = document.getElementById(targetId);
            
            if (targetElement) {
                const offsetTop = targetElement.offsetTop - 100; // Account for fixed header
                window.scrollTo({
                    top: offsetTop,
                    behavior: 'smooth'
                });
                
                // Add a brief highlight to the target heading
                targetElement.style.transition = 'background-color 0.3s ease';
                targetElement.style.backgroundColor = 'var(--accent-color)';
                targetElement.style.color = 'white';
                targetElement.style.padding = '0.5rem';
                targetElement.style.borderRadius = 'var(--border-radius)';
                
                setTimeout(() => {
                    targetElement.style.backgroundColor = '';
                    targetElement.style.color = '';
                    targetElement.style.padding = '';
                    targetElement.style.borderRadius = '';
                }, 1000);
            }
        }
    });
}

// Keyboard Navigation
function initKeyboardNavigation() {
    document.addEventListener('keydown', function(e) {
        // Alt + Left Arrow: Previous chapter
        if (e.altKey && e.key === 'ArrowLeft') {
            const prevLink = document.querySelector('.chapter-nav-prev');
            if (prevLink) {
                window.location.href = prevLink.href;
            }
        }
        
        // Alt + Right Arrow: Next chapter
        if (e.altKey && e.key === 'ArrowRight') {
            const nextLink = document.querySelector('.chapter-nav-next');
            if (nextLink) {
                window.location.href = nextLink.href;
            }
        }
        
        // Alt + T: Toggle theme
        if (e.altKey && e.key === 't') {
            toggleTheme();
        }
    });
}

// Copy code blocks - Initialize buttons on page load and handle copying
function initCopyButtons() {
    const preElements = document.querySelectorAll('pre');
    
    preElements.forEach(pre => {
        if (!pre.querySelector('.copy-button')) {
            const button = document.createElement('button');
            button.className = 'copy-button';
            button.textContent = '📋';
            button.title = 'Copy code';
            button.onclick = function() {
                const code = pre.querySelector('code').textContent;
                navigator.clipboard.writeText(code).then(() => {
                    button.textContent = '✅';
                    setTimeout(() => {
                        button.textContent = '📋';
                    }, 2000);
                });
            };
            pre.style.position = 'relative';
            pre.appendChild(button);
        }
    });
}

// Print functionality
function printChapter() {
    window.print();
}

// Add print styles
const printStyles = `
    @media print {
        .floating-toc, .theme-toggle, 
        .copy-button, .reading-progress {
            display: none !important;
        }
    }
`;

const style = document.createElement('style');
style.textContent = printStyles;
document.head.appendChild(style);
