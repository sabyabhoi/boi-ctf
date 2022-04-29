#!/bin/bash

./a.out
if [[ $? -eq 134 ]]; then
	echo "boiCTF{ABCDEF}"
fi
