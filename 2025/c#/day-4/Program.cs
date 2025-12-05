void Test()
{
  var input = @"..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.";

  var result = Solve(input);
  var expected = 13;

  var resText = expected != result ? "Failed" : "Passed";
  Console.WriteLine($"Test {resText} Result: {result}");
}

int Solve(string input)
{
  var matrix = input
    .Replace("\r", "")
    .Trim()
    .Split("\n")
    .Select(x =>
      x.Select(y => y.ToString()).ToList())
    .ToList();

  var lineLength = matrix[0].Count();
  var offsets = new List<(int x, int y)> { (-1, -1), (-1, 0), (-1, 1), (0, -1), (0, 1), (1, -1), (1, 0), (1, 1) };

  var total = 0;

  for (int i = 0; i < matrix.Count(); i++)
  {
    for (int j = 0; j < lineLength; j++)
    {
      try
      {
        var visiting = matrix[i][j];
        if (visiting != "@") continue;
        var neighbours = 0;

        foreach (var offset in offsets) {
          try {
            var value = matrix[i+offset.x][j+offset.y];
            if (value == "@") neighbours++;
          }
          catch {
            continue;
          }
        }

        if (neighbours < 4) total++;
      }
      catch (ArgumentOutOfRangeException)
      {
        Console.WriteLine($"{i},{j}");
      }
    }
  }

  return total;
}


void Test2()
{
  var input = @"..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.";

  var result = Solve2(input);
  var expected = 43;

  var resText = expected != result ? "Failed" : "Passed";
  Console.WriteLine($"Test {resText} Result: {result}");
}

int Solve2(string input)
{
  var matrix = input
    .Replace("\r", "")
    .Trim()
    .Split("\n")
    .Select(x =>
      x.Select(y => y.ToString()).ToList())
    .ToList();

  var lineLength = matrix[0].Count();
  var offsets = new List<(int x, int y)> { (-1, -1), (-1, 0), (-1, 1), (0, -1), (0, 1), (1, -1), (1, 0), (1, 1) };

  var total = 0;
  var changed = true;

  while (changed)
  {
    changed = false;
    for (int i = 0; i < matrix.Count(); i++)
    {
      for (int j = 0; j < lineLength; j++)
      {
        try
        {
          var visiting = matrix[i][j];
          if (visiting != "@") continue;
          var neighbours = 0;

          foreach (var offset in offsets) {
            try {
              var value = matrix[i+offset.x][j+offset.y];
              if (value == "@") neighbours++;
            }
            catch {
              continue;
            }
          }

          if (neighbours < 4)
          {
            total++;
            matrix[i][j] = ".";
            changed = true;
          }
        }
        catch (ArgumentOutOfRangeException)
        {
          Console.WriteLine($"{i},{j}");
        }
      }
    }
  }

  return total;
}

Test2();
var input = File.ReadAllText("input.txt");
var result = Solve2(input);
Console.WriteLine($"Result: {result}");
