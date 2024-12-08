class Solution:
    def canJump(self, nums: List[int]) -> bool:
        if len(nums) == 1:
            return True
        max_idx = nums[0]
        for idx in range(len(nums)-1):
            if idx <= max_idx:
                max_idx = max(max_idx, idx + nums[idx])
            if max_idx >= len(nums) - 1:
                return True
        return False