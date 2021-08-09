-- Compute the area of a polygon (I've used the formula exposed here:
-- https://www.mathopenref.com/coordpolygonarea.html)
areaPolygon :: (Floating a) => [(a,a)] -> a
areaPolygon xs = abs $ sum (mapTuple cross xs xs)
  where cross x y = (fst x * snd y - fst y * snd x) / 2

-- general function used to traverse a list of tuples
-- used by the function areaPolygon
-- to solve the Compute the perimeter of a polygon and
-- comute the area of a polygon problems
mapTuple :: ( a -> a -> b ) -> [a] -> [a] -> [b]
mapTuple f (x:[]) ys = [f x (head ys)]
mapTuple f (x:xs) ys = (f x (head xs)):mapTuple f xs ys

toIntTuple :: [String] -> [(Double, Double)]
-- toIntTuple [_] = []
toIntTuple [] = []
toIntTuple (x:y:xs) = (read x :: Double, read y :: Double) : toIntTuple xs

main = interact $ show . areaPolygon . toIntTuple . tail . words