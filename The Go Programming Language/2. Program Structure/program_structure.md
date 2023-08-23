# 2 Program Structure

## 2.1 Names

Go has 25 keywords

```
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

If an entity is declared outside of a function, it is visible in all files of the package to which it belongs. 

If the name begins with an upper-case letter, it is exported, which means that it is visible and accessible outside of its own package and may be referred to by other parts of the program

## 2.2 Declarations


## 2.3 Variables
* A short variable declaration must declare at least one new variable 
```
f, err := os.Open(infile)
// ...
g, err := os.open(outfile)
f, err := os.Create(outfile) // compile error: no new variables
```