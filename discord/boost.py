import hikari
import lightbulb

with open("./discord/TOKEN.txt", "r") as f:
    TOKEN = f.readline().strip()

bot = lightbulb.BotApp(
    TOKEN,
    default_enabled_guilds=(755001834418208840,),
    banner=None
)


@bot.command
@lightbulb.command("ping", "ping.")
@lightbulb.implements(lightbulb.commands.SlashCommand)
async def ping_command(ctx):
    pass


@bot.command
@lightbulb.option("question", "The question that the magic-8ball should answer.", str, required=True)
@lightbulb.command("8ball", "Answers a question with the magic-8ball's answer.")
@lightbulb.implements(lightbulb.commands.SlashCommand)
async def ball8_command(ctx):
    pass


@bot.command
@lightbulb.option('dice', 'Number of dice to roll.', int, required=True, min_value=1, max_value=150)
@lightbulb.option('sides', 'Number of sides each dice should have.', int, required=True, min_value=4, max_value=16)
@lightbulb.command('dice', 'Rolls dice.')
@lightbulb.implements(lightbulb.commands.SlashCommand)
async def dice_command(ctx):
    pass


@bot.command
@lightbulb.command('coinflip', 'Flips a coin.')
@lightbulb.implements(lightbulb.commands.SlashCommand)
async def coinflip_command(ctx):
    pass


@bot.command
@lightbulb.command('info', 'Sends some useful information.')
@lightbulb.implements(lightbulb.commands.SlashCommand)
async def info_command(ctx):
    pass


@bot.command
@lightbulb.option('poll', 'The poll.', str, required=True)
@lightbulb.option(
    'channel',
    'The channel to send the poll to.',
    hikari.TextableGuildChannel,
    required=True,
    channel_types=[hikari.ChannelType.GUILD_TEXT]
)
@lightbulb.command('poll', 'Send a yes/no poll to the given channel.')
@lightbulb.implements(lightbulb.SlashCommand)
async def poll_command(ctx):
    pass


@bot.listen(hikari.StartedEvent)
async def on_ready(event) -> None:
    print(
        f"{bot.get_me()} - [python version] has connected to Discord! ( ﾉ ﾟｰﾟ)ﾉ"
    )


if __name__ == "__main__":
    bot.run(
        activity=hikari.Activity(
            type=hikari.ActivityType.WATCHING, name="Avishai doing the magic..."
        )
    )
