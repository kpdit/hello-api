package translation_test

import (
    "testing"
    "hello-api/translation"
)

func TestTranslate(t *testing.T) {
    // Arrange
    tt := []struct{ //
        Word string
        Language string
        Translation string
    }{
        { //
            Word: "hello",
            Language: "english",
            Translation: "hello",
        },
        {
            Word: "hello",
            Language: "german",
            Translation: "hallo",
        },
        {
            Word: "hello",
            Language: "finnish",
            Translation: "hei",
        },
        {
            Word: "hello",
            Language: "dutch",
            Translation: "",
        },
        {
            Word: "bye",
            Language: "dutch",
            Translation: "",
        },
        {
            Word: "bye",
            Language: "german",
            Translation: "",
        },
        {
            Word: "hello",
            Language: "German",
            Translation: "hallo",
        },
        {
            Word: "Hello",
            Language: "german",
            Translation: "hallo",
        },
        {
            Word: "hello ",
            Language: "german",
            Translation: "hallo",
        },
    }

    for _, tl := range tt { //
        // Act
        res := translation.Translate(tl.Word, tl.Language) //

        // Assert
        if res != tl.Translation { //
            t.Errorf(`expected "%s" to be translated to "%s" to be "%s" but received "%s"`, tl.Word, tl.Language, tl.Translation, res)
        }
    }
}

