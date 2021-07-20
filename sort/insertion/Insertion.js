function sort(arr) {
    let key, j
    for (let i = 1; i < arr.length; i++) {
        key = arr[i]
        j = i - 1
        while (j >= 0 && arr[j] > key){
            arr[j+1] = arr[j]
            j = j - 1
        }
        arr[j+1] = key
    }
}

const arr = [5, 2, 9, 41, -6, 58, 12, 6]
sort(arr)
console.log(arr)