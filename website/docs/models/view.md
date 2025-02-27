---

sidebar_position: 2
description: Easily query a database table

---

# View

A view model makes it easy to map an entity to a database table and query it.

To create a view, use the `NewView()` function.

```go
type User struct {
    ID    int
    Name  string
    Email string
}

var userView = psql.NewView[User]("public", "users")
```

:::tip

The `NewViewx()` function takes an extra type parameter to determine how slices are returned.

:::

A view model provides the following methods:

## Name()

This returns a properly quoted name of the table and can be used as a bob [expression](../query-builder/building-queries#expressions). e.g. `"public"."users"`

## NameAs()

Similar to `Name()`, but adds an alias. e.g. `"public"."users" as "public.users"`

## Columns()

Returns an [`orm.Columns`](https://pkg.go.dev/github.com/stephenafamo/bob/orm#Columns) object.  
This is also a bob [expression](../query-builder/building-queries#expressions). Which by default, the expression evaluates to:

```sql
-- table_alias.column_name
"public.users"."id" AS "id",
"public.users"."name" AS "name",
"public.users"."email" AS "email"
```

Learn about how to manipulate a columns list in the [columns documentation](./columns)

## Query()

The `Query()` method starts a SELECT query on the model's database view/table. It can accept [query mods](../query-builder/building-queries#query-mods) to modify the final query.

```go
q := userView.Query(
    ctx, db, 
    sm.Limit(10), // LIMIT 10
)
```

The query can then be executed with `One()`, `All()`, `Cursor()`, `Count()` or `Exists()`.

```go
// SELECT * FROM "users" LIMIT 1
userView.Query(ctx, db).One()

// SELECT * FROM "users"
userView.Query(ctx, db).All()

// Like All, but returns a cursor for moving through large results
userView.Query(ctx, db).Cursor()

// SELECT count(1) FROM "users"
userView.Query(ctx, db).Count()

// Like One(), but only returns a boolean indicating if the model was found
userView.Query(ctx, db).Exists()
```

:::tip

The `Count()` function clones the current query which can be an expensive operation.

:::