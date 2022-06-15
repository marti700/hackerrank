-- Computing the GCD (https://youtu.be/JUzYl1TYMcU)
gcd' :: (Integral a) => a -> a -> a
gcd' a b
  | a < b = gcd' b a
  | a `mod` b == 0 = b
  | otherwise = gcd' b (mod a b)