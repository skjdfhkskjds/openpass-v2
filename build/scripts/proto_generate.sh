#!/bin/bash
# SPDX-License-Identifier: MIT
#
# Copyright (c) 2024 skjdfhkskjds
#
# Permission is hereby granted, free of charge, to any person
# obtaining a copy of this software and associated documentation
# files (the "Software"), to deal in the Software without
# restriction, including without limitation the rights to use,
# copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the
# Software is furnished to do so, subject to the following
# conditions:
#
# The above copyright notice and this permission notice shall be
# included in all copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
# EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
# OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
# NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
# HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
# WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
# FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
# OTHER DEALINGS IN THE SOFTWARE.


# Define the source and destination directories
SRC_DIR="proto"
DEST_DIR="types"

# Path to the `protoc-gen-go` and `protoc-gen-go-grpc` plugins
# Ensure these paths are correct for your system or installed globally
GO_PLUGIN_PATH=$(which protoc-gen-go)
GRPC_PLUGIN_PATH=$(which protoc-gen-go-grpc)

# Create the output directory if it does not exist
mkdir -p "$DEST_DIR"

# Generate Go code for all proto files in services and types directories
array=("password" "user")
for element in "${array[@]}"
do
       echo "Generating code for $element"
       protoc --go_out=$DEST_DIR --go_opt=paths=source_relative \
              "$SRC_DIR/types/v1/$element.proto" 
       
       protoc --go-grpc_out=$DEST_DIR --go-grpc_opt=paths=source_relative \
              "$SRC_DIR/services/v1/$element.proto"
done

echo "Proto files have been processed and output files are in '$DEST_DIR'"
