
# Limos

**Limos** is a Telegram bot designed to provide users with daily university restaurant menus. The bot can automatically send the menu at specified times and respond to various commands to provide helpful information.

## Features

- **Daily Menu Updates**: Automatically sends the university restaurant's menu to all registered users at 11:00 AM and 4:30 PM.
- **Custom Commands**:
    - `/start`: Register as a user and start receiving daily menu updates.
    - `/ajuda`: Provides help and usage instructions.
    - `/cardapio`: Manually request the menu of the day.
    - `/horarios`: Get the restaurant's operating hours.
- **Security Measures**: Rate limiting and input validation to ensure secure operation.

## Getting Started

### Prerequisites

- Go (version 1.18 or higher)
- MongoDB (for user data storage)
- A Telegram Bot Token (obtained from the [BotFather](https://core.telegram.org/bots#6-botfather))

### Installation

1. Clone the repository:

   \`\`\`bash
   git clone https://github.com/your-username/limos.git
   cd limos
   \`\`\`

2. Install dependencies:

   \`\`\`bash
   go mod tidy
   \`\`\`

3. Set up environment variables:

   Create a `.env` file in the project root directory with the following content:

   \`\`\`env
   TELEGRAM_TOKEN=your-telegram-bot-token
   MONGODB_URI=your-mongodb-connection-string
   \`\`\`

4. Run the bot:

   \`\`\`bash
   go run main.go
   \`\`\`

## Project Structure

- **`cmd/limos/`**: The main entry point for the bot.
- **`internal/models/`**: Contains the data models used in the project.
- **`internal/usecase/`**: Business logic for handling user commands, formatting menus, and interacting with the database.
- **`pkg/rate_limit/`**: Implements rate limiting to prevent abuse of the bot's functionality.

## Usage

After starting the bot, users can interact with it using the available commands. The bot will send the university restaurant's menu at 11:00 AM and 4:30 PM every day to all registered users.

## Contributing

If you'd like to contribute to **Limos**, please fork the repository, make your changes, and submit a pull request. All contributions are welcome!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

For any inquiries or support, please contact [techtiribas@gmail.com](mailto:techtiribas@gmail.com).
