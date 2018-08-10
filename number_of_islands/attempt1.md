This solution works for most of the cases except a case where connection to island happens a bit later..
Thus, it is not accepted in lc..

Approach
- Go over the grid, element by element. 
- build a set of islands. 
- for each element, check if there exists an island it is adjoining 
- if it adjoins, then add it as a point in the island. 
- else create a separate island and add this point to that island
- finally count the number of islands in this set. 

Issue:
Consider this test case. 
111
010
111

When checking for element 2,0 - there is no island adjoining it, thus the algorithm will add it as a new island. 
This is actually part of same island but the current check does not do that and thus returns one more island than the answer 
