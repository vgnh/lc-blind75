namespace lc_blind75;

class Program
{
    static void Main(string[] args)
    {
        var p01 = new P01().TwoSum([2, 7, 11, 15], 9);
        Console.WriteLine($"P01: [{string.Join(",", p01)}]");

        var p03 = new P03().LengthOfLongestSubstring("abcabcbb");
        Console.WriteLine($"P03: {p03}");

        var p05 = new P05().LongestPalindrome("racecar");
        Console.WriteLine($"P05: {p05}");

        var p11 = new P11().MaxArea([1, 8, 6, 2, 5, 4, 8, 3, 7]);
        Console.WriteLine($"P11: {p11}");
    }
}
