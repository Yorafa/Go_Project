# Go Project

This repo will collect some projects of go.

## Project

1. Guessing
   - a command line program random a number from 0 to 100 (exclusive) and ask user to guess
2. Command Line Dictionary
   - use caiyun translator's api to translate the command line argument
3. SimpleSocketProxy
   - use sock5 to do proxy

## How to do 2nd project dictionary?

1. go to the translation website, and use inspect to find the network. It will list all the request and response.
2. Find the one with correct response we want (the translation response), and copy its' cUrl.
3. go to [curlconverter](https://curlconverter.com/) and convert it to go (or other language you want)
4. use above step to set up request struct
5. use http module to send request, and get the response
6. use [json2go](https://oktools.net/json2go) convert the response to struct so that we can access easily