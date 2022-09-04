import hikari
import lightbulb
from hikari import ButtonStyle


C_EQUAL_SYMBOL: str = '='
C_CLEAR_SYMBOL: str = 'C'
C_DEL_SYMBOL: str = '⌫'

C_START_MSG: str = '0\u200b'
C_START_OF_BUTTON_CUSTOM_ID: str = 'calculator_button:'
C_SYNTAX_ERROR_MSG: str = 'Syntax error!'
MAX_EMBED_TITLE_LEN: int = 255
C_ERROR_TOO_LONG: str = f'Problem too long! (max {MAX_EMBED_TITLE_LEN} characters)'
C_ERRORS: set[str] = {C_ERROR_TOO_LONG, C_SYNTAX_ERROR_MSG}

C_BUTTON_STYLE_EQUAL: int = ButtonStyle.SUCCESS
C_BUTTON_STYLE_ERASE: int = ButtonStyle.DANGER
C_BUTTON_STYLE_NUMBER: int = ButtonStyle.SECONDARY
C_BUTTON_STYLE_OPERATION: int = ButtonStyle.PRIMARY

C_BUTTONS: list[list[tuple[str, int]]] = [
    [
        ('(', C_BUTTON_STYLE_OPERATION), (')', C_BUTTON_STYLE_OPERATION),
        (C_DEL_SYMBOL, C_BUTTON_STYLE_ERASE), (C_CLEAR_SYMBOL, C_BUTTON_STYLE_ERASE)
    ],
    [
        ('7', C_BUTTON_STYLE_NUMBER), ('8', C_BUTTON_STYLE_NUMBER),
        ('9', C_BUTTON_STYLE_NUMBER), ('/', C_BUTTON_STYLE_OPERATION)
    ],
    [
        ('4', C_BUTTON_STYLE_NUMBER), ('5', C_BUTTON_STYLE_NUMBER),
        ('6', C_BUTTON_STYLE_NUMBER), ('*', C_BUTTON_STYLE_OPERATION)
    ],
    [
        ('1', C_BUTTON_STYLE_NUMBER), ('2', C_BUTTON_STYLE_NUMBER),
        ('3', C_BUTTON_STYLE_NUMBER), ('-', C_BUTTON_STYLE_OPERATION)
    ],
    [
        ('0', C_BUTTON_STYLE_NUMBER), ('.', C_BUTTON_STYLE_NUMBER),
        (C_EQUAL_SYMBOL, C_BUTTON_STYLE_EQUAL), ('+', C_BUTTON_STYLE_OPERATION)
    ]
]

C_OPERATIONS: set[str] = {'+', '-', '*', '/'}
# chars that need to have a \ before them to stop discord formatting
DISCORD_SPECIALS: set[str] = {'*'}


