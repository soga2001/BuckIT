#!/bin/bash
# This script generates the OpenAPI JSON file from the server running on localhost:8080.
# Insure that the server is running before executing this script.
# This script requires the 'json' command-line tool to format the output.

package='json'

# Check if the 'json' package is installed globally and install it if not
package_name='json'
if [[ "$(npm list -g $package_name)" =~ "empty" ]]; then
    echo "Installing $package_name ..."
    npm install -g $package_name
else
    echo "$package_name is already installed"
fi

curl -v http://localhost:8080/openapi.json | json > openapi.json

# If you want to remove the 'json' package after generating the OpenAPI JSON, uncomment the following lines:
# Note: This will remove the package globally, so be cautious if you use it in other projects.
# if [ `npm list -g | grep -c $package` -eq 1 ]; then
#     echo "Removing $package globally..."
#     npm uninstall -g $package
# fi