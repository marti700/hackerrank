-- String reduction
strReduce :: Eq a => [a] -> [a]
strReduce [] = []
strReduce xs = rebuild xs []
  where rebuild [] _ = []
        rebuild (x:xs) ys
          | elem x ys = rebuild xs ys
          | otherwise = x:rebuild xs (x:ys)

main = interact $ unlines . map strReduce . lines
