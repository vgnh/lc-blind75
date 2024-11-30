namespace lc_blind75;

class P11
{
    /// <remarks>
    ///     https://leetcode.com/problems/container-with-most-water/
    ///
    ///     Dynamic Programming:
    ///         Start from the ends.
    ///          Move the shortest of the two ends forward. If start < end, start++ else end--.
    ///          Keep calculating most volume.
    /// </remarks>
    public int MaxArea(int[] height)
    {
        var start = 0;
        var end = height.Length - 1;
        var most = Math.Min(height[start], height[end]) * (end - start);
        while (start < end)
        {
            if (height[start] < height[end])
            {
                start++;
            }
            else
            {
                end--;
            }

            var vol = Math.Min(height[start], height[end]) * (end - start);
            if (vol > most)
            {
                most = vol;
            }
        }
        return most;
    }
}
