import System.IO
import qualified Data.Map as Map

-- Convert string row into list of Int
parseRow :: String -> [Int]
parseRow = map read . words

-- Get a column from matrix
getColumn :: Int -> [[Int]] -> [Int]
getColumn col = map (!! col)

-- Apply operation
applyOp :: String -> [Int] -> Int
applyOp "+" xs = sum xs
applyOp "*" xs = product xs
applyOp op _   = error ("Unknown operator: " ++ op)

main :: IO ()
main = do
    contents <- readFile "inputs.txt"

    let ls = lines contents

    -- Last line = symbols
    let symbols = words (last ls)

    -- Remaining lines = matrix
    let matrix = map parseRow (init ls)

    -- Build mapping
    let mappings =
            Map.fromListWith (++)
            [ (sym, [getColumn col matrix])
            | (col, sym) <- zip [0..] symbols
            ]

    -- Perform operations
    let results =
            [ applyOp sym vals
            | (sym, cols) <- Map.toList mappings
            , vals <- cols
            ]
    print ("Final Total: " ++ show (sum results))