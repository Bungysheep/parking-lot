# PARKING LOT

Parking Lot is an application that handle parking process in a Parking Center with N Slot capacity. The application is written with Golang and can accept commands interactive as well as commands with input file.

## Application Components

#### Car
The Car entity with properties: 1) Registration Number; 2) Colour.

#### Slot
The Slot means the block of parking area with properties: 1) Slot Number; 2) Car object.

#### Parking
The Parking is the parking lot which has N capacity of Slot. It has properties: 1) Capacity; N of Slots.

#### Command
Command process input and execute a proper action. It is responsible to parse action, verify action, parse args, verify args and run the action.

#### ShellProcessor
Provide functionalities where user can interact with shell.

#### FileProcessor
Provide funtionalities where user can run the application with a file input.

## How to Setup+Run the Application

#### Setup
Run the following command under the project directory: ./bin/setup

#### Run
To run in interactive shell command, run the following command under the project directory: ./bin/parking_lot

To run with input filename, run the following command under the project directory: ./bin/parking_lot <filename>
