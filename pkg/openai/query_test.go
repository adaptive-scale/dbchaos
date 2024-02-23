package openai_dbchaos

import (
	"fmt"
	"testing"
)

func TestPrompt(t *testing.T) {
	query := "Generate Schema in SQL for Webshop database. Give me only SQL Commands No text."
	o := &OpenAI{
		APIkey: "",
	}
	val, err := o.Prompt(query)
	if err != nil {
		t.Errorf("Prompt() returned an error: %v", err)
	}

	val, err = o.Prompt("Also give me the some SQL commands insert same data into the " + val + ". Give me only SQL Commands No text.")
	if err != nil {
		t.Errorf("Prompt() returned an error: %v", err)
	}
	fmt.Println(val)
}
