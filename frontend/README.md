Lema Full-Stack Developer Take-Home Assignment
This assignment involves building a user management system where developers must:
Create a backend server using Golang and SQLite for user and post operations.
Use this Nodejs starter template as a reference.
Develop the frontend using React, TypeScript, React Query, and Tailwind CSS to display user data in a paginated table and allow for Post management.
Follow the provided design specifications while emphasizing best practices in performance, security, and maintainability.
Check out the Figma design here

Backend (Golang)
You are required to build the backend using Golang and SQLite.
Requirements
Use Golang to create a RESTful API.
Use SQLite as the database to store users, posts, and addresses.
Use GORM (or any ORM of choice) for database interactions.
Implement proper request validation and error handling.
API Endpoints
User Endpoints
GET /users
Returns a paginated list of users (e.g., /users?pageNumber=0&pageSize=10).
GET /users/count
Returns the total number of users.
GET /users/{id}
Returns details of a specific user, including their address.

Post Endpoints
GET /posts?userId={userId}
Returns all posts for a specific user.
POST /posts
Creates a new post for a user.
Accepts title, body, and userId as input.
Validates input data before saving.
DELETE /posts/{id}
Deletes a post by its ID.
Ensures proper deletion with appropriate HTTP status codes.

Frontend (React & TypeScript)
General Requirements
Implement the UI using React, TypeScript, React Query, and Tailwind CSS.
Ensure graceful handling of API errors and unexpected backend responses.
Use React Query for efficient server-state management.
Follow Figma design specifications closely.
Users Table
Fetch a paginated list of users from the backend.
Display users in a table with the following details:
Full Name
Email Address
Address formatted as: street, state, city, zipcode (keep the column 392px wide, using ellipsis ... for overflow).
User Posts
Clicking on a user row should navigate to a new page displaying the user's posts.
Fetch posts from the /posts?userId={userId} endpoint.
The page should include:
User Summary Header
Total post count
List of all posts (no pagination required)

Post Details
Each post should display:
Title
Body
Delete Icon
Clicking should delete the post via the backend API and update the UI.
Add Post Button
Opens a form to create a new post (with Title and Body).
Upon submission, the new post should be saved without requiring a page refresh.

Development Guidelines
Backend Best Practices
Follow RESTful API principles.
Use GORM (or an ORM of choice) for database operations.
Implement proper input validation to prevent invalid data entry.
Ensure error handling with appropriate HTTP status codes and messages.
State Management with React Query
Use React Query to manage API calls efficiently.
Handle loading and error states properly.
Ensure efficient data caching and synchronization.
Code Reusability and Separation
Structure components for maintainability and reusability.
Extract shared logic into custom hooks or utility functions.
Follow best practices for component composition and prop management.
Responsiveness
Ensure a responsive UI across various devices.
Use Tailwind CSS utilities for flexibility.
Error Handling
Implement robust error handling for both frontend and backend.
Display user-friendly error messages.
Use try-catch blocks and handle promise rejections.
Deliverables
A full-stack application that meets the requirements.
Well-structured source code, ensuring readability and maintainability.
Fully implemented backend with Golang and SQLite.
Unit tests demonstrating component or API testing.
A README.md file with:
Setup instructions for both the backend and frontend.
Steps to install dependencies and run the project locally.

Submission Instructions
Code Repository: Submit via GitHub/GitLab.
Live Deployed Site: Provide a link to the live deployed version of the application.
README File: Include setup and installation instructions.


----------------------------------

# GitHub User Search App with Suggestion List

## Table of contents

- [Overview](#overview)
  - [The challenge](#the-challenge)
  - [Screenshot](#screenshot)
  - [Links](#links)
- [My process](#my-process)
  - [Built with](#built-with)
  - [Key Takeaways](#key-takeaways)
  - [Useful resources](#useful-resources)
- [Author](#author)

## Overview

This project is a GitHub User Search App designed to provide users with an intuitive interface for searching GitHub users by their username. Built with responsiveness and accessibility in mind, the app allows users to search for usernames, view a dropdown suggestion list, and see relevant user information—all while switching between light and dark themes based on user preferences.
Design from

### The challenge

Users should be able to:

- **Responsive Design**: View the optimal layout for the app depending on their device's screen size
- **Hover State**: See hover states for all interactive elements on the page.
- **User Search**: Search for GitHub users by their username.
- **Dropdown Suggestions**: See a dropdown suggestion list as they type.
- **User Information**: See relevant user information based on their search
- **Theme Switching**: Switch between light and dark themes
- **Accessibility**: Navigate the app using keyboard accessibility features.

### Screenshot

![screenshot](src/assets/images/ui-design.jpg)

### Links

- Solution URL: [GitHub Repo](https://github.com/Kellswork/github-user-search-app)
- Live Site URL: [kellswork.github.io/github-user-search-app](https://kellswork.github.io/github-user-search-app/)

## My process

### Built with

- React
- TypeScript
- React Query
- Tailwin CSS
- [React](https://react.dev/)

### Key Takeaways

- **Effective State Management**: Deepened understanding of state management patterns and daa fetching patterns, particularly when handling asynchronous data from multiple sources.
- **User-Centric Design**: Emphasized the importance of an intuitive user interface, ensuring ease of use for diverse users.
- **Custom Hooks**: Developed custom hooks for efficient data fetching.
- **Theme Management**: Implemented dynamic theme switching based on user preferences.
- **Search Optimization**: Utilized debounce and deferred values to minimize unnecessary API calls, improving efficiency and responsiveness of the input search functionality.

### Useful Resources

- [Frontend-Mentor](https://www.frontendmentor.io/challenges/github-user-search-app-Q09YOgaH6) - A fantastic platform for honing front-end skills.
- [React Documentation](https://react.dev) - JS library - For understanding React fundamentals and advanced concepts.

## Author

**Kelechi Ogbonna**

Passionate FullStack(front-end heavy) dedicated to creating user-centric applications.
