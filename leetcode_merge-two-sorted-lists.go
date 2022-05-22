/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// https://leetcode.com/problems/merge-two-sorted-lists/
// https://stackoverflow.com/questions/28447297/how-to-check-for-an-empty-struct
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var merged *ListNode = &ListNode{}
    var mergedHead *ListNode = merged
    for (true){
        if (list1 != nil  && list2 != nil) {
            if list1.Val > list2.Val {
                merged.Val = list2.Val
                list2 = list2.Next
            } else {
                merged.Val = list1.Val
                list1 = list1.Next

            }
        } else if  list1 == nil && list2 != nil {
            merged.Val = list2.Val
            if list2.Next == nil  {
                return mergedHead
            }
            list2 = list2.Next
        } else if list2 == nil  && list1 != nil  {
            merged.Val = list1.Val
            if list1.Next == nil {
                return mergedHead
            }
            list1 = list1.Next
        } else { // both are nil
            return nil
        }
        merged.Next = &ListNode{}
        merged = merged.Next
    }
    return mergedHead // needed for compilation
}