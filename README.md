twitter-action
---
# Auth
About the authentication see: https://developer.twitter.com/en/apps
create an account, create an app
@see https://apps.twitter.com/

# retrieve the access tokens
@see https://developer.twitter.com/en/apps

# Build
```
go get .
go build
```
# Usage
```
export TWITTER_CONSUMER_KEY=xxx
export TWITTER_CONSUMER_SECRET=xxx
export TWITTER_ACCESS_TOKEN=xxx
export TWITTER_ACCESS_SECRET=xxx
./twitter-action -message "Hello Twitter :)"

```

# Docker
```
# If building locally
docker build -t xorilog/twitter-action .

# else:
docker run --rm -e TWITTER_CONSUMER_KEY=${TWITTER_CONSUMER_KEY} \
       -e TWITTER_CONSUMER_SECRET=${TWITTER_CONSUMER_SECRET} \
       -e TWITTER_ACCESS_TOKEN=${TWITTER_ACCESS_TOKEN} \
       -e TWITTER_ACCESS_SECRET=${TWITTER_ACCESS_SECRET} \
       xorilog/twitter-action -message "Hello Twitter :)"
```
