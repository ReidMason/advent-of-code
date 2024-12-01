import Data.Text

main :: IO ()
main = do
  input <- readFile "input.txt"

  let lines = splitLines input
  let leftColumn = Prelude.map getLeftColumn lines
  let rightColumn = Prelude.map getRightColumn lines

  let total = process leftColumn rightColumn 0
  print total

process :: [Integer] -> [Integer] -> Integer -> Integer
process [] [] total = total
process leftColumn rightColumn total = do
  let lowestLeftIndex = getLowestIndex leftColumn 0 (-1) (-1)
  let lowestLeft = leftColumn !! fromIntegral lowestLeftIndex

  let lowestRightIndex = getLowestIndex rightColumn 0 (-1) (-1)
  let lowestRight = rightColumn !! fromIntegral lowestRightIndex

  let diff = difference lowestLeft lowestRight

  let newLeftColumn = removeAt leftColumn lowestLeftIndex
  let newRightColumn = removeAt rightColumn lowestRightIndex

  process newLeftColumn newRightColumn (total + diff)

removeAt :: [a] -> Integer -> [a]
removeAt [] _ = []
removeAt (x : xs) index
  | index == 0 = xs
  | otherwise = x : removeAt xs (index - 1)

difference :: Integer -> Integer -> Integer
difference x y = max x y - min x y

getLowestIndex :: [Integer] -> Integer -> Integer -> Integer -> Integer
getLowestIndex [] _ _ lowestIndex = lowestIndex
getLowestIndex arr index (-1) (-1) = getLowestIndex (Prelude.tail arr) (index + 1) (Prelude.head arr) 0
getLowestIndex arr index lowest lowestIndex =
  do
    let target = Prelude.head arr
    if target < lowest
      then getLowestIndex (Prelude.tail arr) (index + 1) target index
      else getLowestIndex (Prelude.tail arr) (index + 1) lowest lowestIndex

getLeftColumn :: String -> Integer
getLeftColumn "" = 0
getLeftColumn x = read (Prelude.head (getColumns x)) :: Integer

getRightColumn :: String -> Integer
getRightColumn "" = 0
getRightColumn x = read (Prelude.last (getColumns x)) :: Integer

getColumns :: String -> [String]
getColumns x = Prelude.map unpack (splitOn (pack "  ") (pack x))

splitLines :: String -> [String]
splitLines x = Prelude.map unpack (splitOn (pack "\n") (pack x))
