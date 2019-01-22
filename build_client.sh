#!/bin/bash

# Get the directory of this script
# https://stackoverflow.com/questions/59895/getting-the-source-directory-of-a-bash-script-from-within
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# Combine, optimize, and minify CSS
echo "Building the CSS."
python "$DIR/css.py"

# Browserify, combine, and minify JavaScript
echo "Building the JavaScript."
python "$DIR/javascript.py"

echo "Done!"