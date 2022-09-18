module Main exposing (..)

import Array exposing (Array)
import Browser
import Html
import Html.Attributes as Attributes
import Http
import Json.Decode exposing (Decoder, bool, field, int, list, map4)


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
    ( Loading, getRawData )


subscriptions : Model -> Sub Msg
subscriptions _ =
    Sub.none



-- model


type Model
    = Loading
    | Failure
    | Playing GameData


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


gameLen : Int
gameLen =
    7


gameUrl : String
gameUrl =
    "http://localhost:8080/mastermind"



-- update


update : Msg -> Model -> ( Model, Cmd Msg )
update msg _ =
    case msg of
        GotData result ->
            case result of
                Ok rawData ->
                    ( Playing (convertRawDataToGameData rawData), Cmd.none )

                Err _ ->
                    ( Failure, Cmd.none )



-- view


view : Model -> Html.Html Msg
view model =
    case model of
        Loading ->
            Html.text "loading..."

        Failure ->
            Html.text "Failed :("

        Playing game ->
            generateGameHtml game


generateGameHtml : GameData -> Html.Html Msg
generateGameHtml game =
    let
        generateBoard : Html.Html Msg
        generateBoard =
            Html.div [ Attributes.class "game_board" ] (List.range 0 (gameLen - 1) |> List.map (gameRow game))

        generateGuessingButtons : Html.Html Msg
        generateGuessingButtons =
            Html.div [ Attributes.class "game_guessing_buttons" ] [ Html.button [] [ Html.text "Press me!!!!" ] ]
    in
    Html.div [ Attributes.class "game_section" ] [ generateBoard, generateGuessingButtons ]


gameRow : GameData -> Int -> Html.Html Msg
gameRow game i =
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

        displayRound : Html.Html Msg
        displayRound =
            Html.p []
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
                            ++ " - - "
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



-- dataFetch


getRawData : Cmd Msg
getRawData =
    Http.get
        { url = gameUrl
        , expect = Http.expectJson GotData rawDataDecoder
        }


rawDataDecoder : Decoder RawData
rawDataDecoder =
    map4 RawData
        (field "won" bool)
        (field "answer" (list int))
        (field "guesses" (list (list int)))
        (field "results" (list (list int)))


convertRawDataToGameData : RawData -> GameData
convertRawDataToGameData raw =
    let
        getColorSetFromInts : List Int -> ColorSet
        getColorSetFromInts l =
            List.map getColorFromInt l

        getResultSetFromInts : List Int -> GameResultsSet
        getResultSetFromInts l =
            List.map getGameResultFromInt l
    in
    { state = getGameState raw
    , answer = getColorSetFromInts raw.answer
    , guesses = List.map getColorSetFromInts raw.guesses |> Array.fromList
    , results = List.map getResultSetFromInts raw.results |> Array.fromList
    }


getGameState : RawData -> GameState
getGameState raw =
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


getGameResultFromInt : Int -> GameResult
getGameResultFromInt i =
    case i of
        1 ->
            White

        2 ->
            Black

        _ ->
            NoGameResult
