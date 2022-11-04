module Main exposing (..)

import Array exposing (Array)
import Browser
import Html
import Html.Attributes as Attributes
import Html.Events as Events
import Http
import Json.Decode as Decode
import Json.Encode as Encode


main : Program () Model Msg
main =
    Browser.element
        { init = init
        , view = view
        , update = update
        , subscriptions = subscriptions
        }


init : () -> ( Model, Cmd Msg )
init _ =
    ( emptyModel, sendRequest gameUrl Nothing )


subscriptions : Model -> Sub Msg
subscriptions _ =
    Sub.none


type alias Model =
    { dataState : DataState
    , gameData : GameData
    , currentGuess : ColorSet
    }


emptyModel : Model
emptyModel =
    { dataState = Loading, gameData = emptyGameData, currentGuess = [] }


type alias RawGameData =
    { won : Bool
    , answer : List Int
    , guesses : List (List Int)
    , results : List (List Int)
    }


type alias GameData =
    { state : GameState
    , answer : ColorSet
    , guesses : Array ColorSet
    , results : Array GameResultsSet
    }


emptyGameData : GameData
emptyGameData =
    { state = InGame
    , answer = []
    , guesses = Array.fromList []
    , results = Array.fromList []
    }


type DataState
    = Loading
    | Failure
    | Success


type GameState
    = InGame
    | Won
    | Lost


type Color
    = NoColor
    | Red
    | Orange
    | Yellow
    | Green
    | Blue
    | Purple


type alias ColorSet =
    List Color


type GameResult
    = NoGameResult
    | White
    | Black


type alias GameResultsSet =
    List GameResult


type Msg
    = GotData (Result Http.Error RawGameData)
    | GuessBtnClick Color
    | ClearBtnClick
    | NewGame


gameLen : Int
gameLen =
    7


gameUrl : String
gameUrl =
    "http://localhost:8080/mastermind"



-- update


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        GotData result ->
            case result of
                Ok rawData ->
                    ( { model
                        | dataState = Success
                        , gameData = convertRawDataToGameData rawData
                        , currentGuess = []
                      }
                    , Cmd.none
                    )

                Err _ ->
                    ( { model | dataState = Failure }, Cmd.none )

        GuessBtnClick c ->
            case c of
                NoColor ->
                    if model.gameData.state == InGame && List.length model.currentGuess >= 4 then
                        ( { model | dataState = Loading }
                        , sendRequest
                            gameUrl
                            (Just (getBodyForGuessRequest model.gameData model.currentGuess))
                        )

                    else
                        ( model, Cmd.none )

                _ ->
                    if List.length model.currentGuess >= 4 then
                        ( model, Cmd.none )

                    else
                        ( { model | currentGuess = c :: model.currentGuess }, Cmd.none )

        ClearBtnClick ->
            ( { model | currentGuess = [] }, Cmd.none )

        NewGame ->
            init ()



-- view


view : Model -> Html.Html Msg
view model =
    let
        genAppHtml : Html.Html Msg
        genAppHtml =
            Html.div
                [ Attributes.class "app" ]
                [ Html.h1 [ Attributes.class "mastermind_title" ] [ Html.text "Mastermind" ]
                , generateGameHtml model.gameData model.currentGuess
                , Html.hr [ Attributes.class "credits_divider" ] []
                , Html.h2 [ Attributes.class "credits_title" ] [ Html.text "Credits" ]
                , Html.p
                    [ Attributes.class "credits" ]
                    [ Html.text "Mastermind game by "
                    , Html.a
                        [ Attributes.href "https://github.com/48thFlame/", Attributes.class "credits_link" ]
                        [ Html.text "48thFlame" ]
                    , Html.text " "
                    , Html.a
                        [ Attributes.href "https://github.com/48thFlame/Game-hub", Attributes.class "credits_link" ]
                        [ Html.text "Repo" ]
                    ]
                ]
    in
    case model.dataState of
        Failure ->
            Html.text "Something went wrong :( try reloading page."

        Loading ->
            genAppHtml

        Success ->
            genAppHtml


