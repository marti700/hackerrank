f :: (Integral a) => a -> [a] -> [a]
f _ [] = []
f n (x:xs)
  | x < n = x:f n xs
  | otherwise = f n xs

main :: IO ()
main = do
  n <- readLn
  inputData <- getContents
  let
    numbers = map read (lines inputData) :: [Int]
  putStrLn . unlines $ (map show . f n) numbers