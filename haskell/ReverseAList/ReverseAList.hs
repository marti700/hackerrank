rev :: [a] -> [a]
rev [] = []
rev xs = inv $ length xs - 1
  where inv 0 = [head xs]
        inv n = (xs !! n) : inv(n - 1)