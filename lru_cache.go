import "container/list"

type element struct {
    Key int 
    Value int 
}

type LRUCache struct {
    List *list.List 
    Hash map[int]*list.Element
    Capacity int 
}


func Constructor(capacity int) LRUCache {
    l:= list.New() 
    h:= make(map[int]*list.Element, 0)
    return LRUCache{List: l, Hash: h, Capacity: capacity}
}


func (this *LRUCache) Get(key int) int {
    if e, ok := this.Hash[key]; ok  {
        this.List.MoveToBack(e)
        return e.Value.(element).Value
    }    
    return -1 
}

/*
["LRUCache","get","put","get","put","put","get","get"]
[[2],        [2], [2,6], [1], [1,5], [1,2],[1],  [2]]
              -1          -1              
*/

func (this *LRUCache) Put(key int, value int)  {
    if e, ok := this.Hash[key]; ok  {
        e.Value = element{Key: key, Value:value} 
        this.List.MoveToBack(e)
        return 
    }    

    if this.List.Len() >= this.Capacity {
        //LRU is at the front of the list
        delete(this.Hash, this.List.Front().Value.(element).Key) 
        this.List.Remove(this.List.Front())
    } 
    
    e := this.List.PushBack(element{Key: key, Value:value})
    this.Hash[key] = e 
}
