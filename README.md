# countercli
> a cli for [counter](https://github.com/NULLx76/counter)

this is a simple cli for counter, so instead of having to make actual curl requests you can do it with this shitty program!

## usage
```sh
# use get to get a counter
countercli get http://localhost:8080/countername
> 0

# create counters using create 
countercli create http://localhost:8080/counter
> value: 0, token: ff672e05-ae43-4783-b02a-ce480c098e17

# increment counters using increment
countercli increment http://localhost:8080/counter ff672e05-ae43-4783-b02a-ce480c098e17
> 1

# delete counters using delete
countercli delete http://localhost:8080/counter ff672e05-ae43-4783-b02a-ce480c098e17
> deleted successfully
```