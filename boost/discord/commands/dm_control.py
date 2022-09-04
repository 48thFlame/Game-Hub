import subprocess
import hikari
import lightbulb

FEEDBACK_BAN_EXE_NAME = "banFeedback"
FEEDBACK_UNBAN_EXE_NAME = "unbanFeedback"


async def bad_feedback_stuff(event: hikari.DMMessageCreateEvent):
    await event.message.respond("That's not how you do it. :-(\n`+feedback <ban|unban> <user_id>`")


def load_dm_control(bot: lightbulb.BotApp):
    @bot.listen(hikari.DMMessageCreateEvent)
    async def DM_control(event: hikari.DMMessageCreateEvent) -> None:
        if event.author.id == 719948798566072382:  # only if i is who sent the message
            if event.content.startswith("+feedback"):
                msg_data = event.content.split(" ")
                msg_data.pop(0)

                if len(msg_data) != 2:
                    return await bad_feedback_stuff(event)

                user_id: str = msg_data[1]

                try:
                    int(user_id)
                except ValueError:
                    return await bad_feedback_stuff(event)

                ban = msg_data[0] == "ban"
                unban = msg_data[0] == "unban"

                if ban:
                    subprocess.run([f"./{FEEDBACK_BAN_EXE_NAME}", user_id])
                elif unban:
                    subprocess.run([f"./{FEEDBACK_UNBAN_EXE_NAME}", user_id])
                else:
                    return await bad_feedback_stuff(event)

                await event.message.respond(f"Done with {user_id}.")
