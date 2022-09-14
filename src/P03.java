// https://leetcode.com/problems/longest-palindromic-substring/
// Dynamic Programming
/*
   Odd Palindromes
   Start from substring(i, i), <- substring(L, R)
   Keep decreasing L, increasing R
   Store the longest as long as charAt(L) == charAt(R)

   Even Palindromes
   Start from substring(i, i+1), <- substring(L, R)
   Keep decreasing L, increasing R
   Store the longest as long as charAt(L) == charAt(R)
*/

public class P03 {
    public static String longestPalindrome(String s) {
        var longest = "";
        for (var i = 0; i < s.length(); i++) {
            // Odd palindromes
            int l = i, r = i;
            while (l >= 0 && r < s.length() && s.charAt(l) == s.charAt(r)) {
                if (r - l + 1 > longest.length()) {
                    longest = s.substring(l, r + 1);
                }
                l--;
                r++;
            }

            // Even palindromes
            l = i;
            r = i + 1;
            while (l >= 0 && r < s.length() && s.charAt(l) == s.charAt(r)) {
                if (r - l + 1 > longest.length()) {
                    longest = s.substring(l, r + 1);
                }
                l--;
                r++;
            }
        }
        return longest;
    }

    public static void main(String[] args) {
        System.out.println(longestPalindrome("racecar"));
    }
}
