# Abstract Factory Method

1. It is a creational design patterns.
2. It lets you produce families of related objects without specifying its concrete class.

## Problem

Imagine that you’re creating a furniture shop simulator. Your code consists of classes that represent:

1. A family of related products, say: _Chair_ + _Sofa_ + _CoffeeTable_.
2. Several variants of this family. For example, products _Chair_ + _Sofa_ + _CoffeeTable_ are available in these variants: _Modern_, _Victorian_, _ArtDeco_.

You need a way to create individual furniture objects so that they match other objects of the same family. Customers get quite mad when they receive non-matching furniture.

Also, you don’t want to change existing code when adding new products or families of products to the program. Furniture vendors update their catalogs very often, and you wouldn’t want to change the core code each time it happens.

## How to implement

1. Map out a matrix of distinct product types versus variants of these products.

2. Declare abstract product interfaces for all product types. Then make all concrete product classes implement these interfaces.

3. Declare the abstract factory interface with a set of creation methods for all abstract products.

4. Implement a set of concrete factory classes, one for each product variant.

5. Create factory initialization code somewhere in the app. It should instantiate one of the concrete factory classes, depending on the application configuration or the current environment. Pass this factory object to all classes that construct products.

6. Scan through the code and find all direct calls to product constructors. Replace them with calls to the appropriate creation method on the factory object.

## Pros

- You can be sure that the products you’re getting from a factory are compatible with each other.

- You avoid tight coupling between concrete products and client code.

- _Single Responsibility Principle_. You can extract the product creation code into one place, making the code easier to support.

- _Open/Closed Principle_. You can introduce new variants of products without breaking existing client code.

## Cons

- The code may become more complicated than it should be, since a lot of new interfaces and classes are introduced along with the pattern.
