def sort(arr):
    n = len(arr)
    for i in range(n-1):
        minimum = i
        for j in range(i+1, n):
            if arr[j] < arr[minimum]:
                minimum = j
        arr[i], arr[minimum] = arr[minimum], arr[i]


# myList = ['these', 'are', 'unsorted', 'words']
myList = [3.9, 9, -5, 6, 25, -14, 32]

sort(myList)
print('Sorted: ', myList)
