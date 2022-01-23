/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode = nil
    for head != nil {
        //save the temporary next node the current head is pointing to
        var next *ListNode = head.Next
        //reverse the link direction
        head.Next=prev
        //move the prev pointer to the head node
        prev=head
        //move the head pointer to the next node
        head = next
    }
    return prev
}
