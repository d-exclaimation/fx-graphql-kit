# fx-graphql-kit (with Dataloader)
Simple GraphQL-API Reusable module / kit with 
[`fx`](https://github.com/uber-go/fx),
[`gin`](https://github.com/gin-gonic/gin),
[`gqlgen`](https://github.com/99designs/gqlgen), and
[`gorm`](https://github.com/go-gorm/gorm)


## 🚚 Dataloader References (For Relational Data)
One of the biggest downside of GraphQL Resolvers its that it sometimes don't
actually solve the overfetching issues. It instead move it inside the
server for Relational Data instead of over the network and the client.
This is definitely still better than REST but create what's called the
n+1 problem. A way to solve this is using dataloader to resolve 
multiple fetches into one single function or instruction.

This is my implementation of dataloader for gqlgen, but this is not
somthing that should be use right of the bat, rather more of a
reference.

Reasoning includes Dataloaden limitation with no generics. Context
specific middleware, completely different relational data and their
databases.

- Added Dataloaden to handle and resolve multiple queries of the same
type
- Implemented Dataloaden Middleware into Gin
- Updated data entities and databases connection to relational one
- Added UserService for handle queries for User
- Modified ThoughtService to take account for relational data and
  UserService
- Modified Fx dependecies structure to take account for new Providers
  and different dependecies

### Stack of libraries and frameworks
> 1. fx
> 2. gqlgen
> 3. gin
> 4. gorm
> 5. Dataloaden

### What already existed here
> 1. Simple GraphQL CRUD handlers
> 2. Database connection with PostgreSQL
> 3. GraphQL Resolvers and Dataloader
> 4. Services to handle logic
> 6. Fx lifecycle to handle dependencies injection
> 7. Providers and Invokers for fx

```go
func main() {
	fmt.Printf("Thanks for checking this out, %v\n", you)
}
```
