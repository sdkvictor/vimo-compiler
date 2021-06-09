# Vimo Compiler

<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#run-a-file">Run a file</a></li>
      </ul>
    </li>
    <li>
      <a href="#features">Features</a>
      <ul>
        <li><a href="#operations">Operations</a></li>
        <li><a href="#statements">Statements</a></li>
        <li><a href="#basic-types">Basic Types</a></li>
        <li><a href="#predefined-objects">Predefined Objects</a></li>
        <li><a href="#predefined-functions">Predefined Functions</a></li>
      </ul>
    </li>
    <li>
      <a href="#usage">Usage</a>
      <ul>
        <li><a href="#code-syntax">Code Syntax</a></li>
        <li><a href="#your-first-program">Your First Program</a></li>
        <li><a href="#basic-examples">Basic Examples</a></li>
        <li><a href="#game-engine">Game Engine</a></li>
      </ul>
    </li>
    <li><a href="#video-tutorial">Video Tutorial</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#developers">Developers</a></li>
    <li><a href="#acknowledgements">Acknowledgements</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->
## About The Project
Vimo is an imperative programming language based on some famous programming languages, being C++ and C the most representative examples. It uses a similar syntax as those programming languages and as a special feature, it contains a game engine to develop simple 2D videogames such as Pacman and the famous game of the snake.

