# A simple TODO app written in go.

This is a simple command line application to track TODO tasks. 
Currently under development. 


TODO: 

* Create command line parser, the following options should be supported.
    * add : To add a new tasks, the argument should be a quoted string. e.g. gotodo add "Buy some food"
    * list : List all the tasks availables, every tasks will be identified by a index. gotodo list
    * done : Mark the tasks as done, the index should be used as a parameter. gotodo done 2
    * clear : Clear all the tasks, this will delete all the content of the .gotodo.txt file.
* Create and initialize a todo file. This file can be created in ~/.gotodo.txt and will have the following format

```
GOTODO
==========
-     Buy some food
- [x] Send an email to someone@gmail.com
-     Take lunch
```

The done tasks will be mark as "[x]"

