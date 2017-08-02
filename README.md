# jexia-challenge
This challenge consists in translate a name written in English to Klingon and find out its species using http://stapi.co. 

All code was made in Go (Golang - Version: go1.7.4 darwin/amd64) with a tradicional and vanilla packages

Used api:
Base character, returned in search results
POST - http://stapi.co/api/v1/rest/character/search

Full character, returned when queried using UID (for Species purposes)
GET - http://stapi.co/api/v1/rest/character



RULES:
1. This a public Git repository with all commit history included;
2. It is runnable on a Unix bash. The name in English will be passed as the first parameter. It might contain spaces in case of composed names. Example:
```
$ go run main.go Sherman P. Raymond II
```
3. Was considered each Klingon cognate letter a valid correspondence to an English letter.
For example, D is a valid correspondence of d on so on. Was noticed that some letters are missing which means they are not translatable for this test purposes, so was ignored whole input;
4. The output contain:
a. The translated name in Klingon written using the correspondent hexadecimal numbers according to the given table. Format:
i. Each hexadecimal number should be separated from each other  using a single space;
ii. If the translated name has spaces, use 0x0020 for representing each space character;
b. The species of the given Star Trek character name using the API;
c. The translated name and the species name separated by a new line;

Output example:
```
$  go run main.go Ulani Belor
0xF8E5 0xF8D9 0xF8D0 0xF8DB 0xF8D7 0x0020 0xF8D1 0xF8D4 0xF8D9 0xF8DD 0xF8E1 
Cardassian 
```
