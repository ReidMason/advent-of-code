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
process [] _ total = total
process leftColumn rightColumn total = do
    let target = Prelude.head leftColumn
    let occurances = getOccurances rightColumn target
    let newLeftColumn = removeAt leftColumn 0

    process newLeftColumn rightColumn (total + occurances * target)

getOccurances :: [Integer] -> Integer -> Integer
getOccurances [] _ = 0
getOccurances (x : xs) target
    | x == target = 1 + getOccurances xs target
    | otherwise = getOccurances xs target

removeAt :: [a] -> Integer -> [a]
removeAt [] _ = []
removeAt (x : xs) index
    | index == 0 = xs
    | otherwise = x : removeAt xs (index - 1)

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
