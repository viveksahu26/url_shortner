# url_shortner

# How URL Shortner Works !!
**URL shortner endpoint - returns a short URL**
1) go run main.go

2)  
    Enter your URL after *http://localhost:8080/short-url?longURL=*
    And let's say your URL is *http://google.com/1346461234567890123456789/get/vivekkumarsahu*
    Finally you complete URL will look like:
```http://localhost:8080/short-url?longURL=http://google.com/1346461234567890123456789/get/vivekkumarsahu```

3) You will get output
```
{"originalURL":"http://google.com/1346461234567890123456789/get/vivekkumarsahu","shortURL":"http://localhost:8080/RpP^goh8"}
```

# How to run directly from Docker Image !!
## Step:1 
    
    `docker pull viveksahu26/my-url-shortner:latest`

## Step2: 
    
    ```
    docker run --name my-short-url1 -d  -p 8090:8080  viveksahu26/my-url-shortner:latest
    ```

## OR


## Step:2

    Create container from image:

    `docker run -d --name vivek1224 -P viveksahu26/my-url-shortner`

    `docker ps`

    Under PORT section you will see something like: 

    0.0.0.0:49154->8080/tcp
    
    Your PORT Number is: 49154
    And URL is: "http://google.com/1346461234567890123456789/get/viveksahu26"

    Paste it to your Browser : 
    http://localhost:<replace_by_your_port_no>/short-url?longURL=http:<enter_your_URL>
    
    Example:
    Now, copy and paste to your browser:
    
    http://localhost:49154/short-url?longURL=http://google.com/1346461234567890123456789/get/viveksahu26
    


    DockerHub: https://hub.docker.com/repository/docker/viveksahu26/my-url-shortner

    *NOTE*: Do not forget to change port number(49154) and also replace this url(http://google.com/1346461234567890123456789/get/viveksahu26) with your URL.
