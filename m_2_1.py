from heapq import heapify, heappop, heappush


def solution(A:list):
    # write your code in Python (Python 3.6)
    original =  0
    for a in A:
        original += a
    max_heap = [-a for a in A]
    heapify(max_heap)
    filter_num = 0
    lose = 0
    half = original/2
    while lose < half:
        elem = -heappop(max_heap)
        lose += elem/2
        filter_num += 1
        heappush(max_heap, -elem/2)
    return filter_num

print(solution([5,19,8,1]))