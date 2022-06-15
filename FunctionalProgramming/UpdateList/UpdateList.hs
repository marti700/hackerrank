f :: (Num a, Ord a) => [a] -> [a]
f [] = []
f (x:xs) = absv x: f xs
  where absv x
          | x < 0 = x * (-1)
          | otherwise = x

main :: IO ()
main = do
  inputdata <- getContents
  mapM_ print $ f $ map (read :: String -> Int) $ lines inputdata