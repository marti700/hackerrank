--------------------------------- THE ONE HACKERANCK LIKES ---------------------------------------------------
import Text.Printf (printf)
-- solve :: (Floating a, Enum a) => a -> a -> [a] -> [a] -> [a]
solve :: Int -> Int -> [Int] -> [Int] -> [Double]
solve l r a b = [areaUnderTheCurve l r a b, revolvingVolume l r a b]

-- Area under curves and volume of revolving curve
areaUnderTheCurve :: Int -> Int -> [Int] -> [Int] -> Double
areaUnderTheCurve = integrate (* 0.001)
--
revolvingVolume :: Int -> Int -> [Int] -> [Int] -> Double
revolvingVolume = integrate (\x -> (x**2)*0.001*pi)

-- general function using to find the area under a curve of a polynomial function using the riemman
-- sums method.
integrate :: (Double -> Double) -> Int -> Int -> [Int] -> [Int] -> Double
integrate f l r as bs = sum $ map (f . evaluateAt) interval
  where
        interval = tail [fromIntegral l, (fromIntegral l + 0.001)..fromIntegral r]
        evaluateAt x = sum [fromIntegral i * (x^^j) | (i,j) <- zip as bs]

main :: IO ()
main = do
  as <- getLine
  bs <- getLine
  l <- getLine
  let a = map (read :: String -> Int) (words as)
  let b = map (read :: String -> Int) (words bs)
  let c = map (read :: String -> Int) (words l)
  mapM_ (printf "%.1f\n") (solve (head c) (c !! 1) a b)


--------------------------------- THE ONE I LIKE ---------------------------------------------------
-- -- Area under curves and volume of revolving curve
-- areaUnderTheCurve :: (Floating a, Enum a) => a -> a -> [a] -> [a] -> a
-- areaUnderTheCurve = integrate (* 0.001)

-- revolvingVolume :: (Floating a, Enum a) => a -> a -> [a] -> [a] -> a
-- revolvingVolume = integrate (\x -> (x**2)*0.001*pi)

-- -- general function using to find the area under a curve using the riemman
-- -- sums method. This function is used by the function areaUnderTheCurve and
-- -- revolvingVolume to solve the problem Area under curves and volume of
-- -- revolving curve
-- integrate :: (Num a, Floating b, Enum b) => (b -> a) -> b -> b -> [b] -> [b] -> a
-- integrate f l r as bs = sum $ map (f . evaluateAt) interval
--   where
--         interval = tail [l, (l + 0.001)..r]
--         evaluateAt x = sum [i * (x**j) | (i,j) <- zip as bs]

-- main :: IO ()
-- main = do
--   as <- getLine
--   bs <- getLine
--   l <- getLine
--   let a = map (read :: String -> Double) (words as)
--   let b = map (read :: String -> Double) (words bs)
--   let c = map (read :: String -> Double) (words l)
--   print (solve (head c) (c !! 1) a b)

----------------------------------- A MORE VERBOSE ONE IN CASE YOU FORGET THE CURRYING IS HAPPENING ABOVE------------------------
-- Area under curves and volume of revolving curve
-- areaUnderTheCurve l r as bs = integrate (\x -> x*0.001) l r as bs
-- revolvingVolume l r as bs = integrate (\x -> (x**2)*0.001*pi) l r as bs

-- -- general function using to find the area under a curve using the riemman
-- -- sums method. This function is used by the function areaUnderTheCurve and
-- -- revolvingVolume to solve the problem Area under curves and volume of
-- -- revolving curve
-- integrate f l r as bs = sum $ map (f . evaluateAt) (interval)
--   where
--         interval = tail [l, (l + 0.001)..r]
--         evaluateAt x = sum [i * (x**j) | (i,j) <- zip as bs]