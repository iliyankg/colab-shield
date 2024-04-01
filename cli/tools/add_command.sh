#!/bin/bash

commandName=$1

# Check if file already exists
if [ -f "../cmd/$commandName.go" ]; then
    echo "Error: File ../cmd/$commandName.go already exists."
    echo "Command $commandName" already exists. Pick a different name.
    exit 1
fi

# Copy command template to ../cmd directory
cp ./templates/command.template ../cmd/"$commandName".go

# Replace {{commandName}} with the contents of commandName variable
sed -i "s/{{commandName}}/$commandName/g" ../cmd/"$commandName".go

# Add new line to ../cmd/root.go
sed -i "0,/rootCmd\.AddCommand(.*/s//rootCmd.AddCommand($commandName)\n\t&/" ../cmd/root.go