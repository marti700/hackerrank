e :: (Enum a, Eq a, Floating a) => a -> a
e n = 1 + foldl (\acc x -> acc + ((n**x)/fac' x)) 0 [1..9]
  where fac' 1 = 1
        fac' n = n * fac' (n - 1)

main :: IO ()
main = do
  (testCases:cases) <- fmap lines getContents
  let res = map (e . read :: String -> Double) cases
  mapM_ print res
