#!/bin/bash
# ==============================================================================
# Order Service - Git Hooks Installation Script
# ==============================================================================
# Installs git hooks for code quality and standardization enforcement
# ==============================================================================

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# Helper functions
info() { echo -e "${BLUE}ℹ️  $1${NC}"; }
success() { echo -e "${GREEN}✅ $1${NC}"; }
warning() { echo -e "${YELLOW}⚠️  $1${NC}"; }
error() { echo -e "${RED}❌ $1${NC}"; }

echo ""
echo -e "${BLUE}╔══════════════════════════════════════════════════════════════╗${NC}"
echo -e "${BLUE}║         Order Service - Git Hooks Installation          ║${NC}"
echo -e "${BLUE}╚══════════════════════════════════════════════════════════════╝${NC}"
echo ""

# Check if we're in a git repository
if [ ! -d ".git" ]; then
    error "Not in a git repository. Please run this script from the project root."
    exit 1
fi

# Check if hooks directory exists
if [ ! -d "docs/githooks" ]; then
    error "Git hooks directory not found: docs/githooks"
    exit 1
fi

# Create .git/hooks directory if it doesn't exist
if [ ! -d ".git/hooks" ]; then
    info "Creating .git/hooks directory..."
    mkdir -p .git/hooks
fi

# Install hooks
HOOKS=("pre-commit" "commit-msg" "pre-push")
INSTALLED_COUNT=0

for hook in "${HOOKS[@]}"; do
    SOURCE_HOOK="docs/githooks/$hook"
    TARGET_HOOK=".git/hooks/$hook"

    if [ ! -f "$SOURCE_HOOK" ]; then
        warning "Source hook not found: $SOURCE_HOOK"
        continue
    fi

    info "Installing $hook hook..."

    # Backup existing hook if it exists
    if [ -f "$TARGET_HOOK" ]; then
        BACKUP_FILE="$TARGET_HOOK.backup.$(date +%Y%m%d_%H%M%S)"
        warning "Backing up existing hook to: $BACKUP_FILE"
        cp "$TARGET_HOOK" "$BACKUP_FILE"
    fi

    # Copy and make executable
    cp "$SOURCE_HOOK" "$TARGET_HOOK"
    chmod +x "$TARGET_HOOK"

    success "$hook hook installed"
    INSTALLED_COUNT=$((INSTALLED_COUNT + 1))
done

echo ""

if [ $INSTALLED_COUNT -eq 0 ]; then
    error "No hooks were installed"
    exit 1
elif [ $INSTALLED_COUNT -eq ${#HOOKS[@]} ]; then
    success "All $INSTALLED_COUNT hooks installed successfully!"
else
    warning "$INSTALLED_COUNT out of ${#HOOKS[@]} hooks installed"
fi

echo ""
echo -e "${CYAN}Installed hooks:${NC}"
for hook in "${HOOKS[@]}"; do
    if [ -f ".git/hooks/$hook" ]; then
        echo -e "  ✅ $hook"
    else
        echo -e "  ❌ $hook"
    fi
done

echo ""
echo -e "${CYAN}Hook Configuration:${NC}"
echo "You can configure hook behavior with environment variables:"
echo ""
echo "  # Skip specific checks"
echo "  export SKIP_LINT_CHECK=true"
echo "  export SKIP_TYPE_CHECK=true"
echo "  export SKIP_TEST_CHECK=true"
echo "  export SKIP_BUILD_CHECK=true"
echo "  export SKIP_COVERAGE_CHECK=true"
echo "  export SKIP_MODULE_STRUCTURE_CHECK=true"
echo ""
echo "  # Coverage thresholds (default: 90%)"
echo "  export COVERAGE_THRESHOLD_LINES=90"
echo "  export COVERAGE_THRESHOLD_FUNCTIONS=90"
echo "  export COVERAGE_THRESHOLD_BRANCHES=90"
echo "  export COVERAGE_THRESHOLD_STATEMENTS=90"
echo ""
echo "  # Debug mode"
echo "  export DEBUG_HOOKS=true"
echo ""

echo -e "${CYAN}Emergency Bypass:${NC}"
echo "In emergency situations, you can bypass hooks:"
echo ""
echo "  # Skip all pre-commit hooks"
echo "  git commit --no-verify -m \"Emergency fix\""
echo ""
echo "  # Skip pre-push hooks"
echo "  git push --no-verify"
echo ""

echo -e "${CYAN}Testing Hooks:${NC}"
echo "Test the hooks manually:"
echo ""
echo "  # Test pre-commit hook"
echo "  bash .git/hooks/pre-commit"
echo ""
echo "  # Test commit-msg hook"
echo "  echo 'feat(iam): add user registration' > /tmp/test-commit-msg"
echo "  bash .git/hooks/commit-msg /tmp/test-commit-msg"
echo ""
echo "  # Test pre-push hook"
echo "  bash .git/hooks/pre-push origin git@github.com:user/repo.git"
echo ""

success "Git hooks installation completed!"
echo ""