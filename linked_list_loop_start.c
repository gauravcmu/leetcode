/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode *detectCycle(struct ListNode *head) {
    struct ListNode * slow = head;
    struct ListNode * fast = head;
    if (head == NULL) {
        return NULL;
    } 
    while (fast && slow && fast->next) {
        fast = fast->next->next;
        slow = slow->next;
        if (fast == slow) {
            slow = head; 
            while (slow != fast) {
                fast = fast->next;
                slow = slow->next;
            } 
            if (slow == fast) {
                return fast;
            }
        }
    }
    return NULL;
}
