public class LC209 {
    public static int minSubArrayLen(int target, int[] nums) {
        int l = 0, sum = 0;
        var atLeastOne = false;
        var min = nums.length;
        for (var i = 0; i < nums.length; i++) {
            sum += nums[i];
            while (sum >= target) {
                if (!atLeastOne)
                    atLeastOne = true;
                min = Math.min(min, i + 1 - l);
                sum -= nums[l];
                l++;
            }
        }
        return atLeastOne ? min : 0;
    }

    public static void main(String[] args) {
        System.out.println(minSubArrayLen(7, new int[]{2, 3, 1, 2, 4, 3}));
    }
}
