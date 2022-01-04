/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    //create a new list to store the final result
    var list3 *ListNode
    if list1==nil{
        return list2
    }else if list2 == nil {
        return list1
    }else if list1.Val <= list2.Val {
        list3=list1
        list1=list1.Next
    }else {
        list3=list2
        list2=list2.Next
    }
    //traverse through list1, and list2 and combine the two together
    for list1 !=nil || list2 != nil {
        if ( list1 != nil && list2 !=nil && list1.Val <= list2.Val) || list2 == nil{
             list3.Next = list1
             list1=list1.Next
        }ele {
           list3.Next=list2
        }
        //moving list3 head to the next
        list3 = list3.Next
    }
    return list3
}
