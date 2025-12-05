namespace day_5;

public static class Part1
{
    public static void Test()
    {
        var input = @"3-5
10-14
16-20
12-18

1
5
8
11
17
32";

        var result = Solve(input);
        var expected = 3;

        var resText = expected != result ? "Failed" : "Passed";
        Console.WriteLine($"Test {resText} Result: {result}");
    }

    public static int Solve(string input)
    {
        var lines = input
            .Trim()
            .Split("\n");

        var total = 0;
        var ranges = new List<(long start, long end)>();
        var collectingRanges = true;
        foreach (var line in lines)
        {
            if (line.Length == 0)
            {
                collectingRanges = false;
                continue;
            }

            if (collectingRanges)
            {
                var parts = line.Split("-").Select(x => long.Parse(x)).ToList();
                ranges.Add((parts[0], parts[1]));
                continue;
            }

            var ingredientId = long.Parse(line);
            foreach (var range in ranges)
            {
                if (range.start <= ingredientId && range.end >= ingredientId)
                {
                    total++;
                    break;
                }
            }
        }

        return total;
    }
}