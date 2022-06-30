import hikari
import lightbulb


EMBED_COLOR = 0x1F7EAD  # Color for embeds
EMBED_FOOTER_TEXT = "https://discord.gg/ZR2EspdHFQ ● Icons by: icons8.com"


def init_embed() -> hikari.Embed:
    embed = hikari.Embed(color=EMBED_COLOR)
    embed.set_footer(text=EMBED_FOOTER_TEXT, icon='./discord/assets/boost.png')

    return embed


def load_help_command(bot: lightbulb.BotApp):
    @bot.command
    @lightbulb.command('help', 'Sends a list of all available commands.')
    @lightbulb.implements(lightbulb.commands.SlashCommand)
    async def help_callback(ctx: lightbulb.context.Context) -> None:
        _commands = [command for command in ctx.bot.slash_commands]
        _commands.sort()

        embed = init_embed()
        embed.set_author(name='Help', icon='./discord/assets/help.png')
        lines: list[str] = [
            '```adoc',
            '=== Commands ===',
        ]

        for command in _commands:
            lines.append(f'- {command}')

        lines.append('```')
        embed.description = '\n'.join(lines)

        await ctx.interaction.create_initial_response(hikari.ResponseType.DEFERRED_MESSAGE_CREATE)
        await ctx.interaction.edit_initial_response(embed=embed)
