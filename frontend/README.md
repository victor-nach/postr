# postr

## Overview

**postr** is a modern React-based application designed to simplify post management. Built with TypeScript, React Query, and Tailwind CSS, postr provides an intuitive and responsive interface for creating, viewing, and managing users and posts. Seamless integration with a Golang and SQLite backend API ensures that data operations are reliable and efficient.

## Features

- **Post Management:**  View, create, and delete user postsposts through an intuitive user interface.
- **Real-Time Data Sync:** Utilizes React Query for efficient data fetching, caching, and synchronization.
- **Responsive Layout:** Built with Tailwind CSS for a mobile-first design that adapts seamlessly to any device.
- **Robust Error Handling:** Provides clear and user-friendly error messages to enhance the overall experience.
- **Seamless Backend Integration:** Connects with a Golang and SQLite powered backend API to perform reliable CRUD operations.

## Built With

- **React** – A powerful library for building user interfaces.
- **TypeScript** – Adds static typing for better code quality and maintainability.
- **React Query** – Simplifies data fetching and caching.
- **Tailwind CSS** – A utility-first CSS framework for rapid UI development.

## Setup

### Prerequisites

Before you begin, ensure you have the following installed on your system:

- **Node.js** (version 14 or higher)
- **npm** (Node Package Manager)

### Installation

1. **Clone the Repository:**

   ```bash
   git clone git@github.com:victor-nach/postr-frontend.git
   cd postr
   ```

2. **Install Dependencies:**
   ```bash
   npm install
   ```

3. **Running the Application:**
 Start the development server by running:
     ```bash
   npm run dev
   ```
  The app will be available at http://localhost:5173.

4. **Building for Production**
To build the application for production, execute:
     ```bash
   npm run build
   ```


5. **To preview the production build locally:**
     ```bash
   npm run preview
   ```

6. **Running Tests**
 To run the test suite using Vitest, use the following command:
     ```bash
   npm run test
   ```

**Backend**
The postr frontend seamlessly connects to a robust backend API built with Golang and SQLite. This backend is responsible for all data operations related to posts. For more details, check out the postr Backend Repository. https://github.com/victor-nach/postr-backend