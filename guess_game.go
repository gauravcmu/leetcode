/* The guess API is defined in the parent class GuessGame.
   @param num, your guess
   @return -1 if my number is lower, 1 if my number is higher, otherwise return 0
      int guess(int num); */

public class Solution extends GuessGame {
    public int guessNumber(int n) {
        return myguess(1, n);
    }
    
    public int myguess(int start, int end) {
       // int mid = start + (end-start) / 2;
        int mid = (end+start) >>> 1;
        System.out.printf("mid:%d\n", mid);
        if (guess(mid) == 0) {
            return mid;
        } 
        if (guess(mid) == -1) {
            return myguess(start, mid - 1);
        }
        if (guess(mid) == 1) {
            return myguess(mid+1, end);
        }
        return mid;
    }
}
