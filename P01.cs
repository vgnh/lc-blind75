namespace lc_blind75;

class P01
{
    /// <remarks>
    ///     https://leetcode.com/problems/two-sum
    /// </remarks>
    public int[] TwoSum(int[] nums, int target)
    {
        var numIndexMap = new Dictionary<int, int>();
        for (var i = 0; i < nums.Length; i++)
        {
            var a = target - nums[i];
            if (numIndexMap.TryGetValue(a, out var _))
            {
                return [numIndexMap[a], i];
            }
            numIndexMap[nums[i]] = i;
        }
        throw new Exception("No two sum solution");
    }
}
