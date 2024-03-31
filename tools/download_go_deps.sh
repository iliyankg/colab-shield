#!/bin/bash

# Backend folder
cd ../backend
go mod download

# Cli folder
cd ../cli
go mod download

# Protos folder
cd ../protos
go mod download