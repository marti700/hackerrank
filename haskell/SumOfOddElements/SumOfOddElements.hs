f :: (Integral a) => [a] -> a
f [] = 0
f xs = sum [x | x <- xs, x `mod` 2 == 1]


main = do
  inputdata <- getContents
  print (f $ map (read :: String -> Int) $ lines inputdata)