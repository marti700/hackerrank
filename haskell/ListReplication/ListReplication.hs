lreplication :: (Integral a) => a -> [b] -> [b]
lreplication 0 _ = []
lreplication _ [] = []
lreplication n (x:xs) = replicte n x ++ lreplication n xs
  where replicte 0 x = []
        replicte n x = x:replicte (n-1) x