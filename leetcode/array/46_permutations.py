
def permute(nums):
    result = []
    # base case
    if(len(nums)==1):
        return [nums[:]]

    for i in range(len(nums)):
        n = nums.pop(0)        # 1
        perms = permute(nums)  # [[2,3], [3,2]]  
        for perm in perms:
            perm.append(n)     # [[2,3,1], [3,2,1]]  
        result.extend(perms)
        nums.append(n)

    return result


nums = [1,2,3]
print(permute(nums))