generateGameHtml : GameData -> ColorSet -> Html.Html Msg
generateGameHtml game currentGuess =
    let
        playing : Bool
        playing =
            game.state == InGame

        generateBoard : Html.Html Msg
        generateBoard =
            let
                getColorSetFromMaybe : Maybe ColorSet -> ColorSet
                getColorSetFromMaybe mcs =
                    case mcs of
                        Just cs ->
                            cs

                        Nothing ->
                            List.repeat 4 NoColor

                getGameResultSetFromMaybe : Maybe GameResultsSet -> GameResultsSet
                getGameResultSetFromMaybe mgrs =
                    case mgrs of
                        Just grs ->
                            grs

                        Nothing ->
                            List.repeat 4 NoGameResult

                guessResultGameBoardSep : String
                guessResultGameBoardSep =
                    " ------ "

                gameRow : Int -> Html.Html Msg
                gameRow i =
                    Html.div
                        [ Attributes.class "game_board_row" ]
                        [ Html.div
                            []
                            [ Html.span
                                [ Attributes.class "game_row_round_number" ]
                                [ Html.text ("Round " ++ String.fromInt (i + 1) ++ ":")
                                ]
                            , Html.br [] []
                            , Html.span
                                [ Attributes.class "game_row_content" ]
                                [ Html.text
                                    ((Array.get i game.guesses
                                        |> getColorSetFromMaybe
                                        |> List.map colorToString
                                        |> String.join ""
                                     )
                                        ++ guessResultGameBoardSep
                                        ++ (Array.get i game.results
                                                |> getGameResultSetFromMaybe
                                                |> List.map gameResultToString
                                                |> String.join ""
                                           )
                                    )
                                ]
                            ]
                        ]
            in
            Html.div
                [ Attributes.class "game_board" ]
                (List.range 0 (gameLen - 1) |> List.map gameRow)

        generateCurrentGuess : ColorSet -> Html.Html Msg
        generateCurrentGuess cs =
            let
                make4Long : ColorSet
                make4Long =
                    List.append (List.reverse cs) (List.repeat (4 - List.length cs) NoColor)
            in
            Html.div
                [ Attributes.class "game_current_guess" ]
                [ if playing then
                    Html.text
                        ("Current guess: "
                            ++ (make4Long
                                    |> List.map colorToString
                                    |> String.join ""
                               )
                        )

                  else
                    case game.state of
                        Won ->
                            Html.h3
                                [ Attributes.class "game_result" ]
                                [ Html.text "You Won! :)" ]

                        Lost ->
                            Html.h3
                                [ Attributes.class "game_result" ]
                                [ Html.text "You lost :(" ]

                        _ ->
                            Html.text ""
                ]

        generateGameButtons : Html.Html Msg
        generateGameButtons =
            let
                gameGuessingButton : Color -> Html.Html Msg
                gameGuessingButton c =
                    Html.button
                        [ Attributes.class "game_color_button"
                        , Events.onClick (GuessBtnClick c)
                        ]
                        [ Html.text (colorToString c) ]

                guessingColors : List Color
                guessingColors =
                    [ Red, Orange, Yellow, Green, Blue, Purple ]

                gameControlButtons : List (Html.Html Msg)
                gameControlButtons =
                    [ Html.button
                        [ Attributes.class "game_clear_button"
                        , Events.onClick ClearBtnClick
                        ]
                        [ Html.text "Clear 🗙" ]
                    , Html.button
                        [ Attributes.class "game_guess_button"
                        , Events.onClick (GuessBtnClick NoColor)
                        ]
                        [ Html.text "Guess ➔" ]
                    ]

                newGameButton : Html.Html Msg
                newGameButton =
                    Html.button [ Attributes.class "newGame_button", Events.onClick NewGame ] [ Html.text "New game 🗘" ]
            in
            Html.div
                [ Attributes.class "game_buttons" ]
                (if playing then
                    List.concat
                        [ List.map gameGuessingButton guessingColors
                        , List.singleton (Html.br [] [])
                        , gameControlButtons
                        , List.singleton
                            newGameButton
                        ]

                 else
                    List.singleton newGameButton
                )
    in
    Html.div
        [ Attributes.class "game_section" ]
        [ generateBoard, generateCurrentGuess currentGuess, generateGameButtons ]


colorToString : Color -> String
colorToString c =
    case c of
        Red ->
            "🟥"

        Orange ->
            "🟧"

        Yellow ->
            "🟨"

        Green ->
            "🟩"

        Blue ->
            "🟦"

        Purple ->
            "🟪"

        NoColor ->
            -- "🔳"
            "➖"


