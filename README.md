# Bunnies - cuddly functional programming in Golang!
_Bunnies_ is a library to support functional programming with Go 1.18 Generics, providing basic building blocks and types for a more type-safe, boiler-plate free and functional experience.

What is supported, and the roadmap is in the [TODO](#TODO) list below.

## Why the name "Bunnies"?
Scala has the excellent [Cats](https://typelevel.org/cats/) library. Cats are cute. Bunnies are also cute. Why not pick a name for something equally cute, based on the loose inspiration for the library?

# Documentation
Functional programming with generics allows us to express many common algorithms and structures more concisely and DRY'ly.

As mentioned, this library takes inspiration from Scalas Cats, but since Go Generics lacks [Higher-Kinded types](https://en.wikipedia.org/wiki/Kind_(type_theory)), a straight translation is clearly not possible. Furthermore, since a straight translation is not possible, and Go is a distinctly separate language, we aim to implement what makes sense and is useful for Go, in ways which do not undercut the benefits of Go.

To document this library, we intend to split this into three parts:
* Principles
* Types
* Functions

All supported types will have all described functions implemented for them, thus supporting a wide range of FP use-cases. 
Though functions will be implemented separately for each type, we will describe them in this document once, with examples for each type provided in the `godocs`.

Additionally, some types will have certain functions that are only appropriate for them specifically, these will also be described in this document.

## Principles
The two main principles we follow are:
* [Referential transparency](https://en.wikipedia.org/wiki/Referential_transparency)
* Immutability

In short, this means the API's of this library will always produce the same outputs for the same inputs, given pure inputs.
It also means that the library will never mutate function inputs, and return values will always be new values.

## Data Types
### List/Slice
Just a regular Golang slice for all intents and purposes. It's the closest thing to a List and works just fine.
### NonEmptyList
### Optional
`Optional` is a type which demarks a type which can either be empty, or contain a value, and allows a developer to operate typesafely under those assumptions.
### Either
## Functions

### All types
#### Pure
#### Empty
#### Map
#### FlatMap
#### Filter
#### Exists
#### FlatMap2
#### Map2
#### IsEmpty
### List & NonEmpty only
#### FoldLeft

### Option only

#### GetOrElse
### Either only
#### Fold
#### LeftMap
#### GetLeftOrElse


# TODO
- [x] `List` (just slice)
- [x] `Option`
- [ ] `Either`
- [ ] `NonEmptyList`
- [ ] `IO`
- [ ] `Validation`