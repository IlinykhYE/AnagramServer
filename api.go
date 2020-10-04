package main

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
)
var dictionary[] string

// Load an array of words into a dictionary.
func load(c *gin.Context) {
	words := struct {
		Word[] string `json:"words"`
	}{}
	c.BindJSON(&words)

	// Add the value to the dictionary if it was not there before.
	enteredWords: for _, i := range words.Word {
		for _, j := range dictionary {
			i = strings.ToLower(i)
			if j == i {
				continue enteredWords
			}
		}
		dictionary = append(dictionary, i)
	}

	c.JSON(200, gin.H{
		"dictionary": dictionary,
	})
}

// Get anagrams in loaded dictionary.
func get(c *gin.Context){
	word := c.Query("word")
	strings.ToLower(word)
	var anograms[] string
	c1 := GetCharCount(word)

	for i:= 0; i < len(dictionary); i++{
		c2 := GetCharCount(dictionary[i])
		if reflect.DeepEqual(c1, c2) {
			anograms = append(anograms, dictionary[i])
		}
	}

	c.JSON(200, gin.H{
		"anograms": anograms,
	})
}

func GetCharCount(word string) (c map[rune]int) {
	c = make(map[rune]int)
	for _, runeValue := range word {
		if _, ok := c[runeValue]; ok {
			c[runeValue] += 1
		} else {
			c[runeValue] = 1
		}
	}
	return
}
