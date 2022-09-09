module Main exposing (..)

import Browser
import Html exposing (..)
import Http
import Json.Decode exposing (Decoder, bool, field, int, list, map3)


main =
    Browser.element
        { init = init
        , update = update
        , view = view
        , subscriptions = subscriptions
        }


init : () -> ( Model, Cmd Msg )
init _ =
    ( Loading, getData )


type Model
    = Loading
    | Failure
    | Success MastermindRawData


type Msg
    = GotData (Result Http.Error MastermindRawData)


update : Msg -> Model -> ( Model, Cmd Msg )
update msg _ =
    case msg of
        GotData result ->
            case result of
                Ok data ->
                    ( Success data, Cmd.none )

                Err _ ->
                    ( Failure, Cmd.none )


view : Model -> Html Msg
view model =
    viewGame model


viewGame : Model -> Html Msg
viewGame model =
    case model of
        Loading ->
            text "Loading..."

        Failure ->
            text "Failed :("

        Success data ->
            let
                d =
                    Debug.toString data |> Debug.log
            in
            text (dataToString data)


subscriptions : Model -> Sub Msg
subscriptions _ =
    Sub.none


masterUrl : String
masterUrl =
    "http://localhost:8080/mastermind"


type Color
    = Red
    | Orange
    | Yellow
    | Green
    | Blue
    | Purple


colorToLetter : Color -> String
colorToLetter c =
    case c of
        Red ->
            "r"

        Orange ->
            "o"

        Yellow ->
            "y"

        Green ->
            "g"

        Blue ->
            "b"

        Purple ->
            "p"


type alias ColorSet =
    List Color


colorSetToLetters : ColorSet -> String
colorSetToLetters cs =
    List.map colorToLetter cs |> String.join ", "


type alias MastermindRawData =
    { won : Bool
    , answer : List Int
    , guesses : List (List Int)
    }


dataToString : MastermindRawData -> String
dataToString game =
    -- colorSetToLetters game.answer
    -- "data here..."
    List.map String.fromInt game.answer |> String.join ", "


getData : Cmd Msg
getData =
    Http.get
        { url = masterUrl
        , expect = Http.expectJson GotData dataDecoder
        }


dataDecoder : Decoder MastermindRawData
dataDecoder =
    map3 MastermindRawData
        (field "won" bool)
        (field "answer" (list int))
        (field "guesses" (list (list int)))
