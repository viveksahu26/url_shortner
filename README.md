# URL SHORTNER
**URL shortner endpoint - returns a short URL**
It is a service which takes long URL from the user and returns Short URL. 

## How URL Shortner Works !!
It replaces long URL by randomly generated characters of size 10. 
It stores both Short and Long URL in the local file, `url.properties`.
If `url.properties`  file is not present then creates new, otherwise write to the existing ones.
The advantage of saving Short URL and Long URL in the file is to retrieved that same Short URL correspondiing to Long URL. 


## Steps to reproduce it
1) Clone the repo:
    
    `git clone https://github.com/viveksahu26/url_shortner.git`

2) Jump to the directory.

    `cd url_shortner`

3) Execute main program. 

    `go run main.go`

*NOTE:*: Make sure that Port 8080 is free. For customize port you need to make changes in the codebase.

4)  Go to browser:
    
    Enter your URL after *http://localhost:8080/sort-url?longURL=*

    And let's say your Long URL: *http://google.com/1346461234567890123456789/get/vivekkumarsahu*

    Finally you complete URL in the browser will look like:

    http://localhost:8080/enterLongURL?longURL=http://google.com/1346461234567890123456789/get/vivekkumarsahu

3) You will get output

    ```
    {"originalURL":"http://google.com/1346461234567890123456789/get/vivekkumarsahu","shortURL":"http://localhost:8080/RpP^goh8"}
    ```
