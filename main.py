import json
import sys
from llama import llamav2
from telegram import Update
from telegram.ext import Updater,  MessageHandler, filters, CallbackContext, Application, ContextTypes, ConversationHandler

#testing telegram API w python

def main():
    with open("secrets.json") as secrets:
        secret = json.load(secrets)
    secret = secret['secret']
    application = Application.builder().token(secret).build()
    conv_handler = MessageHandler(filters.TEXT, echo)
    application.add_handler(conv_handler)

    # Run the bot until you press Ctrl-C
    application.run_polling(allowed_updates=Update.ALL_TYPES)

async def echo(update: Update, context: ContextTypes.DEFAULT_TYPE) -> int:
    """
    This function would be added to the dispatcher as a handler for messages coming from the Bot API
    """
    if update is not None:
        print(f'{update.message.from_user.first_name} wrote {update.message.text}')
        answer = llamav2(update.message.text)
        print(f'answer:  {answer}')
        # This is equivalent to forwarding, without the sender's name
        await update.message.reply_text(text=answer, reply_to_message_id=update.message.message_id)
        return ConversationHandler.END

if __name__ == "__main__":
    main()