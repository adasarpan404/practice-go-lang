# Single Responsibility

The idea behind the Single Responsibility Principle is that every class, module and function should have only one responsibility. 

```java
public class Student {

     public void registerStudent() {
         // some logic
     }

     public void calculate_Student_Results() {
         // some logic
     }

     public void sendEmail() {
         // some logic
     }
}
```

The Student class violates the single responsibility principle. 

Class `Student` has three purpose. 

1. registerStudent
2. calculate_Student_Results
3. SendEmail

This code will work fine but we cannot make it reusable for classes. The class has a whole lot of logic interconnected that we would have a hard time fixing errors.

Let's fix this

```java
public class StudentRegister {
    public void registerStudent() {
        // some logic
    }
}
```

```java
public class StudentResult {
    public void calculate_Student_Result() {
        // some logic
    }
}
```

```java
public class StudentEmails {
    public void sendEmail() {
        // some logic
    }
}
```

Now we've separated each functionality in our program. We can call the classes anywhere we want to use them in our code.

