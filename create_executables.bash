#!/bin/bash

platforms=("linux/amd64" "linux/arm64" "darwin/amd64" "darwin/arm64")

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	directory_name="dist/oh-heck-$GOOS-$GOARCH"
	output_name="$directory_name/oh-heck"
	# if [ $GOOS = "windows" ]; then
	# 	output_name+='oh-heck.exe'
	# fi	

	env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name
	if [ $? -ne 0 ]; then
   		echo 'An error has occurred! Aborting the script execution...'
		exit 1
	else
		chmod +x "$output_name"
		
		if [ "$GOOS" = "linux" ]; then
			(cd $directory_name && tar -xzvf oh-heck.tar.gz oh-heck)
		elif [ "$GOOS" = "darwin" ]; then
			(cd $directory_name && zip -r oh-heck.zip oh-heck)
		fi
	fi
done

