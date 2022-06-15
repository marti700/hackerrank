-- String compression
strCompress :: String -> String
strCompress msg = concat $ concatMap (\x -> if length x > 1
                                     then [head x]:[show $ length x]
                                     else [head [x]]) (group msg)
  -- can be done with the function group of the Data.List module
  where group [] = []
        group (x:xs) = (x:takeWhile(x ==) xs) : group (dropWhile (== x) xs)

main :: IO ()
main = interact strCompress