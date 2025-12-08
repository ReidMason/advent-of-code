using System.Text.RegularExpressions;

namespace day_6;

public static class Part1
{
    public static void Test()
    {
        var input = @"123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  ";
        var result = Solve(input);
        var expected = 4277556;

        var resText = expected != result ? "Failed" : "Passed";
        Console.WriteLine($"Test {resText} Result: {result}");
    }

    public static long Solve(string input)
    {
        var lines = input.Trim()
            .Split("\n");

        long total = 0;

        var columnValues = new List<List<long>>();
        foreach (var line in lines)
        {
            var columns = line.Split().Where(x => !string.IsNullOrEmpty(x)).ToList();
            for (int i = 0; i < columns.Count(); i++)
            {
                try
                {
                    var rawColumnValue = columns[i];

                    if (rawColumnValue.Trim() == "*")
                    {
                        total += columnValues[i].Aggregate((a, b) => a * b);
                        continue;
                    }

                    if (rawColumnValue.Trim() == "+")
                    {
                        total += columnValues[i].Sum();
                        continue;
                    }

                    var columnValue = int.Parse(rawColumnValue);
                    if (i == columnValues.Count)
                    {
                        columnValues.Add(new List<long> {columnValue});
                    }
                    else
                    {
                        columnValues[i].Add(columnValue);
                    }
                }
                catch (Exception e)
                {
                    Console.WriteLine(e);
                }
            }
        }

        return total;
        // 36961868260
    }
}