#!/bin/sh
echo "------------- GARAGE INIT SCRIPT -----------"

CONTAINER_NAME="$1"

alias garage="MSYS_NO_PATHCONV=1 docker exec -t $CONTAINER_NAME /garage"

echo "Container name: $CONTAINER_NAME"

echo "Check garage status..."
garage status

echo "Get node ID..."
NODE_ID=$(garage status | awk 'NR == 5 {print $1; exit}')
echo "Node ID: $NODE_ID"

echo "Create layout..."
garage layout assign -z garage -c 1G $NODE_ID
echo "Layout created"
echo "Apply layout..."
garage layout apply --version 1
echo "Layout applied"

echo "Create bucket..."
garage bucket create md-drive-bucket
echo "Bucket created"
echo "Create API Key..."
garage key create md-drive-key
echo "API Key created"
echo "Check keys..."
garage key list
garage key info md-drive-key

echo "Don't forget to save KEY ID and SECRET KEY for later use in the .env file of the files-service!"

echo "Set bucket permissions for the API Key..."
garage bucket allow \
  --read \
  --write \
  --owner \
  md-drive-bucket \
  --key md-drive-key

echo "Check bucket info..."
garage bucket info md-drive-bucket

echo "------------- GARAGE INIT SCRIPT END -----------"