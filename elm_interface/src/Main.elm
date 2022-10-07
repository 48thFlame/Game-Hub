module Main exposing (..)

import Array exposing (Array)
import Browser
import Html
import Html.Attributes as Attributes
import Html.Events as Events
import Http
import Json.Decode exposing (Decoder, bool, field, int, list, map4)
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
    ( emptyModel, getRawData )


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


type alias RawData =
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
    = GotData (Result Http.Error RawData)
    | Guess Color
    | Clear


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
                    ( { emptyModel
                        | dataState = Success
                        , gameData = convertRawDataToGameData rawData
                      }
                    , Cmd.none
                    )

                Err _ ->
                    ( { emptyModel | dataState = Failure }, Cmd.none )

        Guess c ->
            case c of
                NoColor ->
                    -- should guess
                    ( model, Cmd.none )

                _ ->
                    if List.length model.currentGuess >= 4 then
                        ( model, Cmd.none )

                    else
                        ( { model | currentGuess = c :: model.currentGuess }, Cmd.none )

        Clear ->
            ( { model | currentGuess = [] }, Cmd.none )



-- view


view : Model -> Html.Html Msg
view model =
    case model.dataState of
        Loading ->
            Html.text "loading..."

        Failure ->
            Html.text "Failed :("

        Success ->
            generateGameHtml model.gameData model.currentGuess


generateGameHtml : GameData -> ColorSet -> Html.Html Msg
generateGameHtml game currentGuess =
    let
        generateBoard : Html.Html Msg
        generateBoard =
            let
                gameRow : Int -> Html.Html Msg
                gameRow i =
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

                        guessGameResultSep : String
                        guessGameResultSep =
                            " ------ "

                        displayRound : Html.Html Msg
                        displayRound =
                            Html.div []
                                [ Html.span [ Attributes.class "game_row_round_number" ]
                                    [ Html.text ("Round " ++ String.fromInt (i + 1) ++ ":")
                                    ]
                                , Html.br [] []
                                , Html.span [ Attributes.class "game_row_content" ]
                                    [ Html.text
                                        ((Array.get i game.guesses
                                            |> getColorSetFromMaybe
                                            |> List.map colorToString
                                            |> String.join ""
                                         )
                                            ++ guessGameResultSep
                                            ++ (Array.get i game.results
                                                    |> getGameResultSetFromMaybe
                                                    |> List.map gameResultToString
                                                    |> String.join ""
                                               )
                                        )
                                    ]
                                ]
                    in
                    Html.div [ Attributes.class "game_board_row" ] [ displayRound ]
            in
            Html.div [ Attributes.class "game_board" ] (List.range 0 (gameLen - 1) |> List.map gameRow)

        generateCurrentGuess : ColorSet -> Html.Html Msg
        generateCurrentGuess cs =
            let
                make4Long : ColorSet
                make4Long =
                    List.append (List.reverse cs) (List.repeat (4 - List.length cs) NoColor)
            in
            Html.div [ Attributes.class "game_current_guess" ]
                [ Html.text ("Current guess: " ++ (make4Long |> List.map colorToString |> String.join ""))
                ]

        generateGameButtons : Html.Html Msg
        generateGameButtons =
            let
                gameGuessingButton : Color -> Html.Html Msg
                gameGuessingButton c =
                    Html.button [ Attributes.class "game_color_button", Events.onClick (Guess c) ] [ Html.text (colorToString c) ]

                guessingColors : List Color
                guessingColors =
                    [ Red, Orange, Yellow, Green, Blue, Purple ]

                gameControlButtons : List (Html.Html Msg)
                gameControlButtons =
                    [ Html.button [ Attributes.class "game_clear_button", Events.onClick Clear ] [ Html.text "Clear 🗙" ]
                    , Html.button [ Attributes.class "game_guess_button", Events.onClick (Guess NoColor) ] [ Html.text "Guess ➔" ]
                    ]
            in
            Html.div [ Attributes.class "game_buttons" ]
                (List.concat [ List.map gameGuessingButton guessingColors, [ Html.br [] [] ], gameControlButtons ])
    in
    Html.div [ Attributes.class "game_section" ] [ generateBoard, generateCurrentGuess currentGuess, generateGameButtons ]


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
            "❎"

        Black ->
            "✅"

        NoGameResult ->
            "➖"



-- data


getRawData : Cmd Msg
getRawData =
    let
        rawDataDecoder : Decoder RawData
        rawDataDecoder =
            map4 RawData
                (field "won" bool)
                (field "answer" (list int))
                (field "guesses" (list (list int)))
                (field "results" (list (list int)))
    in
    Http.post
        { url = gameUrl
        , expect = Http.expectJson GotData rawDataDecoder
        , body = Http.jsonBody (getJsonValueFromGameData emptyGameData)
        }


convertRawDataToGameData : RawData -> GameData
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


getJsonValueFromGameData : GameData -> Encode.Value
getJsonValueFromGameData game =
    let
        getBoolFromGameState : GameState -> Bool
        getBoolFromGameState gs =
            case gs of
                Won ->
                    True

                _ ->
                    False

        getEncodeValueFromColor : Color -> Encode.Value
        getEncodeValueFromColor c =
            let
                getIntFromColor : Int
                getIntFromColor =
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
            in
            getIntFromColor |> Encode.int

        getGuessesValueFromColorSet : ColorSet -> Encode.Value
        getGuessesValueFromColorSet cs =
            Encode.list getEncodeValueFromColor cs

        getResultValueFromGameResultSet : GameResultsSet -> Encode.Value
        getResultValueFromGameResultSet grs =
            let
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
            in
            Encode.list getEncodeValueFromGameResult grs
    in
    Encode.object
        [ ( "won", Encode.bool (getBoolFromGameState game.state) )
        , ( "answer", Encode.list getEncodeValueFromColor game.answer )
        , ( "guesses", Encode.array getGuessesValueFromColorSet game.guesses )
        , ( "results", Encode.array getResultValueFromGameResultSet game.results )
        ]
