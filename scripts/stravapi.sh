#!/bin/bash

# Usage: ./stravapi.sh accessToken resourcePath arguments

if [ -z $1 ]
then
   echo "Usage: ./stravapi.sh accessToken resourcePath arguments" >&2
   exit
fi

# The first argument is the access token to use to make the request
accessToken=$1

# The second argument is the resource to retrieve
resource=$2

# The third argument is a single string with the URL-encoded arguments
# (this doesn't include the access token argument or the question mark)
arguments=$3

# Save some standard curl options into a variable
# Actually don't use any for now.
curlOptions=""

# Protocol
protocol="https"

# Domain
domain="www.strava.com"

# This is the name of the access token attribute
accessTokenAttribute="access_token"

# Stick all of the URL parts together
bigCombinedURL="$protocol://$domain/$resource?$accessTokenAttribute=$accessToken&$arguments"

# For testing, echo the URL to stderr
echo $bigCombinedURL >&2

# Put everything together into a curl call
curl $curlOptions $bigCombinedURL
