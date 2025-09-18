#!/bin/bash

# Color codes for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Emoji for better UX
CHECK="✅"
CROSS="❌"
GEAR="🔧"
PACKAGE="📦"
ROCKET="🚀"
INFO="ℹ️"
WARNING="⚠️"

echo -e "${BLUE}${GEAR} Setting up all dependencies for book building...${NC}"
echo ""

# Function to check if a command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Function to print status
print_status() {
    if [ $1 -eq 0 ]; then
        echo -e "${GREEN}${CHECK} $2${NC}"
    else
        echo -e "${RED}${CROSS} $2${NC}"
    fi
}

# Function to print info
print_info() {
    echo -e "${BLUE}${INFO} $1${NC}"
}

# Function to print warning
print_warning() {
    echo -e "${YELLOW}${WARNING} $1${NC}"
}

# Detect OS
OS="unknown"
if [[ "$OSTYPE" == "darwin"* ]]; then
    OS="macos"
elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
    OS="linux"
elif [[ "$OSTYPE" == "msys" || "$OSTYPE" == "cygwin" ]]; then
    OS="windows"
fi

print_info "Detected OS: $OS"
echo ""

# 1. Check Node.js and npm
echo -e "${BLUE}${GEAR} Checking Node.js and npm...${NC}"

if command_exists node; then
    NODE_VERSION=$(node --version)
    print_status 0 "Node.js is installed ($NODE_VERSION)"
else
    print_status 1 "Node.js is not installed"
    echo "Please install Node.js first:"
    case $OS in
        "macos")
            echo "  brew install node"
            echo "  Or download from: https://nodejs.org/"
            ;;
        "linux")
            echo "  sudo apt-get update && sudo apt-get install nodejs npm"
            echo "  Or use your distribution's package manager"
            ;;
        "windows")
            echo "  Download from: https://nodejs.org/"
            ;;
        *)
            echo "  Download from: https://nodejs.org/"
            ;;
    esac
    echo ""
    read -p "Press Enter when Node.js is installed, or Ctrl+C to exit..."
fi

if command_exists npm; then
    NPM_VERSION=$(npm --version)
    print_status 0 "npm is installed ($NPM_VERSION)"
else
    print_status 1 "npm is not installed (usually comes with Node.js)"
    exit 1
fi

echo ""

# 2. Install npm dependencies
echo -e "${BLUE}${PACKAGE} Installing npm dependencies...${NC}"
if [ -f "package.json" ]; then
    npm install
    if [ $? -eq 0 ]; then
        print_status 0 "npm dependencies installed successfully"
    else
        print_status 1 "Failed to install npm dependencies"
        exit 1
    fi
else
    print_warning "package.json not found. Make sure you're in the project root directory."
    exit 1
fi

echo ""

# 3. Check and install Python 3 and pip
echo -e "${BLUE}${GEAR} Checking Python 3 and pip...${NC}"

if command_exists python3; then
    PYTHON_VERSION=$(python3 --version)
    print_status 0 "Python 3 is installed ($PYTHON_VERSION)"
else
    print_status 1 "Python 3 is not installed"
    echo "Installing Python 3..."
    case $OS in
        "macos")
            if command_exists brew; then
                brew install python3
            else
                echo "Please install Homebrew first: https://brew.sh/"
                echo "Then run: brew install python3"
                exit 1
            fi
            ;;
        "linux")
            sudo apt-get update && sudo apt-get install -y python3 python3-pip
            ;;
        *)
            echo "Please install Python 3 manually from: https://python.org/"
            exit 1
            ;;
    esac
fi

if command_exists pip3; then
    PIP_VERSION=$(pip3 --version)
    print_status 0 "pip3 is installed ($PIP_VERSION)"
else
    print_status 1 "pip3 is not installed"
    case $OS in
        "macos")
            echo "Installing pip3..."
            if command_exists brew; then
                brew install python3
            else
                echo "Please install Homebrew first: https://brew.sh/"
                exit 1
            fi
            ;;
        "linux")
            sudo apt-get install -y python3-pip
            ;;
        *)
            echo "Please install pip3 manually"
            exit 1
            ;;
    esac
fi

echo ""

# 4. Install Pandoc
echo -e "${BLUE}${PACKAGE} Installing Pandoc...${NC}"

