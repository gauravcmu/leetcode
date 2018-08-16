import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string,0)
    //handle empty tree. 
    if root == nil {
        return []string{}
    }
    //pass around path which gets built on each call to recursive bh
    return bh(&res, root, "") 
}

func bh(res *[]string, root *TreeNode, path string) []string {
    if root == nil {
        return *res 
    }
    
    //if first node - add it to path
    if path == "" {
        path = strconv.Itoa(root.Val)
    } else {
        //not first node, add arrow and value
        path = path + "->" + strconv.Itoa(root.Val) 
    }
    
    //Final result gets updated only when leaf is hit
	if root.Left == nil && root.Right == nil  {
		*res = append(*res, path)
		 return *res
    }
    
    //Build path for left and right subtrees 
    bh(res, root.Left, path)
    bh(res, root.Right, path)
    return *res
}
