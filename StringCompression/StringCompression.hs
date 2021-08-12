-- String compression
strCompress :: String -> String
strCompress msg = concat $ concatMap (\x -> if length x > 1
                                     then [head x]:[show $ length x]
                                     else [head [x]]) (group msg)
  where group [] = []
        group (x:xs) = (x:takeWhile(x ==) xs) : group (dropWhile (== x) xs)

main :: IO ()
main = interact strCompress