gameResultToString : GameResult -> String
gameResultToString gr =
    case gr of
        White ->
            -- "❎"
            "⚪"

        Black ->
            -- "✅"
            "⚫"

        NoGameResult ->
            "➖"



-- data


{-| If has body will POST otherwise GET
-}
sendRequest : String -> Maybe Http.Body -> Cmd Msg
sendRequest url mBody =
    let
        rawDataDecoder : Decode.Decoder RawGameData
        rawDataDecoder =
            Decode.map4 RawGameData
                (Decode.field "won" Decode.bool)
                (Decode.field "answer" (Decode.list Decode.int))
                (Decode.field "guesses" (Decode.list (Decode.list Decode.int)))
                (Decode.field "results" (Decode.list (Decode.list Decode.int)))

        expect : Http.Expect Msg
        expect =
            Http.expectJson GotData rawDataDecoder
    in
    case mBody of
        Just body ->
            Http.post
                { url = url
                , expect = expect
                , body = body
                }

        Nothing ->
            Http.get
                { url = url
                , expect = expect
                }


convertRawDataToGameData : RawGameData -> GameData
convertRawDataToGameData raw =
    let
        getGameState : GameState
        getGameState =
            if raw.won then
                Won

            else if not raw.won && List.length raw.guesses < gameLen then
                InGame

            else
                Lost

        getColorFromInt : Int -> Color
        getColorFromInt i =
            case i of
                1 ->
                    Red

                2 ->
                    Orange

                3 ->
                    Yellow

                4 ->
                    Green

                5 ->
                    Blue

                6 ->
                    Purple

                _ ->
                    NoColor

        getColorSetFromInts : List Int -> ColorSet
        getColorSetFromInts l =
            List.map getColorFromInt l

        getGameResultFromInt : Int -> GameResult
        getGameResultFromInt i =
            case i of
                1 ->
                    White

                2 ->
                    Black

                _ ->
                    NoGameResult

        getResultSetFromInts : List Int -> GameResultsSet
        getResultSetFromInts l =
            List.map getGameResultFromInt l
    in
    { state = getGameState
    , answer = getColorSetFromInts raw.answer
    , guesses = List.map getColorSetFromInts raw.guesses |> Array.fromList
    , results = List.map getResultSetFromInts raw.results |> Array.fromList
    }


getBodyForGuessRequest : GameData -> ColorSet -> Http.Body
getBodyForGuessRequest game guess =
    let
        getBoolFromGameState : GameState -> Bool
        getBoolFromGameState gs =
            case gs of
                Won ->
                    True

                _ ->
                    False

        getIntFromColor : Color -> Int
        getIntFromColor c =
            case c of
                Red ->
                    1

                Orange ->
                    2

                Yellow ->
                    3

                Green ->
                    4

                Blue ->
                    5

                Purple ->
                    6

                NoColor ->
                    -1

        getEncodeValueFromColor : Color -> Encode.Value
        getEncodeValueFromColor c =
            getIntFromColor c |> Encode.int

        getGuessesValueFromColorSet : ColorSet -> Encode.Value
        getGuessesValueFromColorSet cs =
            Encode.list getEncodeValueFromColor cs

        getIntFromGameResult : GameResult -> Int
        getIntFromGameResult res =
            case res of
                White ->
                    1

                Black ->
                    2

                NoGameResult ->
                    -1

        getEncodeValueFromGameResult : GameResult -> Encode.Value
        getEncodeValueFromGameResult res =
            getIntFromGameResult res |> Encode.int

        getResultValueFromGameResultSet : GameResultsSet -> Encode.Value
        getResultValueFromGameResultSet grs =
            Encode.list getEncodeValueFromGameResult grs

        getJsonValueFromGameData : Encode.Value
        getJsonValueFromGameData =
            Encode.object
                [ ( "won", Encode.bool (getBoolFromGameState game.state) )
                , ( "answer", Encode.list getEncodeValueFromColor game.answer )
                , ( "guesses", Encode.array getGuessesValueFromColorSet game.guesses )
                , ( "results", Encode.array getResultValueFromGameResultSet game.results )
                ]

        getGuessJsonData : Encode.Value
        getGuessJsonData =
            Encode.object
                [ ( "game", getJsonValueFromGameData )
                , ( "guess", Encode.list getEncodeValueFromColor (List.reverse guess) )
                ]
    in
    Http.jsonBody getGuessJsonData
