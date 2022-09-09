from requests import head


class Node:
    def __init__(self, val=None, next=None):
        self.val = val
        self.next = next
class Solution:
    def insert(self, head: 'Node', insertVal: int) -> 'Node':
        newNode = Node(val=insertVal)
        if head is None:
            head = newNode
            newNode.next = head
            return head
        if head.next is head:
            head.next = newNode
            newNode.next = head
            return head
        cursor = head
        while cursor.next is not head:
            if cursor.val <= insertVal and (cursor.next.val >= insertVal or cursor.val > cursor.next.val):
                temp = cursor.next
                cursor.next = newNode
                newNode.next = temp
                return head
            if cursor.val > cursor.next.val and cursor.next.val > insertVal:
                temp = cursor.next
                cursor.next = newNode
                newNode.next = temp
                return head
            cursor = cursor.next
        cursor.next = newNode
        newNode.next = head
        return head


head = Node(3)
a = Node(4)
b = Node(1)
head.next = a
a.next = b
b.next = head

s = Solution()
s.insert(head, 2)