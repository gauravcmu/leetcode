### problem 
Given a dictionary that contains mapping of employee and his manager as a number of (employee, manager) pairs like below.
{ "A", "C" },
{ "B", "C" },
{ "C", "F" },
{ "D", "E" },
{ "E", "F" },
{ "F", "F" } 

In this example C is manager of A, 
C is also manager of B, F is manager 
of C and so on.

### Solution 

1. Build a mapping which maps manager to employees under her. 
```mapping:map[E:[D] C:[A B] F:[C E]]
2. Find head of the org. 
3. Recurse from Head of the org and call helper. 
4. helper() 
	- if no reports for head in mapping, return
	- else add each report to the result map for this head. 
	- recursively call helper for each employee reporting to this head. 
	- now again add the result slice for each employee to the result set for this head. 
	
