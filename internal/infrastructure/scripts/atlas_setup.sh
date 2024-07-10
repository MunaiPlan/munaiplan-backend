#!/bin/bash

# Load environment variables from .env file
if [ -f .env ]; then
  export $(cat .env | xargs)
fi

# Set the absolute path to the configuration file
CONFIG_PATH="$(pwd)/infrastructure/configs/atlas.hcl"

# Run the Atlas migrate diff command with the specified environment
atlas migrate diff --env gorm --config "file://$CONFIG_PATH"
