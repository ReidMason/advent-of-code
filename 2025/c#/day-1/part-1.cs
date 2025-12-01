public static class Part1
{
  public static int Solve(string input)
  {
    var lines = input.Split('\n');

    var dialValue = 50;
    var zeroCount = 0;

    var values = lines.Select(x => int.Parse(x.Substring(1)) * (x.StartsWith("L") ? -1 : 1));

    foreach (var value in values)
    {
      dialValue += value;
      if (dialValue % 100 == 0) zeroCount++;
    }

    return zeroCount;
  }
}
