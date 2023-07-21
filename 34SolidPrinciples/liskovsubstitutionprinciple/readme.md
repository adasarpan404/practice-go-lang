# Liskov Substitution Principle

According to Barbara Liskov and Jeannette Wing, the Liskov substitution principle states that:

`Let Φ(x) be a property provable about objects x of type T. Then Φ(y) should be true for objects y of type S where S is a subtype of T.`

This Principles implies that when an instance of a class is passed/extended to another class, the inheriting class should have a use case for all the properties and behavior of the inherited class.

Let's say we have a class called `Amphibian` for animals that can live on both land and water. This class has two methods to show the features of an amphibian – `swim()` and `walk()`.

```Java
public class Amphibian {

    public void swim();
    public void walk();

}
```

The `Amphibian` class can extend to a `Frog` class because frogs are amphibians, so they can inherit the properties of the `Amphibian` class without altering the logic and purpose of the class.

```Java
public class Frog extends Amphibian {
    public void swim() {
        System.out.println("The frog is swimming");
    }
    
    public void walk() {
        System.out.println("The frog is walking on land");
    }
}
```

But we cannot extend the `Amphibian` class to a `Dolphin` class because dolphins only live in water which implies that the walk() method would be irrelevant to the `Dolphin` class.

In summary, if a class inherits another, it should do so in a manner that all the properties of the inherited class would remain relevant to its functionality.
