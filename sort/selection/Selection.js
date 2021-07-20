function sort(arr){
    let n = arr.length
    let min
    for (let i = 0; i < n-1; i++) {
        min = i
        for (let j = i+1; j < n; j++) {
            if (arr[j] < arr[min]) min = j
        }
        arr[i] = arr[i]^arr[min]^(arr[min]=arr[i])
    }
}

arr = [2, 5, -4, 3, 25, -14]
console.log(arr)

sort(arr)
console.log(arr)