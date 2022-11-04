import hikari
import lightbulb


def loadNonGamesCommands(bot: lightbulb.BotApp):
    @bot.command
    @lightbulb.command("ping", "Shows bots ping.")
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

    @bot.command
    @lightbulb.option('feedback', 'Your feedback.')
    @lightbulb.command('feedback', 'Sends your feedback to the developers.')
    @lightbulb.implements(lightbulb.commands.SlashCommand)
    async def feedback_command(ctx):
        pass

    @bot.command
    @lightbulb.command("statistics", "Sends your game statistics.")
    @lightbulb.implements(lightbulb.commands.SlashCommand)
    async def statistics_command(ctx):
        pass
