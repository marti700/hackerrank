f ::(Integral a) => [a] -> [a]
f [] = []
f xs = odds False xs
  where odds _ [] = []
        odds n (y:ys)
          | n = y:odds False  ys
          | otherwise = odds True ys