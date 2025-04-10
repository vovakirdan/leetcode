#!/bin/bash
HOOKS_DIR=".git/hooks"
TARGET="$HOOKS_DIR/pre-commit"

mkdir -p "$HOOKS_DIR"
cp scripts/hooks/pre-commit "$TARGET"
chmod +x "$TARGET"

echo "✅ Git pre-commit hook installed at $TARGET"
