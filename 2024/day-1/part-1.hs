import Data.List (sort)
import Data.Text (pack, splitOn, unpack)

main :: IO ()
main = do
    input <- readFile "input.txt"
    let inputs = map (map read . split) (lines input)
    let (left, right) = getColumns inputs
    let total = sum (zipWith (curry diff) (sort left) (sort right))

    print total

split :: String -> [String]
split text = Prelude.map unpack (splitOn (pack "   ") (pack text))

diff :: (Integer, Integer) -> Integer
diff (x, y) = abs (x - y)

getColumns :: [[Integer]] -> ([Integer], [Integer])
getColumns = foldr (\[x, y] (xs, ys) -> (x : xs, y : ys)) ([], [])
