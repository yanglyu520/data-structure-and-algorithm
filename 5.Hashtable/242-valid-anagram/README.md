https://leetcode.com/problems/valid-anagram/

## Question
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

 


## Tips
1. Solution1: using sort to sort string s and t, and then compare if sorted list s1,t1 is the same. Time complexity is O(nlog(n))
2. Solution2: Using hashmap to compare
remember in golang, rune is a character type.
