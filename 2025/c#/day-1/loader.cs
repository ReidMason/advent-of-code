namespace day_1;

public static class Loader
{
  public static string LoadInput()
  {
    var inputPath = "./input.txt";
    return File.ReadAllText(inputPath);
  }
}
