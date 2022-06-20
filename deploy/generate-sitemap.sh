#! /bin/bash

set -e

BASE_URL=$1
BASE_API_URL=$2

API_URL="$BASE_API_URL/api/sitemap?baseUrl=$BASE_URL"

xml=$(curl -X GET --header "Accept: */*" $API_URL)

echo $xml >> public/sitemap.xml