if command_exists pandoc; then
    PANDOC_VERSION=$(pandoc --version | head -1)
    print_status 0 "Pandoc is already installed ($PANDOC_VERSION)"
else
    print_info "Installing Pandoc..."
    case $OS in
        "macos")
            if command_exists brew; then
                brew install pandoc
            else
                echo "Please install Homebrew first: https://brew.sh/"
                echo "Then run: brew install pandoc"
                exit 1
            fi
            ;;
        "linux")
            sudo apt-get update && sudo apt-get install -y pandoc
            ;;
        "windows")
            echo "Please download and install Pandoc from: https://pandoc.org/installing.html"
            exit 1
            ;;
        *)
            echo "Please install Pandoc manually from: https://pandoc.org/installing.html"
            exit 1
            ;;
    esac
    
    if command_exists pandoc; then
        print_status 0 "Pandoc installed successfully"
    else
        print_status 1 "Pandoc installation failed"
        exit 1
    fi
fi

echo ""

# 5. Install WeasyPrint for PDF generation
echo -e "${BLUE}${PACKAGE} Installing WeasyPrint for PDF generation...${NC}"

if command_exists weasyprint; then
    WEASYPRINT_VERSION=$(weasyprint --version)
    print_status 0 "WeasyPrint is already installed ($WEASYPRINT_VERSION)"
else
    print_info "Installing WeasyPrint..."
    
    # Install system dependencies for WeasyPrint
    case $OS in
        "macos")
            if command_exists brew; then
                # Install system dependencies
                brew install cairo pango gdk-pixbuf libffi
                
                # Try different installation methods for WeasyPrint
                print_info "Trying to install WeasyPrint via Homebrew first..."
                if brew install weasyprint 2>/dev/null; then
                    print_status 0 "WeasyPrint installed via Homebrew"
                else
                    print_info "Homebrew installation failed, trying pip with virtual environment..."
                    
                    # Create a virtual environment for WeasyPrint
                    VENV_DIR="$HOME/.weasyprint-venv"
                    python3 -m venv "$VENV_DIR"
                    source "$VENV_DIR/bin/activate"
                    pip install weasyprint
                    deactivate
                    
                    # Create a wrapper script
                    cat > /usr/local/bin/weasyprint << 'EOF'
#!/bin/bash
source "$HOME/.weasyprint-venv/bin/activate"
exec python -m weasyprint "$@"
EOF
                    chmod +x /usr/local/bin/weasyprint
                    
                    if command_exists weasyprint; then
                        print_status 0 "WeasyPrint installed in virtual environment"
                    else
                        print_info "Virtual environment method failed, trying pipx..."
                        
                        # Try pipx installation
                        if command_exists brew; then
                            brew install pipx
                            pipx install weasyprint
                            
                            if command_exists weasyprint; then
                                print_status 0 "WeasyPrint installed via pipx"
                            else
                                print_warning "pipx installation failed"
                            fi
                        fi
                    fi
                fi
            fi
            ;;
        "linux")
            sudo apt-get update
            sudo apt-get install -y build-essential python3-dev python3-pip python3-setuptools python3-wheel python3-cffi libcairo2 libpango-1.0-0 libpangocairo-1.0-0 libgdk-pixbuf2.0-0 libffi-dev shared-mime-info
            
            # Try pip install with user flag first
            if pip3 install --user weasyprint 2>/dev/null; then
                print_status 0 "WeasyPrint installed with --user flag"
            else
                # Try with virtual environment
                VENV_DIR="$HOME/.weasyprint-venv"
                python3 -m venv "$VENV_DIR"
                source "$VENV_DIR/bin/activate"
                pip install weasyprint
                deactivate
                
                # Create wrapper script
                sudo tee /usr/local/bin/weasyprint > /dev/null << 'EOF'
