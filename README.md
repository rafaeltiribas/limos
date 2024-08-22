
# Limos

**Limos** is a Telegram bot developed to provide UFF students with daily university restaurant menus. The bot can automatically send the menu as soon as it gets released and respond to various commands to provide helpful information across University. By the way, Limos was the goddess of starvation in Greek mythology 😅.

<p align="center">
  <img src="assets/short.png" width="700"/>
</p>

## Why Limos?

- ⏳**Get menu of the day earlier**:
  - 🚨 Up to THREE HOURS before the college's most popular updates channel launch time!
  - 🙋‍♂️ Decide earlier if you want to lunch with your friends at University or go out with them!
- 👋 Not a UFF student? No problem!
  - ✅ You can easily adopt Limos code to get the menu of any University you want! :D
  
## Start Using Limos!

Ready to get started with **Limos**? Join the bot on Telegram by clicking the invite link below or scan the QR code to quickly open the chat!

[Join Limos on Telegram](t.me/cardapiouff_bot)

<p align="center">
  <img src="assets/qrcode.png" width="150" alt="QR Code to join Limos"/>
</p>

Just click the link or scan the QR Code and start being the first one to know university restaurant's menu everyday! 😄

## Features

- **Daily Menu Updates**: Automatically sends the university restaurant's menu to all students as soon as it gets released. 🚨 Up to THREE HOURS before the college's most popular updates channel launch time!
- **Custom Commands**:
    - `/start`: Register as a user and start receiving daily menu updates.
    - `/ajuda`: Provides help and usage instructions.
    - `/cardapio`: Manually request the updated menu of the day.
    - `/horarios`: Get the restaurant's operating hours.
    - `/sobre`: Get the bot's description and purpose.
    - `/contato`: Get information about how to reach me.
- **Security Measures**: Rate limiting and input validation to ensure secure operation and avoid flood attacks.

## Getting Started

### Prerequisites

- Go (version 1.18 or higher)
- PostgreSQL (for chatID data storage)
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
   \`\`\`

4. Edit the config.go file with your postgres database info

5. Run the bot:

   \`\`\`bash
   go run main.go
   \`\`\`

## Project Structure

- **`cmd/limos/`**: The main entry point for the bot.
- **`internal/models/`**: Contains the data models used in the project.
- **`internal/usecase/`**: Business logic for handling user commands and formatting menus.
- **`internal/repository/`**: Function for create connections and interacting with the database.
- **`internal/mocks/`**: Mock menu files for tests.
- **`internal/config/`**: For application configuration, including paths and database parameters.
- **`internal/delivery/`**: For starting the bot and receive user requests.
- **`pkg/rate_limit/`**: Implements rate limiting to prevent abuse of the bot's functionality.

## Usage

After starting the bot, users can interact with it using the available commands. The bot will send the university restaurant's menu as soon as it gets released every day to all registered users.

## Contributing

If you'd like to contribute to **Limos**, please fork the repository, make your changes, and submit a pull request. All contributions are welcome!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

For any inquiries or support, please contact [techtiribas@gmail.com](mailto:techtiribas@gmail.com).
