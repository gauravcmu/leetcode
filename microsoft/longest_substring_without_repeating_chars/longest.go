package longest_substring_without_repeating_chars

/*

Given a string, find the longest substring without repeating characters.

For example, for string “abccdefgh”, the longest substring is “cdefgh”.
 */

/*
	abccdefgh

	^
	i

	iterate over the list - build a map. key being character and value being bool.
	as you encounter a value again - maximize length of max and then reset

	abac defgh ch
     ^       ^
	// two pointers start and end.
    for start = 0 to end
		for end = start + 1; to end ; {
			create new map. add key as value at end if not already in map.
			else if value already in map.
				- move the window.
				start = start + 1

		}
 */

