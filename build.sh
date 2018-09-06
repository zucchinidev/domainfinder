#!/bin/bash
echo "Building domainfinder..." && go build -o domainfinder
echo "Building synonyms..." && cd ./synonyms && go build -o ../lib/synonyms && cd ../
echo "Building available..." && cd ./available && go build -o ../lib/available && cd ../
echo "Building sprinkle..." && cd ./sprinkle && go build -o ../lib/sprinkle && cd ../
echo "Building coolify..." && cd ./coolify && go build -o ../lib/coolify && cd ../
echo "Building domainify..." && cd ./domainify && go build -o ../lib/domainify
echo Done.