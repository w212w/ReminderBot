# 📝 **ReminderBot** — Your Personal Page Saver 📚

## 🤖 **What is ReminderBot?**

**ReminderBot** is a **Telegram bot** written in **Go** that helps you save interesting pages and revisit them later! Whether it’s an article, a blog post, or a tutorial, this bot makes it easy to collect and manage your reading list directly in Telegram.  

---

## 📦 **Folder Structure**

📂 ReminderBot
├── 📂 cmd              # Entry point for the bot
├── 📂 clients          # Telegram client implementation
├── 📂 consumer         # Event consumer logic
├── 📂 events           # Event handlers for Telegram updates
├── 📂 storage          # SQLite storage logic
├── .env                # Environment variables (ignored in Git)
└── README.md           # Project documentation

## 🎯 **Features**

✨ **Save Pages with Ease**  
Just send a link to the bot, and it will securely save the page to your personal list.  

🎲 **Get Random Recommendations**  
Want to read something now? Use the `/rnd` command, and the bot will offer you a random page from your list.  
⚠️ *Caution: Once a page is suggested, it will be removed from your list.*  

🔒 **Your Data is Private**  
Everything is stored locally or in a private database—your reading list is only yours!  

---

## 🛠️ **Technologies Used**

- **Programming Language**: [Go](https://go.dev/)  
- **Telegram Bot API**: Seamless integration with Telegram for a smooth user experience.  
- **SQLite**: Lightweight and reliable database for storing your pages.  

---

## 🚀 **How to Use ReminderBot**
### [BOT](https://t.me/ReminderGolangBot)


1. **Start a Chat**: Add the bot to your Telegram and start a conversation.  
2. **Save a Page**: Send a link to the bot.
