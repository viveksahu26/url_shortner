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

4) Check the health of program

    http://localhost:8080/health

5)  Go to browser:
    
    Enter your URL after *http://localhost:8080/sort-url?longURL=*

    And let's say your Long URL: *http://google.com/1346461234567890123456789/get/viveksahu26*

    Finally you complete URL in the browser will look like:

    http://localhost:8080/short-url?longURL=http://google.com/1346461234567890123456789/get/viveksahu26

3) You will get output

    ```
    {"originalURL":"http://google.com/1346461234567890123456789/get/viveksahu26","shortURL":"http://localhost:8080/xtNFxaBwCG"}
    ```

## Steps to reproduce Using Docker Image
### Step:1
Docker Image: viveksahu26/shortingurl:v3.0.0

Pull docker image:

    docker pull viveksahu26/shortingurl:v3.0.0

### Step:2
Run container in non-interactive mode

    docker run -d --name vivek -p 3000:8080 viveksahu26/shortingurl:v3.0.0

Run container in interactive mode

    docker run --name vivek -p 3000:8080 viveksahu26/shortingurl:v3.0.0

### Step:3 
Search on browser

    http://localhost:3000/short-url?longURL=http://google.com/1346461234567890123456789/get/viveksahu26

Where, longURL=http://google.com/1346461234567890123456789/get/viveksahu26
And port: 3000

To check health of app:

    http://localhost:3000/health
