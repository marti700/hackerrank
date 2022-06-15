helloWorlds :: Int -> [String]
helloWorlds 0 = []
helloWorlds n = "Hello World":helloWorlds(n-1)

main = do
  n <- readLn
  mapM_ putStrLn (helloWorlds n)