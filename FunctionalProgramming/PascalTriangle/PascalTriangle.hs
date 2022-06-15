-- Pascal's triangle
pascal :: (Integral a) => a -> [[a]]
pascal n = map triangleRow [1..n]
  -- Calculates the binomial coefficients for the given row of the
  -- pascal triagle
  where
    triangleRow n = [nCr n k | n <- [n-1], k <- [0..n]]
        -- Binomial coefficient formula
    nCr n k = div (fac n) (fac k * fac (n-k))

-- Calculate the factorial of a number
fac :: (Integral a, Eq a) => a -> a
fac 0 = 1
fac n = n * fac  (n - 1)

main = interact $ unlines . map (unwords . map show) . pascal . read