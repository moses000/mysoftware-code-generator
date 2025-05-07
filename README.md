# CRUDGen - PostgreSQL CRUD API Generator

💬 Join the Community Discussion
🚀 I am building code-generator to simplify CRUD code generation from PostgreSQL tables using Go — and I want your input! This is the start of the project, I will ensure to keep increasing it capability and functionality. 

📣 Join the Discussion →

📌 In the discussion, you can:
Ask questions

Share ideas & features you want

Report bugs or edge cases

Contribute templates or improvements

Connect with fellow developers

Let’s collaborate and make code-generator more powerful and flexible for everyone! 💡




**Code-generator** is a CLI tool written in Go that generates boilerplate CRUD code (models, handlers, routes, services, and DTOs) from a PostgreSQL database table. Ideal for kickstarting API development, it simplifies and automates repetitive tasks.

## 🚀 Features

- Generate Go models from PostgreSQL tables.
- Create handlers, routes, services, and DTOs using templates.
- Add custom fields.
- Lightweight and fast.

## 🛠️ Technologies Used

- Go
- Cobra CLI
- PostgreSQL
- YAML for config
- Text/template

## 📦 Installation

1. Clone this repository:
   git clone https://github.com/moses000/mysoftware-code-generator.git
   cd mysoftware-code-generator

2. Install dependencies:

   bash
   go mod tidy
   

3. Run the application:

   go run main.go --conn "<your_postgres_conn>" --table "<table_name>"

## 🧑‍💻 Usage

crudgen --conn "postgres://user:password@localhost:5432/dbname?sslmode=disable" --table "users"

You can also generate specific components:

bash
crudgen -c "<conn_string>" -t "users" --models
crudgen -c "<conn_string>" -t "users" --handlers


## Contributing

## 📄 License

This project is licensed under the GNU License.

## 🙏 Credits

Created by [Imoleayo Moses](https://github.com/moses000)


## CONTRIBUTING

# Contributing to CRUDGen

Thanks for your interest in contributing!

## How to Contribute

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/my-feature`).
3. Make your changes.
4. Run `go fmt` and ensure the project builds.
5. Commit your changes (`git commit -am 'Add new feature'`).
6. Push to your fork (`git push origin feature/my-feature`).
7. Create a Pull Request.

## ✅ Code Style

- Use `gofmt` for formatting.
- Keep logic modular.
- Stick to idiomatic Go practices.

## 🔄 PR Review Process

- All PRs are reviewed before merging.
- Make sure your branch is up to date with `main`.
- Be descriptive in your PR title and message.

## 📞 Need Help?

Open an [Issue](https://github.com/moses000/mysoftware-code-generator/issues) or start a [Discussion](https://github.com/moses000/mysoftware-code-generator/discussions).


## `CODE OF CONDUCT`


# Contributor Covenant Code of Conduct

## Our Pledge

We as members, contributors, and leaders pledge to make participation in our project a harassment-free experience for everyone.

## Our Standards

Examples of behavior that contributes to a positive environment:
- Respectful communication
- Constructive feedback
- Accepting different viewpoints

Examples of unacceptable behavior:
- Harassment or insults
- Discriminatory jokes or comments
- Personal attacks

