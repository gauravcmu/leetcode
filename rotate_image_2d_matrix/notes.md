# Problem statement

Rotate a n x n matrix clockwise. (or counter clockwise)
```
[ 
  1, 2, 3 
  4, 5, 6 
  7, 8, 9 
]

to 

[
   7, 4, 1 
   8, 5, 2 
   9, 6, 3 
]
```
## Solution approach:

### Clockwise rotation by 90 degrees:
- Transpose the matrix (along diagonal, swap elements)
- Reverse elements in each row. 

### Counter clockwise rotation by 90 degrees:
- Transpose as above. 
- Reverse order of rows (not reverse elements in the row yet). 
- reverse each row elements now. 

