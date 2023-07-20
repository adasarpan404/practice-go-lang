# Open–Closed Principle

The open-closed principle states that software entities should be open for extension, but closed for modification.

This implies that such entities - **Classes**, **Function**, and so on - should be created in a way that their core functionalities can be extended to other entities without altering the initial entity's source code.

In the example below, we're going to write the code for calculating the body mass index (BMI) of a person:

```Java
public class Human  {
     public int height;
     public int weight;
}
```

We've created the Human class which provides the height and width properties of the class. Now, let's calculate the first person's BMI.

```Java
public class CalculateBMI {

     public int CALCULATE_JOHN_BMI(Human John) {   
         
         return John.height/John.weight;
         
     }
}
```

We've calculate the BMI of a person named John. We'll go on and calculate the BMI for a person named Jane.

```Java
public class CalculateBMI {

     public int CALCULATE_JOHN_BMI(Human John) {   
         
         return John.height/John.weight;
         
     }
     
     public int CALCULATE_JANE_BMI(Human Jane) {   
         
         return Jane.height/Jane.weight;
         
     }
}
```

The problem with this is that we keep modifying the code every time we need to calculate the BMI of another person.

This also violates the SRP because the class now has more than one reason to change.

Although the code above may work perfectly, it's not efficient. We modify the code constantly which may lead to bugs. And the code only has provision for humans. What if we have to calculate for an animal or an object?

Let's fix the problem using the open-closed principle.

```Java
public interface Entity {

     public int CalculateBMI();

}

// John entity
public class John implements Entity {

     int height;
     int weight;

     public double CalculateBMI() {

           return John.height/John.weight;
     }
}

// Jane entity
public class Jane implements Entity {

     int height;
     int weight;

     public double CalculateBMI() {

           return Jane.height/Jane.weight;
     }
}

// Dog entity
public class Dog implements Entity {

     int height;
     int weight;

     public double CalculateBMI() {

           return Dog.height/Dog.weight;
     }
}
```

In the code above, we have created an interface called `Entity` with a `CalculateBMI()` method.

Each entity – `John`, `Jane` and `Dog` – extends the functionality of the `Entity` interface.

Now we no longer have to modify existing code when we create a new entity - we just extend the functionality we need and apply it to the new entity.
