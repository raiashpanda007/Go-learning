<h1>When to use pointers and when to not</h1>

| When to use value (no pointer) | When to use pointer `*T`      |
| ------------------------------ | ----------------------------- |
| Small structs with fixed size  | Big structs / nested models   |
| No need to modify original     | Need to update original value |
| No `nil` state needed          | Field can be optional         |
| Quick, simple data             | Entities, DB models, JSON     |
