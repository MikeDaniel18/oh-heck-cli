#!/bin/bash

platforms=("linux/amd64" "linux/arm64" "darwin/amd64" "darwin/arm64")

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	output_name='dist/oh-heck-'$GOOS'-'$GOARCH'/oh-heck'
	# if [ $GOOS = "windows" ]; then
	# 	output_name+='oh-heck.exe'
	# fi	

	env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name
	if [ $? -ne 0 ]; then
   		echo 'An error has occurred! Aborting the script execution...'
		exit 1
	else
		chmod +x "$output_name"
	fi
done

