public static class Part1
{
  public static int Solve(string input)
  {
    var lines = input.Split('\n');

    var dialValue = 50;
    var zeroCount = 0;

    foreach (var line in lines)
    {
      var value = int.Parse(line.Substring(1));
      var negative = line.StartsWith("L");

      for (var i = 0; i < value; i++)
      {
        dialValue += negative ? -1 : 1;

        if (dialValue < 0) dialValue = 100 + dialValue;
        if (dialValue > 99) dialValue = dialValue - 100;
      }

      if (dialValue == 0) zeroCount++;
    }

    return zeroCount;
  }
}
