namespace lc_blind75;

class P05
{
    /// <remarks>
    ///     https://leetcode.com/problems/longest-palindromic-substring/
    ///
    ///     Dynamic Programming:
    ///         Odd Palindromes
    ///         Start from Substring(i, i), <- Substring(L, R)
    ///         Keep decreasing L, increasing R
    ///         Store the longest as long as charAt(L) == charAt(R)
    ///         Even Palindromes
    ///         Start from Substring(i, i+1), <- Substring(L, R)
    ///         Keep decreasing L, increasing R
    ///         Store the longest as long as charAt(L) == charAt(R)
    /// </remarks>
    public string LongestPalindrome(string s)
    {
        var longest = "";

        for (var i = 0; i < s.Length; i++)
        {
            // Odd palindromes
            int l = i, r = i;
            while (l >= 0 && r < s.Length && s[l] == s[r])
            {
                if (r - l + 1 > longest.Length)
                {
                    longest = s.Substring(l, r + 1 - l);
                }
                l--;
                r++;
            }

            // Even palindromes
            l = i;
            r = i + 1;
            while (l >= 0 && r < s.Length && s[l] == s[r])
            {
                if (r - l + 1 > longest.Length)
                {
                    longest = s.Substring(l, r + 1 - l);
                }
                l--;
                r++;
            }
        }

        return longest;
    }
}
