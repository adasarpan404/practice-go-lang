# Interface Seggregation Principle

The **interface segregation principle** states that the interface of a program should be split in a way that the user/client would only have access to the necessary methods related to their needs.

Example

```java
public interface Teacher {

    void English();

    void Biology();

    void Chemistry();
    
    void Mathematics();

}
```

We've created an interface called `Teacher` which has various subjects as its methods. Let's extend this interface to our first teacher.

```Java
public class Jane implements Teacher {

    @Override
    public void English() {
        System.out.println("Jane is teaching the students English language.");
    }

    @Override
    public void Biology() {
    }

    @Override
    public void Chemistry() {
    }

    @Override
    public void Mathematics() {
    }
}
```

From the code above, you can tell that Jane is an English teacher who has no business with the other subjects. But these other methods are extended by default with the Teacher interface.

Do not confuse the Liskov substitution principle and the interface segregation principle. They may seem similar but they are not entirely the same.

Liskov substitution principle gives us the idea that when a new class has the need to inherit an existing class, it should do so because this new class has a need for the methods the existing class has.

On the other hand, the interface segregation principle makes us understand that it is unnecessary and unreasonable to create an interface with a lot of methods as some of these methods may be irrelevant to the needs of a particular user when extended.

Example

```java
public interface Teacher {

    void Teach();

}
```

The `Teacher` interface now has only one method. Let's go on and extend this interface to support the different subjects.

```java
public interface Teacher {

    void Teach();

}
```

The `Teacher` interface now has only one method. Let's go on and extend this interface to support the different subjects.

```java
public interface EnglishTeacher extends Teacher {

    void English();

}
```

```java
// Mathematics teacher interface

public interface MathematicsTeacher extends Teacher {

    void Mathematics();

}
```

```java
// Chemistry teacher interface

public interface ChemistryTeacher extends Teacher {

    void Chemistry();

}
```

```java
// Biology teacher interface

public interface BiologyTeacher extends Teacher {

    void Bilogy();

}
```

We have created different interfaces for every subject. Now Jane can teach English without carrying the other methods with them. Here is an example:

```java
public class Jane implements EnglishTeacher {
    
    @Override
    public void Teach() {
        System.out.println("Jane has started teaching.");
    }

    @Override
    public void English() {
        System.out.println("Jane is teaching the students English language.");
    }

}```
