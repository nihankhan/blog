# Blog Application

This is a simple web application written in Go that provides basic functionalities for a blog. It allows users to create accounts, log in, log out, view the dashboard, and manage blog posts.

## Features

- User Registration: Users can sign up for an account by providing their information.
- User Login and Logout: Registered users can log in to the system and log out when they're done.
- Dashboard: Users have access to a dashboard where they can manage their blog posts.
- Blog Post Management: Users can create, edit, and delete their blog posts.
- Static Assets: The application serves static assets such as CSS files, images, etc.

## Prerequisites

Before running the application, ensure you have the following:

- Go programming language (https://golang.org/dl/)
- MySQL database server

## Getting Started

1. Clone the repository:

   ```
   git clone <repository_url>
   ```

2. Navigate to the project directory:

   ```
   cd blog
   ```

3. Install the required Go packages:

   ```
   go get github.com/go-sql-driver/mysql
   ```

4. Set up the MySQL database:

   - Create a new MySQL database named "blog".
   - Update the database connection configuration in the `main.go` file (`init()` function) to match your MySQL configuration:

     ```go
     models.Db, models.Err = sql.Open("mysql", "your_username:your_password@tcp(your_database_host:your_database_port)/blog")
     ```

5. Run the application:

   ```
   go run main.go
   ```

6. Access the application in your web browser:

   ```
   http://localhost:8000
   ```

## Usage

- **Registration**: Visit `/signup` to create a new account.
- **Login**: Go to `/login` to log in with your account credentials.
- **Dashboard**: Access `/dashboard` to manage your blog posts.
- **Create Post**: From the dashboard, you can create a new blog post.
- **Edit and Delete Post**: On the dashboard, you can edit or delete your existing blog posts.
- **Logout**: Click on the "Logout" link to log out of the application.

## Directory Structure

- `assets/`: Contains static assets like CSS files, images, etc.
- `controllers/`: Contains controller functions that handle HTTP requests and responses.
- `models/`: Contains database models and initialization code.
- `main.go`: The main application entry point.

## Contributions

Contributions to this project are welcome. You can contribute by submitting bug reports, feature requests, or pull requests.

## License

This project is licensed under the [MIT License](LICENSE).

---

Feel free to customize this `readme.md` to provide more details, add screenshots, or expand on specific features as needed. Make sure to update the repository URL and other details to match your actual project.
