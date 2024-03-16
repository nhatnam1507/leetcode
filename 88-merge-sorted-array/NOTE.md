We start from last of nums1 array
We will have 3 pointer: 
- The `last` point to last of nums1 array. This pointer is where we will put the value into. It's value should be `m + n -1`
- The `i` point to the last member of `nums1` array. Here it's value should be `m - 1`
- The `j` point to the last member of `nums2` array. Here it's value should be `n - 1`
We travel from the `last` to begin. 
At each `last`, we will compare the value of `i` and `j`. We will take the larger value to `nums1[last]`  and increase the pointer point to this value by 1
At the end, when all value of `nums1` has been filled, we just need fill all the rest member of `nums2` to `nums1`. Because these arrays already sorted, the minium of `nums1` will alway larger than the maxium of the rest of `nums2`