#!/bin/bash
source "$HOME/.weasyprint-venv/bin/activate"
exec python -m weasyprint "$@"
EOF
                sudo chmod +x /usr/local/bin/weasyprint
                
                if command_exists weasyprint; then
                    print_status 0 "WeasyPrint installed in virtual environment"
                fi
            fi
            ;;
    esac
    
    if ! command_exists weasyprint; then
        print_status 1 "WeasyPrint installation failed"
        print_warning "Falling back to alternative PDF generators..."
        
        # Try to install wkhtmltopdf as fallback (note: discontinued on macOS Homebrew)
        echo -e "${BLUE}${PACKAGE} Installing wkhtmltopdf as fallback...${NC}"
        case $OS in
            "macos")
                print_warning "wkhtmltopdf has been discontinued in Homebrew as of 2024-12-16"
                print_info "Trying alternative installation methods..."
                
                # Try installing via direct download or other methods
                if ! command_exists wkhtmltopdf; then
                    print_info "You can manually install wkhtmltopdf from: https://wkhtmltopdf.org/downloads.html"
                    print_info "Or use Puppeteer/Chrome headless as an alternative"
                    
                    # Install Puppeteer as a fallback for PDF generation
                    print_info "Installing Puppeteer as PDF generation alternative..."
                    npm install --save-dev puppeteer
                    
                    if [ $? -eq 0 ]; then
                        print_status 0 "Puppeteer installed as PDF fallback"
                    else
                        print_warning "Puppeteer installation failed"
                    fi
                fi
                ;;
            "linux")
                sudo apt-get install -y wkhtmltopdf
                if command_exists wkhtmltopdf; then
                    print_status 0 "wkhtmltopdf installed as PDF fallback"
                fi
                ;;
        esac
        
        if ! command_exists wkhtmltopdf && [ "$OS" != "macos" ]; then
            print_warning "wkhtmltopdf installation failed"
        fi
        
        if ! command_exists wkhtmltopdf && ! npm list puppeteer &>/dev/null; then
            print_warning "No PDF generator available. You can still generate HTML and use browser print-to-PDF"
        fi
    fi
fi

echo ""

# 6. Install pdfjam (for PDF manipulation)
echo -e "${BLUE}${PACKAGE} Installing pdfjam (for PDF manipulation)...${NC}"

if command_exists pdfjam; then
    print_status 0 "pdfjam is already installed"
else
    print_info "pdfjam not found. Installing..."
    case $OS in
        "macos")
            if command_exists brew; then
                # Install TeX Live via BasicTeX (smaller) or MacTeX (full)
                print_info "Installing BasicTeX (includes pdfjam)..."
                brew install basictex
                
                # Add TeX Live to PATH
                export PATH="/usr/local/texlive/2025basic/bin/universal-darwin:$PATH"
                
                # Update tlmgr and install pdfjam if not included
                if command_exists tlmgr; then
                    sudo tlmgr update --self
                    sudo tlmgr install pdfjam
                fi
            else
                print_warning "Homebrew not found. Please install BasicTeX manually from:"
                print_info "https://www.tug.org/mactex/morepackages.html"
                print_info "Or install full MacTeX from: https://www.tug.org/mactex/"
            fi
            ;;
        "linux")
            print_info "Installing TeX Live and pdfjam..."
            sudo apt-get update && sudo apt-get install -y texlive-extra-utils
            ;;
        "windows")
            print_info "Please install MiKTeX or TeX Live for Windows:"
            print_info "https://miktex.org/ or https://www.tug.org/texlive/"
            ;;
        *)
            print_info "Please install TeX Live or MiKTeX for your system:"
            print_info "https://www.tug.org/texlive/"
            ;;
    esac
    
    # Check if pdfjam is now available (may need PATH update)
    if command_exists pdfjam; then
        print_status 0 "pdfjam installed successfully"
    else
        print_warning "pdfjam installation failed or not in PATH."
        print_info "You may need to restart your terminal or update your PATH."
        print_info "pdfjam is typically located in your TeX Live bin directory."
    fi
fi

echo ""

# 7. Install Calibre (optional, for MOBI generation)
echo -e "${BLUE}${PACKAGE} Installing Calibre (optional, for MOBI generation)...${NC}"

if command_exists ebook-convert; then
    print_status 0 "Calibre is already installed"
else
    print_info "Calibre not found. Installing..."
    case $OS in
        "macos")
            if command_exists brew; then
                brew install calibre
            else
                print_warning "Homebrew not found. Please install Calibre manually from: https://calibre-ebook.com/"
            fi
            ;;
        "linux")
            sudo apt-get update && sudo apt-get install -y calibre
            ;;
        "windows")
            print_info "Please download and install Calibre from: https://calibre-ebook.com/"
            ;;
        *)
            print_info "Please install Calibre manually from: https://calibre-ebook.com/"
            ;;
    esac
    
    if command_exists ebook-convert; then
        print_status 0 "Calibre installed successfully"
    else
        print_warning "Calibre installation failed or not in PATH. MOBI generation will be skipped."
    fi
