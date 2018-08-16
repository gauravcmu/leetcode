/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode *getIntersectionNode(struct ListNode *headA, struct ListNode *headB) {
    int len1 = 0;
    int len2 = 0;
    int i = 0; 
    struct ListNode * temp = headA;

    
    if (headA == NULL || headB == NULL ) {
       return NULL; 
    }
   
    while (temp != NULL) {
        temp = temp->next; 
        len1++;
    }
    printf("len1 %d\n", len1); 
    temp = headB;     
    while (temp != NULL) {
        temp = temp->next; 
        len2++;
    }
    printf("len2 %d\n", len2);
        
    if (len1 > len2) {
        for (i=0; i<len1-len2; i++) {
            headA = headA->next; 
        }
    } 
    if (len2 > len1) {
        for (i=0; i<len2-len1; i++) {
            headB = headB->next; 
        }
    }    
    while (headA && headB) {
       if (headA == headB) {
           return headA;
       }  else {
            headA = headA->next; 
            headB = headB->next; 
       }
    }
    return NULL ;
    
}
