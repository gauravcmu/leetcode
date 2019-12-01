Input: ")()())"
Output: 4
Explanation: The longest valid parentheses substring is "()()"


Use stack to find if parentheses match. 

Push index of the opening bracket onto stack. This way you can find length at any point using the index. 


Go over the string, byte by byte. 
	- If the byte is ( - then push the index of this byte onto the stack. 
	- if it is closing bracket ) - then do the following. 
		- Pop the top of the stack. 
		- till this point we have a balanced string. 
		- length of this string is current index - stack top 
                - or it is index + 1 if the stack is empty - eg. ()() 
		- finally maximize length - between lenght and temp length from above. 
	- 			

Approach 4: Without extra space
Algorithm

In this approach, we make use of two counters leftleft and rightright. First, we start traversing the string from the left towards the right and for every \text{‘(’}‘(’ encountered, we increment the leftleft counter and for every \text{‘)’}‘)’ encountered, we increment the rightright counter. Whenever leftleft becomes equal to rightright, we calculate the length of the current valid string and keep track of maximum length substring found so far. If rightright becomes greater than leftleft we reset leftleft and rightright to 00.

Next, we start traversing the string from right to left and similar procedure is applied.

public class Solution {
    public int longestValidParentheses(String s) {
        int left = 0, right = 0, maxlength = 0;
        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) == '(') {
                left++;
            } else {
                right++;
            }
            if (left == right) {
                maxlength = Math.max(maxlength, 2 * right);
            } else if (right >= left) {
                left = right = 0;
            }
        }
        left = right = 0;
        for (int i = s.length() - 1; i >= 0; i--) {
            if (s.charAt(i) == '(') {
                left++;
            } else {
                right++;
            }
            if (left == right) {
                maxlength = Math.max(maxlength, 2 * left);
            } else if (left >= right) {
                left = right = 0;
            }
        }
        return maxlength;
    }
}

- why two passes needed: 

Consider this string "((()" - when you go left to right, you will never calculate maxlength as left is always more than right. Thus right to left calculates such a case and vice-versa. 