fi

echo ""

# 8. Install markdownlint-cli globally (for linting)
echo -e "${BLUE}${PACKAGE} Installing markdownlint-cli globally...${NC}"

if command_exists markdownlint; then
    print_status 0 "markdownlint-cli is already installed"
else
    print_info "Installing markdownlint-cli globally..."
    npm install -g markdownlint-cli
    
    if command_exists markdownlint; then
        print_status 0 "markdownlint-cli installed successfully"
    else
        print_warning "markdownlint-cli installation failed. Linting may not work."
    fi
fi

echo ""

# 9. Final summary and test
echo -e "${GREEN}${ROCKET} Setup Complete! Summary:${NC}"
echo ""

echo "📋 Installed Dependencies:"
echo "=========================="

# Check each dependency
command_exists node && echo -e "${GREEN}${CHECK} Node.js: $(node --version)${NC}" || echo -e "${RED}${CROSS} Node.js: Not installed${NC}"
command_exists npm && echo -e "${GREEN}${CHECK} npm: $(npm --version)${NC}" || echo -e "${RED}${CROSS} npm: Not installed${NC}"
command_exists python3 && echo -e "${GREEN}${CHECK} Python 3: $(python3 --version)${NC}" || echo -e "${RED}${CROSS} Python 3: Not installed${NC}"
command_exists pip3 && echo -e "${GREEN}${CHECK} pip3: Available${NC}" || echo -e "${RED}${CROSS} pip3: Not installed${NC}"
command_exists pandoc && echo -e "${GREEN}${CHECK} Pandoc: $(pandoc --version | head -1 | cut -d' ' -f2)${NC}" || echo -e "${RED}${CROSS} Pandoc: Not installed${NC}"
command_exists weasyprint && echo -e "${GREEN}${CHECK} WeasyPrint: Available${NC}" || echo -e "${YELLOW}${WARNING} WeasyPrint: Not installed${NC}"

# Check for PDF generation alternatives
if command_exists wkhtmltopdf; then
    echo -e "${GREEN}${CHECK} wkhtmltopdf: Available${NC}"
elif npm list puppeteer &>/dev/null; then
    echo -e "${GREEN}${CHECK} Puppeteer: Available (PDF alternative)${NC}"
else
    echo -e "${YELLOW}${WARNING} PDF generators: Limited (use browser print-to-PDF)${NC}"
fi

command_exists pdfjam && echo -e "${GREEN}${CHECK} pdfjam: Available${NC}" || echo -e "${YELLOW}${WARNING} pdfjam: Not installed (PDF manipulation will be limited)${NC}"

command_exists ebook-convert && echo -e "${GREEN}${CHECK} Calibre: Available${NC}" || echo -e "${YELLOW}${WARNING} Calibre: Not installed (MOBI generation will be skipped)${NC}"
command_exists markdownlint && echo -e "${GREEN}${CHECK} markdownlint-cli: Available${NC}" || echo -e "${YELLOW}${WARNING} markdownlint-cli: Not installed${NC}"

echo ""
echo "🎯 Available Build Commands:"
echo "============================"
echo "npm run build          - Build all formats"
echo "npm run build:leanpub  - Build Leanpub format"
echo "npm run build:kindle   - Build Kindle format (EPUB + MOBI if Calibre available)"
echo "npm run build:web      - Build web format"
echo "npm run build:pdf      - Build PDF format"
echo "npm run dev            - Start development server"
echo "npm run lint           - Lint markdown files"
echo "npm run word-count     - Count words in all chapters"
echo "npm run validate       - Validate project structure"
echo "npm run clean          - Clean build artifacts"

echo ""
echo -e "${GREEN}${ROCKET} Ready to build! Try: npm run build${NC}"

# Test that we can at least validate the project
if [ -f "scripts/validate.js" ]; then
    echo ""
    print_info "Running project validation..."
    node scripts/validate.js
fi
