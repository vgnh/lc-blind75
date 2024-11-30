namespace lc_blind75;

class P03
{
    /// <remarks>
    ///     https://leetcode.com/problems/longest-substring-without-repeating-characters
    /// </remarks>
    public int LengthOfLongestSubstring(string s)
    {
        if (s.Length == 0)
        {
            return 0;
        }

        var start = 0;
        var max = 0;
        for (var i = 0; i < s.Length - 1; i++)
        {
            var s2 = s.Substring(start, i + 1 - start);
            if (s2.Contains(s[i + 1]))
            {
                start += s2.LastIndexOf(s[i + 1]) + 1;
            }
            else
            {
                max = Math.Max(s2.Length + 1, max);
            }
        }

        return Math.Max(max, 1);
    }
}
