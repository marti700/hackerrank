-- Compute the perimeter of a Polygon
perimetrePolygon :: (Floating a) => [(a,a)] -> a
perimetrePolygon xs = sum $ mapTuple distance xs xs
  where distance c k = sqrt ( ((abs (fst k) - fst c)**2) + ((abs (snd k) - snd c)**2 ))


-- general function used to traverse a list of tuples
-- used by the functions perimeterPolygon
-- to solve the Compute the perimeter of a polygon and
-- comute the area of a polygon problems
mapTuple :: ( a -> a -> b ) -> [a] -> [a] -> [b]
mapTuple f (x:[]) ys = [f x (head ys)]
mapTuple f (x:xs) ys = (f x (head xs)):mapTuple f xs ys

toIntTuple :: [String] -> [(Double, Double)]
-- toIntTuple [_] = []
toIntTuple [] = []
toIntTuple (x:y:xs) = (read x :: Double, read y :: Double) : toIntTuple xs

main :: IO ()
main = interact $ show . perimetrePolygon . toIntTuple . tail . words