def load_calculator(bot: lightbulb.BotApp):
    class CalculatorCommand:
        @staticmethod
        def is_blank(title: str) -> bool:
            return title.endswith(C_START_MSG)

        @staticmethod
        def is_error(title: str) -> bool:
            return title in C_ERRORS

        @staticmethod
        def is_operation(obj: str) -> bool:
            return obj in C_OPERATIONS

        @staticmethod
        def is_discord_special(obj: str) -> bool:
            return obj in DISCORD_SPECIALS

        @staticmethod
        def gave_answer(title: str) -> bool:
            return title.endswith('\u200b')

        @staticmethod
        def clear_to_answer(title: str) -> str:
            return title[title.find('=') + 1:-1]

        def build_calculator_buttons(self) -> list:
            buttons = []
            for row in C_BUTTONS:
                action_row = bot.rest.build_action_row()
                for button_face, style in row:
                    (
                        action_row.add_button(
                            style,
                            f'{C_START_OF_BUTTON_CUSTOM_ID}{button_face}'
                        )
                        .set_label(button_face)
                        .add_to_container()
                    )
                buttons.append(action_row)

            return buttons

        @staticmethod
        def button_clear() -> str:
            return C_START_MSG

        def button_del(self, embed_title_input: str) -> str | None:
            if self.is_blank(embed_title_input):
                return C_START_MSG
            elif self.gave_answer(embed_title_input):
                embed_title = self.clear_to_answer(embed_title_input)
            elif self.is_error(embed_title_input) or len(embed_title_input) == 1:
                embed_title = C_START_MSG
            elif self.is_discord_special(embed_title_input[-1]):
                embed_title = embed_title_input[:-2]
            else:
                embed_title = embed_title_input[:-1]

            return embed_title

        def button_equal(self, embed_title_input: str) -> str:
            if self.is_blank(embed_title_input) or self.gave_answer(embed_title_input) or self.is_error(embed_title_input):
                return C_START_MSG
            else:
                answer = str(embed_title_input).replace('\\', '')
                # noinspection PyBroadException
                try:
                    answer = str(eval(answer))
                except Exception:
                    return C_SYNTAX_ERROR_MSG

            embed_title = f'{embed_title_input}={answer}\u200b'

            if len(embed_title) > MAX_EMBED_TITLE_LEN:
                return C_ERROR_TOO_LONG

            return embed_title

        def button_any(self, embed_title_input: str, value: str) -> str:
            embed_title = embed_title_input

            if self.gave_answer(embed_title_input):
                if self.is_operation(value):
                    embed_title = self.clear_to_answer(embed_title_input)
                else:  # if didn't give operation after answer, should clear title
                    embed_title = ''
            elif self.is_blank(embed_title_input):
                if value == '0':
                    return C_START_MSG
                embed_title = ''
            elif self.is_error(embed_title_input):
                embed_title = ''
            elif self.is_discord_special(value):
                embed_title += '\\'

            embed_title += value

            return embed_title

        async def callback(self, ctx: lightbulb.context.Context) -> None:
            embed = hikari.Embed(color=0x1F7EAD, title=C_START_MSG)
            buttons = self.build_calculator_buttons()

            await ctx.interaction.create_initial_response(hikari.ResponseType.DEFERRED_MESSAGE_CREATE)
            await ctx.interaction.edit_initial_response(embed=embed, components=buttons)

        async def button_press(self, event: hikari.InteractionCreateEvent) -> None:
            embed: hikari.Embed = event.interaction.message.embeds[0]
            embed_title: str = embed.title
            value: str = str(event.interaction.custom_id).removeprefix(
                C_START_OF_BUTTON_CUSTOM_ID)

            if value == C_CLEAR_SYMBOL:
                embed_title = self.button_clear()
            elif value == C_DEL_SYMBOL:
                embed_title = self.button_del(embed_title)
            elif len(embed_title) >= MAX_EMBED_TITLE_LEN:
                embed_title = C_ERROR_TOO_LONG
            elif value == C_EQUAL_SYMBOL:
                embed_title = self.button_equal(embed_title)
            else:  # any other button
                embed_title = self.button_any(embed_title, value)

            embed.title = embed_title
            await event.interaction.create_initial_response(hikari.ResponseType.MESSAGE_UPDATE, embed=embed)

    calculator_command_manager = CalculatorCommand()

    # calculator command
    @bot.command
    @lightbulb.command('calculator', 'Builds a calculator that you can use.')
    @lightbulb.implements(lightbulb.commands.SlashCommand)
    async def calculator_callback(ctx: lightbulb.context.Context) -> None:
        await calculator_command_manager.callback(ctx)

    # buttons events
    @bot.listen(hikari.InteractionCreateEvent)
    async def on_button_press(event: hikari.InteractionCreateEvent) -> None:
        if isinstance(event.interaction, hikari.ComponentInteraction) and \
                event.interaction.custom_id.startswith(C_START_OF_BUTTON_CUSTOM_ID):
            # if is button press event and not other interaction event
            await calculator_command_manager.button_press(event)
