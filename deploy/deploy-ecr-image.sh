#! /bin/bash

set -e

IMAGE_TAG="$GITHUB_SHA"
IMAGE_ARN=$1

aws ecr get-login-password --region ap-south-1 | docker login --username AWS --password-stdin 569704406482.dkr.ecr.ap-south-1.amazonaws.com

docker build -t canopas-website-ssr-app:$IMAGE_TAG .

docker tag canopas-website-ssr-app:$IMAGE_TAG $IMAGE_ARN:$IMAGE_TAG

docker push $IMAGE_ARN:$IMAGE_TAG

# delete untagged images
aws ecr describe-repositories --output text | awk '{print $5}' | egrep -v '^$' | while read line; do
    repo=$(echo $line | sed -e "s/arn:aws:ecr.*\///g")
    aws ecr list-images --repository-name $repo --filter tagStatus=UNTAGGED --query 'imageIds[*]' --output text | while read imageId; do
        aws ecr batch-delete-image --repository-name $repo --image-ids imageDigest=$imageId
    done
done
