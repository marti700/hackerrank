-- String mingling
mingle :: String -> String -> String
mingle [] [] = []
mingle (p:ps) (s:ss) = p:s:mingle ps ss

buildResult :: [String] -> String
-- There are always two lines of input
buildResult xs = mingle (head xs) (xs !! 1)

main = interact $ buildResult . lines
