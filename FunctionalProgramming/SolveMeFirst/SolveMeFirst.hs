solveMeFirst :: Int -> Int -> Int
solveMeFirst a b = a + b

main :: IO ()
main = do
  a <- readLn
  b <- readLn
  print $ solveMeFirst a b