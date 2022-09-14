# url_shortner

# How URL Shortner Works !!
**URL shortner endpoint - returns a short URL**
1) go run main.go

2)  
    Enter your URL after *http://localhost:8080/short-url?longURL=*
    And let's say your URL is *http://google.com/1346461234567890123456789/get/vivekkumarsahu*
    Finally you complete URL will look like:
```http://localhost:8080/short-url?longURL=http://google.com/1346461234567890123456789/get/who_is_who```

3) You will get output
```
{"originalURL":"http://google.com/1346461234567890123456789/get/vivekkumarsahu","shortURL":"http://localhost:8080/RpP^goh8"}
```
