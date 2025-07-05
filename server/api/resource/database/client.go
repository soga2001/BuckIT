package database

import (
	"errors"
	"os"

	"github.com/supabase-community/supabase-go"
)

func GetClient() (*supabase.Client, error) {
	projectURL, ok := os.LookupEnv("SUPABASE_URL")
	if !ok {
		return nil, errors.New("SUPABASE_URL not set")
	}

	anonKey, ok := os.LookupEnv("SUPABASE_ANON_KEY")
	if !ok {
		return nil, errors.New("SUPABASE_ANON_KEY not set")
	}

	client, err := supabase.NewClient(projectURL, anonKey, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
