# go-todo
Command line todo-list written in Go


Example console output and demo of functions:

```
mattylee@Matthews-MacBook-Pro todo % go run main.go
Error opening file
Creating new todo-list


1. New item | 2. Edit item | 3. Remove item | 4. Edit item order | 5. Save & Exit | 6. Save
1
What is the new item?
Item 1

0 .  Item 1

1. New item | 2. Edit item | 3. Remove item | 4. Edit item order | 5. Save & Exit | 6. Save
1
What is the new item?
Item 2

0 .  Item 1
1 .  Item 2

1. New item | 2. Edit item | 3. Remove item | 4. Edit item order | 5. Save & Exit | 6. Save
2
What item do you want to edit?

0 .  Item 1
1 .  Item 2

0
What do you want to change it to? (blank to cancel)
Item 1 edited

0 .  Item 1 edited
1 .  Item 2

1. New item | 2. Edit item | 3. Remove item | 4. Edit item order | 5. Save & Exit | 6. Save
4
What numbers do you want to swap? (separate with comma

0 .  Item 1 edited
1 .  Item 2

0,1

0 .  Item 2
1 .  Item 1 edited

1. New item | 2. Edit item | 3. Remove item | 4. Edit item order | 5. Save & Exit | 6. Save
3
What item do you want to remove? (invalid input to cancel)

0 .  Item 2
1 .  Item 1 edited

1

0 .  Item 2

1. New item | 2. Edit item | 3. Remove item | 4. Edit item order | 5. Save & Exit | 6. Save
5
Saved```
