import "container/list"

type LRUCache struct {
    nodelist * list.List 
    hashmap map[int]*list.Element 
    capacity int 
}

type info struct {
    value int 
    key int 
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
        nodelist: list.New(),
        hashmap: make(map[int]*list.Element, capacity),
        capacity: capacity, 
    }
}


func (this *LRUCache) Get(key int) int {
    if value, ok := this.hashmap[key]; ok {
        //move this node to the back. 
        this.nodelist.MoveToBack(value)
        return value.Value.(info).value 
    }
    return -1 
}


func (this *LRUCache) Put(key int, value int)  {
     if e, ok := this.hashmap[key]; ok {
        e.Value = info {key:key, value:value}
        this.nodelist.MoveToBack(e)
         return 
     }
    
    if this.capacity <= this.nodelist.Len() {
        //reached capacity 
        element := this.nodelist.Front()
        this.nodelist.Remove(this.nodelist.Front())
        delete(this.hashmap, element.Value.(info).key) 
    }
    
    e := this.nodelist.PushBack(info{value: value, key: key})
    this.hashmap[key] = e 
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
