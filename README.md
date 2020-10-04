# AnagramServer
 HTTP service for fast search of anagrams in the dictionary.
 
### Endpoint description
***1. load*** <br>
 Load an array of words into a dictionary. <br>
 duplicate values cannot be loaded. <br> 
  *POST localhost:8080/load* <br>
  
 anagram loading example:
 
 body JSON:
 {
 
    "words": ["foobar", "aabb", "baba", "boofar", "test"]
    
 }

 response:
	
 {
	
    "dictionary": [
        "foobar",
        "aabb", 
        "baba", 
        "boofar",
        "test"
    ]
}

***2. get***

Search for anagrams by word in the loaded dictionary.

*GET localhost:8080/get?word=abba*

response:

{

    "anograms": [
        "aabb",
        "baba"
    ]
}

 
