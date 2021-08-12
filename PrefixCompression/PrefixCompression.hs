-- Prefix compression
prefixCompression :: String -> String -> [(Int, String)]
prefixCompression xs ys = [(prefixLength, prefix xs ys), (length postFixx, postFixx), (length postfixy, postfixy)]
  where prefix [] _ = []
        prefix _ [] = []
        prefix(x:xs) (y:ys)
          | x == y = x:prefix xs ys
          | otherwise = []
        prefixLength = length (prefix xs ys)
        postFixx = drop prefixLength xs
        postfixy = drop prefixLength ys

main = do
  s1 <- getLine
  s2 <- getLine
  putStrLn $ unlines $ fmap (\x -> show (fst x) ++ " " ++ snd x) (prefixCompression s1 s2)
