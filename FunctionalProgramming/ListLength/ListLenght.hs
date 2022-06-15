len :: (Num a) => [b] -> a
len lst = sum [1 | x <- lst]