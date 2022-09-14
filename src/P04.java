// https://leetcode.com/problems/container-with-most-water/
// Dynamic Programming
/*
   Start from the ends.
   Move the shortest of the two ends forward. If start < end, start++ else end--.
   Keep calculating most volume.
*/

public class P04 {
    public static int maxArea(int[] height) {
        int start = 0, end = height.length - 1;
        var most = Math.min(height[start], height[end]) * (end - start);
        while (start < end) {
            if (height[start] < height[end]) {
                start++;
            } else {
                end--;
            }
            var vol = Math.min(height[start], height[end]) * (end - start);
            if (vol > most) {
                most = vol;
            }
        }
        return most;
    }

    public static void main(String[] args) {
        System.out.println(maxArea(new int[]{1, 8, 6, 2, 5, 4, 8, 3, 7}));
    }
}
