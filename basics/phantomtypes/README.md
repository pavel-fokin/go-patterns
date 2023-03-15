# Phantom Types

__Phantom Type__ is a type that doesn’t use at least one of its generic type parameters.

For example,

```go
// Roles.
type Basic struct{}
type Admin struct{}

// User.
type User[T Basic | Admin] struct {
	name string
}
```

And even we don't use the type parameter `T` in the `User` struct,
it allows to check type rules more strictly.