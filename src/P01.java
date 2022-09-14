// https://leetcode.com/problems/two-sum
//
/*

 */

import java.util.Arrays;
import java.util.HashMap;

public class P01 {
    public static int[] twoSum(int[] nums, int target) {
        var numIndexMap = new HashMap<Integer, Integer>();
        for (var i = 0; i < nums.length; i++) {
            var complement = target - nums[i];
            if (numIndexMap.containsKey(complement)) {
                return new int[]{numIndexMap.get(complement), i};
            }
            numIndexMap.put(nums[i], i);
        }
        return null;
    }

    public static void main(String[] args) {
        System.out.println(Arrays.toString(twoSum(new int[]{2, 7, 11, 15}, 9)));
    }
}
