import Control.Monad
-- Functions or Not
isFunction :: (Integral a, Eq a) => [(a,a)] -> Bool
isFunction [] = True
isFunction (x:xs) = if isRepeated xs then False else isFunction xs
  where isRepeated [] = False
        isRepeated (y:ys)
          | (fst x == fst y) && (snd x /= snd y) = True
          | otherwise = isRepeated ys

-- Clean Input
clean [] = []
clean (x:xs) = map (makeTuple . words) (take (read x :: Int) xs) : clean xs
  where
    makeTuple ys = (read (head ys) :: Int, read (ys !! 1) :: Int)

main = do
    t <- fmap (read::String->Int) getLine
    forM [1..t] (\_->do
        n <- fmap (read::String->Int) getLine
        func <- forM [1..n] (\_->do fmap ((\[a, b]->(a,b)).map (read::String->Int).words) getLine :: IO (Int, Int))
        putStrLn $ if isFunction func then "YES" else "NO")


--- COME BACK TO THIS!
-- -- Functions or Not
-- isFunction :: (Integral a, Eq a) => [(a,a)] -> String
-- isFunction [] = "YES"
-- isFunction (x:xs) = if isRepeated xs then "NO" else isFunction xs
--   where isRepeated [] = False
--         isRepeated (y:ys)
--           | (fst x == fst y) && (snd x /= snd y) = True
--           | otherwise = isRepeated ys

-- -- Clean Input
-- clean :: [String] -> [[(Int, Int)]]
-- clean [] = []
-- clean (x:xs) = map (makeTuple . words) (take (read x :: Int) xs) : clean xs
--   where
--     makeTuple ys = (read (head ys) :: Int, read (ys !! 1) :: Int)

-- main = do
--   input <- getContents
--   mapM_ (print . isFunction) (clean $ tail (lines input))