<!-- BUILT WITH -->
### Built With
Vimo was built using the programming language Go, Gocc as a compiler kit for Go and Pixel for the game engine.
* [Go](https://golang.org/)
* [Gocc](https://github.com/goccmack/gocc)
* [Pixel](https://github.com/faiface/pixel)

<!-- GETTING STARTED -->
## Getting Started
This section describes how to setup the programming environment to create and run a vimo program.

<!-- PREREQUISITES -->
### Prerequisites
You must have installed Go first and set up your $GOPATH. Following are some guides on how to set up the aforementioned prerequisites.
* [Go](https://golang.org/doc/install) - Install Go
* [Complete Setup Guide](https://www.callicoder.com/golang-installation-setup-gopath-workspace/) - Complete Go and $GOPATH setup guide for **Windows**, **MacOS** and **Linux**
* [Windows Go Setup](https://www.freecodecamp.org/news/setting-up-go-programming-language-on-windows-f02c8c14e2f/#:~:text=Create%20the%20GOPATH%20variable%20and,System%20Variables%20click%20on%20New.&text=To%20check%20that%20your%20path,%25%E2%80%9D%20on%20the%20command%20line) - More detailed guide to set up the Go environment for Windows

<!-- RUN A FILE -->
### Run A File
Once the Go setup is done, to run a file you must first be inside the run directory.

```sh
$ ./golang-compiler/run
```

Once you are inside the correct directory, to run a file you must run the following command.

```sh
$ go run run.go <path of your file>
```

<!-- FEATURES -->
## Features
Vimo has the basic operations and data types of programming, as well as predefined functions and objects with their attributes and methods to use its game engine to create 2D videogames and for different uses, which is explained in more detail below.

<!-- OPERATIONS -->
### Operations
* `+` - Add
* `-` - Subtract
* `*` - Multiply
* `/` - Divide
* `&&` - And
* `||` - Or
* `==` - Equal
* `<` - Less than
* `>` - Greater than

<!-- STATEMENTS -->
### Statements
* `for` - For
* `while` - While
* `if` - If
* `if-else` - If-else
* `return` - Return
* `print` - Print

<!-- BASIC TYPES -->
### Basic Types
Vimo supports the usage of the following basic types for variables or as a return type for functions.
| Type | Description |
:-: | -----------
int | An integer data type stores whole numbers, positive or negative.
float | A float data type stores floating precision numbers, positive or negative.
char | A char or character data type stores a single character or a letter.
string | A string data type stores a sequence of characters.
bool | A boolean data type stores only two possible values, **true** or **false**.
void | A void type is a special type used only for functions as a return type when the function doesn't return a result value to the function call.

<!-- PREDEFINED OBJECTS -->
### Predefined Objects
Vimo includes the following predefined objects with its attributes and methods to develop 2D videogames.

#### Basic Objects
* `Square` - A Square object is used to store a square that can be drawn in the window.
* `Circle` - A Circle object is used to store a circle that can be drawn in the window.

##### These object types share the following attributes.
| Attribute | Description |
:-: | -----------
x | Represents the position of the object on the x-axis of the window.
y | Represents the position of the object on the y-axis of the window.
width | Stores the width size of the object.
height | Stores the height size of the object.
color | Store the color of the object in hexadecimal format without '#', e.g. "ffffff".

#### Special Objects
* `Image` - A Image object is used to store an image that can be used as a sprite in the game.
##### This object type has the following attributes.
| Attribute | Description |
:-: | -----------
x | Represents the position of the object on the x-axis of the window.
y | Represents the position of the object on the y-axis of the window.negative.
width | Stores the width size of the object.
height | Stores the height size of the object.
image | Stores the path to the image on the project.

* `Text` - A Text object is used to store a message that can be drawn on the window.
##### This object type has the following attributes.
| Attribute | Description |
:-: | -----------
x | Represents the position of the object on the x-axis of the window.
y | Represents the position of the object on the y-axis of the window.negative.
message | Stores the message that can be drawn on the window.

* `Background` - A Background is used to store the image that can be used as background for the window.
##### This object type has the following attributes.
| Attribute | Description |
:-: | -----------
width | Stores the width size of the object.
height | Stores the height size of the object.
image | Stores the path to the image on the project.

<!-- PREDEFINED FUNCTIONS -->
### Predefined Functions
The following functions were predefined to help the developer with the creation of 2D videogames and some of them for diverse uses.

#### For game development
| Attribute | Description |
:-: | -----------
Render() | Draws the object on the window.
KeyPressed() | Checks if a key that was pressed was already released.
CheckCollision() | Checks if two objects collide with each other.
Update() | Updates the window with its rendered objects.
Clear() | Clears the rendered objects of the window.

#### For diverse uses
| Function | Description |
:-: | -----------
Print() | Prints a basic type on the console.
Pow() | Returns the power of a number.
Sqrt() | Returns the square root of a number.

#### Predefined Functions important notes
| Function | Important Note |
:-: | -----------
Render() | This method can be used with any object data type (Square, Circle, etc.).
KeyPressed() | This method can be used only with the keys specified in the Game Engine section.
CheckCollision() | This method can be used only with the objects Square and Circle, you can check collisions between Circle and Circle, but also between Circle and Square and so on.
Update() | This method must be always present after rendering the objects. This is shown in more detail in the example of the section Game Engine.
Clear() | This method must be always present before rendering the objects. This is shown in more detail in the example of the section Game Engine.
Print() | This method can receive as well a function call in case the function that is called returns a basic type.
Pow() | This method can only be used with float numbers. It receives two parameters, the first one is the base and the second one is the power.
Sqrt() | This method can only be used with a float number. It receives only one parameter, which is the number you want to get the square root from.

<!-- USAGE -->
## Usage
This section explains Vimo syntax step by step to create a basic program with examples and also how to use the predefined objects and methods to develop a simple game.

<!-- CODE SYNTAX -->
### Code Syntax
This section explains Vimo code syntax, it is shown all what you need to learn to create your first program using this programming language.

#### Variables and functions declaration
```sh
// You must declare global variables inside the curly brackets if needed
{
  int a, b, c;
  float f;
  char c;
  string s;
  bool b;
}

// You can also declare local variables inside functions
// Void is the return type, it can also be any of the basic types
void functionExample() {
  int localVariable;
}

int sumTwoNumbers(int number1, int number2) {
  return number1, number2;
}
```

#### Arrays declaration
```sh
  int[10] arrInt;
  float[5] arrFloat;
```
#### Important notes
* Arrays must be declared with its size.
* Arrays can be of any type, even our predefined object data types.

#### Expressions and Assignment
```sh
{
  int GlobalInt;
}

void main() {
  float f;
  GlobalInt = 10;
  f = 10.0;
  int a, b, c, d;
  a = 1;
  b = 2;
  c = 3;
  d = 4;
  int arrInt[10];
  int arrElement;

  // Array assignment
  arrInt[0] = 1;

  // Array access
  arrElement = arrInt[0];
  
  int arExpressionResult;
  // Arithmetic expressions
  arExpressionResult = a + 1 * (b + 100 / (c / 2)) - d

  bool andExpressionResult, orExpressionResult, eqExpressionResult;
  // Logic expressions
    // And (can only be used for booleans)
  andExpressionResult = (true && false);
    // Or (can only be used for booleans)
  orExpressionResult = (true || false);
    // Equal (can be used for all basic types)
  eqExpressionResult = (5 == 1);

  bool relExpressionResult;
  // Relational expressions
  relExpressionResult = ((5 < 2) && (10 < 9) || (5 > 1));
}
```
Note: To assign a value to a variable, you must use the '=' operator after and the value must be the same type as the variable you are trying to assign it to.

#### For loop
```sh
for(i = 0; i < 5; i = i + 1) {
  // Code block
}
```

#### While loop
```sh
while(i < 5) {
  // Code block
}
```

#### If statement
```sh
if(i < 5) {
  // Code block
}
```

#### If-else statement
```sh
if (i == 10) {
  // Code block
}
else{
  // Code block
}
```

#### Return statement
```sh
int functionReturn() {
  return 1;
}
```

#### Print statement
```sh
void functionPrint() {
  print("Print example");
}
```

Now that you have learned the code syntax for Vimo, the next section explains how to create your first program, where it is shown how to print "Hello World" using a global variable.

### Your First Program
The first thing you must to do when creating a Vimo program is to declare the header.
```sh
program HelloWorld;
```

After that, you may declare the global variables that you need for your program inside curly brackets.
```sh
{
  string s;
}
```

In case you don't need global variables, you can simply leave it empty.
```sh
{

}
```

Once the global variables are declared, you can declare as many functions as you wish.
```sh
void printHelloWorld() {
  s = "Hello World";
  print(s);
}
```

In case you have functions declared, you must call them inside the main function, which is shown in the following example as the final code.
sh
```sh
program HelloWorld;

{
  string s;
}

void printHelloWorld() {
  s = "Hello World";
  print(s);
}

void main() {
  printHelloWorld();
}
```
Above code output.
```sh
"Hello World"
```

Congratulations, you just coded and executed successfully your first Vimo program! Following the next section shows you some algorithms implementations using Vimo.

<!-- BASIC EXAMPLES -->
### Basic Examples
This section shows basic examples of code that can be done using Vimo, as well as the implementation of some well-known algorithms.

#### A program that sums two numbers
```sh
program AddTwoNumbers;

{

}

int addTwoNumbers(int number1, int number2) {
  return number1 + number2;
}

void main() {
  int result;
  result = addTwoNumbers(1, 2);
  print(result); // result = 3
}
```

#### A program that sums all of the elements inside an array of integers
```sh
program SumArrayElements;

{
  int i;
}

int sumArrElem(int[10] arr) {
  int temp;
  temp = 0;
  for(i = 0; i < 10; i = i + 1) {
    temp = temp + arr[i];
  }

  return temp;
}

void main() {
  int[10] arr;
  arr[0] = 5;
  arr[1] = 2;
  arr[2] = 3;
  arr[3] = 10;
  arr[4] = 7;
  arr[5] = 1;
  arr[6] = 4;
  arr[7] = 8;
  arr[8] = 6;
  arr[9] = 9;
  int result;
  result = sumArrElem(arr);
  print(result); // result = 55
}
```

#### A program that finds an element inside an array of integers
```sh
program FindElement;

{
  int i;
}

int find(int n, int[10] arr){
    for(i = 0; i < 10; i = i + 1) {
        if(arr[i] == n) {
            return i;
        }
    }
    return -1;
}

void main() {
  int[10] unsortedList;
  unsortedList[0] = 5;
  unsortedList[1] = 2;
  unsortedList[2] = 3;
  unsortedList[3] = 10;
  unsortedList[4] = 7;
  unsortedList[5] = 1;
  unsortedList[6] = 4;
  unsortedList[7] = 8;
  unsortedList[8] = 6;
  unsortedList[9] = 9;

  print(find(4, unsortedList));
}
```

#### Fibonacci iterative version
```sh
program FiboIterative;

{
    int i;
}

int fibIt(int n) {
    int a, b, tmp;
    a = 0;
    b = 1;

    if (n < 2) {
        return n;
    }

    for (i = 2; i < n; i = i+1) {
        tmp = a + b;
        a = b;
        b = tmp;
    }

    return b;
}

void main() {
  print(fibIt(5));
}
```

#### Fibonacci recursive version
```sh
program FiboRecursive;

{
    int i;
}

int fib(int n) {
    if (n == 2) {
        return 1;
    }

    if (n == 1) {
        return 0;
    }

    int y;
    y = fib(n - 2);
    return fib(n - 1) + y;
}

void main() {
  print(fib(5));
}
```

#### Bubble sort implementation
```sh
program BubbleSort;

{
  int i;
}

void printList(int[10] arr){
    int i;
    for (i = 0; i < 10; i = i + 1) {
        print(arr[i]);
    }
}

void bubbleSort(int[10] arr){
  int i, j;
  for(i=0; i<9; i=i+1){
      for(j=0; j<10 - i - 1; j=j+1){
          if(arr[j]>arr[j+1]){
              int tmp, tmp2;
              tmp = arr[j+1];
              tmp2 = arr[j];
              arr[j+1] = tmp2;
              arr[j] = tmp;
          }
      }
  }
  printList(arr);
}

void main() {
  int[10] unsortedList;
  unsortedList[0] = 5;
  unsortedList[1] = 2;
  unsortedList[2] = 3;
  unsortedList[3] = 10;
  unsortedList[4] = 7;
  unsortedList[5] = 1;
  unsortedList[6] = 4;
  unsortedList[7] = 8;
  unsortedList[8] = 6;
  unsortedList[9] = 9;
  
  print("Unsorted list");
  printList(unsortedList);
  print("Sorted list");
  bubbleSort(unsortedList);
}
```

* More examples can be found [here](https://github.com/sdkvictor/golang-compiler/tree/main/run/examples).

<!-- GAME ENGINE -->
### Game Engine
This section contains the code syntax, the detailed usage of the predefined methods dedicated for game development and how to implement some features in a videogame.

#### Keys supported by Vimo that can be used for game development.
| Key | Description |
:-: | -----------
"Up" | Represents the up arrow key.
"Left" | Represents the left arrow key.
"Down" | Represents the down arrow key.
"Right | Represents the right arrow key.
"Space" | Represents the space key.
"Enter" | Represents the enter key.
"Backspace" | Represents the backspace key.
"Esc" | Represents the Esc key.
"A - Z" | Any letter from the english alphabet can be represented by a capital letter (Example: "A").
"MouseLeft" | Represents the left mouse button.


#### Basic structure of a program.
It is the same as in a normal program as shown before, the only difference is that the predefined objects must be used to draw them on the window as shown in the following example.
```sh
program FirstVideogame;

// Global variables declaration
{
    
}

// Method to update all objects each frame
int tick() {
    
}

// Main function, where you call all the needed functions
void main() {
  
}
```
Important Note
* We highly recommend having a tick() method, or as you wish to name it, which will be run inside a loop to update all the objects of your game. More details are shown in the next code blocks on how to implement this function.

#### How to declare the predefined objects and use its respective attributes.
```sh
program FirstVideogame;

// Global variables declaration
{
    // Declaring the objects as global objects, you can also
    // declare them inside a function as local variables, 
    // that depends on you.
    Square mySquare;
    Circle myCircle;
    Image myImage;
    Text myText;
    bool gameRunning;
}

// Method to update all objects each frame
int tick() {
    
}

// Method to assign all the attributes of all the objects
void objectsAssignment() {
    // Square attributes setup, only has the following attributes
    mySquare.x = 10.0;
    mySquare.y = 200.0;
    mySquare.width = 50.0;
    mySquare.height = 50.0;
    mySquare.color = "ffffff";

    // Circle attributes setup, only has the following attributes
    myCircle.x = 700.0;
    myCircle.y = 200.0;
    myCircle.width = 50.0;
    myCircle.height = 50.0;
    myCircle.color = "808080";

    // Image attributes setup, only has the following attributes
    myImage.image = "smileyface.png";
    myImage.x = 350.0;
    myImage.y = 200.0;
    myImage.width = 50.0;
    myImage.height = 50.0;

    // Text attributes setup, only has the following attributes
    myText.message = "GAME RUNNING";
    myText.x = 100;
    myText.y = 100;
}

// Main function, where you call all the needed functions
void main() {
    // Call of the function to declare all the attributes of the objects
    objectsDeclaration();
}
```

#### How to draw the predefined objects on the window.
To draw the objects, use a while loop to run the game and inside you call the function tick() or however you want to name it, to update all the objects that you wish to draw on the window. Inside the tick() method, to update all the objects each frame, first you have to call the method Clear() which is in charge of clearing all the objects. Then you may insert all the code you need to update the objects as you wish but after that, you must render all the objects. That is done in this example inside the method render() where we use the predefined method Render() with all the objects that are drawn on the window. Finally, the method Update() must be used to update all the objects on each frame, this is specially important when you want to move an object on the window.

```sh
program FirstVideogame;

// Global variables declaration
{
    // Declaring the objects as global objects, you can also
    // declare them inside a function as local variables, 
    // that depends on you.
    Square mySquare;
    Circle myCircle;
    Image myImage;
    Text myText;
    bool gameRunning;
}

// Method to update all objects each frame
void tick() {
    // Method to clear all the elements on the window
    Clear();

    // Here you can write code that updates the objects

    // Call of the function that draws the objects on the window
    render();

    // Method to update the window with the new rendered objects in the next frame
    Update();
}

// Method to assign all the attributes of all the objects
void objectsAssignment() {
    // Square attributes setup, only has the following attributes
    mySquare.x = 10.0;
    mySquare.y = 200.0;
    mySquare.width = 50.0;
    mySquare.height = 50.0;
    mySquare.color = "ffffff";

    // Circle attributes setup, only has the following attributes
    myCircle.x = 700.0;
    myCircle.y = 200.0;
    myCircle.width = 50.0;
    myCircle.height = 50.0;
    myCircle.color = "808080";

    // Image attributes setup, only has the following attributes
    myImage.image = "discord.png";
    myImage.x = 350.0;
    myImage.y = 200.0;
    myImage.width = 50.0;
    myImage.height = 50.0;

    // Text attributes setup, only has the following attributes
    myText.message = "GAME RUNNING";
    myText.x = 100.0;
    myText.y = 100.0;
}

// Method to draw all the objects on the window
void render() {
    Render(mySquare);
    Render(myCircle);
    Render(myImage);
    Render(myText);
}

// Main function, where you call all the needed functions
void main() {
    // Call of the function to declare all the attributes of the objects
    objectsAssignment();

    gameRunning = true;

    // While loop to update all the objects each frame
    while(gameRunning){
        tick();
    }
    
}
```
#### The above code shows the following window
![image](https://raw.githubusercontent.com/sdkvictor/golang-compiler/main/example-render.png)

#### How to move the objects created.
This section explains how to move objects on the window. To do this, two methods are created, each one in charge of moving its respective object. Inside each function, the predefined method KeyPressed() is used to check which key is pressed and depending on that, the x and y attributes are modified for the correct movement of the objects. Both functions must be called after inside the tick() method or however you named it, after using the method Clear() and before using the methods render(), where we render all the objects, and the method Update(), which is shown the code below. 
* Note: The render of the image was set as a comment for the next block of code because it was used only as a example on how to use images as sprites, but it won't be used anymore.
* Note: The key space isn't used in this example, but you can use it if you wish, it works the same as shown in the example with the predefined method KeyPressed().
```sh
program FirstVideogame;

// Global variables declaration
{
    // Declaring the objects as global objects, you can also
    // declare them inside a function as local variables, 
    // that depends on you.
    Square mySquare;
    Circle myCircle;
    Image myImage;
    Text myText;
}

// Method to update all objects each frame
void tick() {
    // Method to clear all the elements on the window
    Clear();

    // Here you can write code that updates the objects

    // Call of the function in charge of moving the square
    moveSquare();

    // Call of the function in charge of moving the circle
    moveCircle();

    // Call of the function that draws the objects on the window
    render();

    // Method to update the window with the new rendered objects in the next frame
    Update();
}

// Method to assign all the attributes of all the objects
void objectsAssignment() {
    // Square attributes setup, only has the following attributes
    mySquare.x = 10.0;
    mySquare.y = 200.0;
    mySquare.width = 50.0;
    mySquare.height = 50.0;
    mySquare.color = "ffffff";

    // Circle attributes setup, only has the following attributes
    myCircle.x = 700.0;
    myCircle.y = 200.0;
    myCircle.width = 50.0;
    myCircle.height = 50.0;
    myCircle.color = "808080";

    // Image attributes setup, only has the following attributes
    myImage.image = "discord.png";
    myImage.x = 350.0;
    myImage.y = 200.0;
    myImage.width = 50.0;
    myImage.height = 50.0;

    // Text attributes setup, only has the following attributes
    myText.message = "GAME RUNNING";
    myText.x = 310.0;
    myText.y = 450.0;
}

// Method to draw all the objects on the window
void render() {
    Render(mySquare);
    Render(myCircle);
    //Render(myImage);
    Render(myText);
}

// Method to move the square
void moveSquare() {
    if(KeyPressed("W")){
        mySquare.y = mySquare.y + 10.0;
    }
    if(KeyPressed("A")){
        mySquare.x = mySquare.x - 10.0;
    }
    if(KeyPressed("S")){
        mySquare.y = mySquare.y - 10.0;
    }
    if(KeyPressed("D")){
        mySquare.x = mySquare.x + 10.0;
    }
}

// Method to move the circle
void moveCircle() {
    if(KeyPressed("Up")){
        myCircle.y = myCircle.y + 10.0;
    }
    if(KeyPressed("Left")){
        myCircle.x = myCircle.x - 10.0;
    }
    if(KeyPressed("Down")){
        myCircle.y = myCircle.y - 10.0;
    }
    if(KeyPressed("Right")){
        myCircle.x = myCircle.x + 10.0;
    }
}

// Main function, where you call all the needed functions
void main() {
    // Call of the function to declare all the attributes of the objects
    objectsAssignment();

    // While loop to update all the objects each frame
    while(true){
        tick();
    }
    
}
```
#### The above code shows the following window
![2021-06-02](https://raw.githubusercontent.com/sdkvictor/golang-compiler/main/example-move.gif)

#### How to check collisions between objects.
This section shows how to check when objects collide. To do this, a method was added where the predefined method CheckCollision() is implemented. It receives the two objects as parameter and returns a boolean, true in case they collided, false others. This can be seen in the finished code example below.
```sh
program FirstVideogame;

// Global variables declaration
{
    // Declaring the objects as global objects, you can also
    // declare them inside a function as local variables, 
    // that depends on you.
    Square mySquare;
    Circle myCircle;
    Image myImage;
    Text myText;
    bool gameRunning;
}

// Method to update all objects each frame
void tick() {
    // Method to clear all the elements on the window
    Clear();

    // Here you can write code that updates the objects 

    // Call of the function in charge of moving the square
    moveSquare();

    // Call of the function to check if the square and circle collide
    objectsCollide();

    // Call of the function in charge of moving the circle
    moveCircle();

    // Call of the function that draws the objects on the window
    render();

    // Method to update the window with the new rendered objects in the next frame
    Update();
}

// Method to assign all the attributes of all the objects
void objectsAssignment() {
    // Square attributes setup, only has the following attributes
    mySquare.x = 10.0;
    mySquare.y = 200.0;
    mySquare.width = 50.0;
    mySquare.height = 50.0;
    mySquare.color = "ffffff";

    // Circle attributes setup, only has the following attributes
    myCircle.x = 700.0;
    myCircle.y = 200.0;
    myCircle.width = 50.0;
    myCircle.height = 50.0;
    myCircle.color = "808080";

    // Image attributes setup, only has the following attributes
    myImage.image = "discord.png";
    myImage.x = 350.0;
    myImage.y = 200.0;
    myImage.width = 50.0;
    myImage.height = 50.0;

    // Text attributes setup, only has the following attributes
    myText.message = "GAME RUNNING";
    myText.x = 310.0;
    myText.y = 450.0;
}

// Method to draw all the objects on the window
void render() {
    Render(mySquare);
    Render(myCircle);
    //Render(myImage);
    Render(myText);
}

// Method to move the square
void moveSquare() {
    if(KeyPressed("W")){
        mySquare.y = mySquare.y + 10.0;
    }
    if(KeyPressed("A")){
        mySquare.x = mySquare.x - 10.0;
    }
    if(KeyPressed("S")){
        mySquare.y = mySquare.y - 10.0;
    }
    if(KeyPressed("D")){
        mySquare.x = mySquare.x + 10.0;
    }
}

// Method to move the circle
void moveCircle() {
    if(KeyPressed("Up")){
        myCircle.y = myCircle.y + 10.0;
    }
    if(KeyPressed("Left")){
        myCircle.x = myCircle.x - 10.0;
    }
    if(KeyPressed("Down")){
        myCircle.y = myCircle.y - 10.0;
    }
    if(KeyPressed("Right")){
        myCircle.x = myCircle.x + 10.0;
    }
}

// Method in charge of detecting a collision between the circle 
// and the square and ending the game in case that happens
void objectsCollide() {
    if (CheckCollision(myCircle, mySquare)) {
        myText.message = "GAME OVER";
        gameRunning = false;
    }
}

// Main function, where you call all the needed functions
void main() {
    // Call of the function to declare all the attributes of the objects
    objectsAssignment();

    gameRunning = true;

    // While loop to update all the objects each frame
    while(gameRunning){
        tick();
    }
}
```

#### The above code shows the following window
![2021-06-02](https://raw.githubusercontent.com/sdkvictor/golang-compiler/main/example-collision.gif)

<!-- VIDEO TUTORIAL -->
## Video Tutorial
A video tutorial that shows how to create a file, write code and run it to see the output can be found [here](https://www.youtube.com/watch?v=hwT5YB1-V-s).

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<!-- DEVELOPERS -->
## Developers:
* [Moisés Fernández](https://github.com/moyfdzz)
* [Víctor Villarreal](https://github.com/sdkvictor)

<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements
* [README Template](https://github.com/othneildrew/Best-README-Template/blob/master/README.md) - the template for the README of this project
* [Hex to RGB](https://gist.github.com/CraigChilds94/6514edbc6a2db5e434a245487c525c75) - gist of code to convert from hexadecimal to RGB
