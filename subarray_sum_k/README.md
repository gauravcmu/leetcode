Find cumulative sum till index i. 

If mapping has sum_so_far - k -> this implies sum exists. 
Count how many times you get ^ and that will be the answer. 
Ensure 0 is also an entry in mapping as this would be the case when a number is = k 
For example [3, 1, 2] and k = 3. So for first element sum_so_far = 3 and sum_so_far - k = 0. 


