#!/bin/bash

# Generate Bcrypt Password Hash
# This script generates a bcrypt hash for a given password

echo "====================================="
echo "Bcrypt Password Generator"
echo "====================================="

# Read password
read -s -p "Enter password: " password
echo ""
read -s -p "Confirm password: " password_confirm
echo ""

if [ "$password" != "$password_confirm" ]; then
    echo "Error: Passwords do not match"
    exit 1
fi

# Generate hash using Python
if command -v python3 &> /dev/null; then
    echo "Generating bcrypt hash..."
    python3 -c "import bcrypt; print(bcrypt.hashpw(b'$password', bcrypt.gensalt()).decode())"
else
    echo "Error: Python 3 is required"
    echo "Install with: apt-get install python3-bcrypt"
    echo "Or use: go run scripts/generate-password.go"
    exit 1
fi
