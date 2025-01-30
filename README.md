
# Currency Converter CLI 💰

A sleek and user-friendly **terminal-based currency converter** built with **Golang**! This project leverages the power of the [HUH](https://github.com/charmbracelet/huh) library for a beautiful command-line interface (CLI) experience and integrates with the [Open Exchange Rates API](https://openexchangerates.org/) to fetch real-time currency conversion data.

Whether you're a developer looking to learn Golang or someone who needs quick currency conversions, this tool is for you! 🚀

---

## Features ✨

- **Beautiful Terminal UX**: Built with the `HUH` library for an intuitive and visually appealing interface.
- **Real-Time Currency Conversion**: Fetches up-to-date exchange rates using the Open Exchange Rates API.
- **Simple and Fast**: Convert currencies directly from your terminal with just a few keystrokes.
- **Golang-Powered**: A great example of using Go for building practical, efficient, and maintainable CLI tools.

---

## How It Works 🛠️

1. **User Input**: The user selects the source currency, target currency, and amount to convert via a clean form interface.
2. **API Integration**: The tool fetches real-time exchange rates from the Open Exchange Rates API.
3. **Conversion Logic**: The amount is converted using the fetched exchange rates.
4. **Output**: The converted amount is displayed in a user-friendly format.

---

## Prerequisites 📋

Before running the project, make sure you have the following:

- **Go installed**: Download and install Go from [golang.org](https://golang.org/doc/install).
- **Open Exchange Rates API Key**: Sign up for a free API key at [Open Exchange Rates](https://openexchangerates.org/signup).
- **HUH Library**: Install the `HUH` library by running:
  ```bash
  go get github.com/charmbracelet/huh
  ```

---

## Installation 🚀

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/currency-converter-cli.git
   cd currency-converter-cli
   ```

2. Set up your API key:
   - Create a `.env` file in the root directory.
   - Add your Open Exchange Rates API key:
     ```env
     API_KEY=your_api_key_here
     ```

3. Run the project:
   ```bash
   go run main.go
   ```

---

## Usage 🖥️

Once the program is running, follow the prompts:

1. **Select the source currency** (e.g., USD).
2. **Select the target currency** (e.g., EUR).
3. **Enter the amount** you want to convert.
4. **View the result** in the terminal!

Example:
```
? Select source currency: USD
? Select target currency: EUR
? Enter amount to convert: 100
✅ 100 USD = 92.50 EUR
```

---

## Why This Project? 🤔

- **Learn Golang**: This project is a great way to practice Go fundamentals, including HTTP requests, JSON parsing, and CLI development.
- **Explore HUH**: Discover how to create beautiful terminal interfaces with the `HUH` library.
- **Real-World Application**: Build something practical while learning new technologies.

---

## Future Improvements 🔮

- [ ] Add support for historical exchange rates.
- [ ] Implement a currency list with search functionality.
- [ ] Add error handling for invalid API keys or network issues.
- [ ] Allow users to save favorite currency pairs for quick access.

---

## Contributing 🤝

Contributions are welcome! If you'd like to improve this project, feel free to open an issue or submit a pull request. Let's make this tool even better together!

---

## License 📜

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Acknowledgments 🙏

- [Open Exchange Rates](https://openexchangerates.org/) for providing the currency data.
- [HUH](https://github.com/charmbracelet/huh) for making terminal interfaces beautiful.
- The Go community for creating such an amazing programming language.

---

Happy coding! 💻✨

---

This `README.md` is designed to be visually appealing, informative, and professional. It highlights the purpose of your project, its features, and how to use it, while also encouraging contributions and acknowledging dependencies. Let me know if you’d like to tweak it further! 😊
