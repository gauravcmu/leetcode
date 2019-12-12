Construct quad trees given a grid of N*N boolean values 
`
/*
// Definition for a QuadTree node.
class Node {
    public boolean val;
    public boolean isLeaf;
    public Node topLeft;
    public Node topRight;
    public Node bottomLeft;
    public Node bottomRight;

    public Node() {}

    public Node(boolean _val,boolean _isLeaf,Node _topLeft,Node _topRight,Node _bottomLeft,Node _bottomRight) {
        val = _val;
        isLeaf = _isLeaf;
        topLeft = _topLeft;
        topRight = _topRight;
        bottomLeft = _bottomLeft;
        bottomRight = _bottomRight;
    }
};
*/`

Solution: 

Write a recursive function helper, that takes as input the start point of the grid in the form of row and col and the 
length of the grid (N)

Pseudocode: 
```
   func helper(grid [][]bool, row, col, len int) Node {
        // check if this grid has only one element, then it is a leaf. 
        // build a leaf node and return that node. 

        // make 4 recursive calls to get leftUpper/Lower, RightUpper/Lower nodes. 
        // To make calls, use the row and col, modify row / col with row+len/2 and len = len/2 

        // Now, check all 4 nodes returned are leaves. If so, do they all have same values. If so:
             create a node, set value and return as a leaf node. 
          else
             create a node, set its leftUpper/Lower/Right etc. with the above 4 nodes. 
   }
```