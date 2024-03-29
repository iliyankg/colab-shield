#!/bin/bash

serverAddress=$1
fileTypes=$2

echo "This script is used to configure the server address and file types for the Colab Shield tool." 
echo "Usage: ./install.sh " 
echo "Arguments:" 
echo " server_address: The address of the server where Colab Shield is deployed." 
echo " file_types: A comma-separated list of file types to be protected by Colab Shield." echo "Example: ./install.sh http://example.com txt,csv"

if [ -z $serverAddress ]; then
    echo "Please provide the server address."
    exit 1
fi

if [ -z $fileTypes ]; then
    echo "Please provide the file types."
    exit 1
fi

if [ ! -d .git ]; then
    echo "We are not in a git repository."
    exit 1
fi

git config colab.shield.server $serverAddress
git config colab.shield.filetypes $fileTypes
