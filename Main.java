class Person {
    String name;
    int age;

    public Person(String name, int age) {
        this.name = name;
        this.age = age;
    }

    @Override
    public String toString() {
        return name + ": " + age;
    }
}

public class Main {
    public static void swapPersons(Person[] people, int index1, int index2) {
        // Swap references
        Person temp = people[index1];
        people[index1] = people[index2];
        people[index2] = temp;
    }

    public static void main(String[] args) {
        Person[] people = {
                new Person("Alice", 30),
                new Person("Bob", 25),
                new Person("Charlie", 35)
        };

        System.out.println("Before swap:");
        for (Person p : people) {
            System.out.println(p);
        }

        swapPersons(people, 0, 2);

        System.out.println("After swap:");
        for (Person p : people) {
            System.out.println(p);
        }
    }
}
