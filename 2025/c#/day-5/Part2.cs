using System.Numerics;

namespace day_5;

public class Part2
{
    public static void Test()
    {
        var input = @"2-2
3-5
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
        long expected = 15;

        var resText = expected != result ? "Failed" : "Passed";
        Console.WriteLine($"Test {resText} Result: {result}");
    }

    public static long Solve(string input)
    {
        var lines = input
            .Trim()
            .Split("\n");

        var rangeNumbers = new List<(long value, int id)>();
        for (var i = 0; i < lines.Length; i++)
        {
            var line = lines[i];
            if (line.Length == 0) break;

            var parts = line.Split("-").Select(x => long.Parse(x)).ToList();
            rangeNumbers.Add((parts[0], i));
            rangeNumbers.Add((parts[1], i));
        }

        rangeNumbers = rangeNumbers.OrderBy(x => x.value).ToList();

        long total = 0;
        var bank = new List<long>();
        long start = 0;
        long prevEnd = 0;
        foreach (var number in rangeNumbers)
        {
            if (bank.Count == 0) start = number.value;
            if (!bank.Remove(number.id)) bank.Add(number.id);
            if (bank.Count == 0)
            {
                total += (number.value - start);
                if (prevEnd != start) total++;
                prevEnd = number.value;
            }
        }

        return total;
    }
}