Assumptions:

1) File Input
   a) all will be GET URLs
   b) url inputs in a file are seperated by new line; i.e there wont be two input urls in the single line
   c) there can be empty line inbetween the set of urls
   d) url string can be malformed as well
   e) length of both files might not be same. i.e input values in both files are not equal for comparision. Adding empty lines in the end as well to make it equal
   f) one url wont be split into two lines
   g) url can be on any length but placed in only one line
   h) file can of any size
   i) file format can be of txt or csv type following above assumptions
   j) file will be readable to all users
   k) any characters (like space, ", &, %, =, ?, @) which needs encoding is already been encoded
   l) there can be empty spaces before start of url or end of url (we wont remove these spaces while giving the input)
   m) No headers/authentication is needed for making the request

   JSON COMPARE assumptions:
   a) order of array should be maintained if not its treated as a mismatch


   How to execute?

   1) install latest go version
   2) under go workspace create src, bin, lib (three directories)
   3) unzip the code and put it under src
   4) go get https://github.com/valyala/fasthttp
   5) run the main file -> go run main.go -file1 /tmp/sample.txt -file2 /tmp/sample2.txt

   output file with name stdout.log will created under root directory


   For Running Unit test:

   1) Go to test folder
   2) go test -v (for executing all the test)


   Note: if you are running from your local change the file descriptors (if needed)
   https://wilsonmar.github.io/maximum-limits/

