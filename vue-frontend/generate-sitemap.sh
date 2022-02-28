#! /bin/bash

set -e

BASE_URL=$1
BASE_API_URL=$2

API_URL="$BASE_API_URL/sitemap?baseUrl=$BASE_URL"

$(curl -X GET --header "Accept: */*" $API_URL)