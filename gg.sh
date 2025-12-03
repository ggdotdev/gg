#!/bin/sh
set -e
echo "Installing gg — the 2-letter agent-native git client"
ARCH=$(uname -m)
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
curl -L "https://github.com/ggdotdev/gg/releases/download/v0.1.0/gg_${OS}_${ARCH}" -o /tmp/gg
sudo mv /tmp/gg /usr/local/bin/gg
sudo chmod +x /usr/local/bin/gg
echo "✓ gg installed! Try: gg maaza"
