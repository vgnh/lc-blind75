// https://leetcode.com/problems/longest-substring-without-repeating-characters
//
/*

 */

public class P02 {
    public static int lengthOfLongestSubstring(String s) {
        if (s.length() == 0) {
            return 0;
        }

        int start = 0, max = 0;
        for (var i = 0; i < s.length() - 1; i++) {
            var s2 = s.substring(start, i + 1);
            if (s2.indexOf(s.charAt(i + 1)) != -1) {
                start += s2.lastIndexOf(s.charAt(i + 1)) + 1;
            } else {
                max = Math.max(s2.length() + 1, max);
            }
        }

        return Math.max(max, 1);
    }

    public static void main(String[] args) {
        System.out.println(lengthOfLongestSubstring("abcabcbb"));
    }
}
