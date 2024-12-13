# Flight Booking Application

A simple console-based flight booking application written in Go.

## Overview

This application allows users to search for flights based on their location, destination, and date of travel. Admins can add new flights and delete existing ones.

## Features

- User can search for flights based on location, destination, and date of travel
- User can view available flights with their respective prices, departure and arrival times
- Admin can add new flights to the system
- Admin can delete existing flights from the system

## Usage

To run the application, navigate to the project directory and execute the following command:

```bash
go run main.go
```

# Launching the Application

This will launch the application and present the user with a menu to choose from.

## Menu Options

The application has two main menus: one for users and one for admins.

## User Menu

- Search for flights
- Exit

## Admin Menu

- Add a new flight
- Delete a flight
- Exit

## Code Structure

The application is divided into two packages: flight and main.

The flight package contains the logic for the flight booking application, including the data structures for flights, search functionality, and admin functionality.
The main package contains the entry point of the application and handles the menu system.
Dependencies

## This application uses the following dependencies:

- Go 1.x (latest version recommended)
- No external dependencies required
