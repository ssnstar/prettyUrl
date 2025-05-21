# prettyurl is a url shortner that exposes following APIs
#   http://{{HOST}}/shorturl 
#   Moto : To create a tinyUrl (8 char long) from a long URL. 
#   Method : POST
#   Body it takes : {"url": "long_url"}
#   TEST : curl -k http://localhost/shorturl -d '{"url":"http://www.youtube.com"}'

#   http://{{HOST}}/redirect
#   Moto : To use a tinyUrl (8 char long) to redirect back a long URL.
#   Method : POST
#   Body it takes : {"url": "short_url"}
#   TEST : curl -k http://localhost/shorturl -d '{"url":"http://www.youtube.com"}'

## ASSUMPOTIONS : 
## Short URL length ( 8 – 10 ) + domain
## Long URL length ( 2 Bytes) 
## SIMPLE STORAGE OR CACHE WHICH IS USED ONLY IN RUNTIME. RESTART OF SERVICE WILL LOOSE THE DATA.

### STEPS to BUILD
### Go build
    # Target is alpine linux, so build as rquired.
    env GOOS=linux GOARCH=386 go build main.go
### Docker Build 
    docker build -t prettyurl:shan .
### Docker Run
    docker run -p80:80 -it prettyurl:shan
